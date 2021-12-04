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

// ChatStatisticsInviterInfo represents TL type `chatStatisticsInviterInfo#2583d48b`.
type ChatStatisticsInviterInfo struct {
	// User identifier
	UserID int64
	// Number of new members invited by the user
	AddedMemberCount int32
}

// ChatStatisticsInviterInfoTypeID is TL type id of ChatStatisticsInviterInfo.
const ChatStatisticsInviterInfoTypeID = 0x2583d48b

// Ensuring interfaces in compile-time for ChatStatisticsInviterInfo.
var (
	_ bin.Encoder     = &ChatStatisticsInviterInfo{}
	_ bin.Decoder     = &ChatStatisticsInviterInfo{}
	_ bin.BareEncoder = &ChatStatisticsInviterInfo{}
	_ bin.BareDecoder = &ChatStatisticsInviterInfo{}
)

func (c *ChatStatisticsInviterInfo) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.UserID == 0) {
		return false
	}
	if !(c.AddedMemberCount == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatStatisticsInviterInfo) String() string {
	if c == nil {
		return "ChatStatisticsInviterInfo(nil)"
	}
	type Alias ChatStatisticsInviterInfo
	return fmt.Sprintf("ChatStatisticsInviterInfo%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatStatisticsInviterInfo) TypeID() uint32 {
	return ChatStatisticsInviterInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatStatisticsInviterInfo) TypeName() string {
	return "chatStatisticsInviterInfo"
}

// TypeInfo returns info about TL type.
func (c *ChatStatisticsInviterInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatStatisticsInviterInfo",
		ID:   ChatStatisticsInviterInfoTypeID,
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
			Name:       "AddedMemberCount",
			SchemaName: "added_member_count",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatStatisticsInviterInfo) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatStatisticsInviterInfo#2583d48b as nil")
	}
	b.PutID(ChatStatisticsInviterInfoTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatStatisticsInviterInfo) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatStatisticsInviterInfo#2583d48b as nil")
	}
	b.PutInt53(c.UserID)
	b.PutInt32(c.AddedMemberCount)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatStatisticsInviterInfo) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatStatisticsInviterInfo#2583d48b to nil")
	}
	if err := b.ConsumeID(ChatStatisticsInviterInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode chatStatisticsInviterInfo#2583d48b: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatStatisticsInviterInfo) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatStatisticsInviterInfo#2583d48b to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode chatStatisticsInviterInfo#2583d48b: field user_id: %w", err)
		}
		c.UserID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatStatisticsInviterInfo#2583d48b: field added_member_count: %w", err)
		}
		c.AddedMemberCount = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatStatisticsInviterInfo) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatStatisticsInviterInfo#2583d48b as nil")
	}
	b.ObjStart()
	b.PutID("chatStatisticsInviterInfo")
	b.FieldStart("user_id")
	b.PutInt53(c.UserID)
	b.FieldStart("added_member_count")
	b.PutInt32(c.AddedMemberCount)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatStatisticsInviterInfo) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatStatisticsInviterInfo#2583d48b to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatStatisticsInviterInfo"); err != nil {
				return fmt.Errorf("unable to decode chatStatisticsInviterInfo#2583d48b: %w", err)
			}
		case "user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode chatStatisticsInviterInfo#2583d48b: field user_id: %w", err)
			}
			c.UserID = value
		case "added_member_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatStatisticsInviterInfo#2583d48b: field added_member_count: %w", err)
			}
			c.AddedMemberCount = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUserID returns value of UserID field.
func (c *ChatStatisticsInviterInfo) GetUserID() (value int64) {
	return c.UserID
}

// GetAddedMemberCount returns value of AddedMemberCount field.
func (c *ChatStatisticsInviterInfo) GetAddedMemberCount() (value int32) {
	return c.AddedMemberCount
}
