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

// PaymentsGetUserStarGiftsRequest represents TL type `payments.getUserStarGifts#5e72c7e1`.
//
// See https://core.telegram.org/method/payments.getUserStarGifts for reference.
type PaymentsGetUserStarGiftsRequest struct {
	// UserID field of PaymentsGetUserStarGiftsRequest.
	UserID InputUserClass
	// Offset field of PaymentsGetUserStarGiftsRequest.
	Offset string
	// Limit field of PaymentsGetUserStarGiftsRequest.
	Limit int
}

// PaymentsGetUserStarGiftsRequestTypeID is TL type id of PaymentsGetUserStarGiftsRequest.
const PaymentsGetUserStarGiftsRequestTypeID = 0x5e72c7e1

// Ensuring interfaces in compile-time for PaymentsGetUserStarGiftsRequest.
var (
	_ bin.Encoder     = &PaymentsGetUserStarGiftsRequest{}
	_ bin.Decoder     = &PaymentsGetUserStarGiftsRequest{}
	_ bin.BareEncoder = &PaymentsGetUserStarGiftsRequest{}
	_ bin.BareDecoder = &PaymentsGetUserStarGiftsRequest{}
)

func (g *PaymentsGetUserStarGiftsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.UserID == nil) {
		return false
	}
	if !(g.Offset == "") {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *PaymentsGetUserStarGiftsRequest) String() string {
	if g == nil {
		return "PaymentsGetUserStarGiftsRequest(nil)"
	}
	type Alias PaymentsGetUserStarGiftsRequest
	return fmt.Sprintf("PaymentsGetUserStarGiftsRequest%+v", Alias(*g))
}

// FillFrom fills PaymentsGetUserStarGiftsRequest from given interface.
func (g *PaymentsGetUserStarGiftsRequest) FillFrom(from interface {
	GetUserID() (value InputUserClass)
	GetOffset() (value string)
	GetLimit() (value int)
}) {
	g.UserID = from.GetUserID()
	g.Offset = from.GetOffset()
	g.Limit = from.GetLimit()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentsGetUserStarGiftsRequest) TypeID() uint32 {
	return PaymentsGetUserStarGiftsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentsGetUserStarGiftsRequest) TypeName() string {
	return "payments.getUserStarGifts"
}

// TypeInfo returns info about TL type.
func (g *PaymentsGetUserStarGiftsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "payments.getUserStarGifts",
		ID:   PaymentsGetUserStarGiftsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "Offset",
			SchemaName: "offset",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *PaymentsGetUserStarGiftsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode payments.getUserStarGifts#5e72c7e1 as nil")
	}
	b.PutID(PaymentsGetUserStarGiftsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *PaymentsGetUserStarGiftsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode payments.getUserStarGifts#5e72c7e1 as nil")
	}
	if g.UserID == nil {
		return fmt.Errorf("unable to encode payments.getUserStarGifts#5e72c7e1: field user_id is nil")
	}
	if err := g.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.getUserStarGifts#5e72c7e1: field user_id: %w", err)
	}
	b.PutString(g.Offset)
	b.PutInt(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *PaymentsGetUserStarGiftsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode payments.getUserStarGifts#5e72c7e1 to nil")
	}
	if err := b.ConsumeID(PaymentsGetUserStarGiftsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.getUserStarGifts#5e72c7e1: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *PaymentsGetUserStarGiftsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode payments.getUserStarGifts#5e72c7e1 to nil")
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode payments.getUserStarGifts#5e72c7e1: field user_id: %w", err)
		}
		g.UserID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode payments.getUserStarGifts#5e72c7e1: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode payments.getUserStarGifts#5e72c7e1: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// GetUserID returns value of UserID field.
func (g *PaymentsGetUserStarGiftsRequest) GetUserID() (value InputUserClass) {
	if g == nil {
		return
	}
	return g.UserID
}

// GetOffset returns value of Offset field.
func (g *PaymentsGetUserStarGiftsRequest) GetOffset() (value string) {
	if g == nil {
		return
	}
	return g.Offset
}

// GetLimit returns value of Limit field.
func (g *PaymentsGetUserStarGiftsRequest) GetLimit() (value int) {
	if g == nil {
		return
	}
	return g.Limit
}

// PaymentsGetUserStarGifts invokes method payments.getUserStarGifts#5e72c7e1 returning error if any.
//
// See https://core.telegram.org/method/payments.getUserStarGifts for reference.
func (c *Client) PaymentsGetUserStarGifts(ctx context.Context, request *PaymentsGetUserStarGiftsRequest) (*PaymentsUserStarGifts, error) {
	var result PaymentsUserStarGifts

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}