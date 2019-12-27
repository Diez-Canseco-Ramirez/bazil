// Package clock implements a logical clock, tracking changes at multiple peers.
//
// Is is inspired by the paper "File Synchronization with Vector Time
// Pairs" and the Tra software. This structure avoids the need for
// tombstones.
//
//  http://publications.csail.mit.edu/tmp/MIT-CSAIL-TR-2005-014.pdf
//  http://swtch.com/tra/
package clock

import (
	"encoding"
	"encoding/binary"
	"errors"
	"fmt"
)

// A Peer is a replica that may create new versions of the tracked
// data. Peers are identified by small unsigned integers, for
// efficiency.
type Peer uint32

// MaxPeer is the largest possible value a Peer can have.
const MaxPeer = ^Peer(0)

// Epoch is a logical clock timestamp. Time 0 is never valid.
type Epoch uint64

var _ encoding.BinaryMarshaler = (*Epoch)(nil)

func (e *Epoch) MarshalBinary() ([]byte, error) {
	var tmp [8]byte
	binary.BigEndian.PutUint64(tmp[:], uint64(*e))
	return tmp[:], nil
}

var _ encoding.BinaryUnmarshaler = (*Epoch)(nil)

func (e *Epoch) UnmarshalBinary(data []byte) error {
	if len(data) != 8 {
		return errors.New("binary epoch is wrong length")
	}
	*e = Epoch(binary.BigEndian.Uint64(data))
	return nil
}

// Clock is a logical clock.
//
// The zero value is a valid empty clock, but most callers should call
// Create to get the creation time set up.
type Clock struct {
	sync vector
	mod  vector
	// create is always of length 1
	//
	// TODO could make it an item not a vector, but then we can't
	// use compareLE.
	create vector
}

// Create returns a new Vector Pair that knows it was created by id at
// time now.
func Create(id Peer, now Epoch) *Clock {
	c := &Clock{
		sync:   vector{list: []item{{id: id, t: now}}},
		mod:    vector{list: []item{{id: id, t: now}}},
		create: vector{list: []item{{id: id, t: now}}},
	}
	return c
}

// Update adds or updates the version vector entry for id to point to
// time now.
//
// As an optimization, it removes all the other modification time
// entries. This is only safe for files, not directories; see section
// 3.5.2 "Encoding Modification Times" of the Tra paper.
//
// Caller guarantees that one of the following is true:
//     - now is greater than any old value seen for peer id
//     - now is equal to an earlier value for this peer id, and no
//       other peer ids have been updated since that Update
func (c *Clock) Update(id Peer, now Epoch) {
	c.mod.updateSimplify(id, now)
	c.sync.update(id, now)
}

// UpdateParent is like Update, but does not simplify the modification
// time version vector. It is safe to use for directories and other
// entities where updates are not necessarily sequenced.
//
// Caller guarantees if an entry exists for id already, now is greater
// than or equal to the old value.
func (c *Clock) UpdateParent(id Peer, now Epoch) {
	c.mod.update(id, now)
	c.sync.update(id, now)
}

// UpdateSync updates the sync time only for id to point to time now.
//
// Caller guarantees if an entry exists for id already, now is greater
// than or equal to the old value.
func (c *Clock) UpdateSync(id Peer, now Epoch) {
	c.sync.update(id, now)
}

// UpdateFromChild tracks child modification times in the parent.
//
// Return value reports whether s changed.
func (c *Clock) UpdateFromChild(child *Clock) bool {
	changed := c.mod.merge(child.mod)
	return changed
}

// UpdateFromParent simplifies child sync times based on the parent.
func (c *Clock) UpdateFromParent(parent *Clock) {
	c.sync.rebase(parent.sync)
}

// Tombstone changes clock into a tombstone.
func (c *Clock) Tombstone() {
	// vpair paper section 3.3.2
	c.mod = vector{}
	c.create = vector{}
}

// ResolveTheirs records a conflict resolution in favor of other.
func (c *Clock) ResolveTheirs(other *Clock) {
	c.mod = vector{}
	c.mod.merge(other.mod)
	c.sync.merge(other.sync)
	// vpair paper is silent on what to do with create times; if we
	// don't do this, c.create can remain empty
	if len(c.create.list) == 0 {
		c.create.merge(other.create)
	}
}

// ResolveOurs records a conflict resolution in favor of us.
func (c *Clock) ResolveOurs(other *Clock) {
	// no change to c.mod
	c.sync.merge(other.sync)
	// no change to c.create
}

// ResolveNew records a conflict resolution in favor of newly created
// content.
func (c *Clock) ResolveNew(other *Clock) {
	c.mod.merge(other.mod)
	c.sync.merge(other.sync)
	// no change to c.create
}

func (c Clock) String() string {
	return fmt.Sprintf("{sync%s mod%s create%s}", c.sync, c.mod, c.create)
}

var (
	ErrRewritePeerNotMapped = errors.New("cannot rewrite peer id for an unknown peer")
)

// RewritePeers updates the peer identifiers in the clock based on the
// given mapping. This is useful because the short identifiers are not
// globally allocated.
//
// Returns ErrRewritePeerNotMapped if the clock contains a peer not
// present in the map. If an error occurs, the clock is in an
// undefined state and must not be used.
func (c *Clock) RewritePeers(m map[Peer]Peer) error {
	if err := c.sync.rewritePeers(m); err != nil {
		return err
	}
	if err := c.mod.rewritePeers(m); err != nil {
		return err
	}
	if err := c.create.rewritePeers(m); err != nil {
		return err
	}
	return nil
}

// TombstoneFromParent returns a new tombstone clock based on the parent.
func TombstoneFromParent(parent *Clock) *Clock {
	c := &Clock{}
	c.sync.list = append(c.sync.list, parent.sync.list...)
	c.mod.list = make([]item, 0)
	c.create.list = make([]item, 0)
	return c
}

// Action is a suggested action to take to combine two data items.
//
// The zero value of Action is not valid.
type Action int

//go:generate go build -o ../../task/tools/bin/stringer golang.org/x/tools/cmd/stringer
//go:generate ../../task/tools/bin/stringer -type=Action

const (
	invalidAction Action = iota
	// Copy means that the incoming version is newer, and its data
	// should be used.
	Copy
	// Nothing means the local version is newer (or same), and data
	// should not change.
	Nothing
	// Conflict means the two versions have diverged.
	Conflict
)

// Sync returns what action receiving state from A to B should cause
// us to take.
func Sync(a, b *Clock) Action {
	mod := a.mod
	if len(mod.list) == 0 {
		// it's a tombstone optimized clock, pretend mod is same as
		// sync
		mod = a.sync
	}
	if compareLE(mod, b.sync) {
		return Nothing
	}

	if compareLE(b.mod, a.sync) {
		return Copy
	}

	return Conflict
}

// SyncToMissing returns what action receiving state from A to B
// should cause us to take. B does not exist currently.
func SyncToMissing(a, b *Clock) Action {
	if compareLE(a.mod, b.sync) {
		return Nothing
	}

	if !compareLE(a.create, b.sync) {
		return Copy
	}

	return Conflict
}
