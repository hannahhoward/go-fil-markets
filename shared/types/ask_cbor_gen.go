package types

import (
	"fmt"
	"io"

	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

var _ = xerrors.Errorf

func (t *SignedStorageAsk) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{130}); err != nil {
		return err
	}

	// t.Ask (types.StorageAsk) (struct)
	if err := t.Ask.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Signature (types.Signature) (struct)
	if err := t.Signature.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *SignedStorageAsk) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Ask (types.StorageAsk) (struct)

	{

		pb, err := br.PeekByte()
		if err != nil {
			return err
		}
		if pb == cbg.CborNull[0] {
			var nbuf [1]byte
			if _, err := br.Read(nbuf[:]); err != nil {
				return err
			}
		} else {
			t.Ask = new(StorageAsk)
			if err := t.Ask.UnmarshalCBOR(br); err != nil {
				return err
			}
		}

	}
	// t.Signature (types.Signature) (struct)

	{

		pb, err := br.PeekByte()
		if err != nil {
			return err
		}
		if pb == cbg.CborNull[0] {
			var nbuf [1]byte
			if _, err := br.Read(nbuf[:]); err != nil {
				return err
			}
		} else {
			t.Signature = new(Signature)
			if err := t.Signature.UnmarshalCBOR(br); err != nil {
				return err
			}
		}

	}
	return nil
}

func (t *StorageAsk) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{134}); err != nil {
		return err
	}

	// t.Price (tokenamount.TokenAmount) (struct)
	if err := t.Price.MarshalCBOR(w); err != nil {
		return err
	}

	// t.MinPieceSize (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.MinPieceSize))); err != nil {
		return err
	}

	// t.Miner (address.Address) (struct)
	if err := t.Miner.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Timestamp (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Timestamp))); err != nil {
		return err
	}

	// t.Expiry (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Expiry))); err != nil {
		return err
	}

	// t.SeqNo (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.SeqNo))); err != nil {
		return err
	}
	return nil
}

func (t *StorageAsk) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 6 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Price (tokenamount.TokenAmount) (struct)

	{

		if err := t.Price.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.MinPieceSize (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.MinPieceSize = uint64(extra)
	// t.Miner (address.Address) (struct)

	{

		if err := t.Miner.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.Timestamp (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.Timestamp = uint64(extra)
	// t.Expiry (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.Expiry = uint64(extra)
	// t.SeqNo (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.SeqNo = uint64(extra)
	return nil
}
