// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// MessagesReorderPinnedSavedDialogsRequest represents TL type `messages.reorderPinnedSavedDialogs#8b716587`.
//
// See https://core.telegram.org/method/messages.reorderPinnedSavedDialogs for reference.
type MessagesReorderPinnedSavedDialogsRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Force field of MessagesReorderPinnedSavedDialogsRequest.
	Force bool
	// Order field of MessagesReorderPinnedSavedDialogsRequest.
	Order []InputDialogPeerClass
}

// MessagesReorderPinnedSavedDialogsRequestTypeID is TL type id of MessagesReorderPinnedSavedDialogsRequest.
const MessagesReorderPinnedSavedDialogsRequestTypeID = 0x8b716587

// Ensuring interfaces in compile-time for MessagesReorderPinnedSavedDialogsRequest.
var (
	_ bin.Encoder     = &MessagesReorderPinnedSavedDialogsRequest{}
	_ bin.Decoder     = &MessagesReorderPinnedSavedDialogsRequest{}
	_ bin.BareEncoder = &MessagesReorderPinnedSavedDialogsRequest{}
	_ bin.BareDecoder = &MessagesReorderPinnedSavedDialogsRequest{}
)

func (r *MessagesReorderPinnedSavedDialogsRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Flags.Zero()) {
		return false
	}
	if !(r.Force == false) {
		return false
	}
	if !(r.Order == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *MessagesReorderPinnedSavedDialogsRequest) String() string {
	if r == nil {
		return "MessagesReorderPinnedSavedDialogsRequest(nil)"
	}
	type Alias MessagesReorderPinnedSavedDialogsRequest
	return fmt.Sprintf("MessagesReorderPinnedSavedDialogsRequest%+v", Alias(*r))
}

// FillFrom fills MessagesReorderPinnedSavedDialogsRequest from given interface.
func (r *MessagesReorderPinnedSavedDialogsRequest) FillFrom(from interface {
	GetForce() (value bool)
	GetOrder() (value []InputDialogPeerClass)
}) {
	r.Force = from.GetForce()
	r.Order = from.GetOrder()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesReorderPinnedSavedDialogsRequest) TypeID() uint32 {
	return MessagesReorderPinnedSavedDialogsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesReorderPinnedSavedDialogsRequest) TypeName() string {
	return "messages.reorderPinnedSavedDialogs"
}

// TypeInfo returns info about TL type.
func (r *MessagesReorderPinnedSavedDialogsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.reorderPinnedSavedDialogs",
		ID:   MessagesReorderPinnedSavedDialogsRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Force",
			SchemaName: "force",
			Null:       !r.Flags.Has(0),
		},
		{
			Name:       "Order",
			SchemaName: "order",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (r *MessagesReorderPinnedSavedDialogsRequest) SetFlags() {
	if !(r.Force == false) {
		r.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (r *MessagesReorderPinnedSavedDialogsRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.reorderPinnedSavedDialogs#8b716587 as nil")
	}
	b.PutID(MessagesReorderPinnedSavedDialogsRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *MessagesReorderPinnedSavedDialogsRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.reorderPinnedSavedDialogs#8b716587 as nil")
	}
	r.SetFlags()
	if err := r.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.reorderPinnedSavedDialogs#8b716587: field flags: %w", err)
	}
	b.PutVectorHeader(len(r.Order))
	for idx, v := range r.Order {
		if v == nil {
			return fmt.Errorf("unable to encode messages.reorderPinnedSavedDialogs#8b716587: field order element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.reorderPinnedSavedDialogs#8b716587: field order element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *MessagesReorderPinnedSavedDialogsRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.reorderPinnedSavedDialogs#8b716587 to nil")
	}
	if err := b.ConsumeID(MessagesReorderPinnedSavedDialogsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.reorderPinnedSavedDialogs#8b716587: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *MessagesReorderPinnedSavedDialogsRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.reorderPinnedSavedDialogs#8b716587 to nil")
	}
	{
		if err := r.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.reorderPinnedSavedDialogs#8b716587: field flags: %w", err)
		}
	}
	r.Force = r.Flags.Has(0)
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.reorderPinnedSavedDialogs#8b716587: field order: %w", err)
		}

		if headerLen > 0 {
			r.Order = make([]InputDialogPeerClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputDialogPeer(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.reorderPinnedSavedDialogs#8b716587: field order: %w", err)
			}
			r.Order = append(r.Order, value)
		}
	}
	return nil
}

// SetForce sets value of Force conditional field.
func (r *MessagesReorderPinnedSavedDialogsRequest) SetForce(value bool) {
	if value {
		r.Flags.Set(0)
		r.Force = true
	} else {
		r.Flags.Unset(0)
		r.Force = false
	}
}

// GetForce returns value of Force conditional field.
func (r *MessagesReorderPinnedSavedDialogsRequest) GetForce() (value bool) {
	if r == nil {
		return
	}
	return r.Flags.Has(0)
}

// GetOrder returns value of Order field.
func (r *MessagesReorderPinnedSavedDialogsRequest) GetOrder() (value []InputDialogPeerClass) {
	if r == nil {
		return
	}
	return r.Order
}

// MapOrder returns field Order wrapped in InputDialogPeerClassArray helper.
func (r *MessagesReorderPinnedSavedDialogsRequest) MapOrder() (value InputDialogPeerClassArray) {
	return InputDialogPeerClassArray(r.Order)
}

// MessagesReorderPinnedSavedDialogs invokes method messages.reorderPinnedSavedDialogs#8b716587 returning error if any.
//
// See https://core.telegram.org/method/messages.reorderPinnedSavedDialogs for reference.
// Can be used by bots.
func (c *Client) MessagesReorderPinnedSavedDialogs(ctx context.Context, request *MessagesReorderPinnedSavedDialogsRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
