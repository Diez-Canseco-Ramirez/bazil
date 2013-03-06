package tokens

const (
	// The DB bucket that contains general-purpose Bazil data not tied to
	// any specific volume.
	BucketBazil = "bazil"

	// The DB bucket that contains a sub-bucket per volume, named by
	// the volume ID.
	BucketVolume = "volume"
)
