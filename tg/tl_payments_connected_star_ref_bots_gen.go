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

// PaymentsConnectedStarRefBots represents TL type `payments.connectedStarRefBots#98d5ea1d`.
//
// See https://core.telegram.org/constructor/payments.connectedStarRefBots for reference.
type PaymentsConnectedStarRefBots struct {
	// Count field of PaymentsConnectedStarRefBots.
	Count int
	// ConnectedBots field of PaymentsConnectedStarRefBots.
	ConnectedBots []ConnectedBotStarRef
	// Users field of PaymentsConnectedStarRefBots.
	Users []UserClass
}

// PaymentsConnectedStarRefBotsTypeID is TL type id of PaymentsConnectedStarRefBots.
const PaymentsConnectedStarRefBotsTypeID = 0x98d5ea1d

// Ensuring interfaces in compile-time for PaymentsConnectedStarRefBots.
var (
	_ bin.Encoder     = &PaymentsConnectedStarRefBots{}
	_ bin.Decoder     = &PaymentsConnectedStarRefBots{}
	_ bin.BareEncoder = &PaymentsConnectedStarRefBots{}
	_ bin.BareDecoder = &PaymentsConnectedStarRefBots{}
)

func (c *PaymentsConnectedStarRefBots) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Count == 0) {
		return false
	}
	if !(c.ConnectedBots == nil) {
		return false
	}
	if !(c.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *PaymentsConnectedStarRefBots) String() string {
	if c == nil {
		return "PaymentsConnectedStarRefBots(nil)"
	}
	type Alias PaymentsConnectedStarRefBots
	return fmt.Sprintf("PaymentsConnectedStarRefBots%+v", Alias(*c))
}

// FillFrom fills PaymentsConnectedStarRefBots from given interface.
func (c *PaymentsConnectedStarRefBots) FillFrom(from interface {
	GetCount() (value int)
	GetConnectedBots() (value []ConnectedBotStarRef)
	GetUsers() (value []UserClass)
}) {
	c.Count = from.GetCount()
	c.ConnectedBots = from.GetConnectedBots()
	c.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentsConnectedStarRefBots) TypeID() uint32 {
	return PaymentsConnectedStarRefBotsTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentsConnectedStarRefBots) TypeName() string {
	return "payments.connectedStarRefBots"
}

// TypeInfo returns info about TL type.
func (c *PaymentsConnectedStarRefBots) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "payments.connectedStarRefBots",
		ID:   PaymentsConnectedStarRefBotsTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Count",
			SchemaName: "count",
		},
		{
			Name:       "ConnectedBots",
			SchemaName: "connected_bots",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *PaymentsConnectedStarRefBots) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode payments.connectedStarRefBots#98d5ea1d as nil")
	}
	b.PutID(PaymentsConnectedStarRefBotsTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *PaymentsConnectedStarRefBots) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode payments.connectedStarRefBots#98d5ea1d as nil")
	}
	b.PutInt(c.Count)
	b.PutVectorHeader(len(c.ConnectedBots))
	for idx, v := range c.ConnectedBots {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode payments.connectedStarRefBots#98d5ea1d: field connected_bots element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(c.Users))
	for idx, v := range c.Users {
		if v == nil {
			return fmt.Errorf("unable to encode payments.connectedStarRefBots#98d5ea1d: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode payments.connectedStarRefBots#98d5ea1d: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *PaymentsConnectedStarRefBots) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode payments.connectedStarRefBots#98d5ea1d to nil")
	}
	if err := b.ConsumeID(PaymentsConnectedStarRefBotsTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.connectedStarRefBots#98d5ea1d: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *PaymentsConnectedStarRefBots) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode payments.connectedStarRefBots#98d5ea1d to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode payments.connectedStarRefBots#98d5ea1d: field count: %w", err)
		}
		c.Count = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode payments.connectedStarRefBots#98d5ea1d: field connected_bots: %w", err)
		}

		if headerLen > 0 {
			c.ConnectedBots = make([]ConnectedBotStarRef, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value ConnectedBotStarRef
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode payments.connectedStarRefBots#98d5ea1d: field connected_bots: %w", err)
			}
			c.ConnectedBots = append(c.ConnectedBots, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode payments.connectedStarRefBots#98d5ea1d: field users: %w", err)
		}

		if headerLen > 0 {
			c.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode payments.connectedStarRefBots#98d5ea1d: field users: %w", err)
			}
			c.Users = append(c.Users, value)
		}
	}
	return nil
}

// GetCount returns value of Count field.
func (c *PaymentsConnectedStarRefBots) GetCount() (value int) {
	if c == nil {
		return
	}
	return c.Count
}

// GetConnectedBots returns value of ConnectedBots field.
func (c *PaymentsConnectedStarRefBots) GetConnectedBots() (value []ConnectedBotStarRef) {
	if c == nil {
		return
	}
	return c.ConnectedBots
}

// GetUsers returns value of Users field.
func (c *PaymentsConnectedStarRefBots) GetUsers() (value []UserClass) {
	if c == nil {
		return
	}
	return c.Users
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (c *PaymentsConnectedStarRefBots) MapUsers() (value UserClassArray) {
	return UserClassArray(c.Users)
}
