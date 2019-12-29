package control

import (
	"context"
	"log"

	"bazil.org/bazil/db"
	"bazil.org/bazil/server/control/wire"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

const sharingKeySize = 32

func (c controlRPC) SharingKeyAdd(ctx context.Context, req *wire.SharingKeyAddRequest) (*wire.SharingKeyAddResponse, error) {
	if len(req.Secret) != sharingKeySize {
		return nil, grpc.Errorf(codes.InvalidArgument, "sharing key must be exactly 32 bytes")
	}

	var secret [32]byte
	copy(secret[:], req.Secret)

	update := func(tx *db.Tx) error {
		if _, err := tx.SharingKeys().Add(req.Name, &secret); err != nil {
			return err
		}
		return nil
	}
	if err := c.app.DB.Update(update); err != nil {
		switch err {
		case db.ErrSharingKeyExist:
			return nil, grpc.Errorf(codes.AlreadyExists, "sharing key exists already: %x", req.Name)
		case db.ErrSharingKeyNameInvalid:
			return nil, grpc.Errorf(codes.InvalidArgument, "%v", err)
		}
		if grpc.Code(err) != codes.Unknown {
			return nil, err
		}
		log.Printf("db update error: put sharing key %q: %v", req.Name, err)
		return nil, grpc.Errorf(codes.Internal, "database error")
	}
	return &wire.SharingKeyAddResponse{}, nil
}
