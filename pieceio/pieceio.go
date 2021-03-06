package pieceio

import (
	"context"
	"fmt"
	"io"

	"github.com/ipfs/go-cid"
	blockstore "github.com/ipfs/go-ipfs-blockstore"
	"github.com/ipld/go-ipld-prime"

	"github.com/filecoin-project/go-fil-markets/filestore"
)

type SectorCalculator interface {
	// GeneratePieceCommitment takes a PADDED io stream and a total size and generates a commP
	GeneratePieceCommitment(piece io.Reader, pieceSize uint64) ([]byte, error)
}

type PadReader interface {
	// PaddedSize returns the expected size of a piece after it's been padded
	PaddedSize(size uint64) uint64
}

type CarIO interface {
	// WriteCar writes a given payload to a CAR file and into the passed IO stream
	WriteCar(ctx context.Context, bs ReadStore, payloadCid cid.Cid, selector ipld.Node, w io.Writer) error
	// LoadCar loads blocks into the a store from a given CAR file
	LoadCar(bs WriteStore, r io.Reader) (cid.Cid, error)
}

type pieceIO struct {
	padReader        PadReader
	carIO            CarIO
	sectorCalculator SectorCalculator
	store            filestore.FileStore
	bs               blockstore.Blockstore
}

func NewPieceIO(padReader PadReader, carIO CarIO, sectorCalculator SectorCalculator, store filestore.FileStore, bs blockstore.Blockstore) PieceIO {
	return &pieceIO{padReader, carIO, sectorCalculator, store, bs}
}

func (pio *pieceIO) GeneratePieceCommitment(payloadCid cid.Cid, selector ipld.Node) ([]byte, filestore.File, error) {
	f, err := pio.store.CreateTemp()
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		f.Close()
		_ = pio.store.Delete(f.Path())
	}
	err = pio.carIO.WriteCar(context.Background(), pio.bs, payloadCid, selector, f)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	size := f.Size()
	pieceSize := uint64(size)
	paddedSize := pio.padReader.PaddedSize(pieceSize)
	remaining := paddedSize - pieceSize
	padbuf := make([]byte, remaining)
	padded, err := f.Write(padbuf)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	if uint64(padded) != remaining {
		cleanup()
		return nil, nil, fmt.Errorf("wrote %d byte of padding while expecting %d to be written", padded, remaining)
	}
	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	commitment, err := pio.sectorCalculator.GeneratePieceCommitment(f, paddedSize)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	return commitment, f, nil
}

func (pio *pieceIO) ReadPiece(r io.Reader) (cid.Cid, error) {
	return pio.carIO.LoadCar(pio.bs, r)
}
