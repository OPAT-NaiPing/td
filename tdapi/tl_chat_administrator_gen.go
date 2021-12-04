// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// ChatAdministrator represents TL type `chatAdministrator#7277bd2c`.
type ChatAdministrator struct {
	// User identifier of the administrator
	UserID int64
	// Custom title of the administrator
	CustomTitle string
	// True, if the user is the owner of the chat
	IsOwner bool
}

// ChatAdministratorTypeID is TL type id of ChatAdministrator.
const ChatAdministratorTypeID = 0x7277bd2c

// Ensuring interfaces in compile-time for ChatAdministrator.
var (
	_ bin.Encoder     = &ChatAdministrator{}
	_ bin.Decoder     = &ChatAdministrator{}
	_ bin.BareEncoder = &ChatAdministrator{}
	_ bin.BareDecoder = &ChatAdministrator{}
)

func (c *ChatAdministrator) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.UserID == 0) {
		return false
	}
	if !(c.CustomTitle == "") {
		return false
	}
	if !(c.IsOwner == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatAdministrator) String() string {
	if c == nil {
		return "ChatAdministrator(nil)"
	}
	type Alias ChatAdministrator
	return fmt.Sprintf("ChatAdministrator%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatAdministrator) TypeID() uint32 {
	return ChatAdministratorTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatAdministrator) TypeName() string {
	return "chatAdministrator"
}

// TypeInfo returns info about TL type.
func (c *ChatAdministrator) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatAdministrator",
		ID:   ChatAdministratorTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "CustomTitle",
			SchemaName: "custom_title",
		},
		{
			Name:       "IsOwner",
			SchemaName: "is_owner",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatAdministrator) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatAdministrator#7277bd2c as nil")
	}
	b.PutID(ChatAdministratorTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatAdministrator) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatAdministrator#7277bd2c as nil")
	}
	b.PutInt53(c.UserID)
	b.PutString(c.CustomTitle)
	b.PutBool(c.IsOwner)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatAdministrator) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatAdministrator#7277bd2c to nil")
	}
	if err := b.ConsumeID(ChatAdministratorTypeID); err != nil {
		return fmt.Errorf("unable to decode chatAdministrator#7277bd2c: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatAdministrator) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatAdministrator#7277bd2c to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode chatAdministrator#7277bd2c: field user_id: %w", err)
		}
		c.UserID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode chatAdministrator#7277bd2c: field custom_title: %w", err)
		}
		c.CustomTitle = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatAdministrator#7277bd2c: field is_owner: %w", err)
		}
		c.IsOwner = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatAdministrator) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatAdministrator#7277bd2c as nil")
	}
	b.ObjStart()
	b.PutID("chatAdministrator")
	b.FieldStart("user_id")
	b.PutInt53(c.UserID)
	b.FieldStart("custom_title")
	b.PutString(c.CustomTitle)
	b.FieldStart("is_owner")
	b.PutBool(c.IsOwner)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatAdministrator) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatAdministrator#7277bd2c to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatAdministrator"); err != nil {
				return fmt.Errorf("unable to decode chatAdministrator#7277bd2c: %w", err)
			}
		case "user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode chatAdministrator#7277bd2c: field user_id: %w", err)
			}
			c.UserID = value
		case "custom_title":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode chatAdministrator#7277bd2c: field custom_title: %w", err)
			}
			c.CustomTitle = value
		case "is_owner":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatAdministrator#7277bd2c: field is_owner: %w", err)
			}
			c.IsOwner = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUserID returns value of UserID field.
func (c *ChatAdministrator) GetUserID() (value int64) {
	return c.UserID
}

// GetCustomTitle returns value of CustomTitle field.
func (c *ChatAdministrator) GetCustomTitle() (value string) {
	return c.CustomTitle
}

// GetIsOwner returns value of IsOwner field.
func (c *ChatAdministrator) GetIsOwner() (value bool) {
	return c.IsOwner
}
