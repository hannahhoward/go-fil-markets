package retrievalimpl

import (
	"fmt"
	"io"

	"github.com/filecoin-project/go-fil-markets/shared/types"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

var _ = xerrors.Errorf

func (t *RetParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{129}); err != nil {
		return err
	}

	// t.Unixfs0 (retrievalimpl.Unixfs0Offer) (struct)
	if err := t.Unixfs0.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *RetParams) UnmarshalCBOR(r io.Reader) error {
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

	// t.Unixfs0 (retrievalimpl.Unixfs0Offer) (struct)

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
			t.Unixfs0 = new(Unixfs0Offer)
			if err := t.Unixfs0.UnmarshalCBOR(br); err != nil {
				return err
			}
		}

	}
	return nil
}

func (t *Unixfs0Offer) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{130}); err != nil {
		return err
	}

	// t.Offset (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Offset))); err != nil {
		return err
	}

	// t.Size (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Size))); err != nil {
		return err
	}
	return nil
}

func (t *Unixfs0Offer) UnmarshalCBOR(r io.Reader) error {
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

	// t.Offset (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.Offset = uint64(extra)
	// t.Size (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.Size = uint64(extra)
	return nil
}

func (t *OldDealProposal) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{131}); err != nil {
		return err
	}

	// t.Payment (retrievalimpl.OldPaymentInfo) (struct)
	if err := t.Payment.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Ref (cid.Cid) (struct)

	if err := cbg.WriteCid(w, t.Ref); err != nil {
		return xerrors.Errorf("failed to write cid field t.Ref: %w", err)
	}

	// t.Params (retrievalimpl.RetParams) (struct)
	if err := t.Params.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *OldDealProposal) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Payment (retrievalimpl.OldPaymentInfo) (struct)

	{

		if err := t.Payment.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.Ref (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.Ref: %w", err)
		}

		t.Ref = c

	}
	// t.Params (retrievalimpl.RetParams) (struct)

	{

		if err := t.Params.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	return nil
}

func (t *OldDealResponse) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{130}); err != nil {
		return err
	}

	// t.Status (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Status))); err != nil {
		return err
	}

	// t.Message (string) (string)
	if len(t.Message) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Message was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len(t.Message)))); err != nil {
		return err
	}
	if _, err := w.Write([]byte(t.Message)); err != nil {
		return err
	}
	return nil
}

func (t *OldDealResponse) UnmarshalCBOR(r io.Reader) error {
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

	// t.Status (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.Status = uint64(extra)
	// t.Message (string) (string)

	{
		sval, err := cbg.ReadString(br)
		if err != nil {
			return err
		}

		t.Message = string(sval)
	}
	return nil
}

func (t *Block) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{130}); err != nil {
		return err
	}

	// t.Prefix ([]uint8) (slice)
	if len(t.Prefix) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Prefix was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajByteString, uint64(len(t.Prefix)))); err != nil {
		return err
	}
	if _, err := w.Write(t.Prefix); err != nil {
		return err
	}

	// t.Data ([]uint8) (slice)
	if len(t.Data) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Data was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajByteString, uint64(len(t.Data)))); err != nil {
		return err
	}
	if _, err := w.Write(t.Data); err != nil {
		return err
	}
	return nil
}

func (t *Block) UnmarshalCBOR(r io.Reader) error {
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

	// t.Prefix ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Prefix: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}
	t.Prefix = make([]byte, extra)
	if _, err := io.ReadFull(br, t.Prefix); err != nil {
		return err
	}
	// t.Data ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Data: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}
	t.Data = make([]byte, extra)
	if _, err := io.ReadFull(br, t.Data); err != nil {
		return err
	}
	return nil
}

func (t *OldQueryResponse) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{131}); err != nil {
		return err
	}

	// t.Status (retrievalimpl.OldQueryResponseStatus) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Status))); err != nil {
		return err
	}

	// t.Size (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Size))); err != nil {
		return err
	}

	// t.MinPrice (tokenamount.TokenAmount) (struct)
	if err := t.MinPrice.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *OldQueryResponse) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Status (retrievalimpl.OldQueryResponseStatus) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.Status = OldQueryResponseStatus(extra)
	// t.Size (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.Size = uint64(extra)
	// t.MinPrice (tokenamount.TokenAmount) (struct)

	{

		if err := t.MinPrice.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	return nil
}

func (t *OldQuery) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{129}); err != nil {
		return err
	}

	// t.Piece (cid.Cid) (struct)

	if err := cbg.WriteCid(w, t.Piece); err != nil {
		return xerrors.Errorf("failed to write cid field t.Piece: %w", err)
	}

	return nil
}

func (t *OldQuery) UnmarshalCBOR(r io.Reader) error {
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

	// t.Piece (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.Piece: %w", err)
		}

		t.Piece = c

	}
	return nil
}

func (t *OldPaymentInfo) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{131}); err != nil {
		return err
	}

	// t.Channel (address.Address) (struct)
	if err := t.Channel.MarshalCBOR(w); err != nil {
		return err
	}

	// t.ChannelMessage (cid.Cid) (struct)

	if t.ChannelMessage == nil {
		if _, err := w.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCid(w, *t.ChannelMessage); err != nil {
			return xerrors.Errorf("failed to write cid field t.ChannelMessage: %w", err)
		}
	}

	// t.Vouchers ([]*types.SignedVoucher) (slice)
	if len(t.Vouchers) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Vouchers was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajArray, uint64(len(t.Vouchers)))); err != nil {
		return err
	}
	for _, v := range t.Vouchers {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}
	return nil
}

func (t *OldPaymentInfo) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Channel (address.Address) (struct)

	{

		if err := t.Channel.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.ChannelMessage (cid.Cid) (struct)

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

			c, err := cbg.ReadCid(br)
			if err != nil {
				return xerrors.Errorf("failed to read cid field t.ChannelMessage: %w", err)
			}

			t.ChannelMessage = &c
		}

	}
	// t.Vouchers ([]*types.SignedVoucher) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Vouchers: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}
	if extra > 0 {
		t.Vouchers = make([]*types.SignedVoucher, extra)
	}
	for i := 0; i < int(extra); i++ {

		var v types.SignedVoucher
		if err := v.UnmarshalCBOR(br); err != nil {
			return err
		}

		t.Vouchers[i] = &v
	}

	return nil
}
