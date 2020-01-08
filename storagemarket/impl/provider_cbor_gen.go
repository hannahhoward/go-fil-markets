package storageimpl

import (
	"fmt"
	"io"

	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

var _ = xerrors.Errorf

func (t *MinerDeal) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{129}); err != nil {
		return err
	}

	// t.MinerDeal (storagemarket.MinerDeal) (struct)
	if err := t.MinerDeal.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *MinerDeal) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 1 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.MinerDeal (storagemarket.MinerDeal) (struct)

	{

		if err := t.MinerDeal.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	return nil
}