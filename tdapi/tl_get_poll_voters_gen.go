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

// GetPollVotersRequest represents TL type `getPollVoters#7bb2649e`.
type GetPollVotersRequest struct {
	// Identifier of the chat to which the poll belongs
	ChatID int64
	// Identifier of the message containing the poll
	MessageID int64
	// 0-based identifier of the answer option
	OptionID int32
	// Number of users to skip in the result; must be non-negative
	Offset int32
	// The maximum number of users to be returned; must be positive and can't be greater than
	// 50. For optimal performance, the number of returned users is chosen by TDLib and can
	// be smaller than the specified limit, even if the end of the voter list has not been
	// reached
	Limit int32
}

// GetPollVotersRequestTypeID is TL type id of GetPollVotersRequest.
const GetPollVotersRequestTypeID = 0x7bb2649e

// Ensuring interfaces in compile-time for GetPollVotersRequest.
var (
	_ bin.Encoder     = &GetPollVotersRequest{}
	_ bin.Decoder     = &GetPollVotersRequest{}
	_ bin.BareEncoder = &GetPollVotersRequest{}
	_ bin.BareDecoder = &GetPollVotersRequest{}
)

func (g *GetPollVotersRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}
	if !(g.MessageID == 0) {
		return false
	}
	if !(g.OptionID == 0) {
		return false
	}
	if !(g.Offset == 0) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetPollVotersRequest) String() string {
	if g == nil {
		return "GetPollVotersRequest(nil)"
	}
	type Alias GetPollVotersRequest
	return fmt.Sprintf("GetPollVotersRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetPollVotersRequest) TypeID() uint32 {
	return GetPollVotersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetPollVotersRequest) TypeName() string {
	return "getPollVoters"
}

// TypeInfo returns info about TL type.
func (g *GetPollVotersRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getPollVoters",
		ID:   GetPollVotersRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageID",
			SchemaName: "message_id",
		},
		{
			Name:       "OptionID",
			SchemaName: "option_id",
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
func (g *GetPollVotersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getPollVoters#7bb2649e as nil")
	}
	b.PutID(GetPollVotersRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetPollVotersRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getPollVoters#7bb2649e as nil")
	}
	b.PutInt53(g.ChatID)
	b.PutInt53(g.MessageID)
	b.PutInt32(g.OptionID)
	b.PutInt32(g.Offset)
	b.PutInt32(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetPollVotersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getPollVoters#7bb2649e to nil")
	}
	if err := b.ConsumeID(GetPollVotersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getPollVoters#7bb2649e: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetPollVotersRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getPollVoters#7bb2649e to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getPollVoters#7bb2649e: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getPollVoters#7bb2649e: field message_id: %w", err)
		}
		g.MessageID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getPollVoters#7bb2649e: field option_id: %w", err)
		}
		g.OptionID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getPollVoters#7bb2649e: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getPollVoters#7bb2649e: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetPollVotersRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getPollVoters#7bb2649e as nil")
	}
	b.ObjStart()
	b.PutID("getPollVoters")
	b.FieldStart("chat_id")
	b.PutInt53(g.ChatID)
	b.FieldStart("message_id")
	b.PutInt53(g.MessageID)
	b.FieldStart("option_id")
	b.PutInt32(g.OptionID)
	b.FieldStart("offset")
	b.PutInt32(g.Offset)
	b.FieldStart("limit")
	b.PutInt32(g.Limit)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetPollVotersRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getPollVoters#7bb2649e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getPollVoters"); err != nil {
				return fmt.Errorf("unable to decode getPollVoters#7bb2649e: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getPollVoters#7bb2649e: field chat_id: %w", err)
			}
			g.ChatID = value
		case "message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getPollVoters#7bb2649e: field message_id: %w", err)
			}
			g.MessageID = value
		case "option_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getPollVoters#7bb2649e: field option_id: %w", err)
			}
			g.OptionID = value
		case "offset":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getPollVoters#7bb2649e: field offset: %w", err)
			}
			g.Offset = value
		case "limit":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getPollVoters#7bb2649e: field limit: %w", err)
			}
			g.Limit = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetPollVotersRequest) GetChatID() (value int64) {
	return g.ChatID
}

// GetMessageID returns value of MessageID field.
func (g *GetPollVotersRequest) GetMessageID() (value int64) {
	return g.MessageID
}

// GetOptionID returns value of OptionID field.
func (g *GetPollVotersRequest) GetOptionID() (value int32) {
	return g.OptionID
}

// GetOffset returns value of Offset field.
func (g *GetPollVotersRequest) GetOffset() (value int32) {
	return g.Offset
}

// GetLimit returns value of Limit field.
func (g *GetPollVotersRequest) GetLimit() (value int32) {
	return g.Limit
}

// GetPollVoters invokes method getPollVoters#7bb2649e returning error if any.
func (c *Client) GetPollVoters(ctx context.Context, request *GetPollVotersRequest) (*Users, error) {
	var result Users

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
