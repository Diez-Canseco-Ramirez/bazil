package fs

import (
	"context"
	"os"

	"bazil.org/bazil/util/env"
	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

type dotBazil struct {
	fs     *Volume
	parent *dir
}

var _ = fs.Node(&dotBazil{})

func (d *dotBazil) Attr(ctx context.Context, a *fuse.Attr) error {
	a.Mode = os.ModeDir | 0555
	a.Uid = env.MyUID
	a.Uid = env.MyGID
	return nil
}

var _ = fs.NodeStringLookuper(&dotBazil{})

func (d *dotBazil) Lookup(ctx context.Context, name string) (fs.Node, error) {
	switch name {
	case "pending":
		child := &pendingList{
			dir: d.parent,
		}
		return child, nil
	default:
		return nil, fuse.ENOENT
	}
}

var _ fs.HandleReadDirAller = (*dotBazil)(nil)

func (d *dotBazil) ReadDirAll(ctx context.Context) ([]fuse.Dirent, error) {
	r := []fuse.Dirent{
		{Name: "pending", Type: fuse.DT_Dir},
	}
	return r, nil
}
