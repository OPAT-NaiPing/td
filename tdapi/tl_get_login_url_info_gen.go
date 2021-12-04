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

// GetLoginURLInfoRequest represents TL type `getLoginUrlInfo#ccc99db3`.
type GetLoginURLInfoRequest struct {
	// Chat identifier of the message with the button
	ChatID int64
	// Message identifier of the message with the button
	MessageID int64
	// Button identifier
	ButtonID int64
}

// GetLoginURLInfoRequestTypeID is TL type id of GetLoginURLInfoRequest.
const GetLoginURLInfoRequestTypeID = 0xccc99db3

// Ensuring interfaces in compile-time for GetLoginURLInfoRequest.
var (
	_ bin.Encoder     = &GetLoginURLInfoRequest{}
	_ bin.Decoder     = &GetLoginURLInfoRequest{}
	_ bin.BareEncoder = &GetLoginURLInfoRequest{}
	_ bin.BareDecoder = &GetLoginURLInfoRequest{}
)

func (g *GetLoginURLInfoRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}
	if !(g.MessageID == 0) {
		return false
	}
	if !(g.ButtonID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetLoginURLInfoRequest) String() string {
	if g == nil {
		return "GetLoginURLInfoRequest(nil)"
	}
	type Alias GetLoginURLInfoRequest
	return fmt.Sprintf("GetLoginURLInfoRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetLoginURLInfoRequest) TypeID() uint32 {
	return GetLoginURLInfoRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetLoginURLInfoRequest) TypeName() string {
	return "getLoginUrlInfo"
}

// TypeInfo returns info about TL type.
func (g *GetLoginURLInfoRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getLoginUrlInfo",
		ID:   GetLoginURLInfoRequestTypeID,
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
			Name:       "ButtonID",
			SchemaName: "button_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetLoginURLInfoRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getLoginUrlInfo#ccc99db3 as nil")
	}
	b.PutID(GetLoginURLInfoRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetLoginURLInfoRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getLoginUrlInfo#ccc99db3 as nil")
	}
	b.PutInt53(g.ChatID)
	b.PutInt53(g.MessageID)
	b.PutInt53(g.ButtonID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetLoginURLInfoRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getLoginUrlInfo#ccc99db3 to nil")
	}
	if err := b.ConsumeID(GetLoginURLInfoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getLoginUrlInfo#ccc99db3: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetLoginURLInfoRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getLoginUrlInfo#ccc99db3 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getLoginUrlInfo#ccc99db3: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getLoginUrlInfo#ccc99db3: field message_id: %w", err)
		}
		g.MessageID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getLoginUrlInfo#ccc99db3: field button_id: %w", err)
		}
		g.ButtonID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetLoginURLInfoRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getLoginUrlInfo#ccc99db3 as nil")
	}
	b.ObjStart()
	b.PutID("getLoginUrlInfo")
	b.FieldStart("chat_id")
	b.PutInt53(g.ChatID)
	b.FieldStart("message_id")
	b.PutInt53(g.MessageID)
	b.FieldStart("button_id")
	b.PutInt53(g.ButtonID)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetLoginURLInfoRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getLoginUrlInfo#ccc99db3 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getLoginUrlInfo"); err != nil {
				return fmt.Errorf("unable to decode getLoginUrlInfo#ccc99db3: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getLoginUrlInfo#ccc99db3: field chat_id: %w", err)
			}
			g.ChatID = value
		case "message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getLoginUrlInfo#ccc99db3: field message_id: %w", err)
			}
			g.MessageID = value
		case "button_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getLoginUrlInfo#ccc99db3: field button_id: %w", err)
			}
			g.ButtonID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetLoginURLInfoRequest) GetChatID() (value int64) {
	return g.ChatID
}

// GetMessageID returns value of MessageID field.
func (g *GetLoginURLInfoRequest) GetMessageID() (value int64) {
	return g.MessageID
}

// GetButtonID returns value of ButtonID field.
func (g *GetLoginURLInfoRequest) GetButtonID() (value int64) {
	return g.ButtonID
}

// GetLoginURLInfo invokes method getLoginUrlInfo#ccc99db3 returning error if any.
func (c *Client) GetLoginURLInfo(ctx context.Context, request *GetLoginURLInfoRequest) (LoginURLInfoClass, error) {
	var result LoginURLInfoBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.LoginUrlInfo, nil
}
