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

// ChatStatisticsMessageSenderInfo represents TL type `chatStatisticsMessageSenderInfo#690a7e4b`.
type ChatStatisticsMessageSenderInfo struct {
	// User identifier
	UserID int64
	// Number of sent messages
	SentMessageCount int32
	// Average number of characters in sent messages; 0 if unknown
	AverageCharacterCount int32
}

// ChatStatisticsMessageSenderInfoTypeID is TL type id of ChatStatisticsMessageSenderInfo.
const ChatStatisticsMessageSenderInfoTypeID = 0x690a7e4b

// Ensuring interfaces in compile-time for ChatStatisticsMessageSenderInfo.
var (
	_ bin.Encoder     = &ChatStatisticsMessageSenderInfo{}
	_ bin.Decoder     = &ChatStatisticsMessageSenderInfo{}
	_ bin.BareEncoder = &ChatStatisticsMessageSenderInfo{}
	_ bin.BareDecoder = &ChatStatisticsMessageSenderInfo{}
)

func (c *ChatStatisticsMessageSenderInfo) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.UserID == 0) {
		return false
	}
	if !(c.SentMessageCount == 0) {
		return false
	}
	if !(c.AverageCharacterCount == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatStatisticsMessageSenderInfo) String() string {
	if c == nil {
		return "ChatStatisticsMessageSenderInfo(nil)"
	}
	type Alias ChatStatisticsMessageSenderInfo
	return fmt.Sprintf("ChatStatisticsMessageSenderInfo%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatStatisticsMessageSenderInfo) TypeID() uint32 {
	return ChatStatisticsMessageSenderInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatStatisticsMessageSenderInfo) TypeName() string {
	return "chatStatisticsMessageSenderInfo"
}

// TypeInfo returns info about TL type.
func (c *ChatStatisticsMessageSenderInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatStatisticsMessageSenderInfo",
		ID:   ChatStatisticsMessageSenderInfoTypeID,
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
			Name:       "SentMessageCount",
			SchemaName: "sent_message_count",
		},
		{
			Name:       "AverageCharacterCount",
			SchemaName: "average_character_count",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatStatisticsMessageSenderInfo) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatStatisticsMessageSenderInfo#690a7e4b as nil")
	}
	b.PutID(ChatStatisticsMessageSenderInfoTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatStatisticsMessageSenderInfo) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatStatisticsMessageSenderInfo#690a7e4b as nil")
	}
	b.PutInt53(c.UserID)
	b.PutInt32(c.SentMessageCount)
	b.PutInt32(c.AverageCharacterCount)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatStatisticsMessageSenderInfo) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatStatisticsMessageSenderInfo#690a7e4b to nil")
	}
	if err := b.ConsumeID(ChatStatisticsMessageSenderInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode chatStatisticsMessageSenderInfo#690a7e4b: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatStatisticsMessageSenderInfo) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatStatisticsMessageSenderInfo#690a7e4b to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode chatStatisticsMessageSenderInfo#690a7e4b: field user_id: %w", err)
		}
		c.UserID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatStatisticsMessageSenderInfo#690a7e4b: field sent_message_count: %w", err)
		}
		c.SentMessageCount = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatStatisticsMessageSenderInfo#690a7e4b: field average_character_count: %w", err)
		}
		c.AverageCharacterCount = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatStatisticsMessageSenderInfo) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatStatisticsMessageSenderInfo#690a7e4b as nil")
	}
	b.ObjStart()
	b.PutID("chatStatisticsMessageSenderInfo")
	b.FieldStart("user_id")
	b.PutInt53(c.UserID)
	b.FieldStart("sent_message_count")
	b.PutInt32(c.SentMessageCount)
	b.FieldStart("average_character_count")
	b.PutInt32(c.AverageCharacterCount)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatStatisticsMessageSenderInfo) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatStatisticsMessageSenderInfo#690a7e4b to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatStatisticsMessageSenderInfo"); err != nil {
				return fmt.Errorf("unable to decode chatStatisticsMessageSenderInfo#690a7e4b: %w", err)
			}
		case "user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode chatStatisticsMessageSenderInfo#690a7e4b: field user_id: %w", err)
			}
			c.UserID = value
		case "sent_message_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatStatisticsMessageSenderInfo#690a7e4b: field sent_message_count: %w", err)
			}
			c.SentMessageCount = value
		case "average_character_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatStatisticsMessageSenderInfo#690a7e4b: field average_character_count: %w", err)
			}
			c.AverageCharacterCount = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUserID returns value of UserID field.
func (c *ChatStatisticsMessageSenderInfo) GetUserID() (value int64) {
	return c.UserID
}

// GetSentMessageCount returns value of SentMessageCount field.
func (c *ChatStatisticsMessageSenderInfo) GetSentMessageCount() (value int32) {
	return c.SentMessageCount
}

// GetAverageCharacterCount returns value of AverageCharacterCount field.
func (c *ChatStatisticsMessageSenderInfo) GetAverageCharacterCount() (value int32) {
	return c.AverageCharacterCount
}
