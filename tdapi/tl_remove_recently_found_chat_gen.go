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

// RemoveRecentlyFoundChatRequest represents TL type `removeRecentlyFoundChat#2ac1bf1c`.
type RemoveRecentlyFoundChatRequest struct {
	// Identifier of the chat to be removed
	ChatID int64
}

// RemoveRecentlyFoundChatRequestTypeID is TL type id of RemoveRecentlyFoundChatRequest.
const RemoveRecentlyFoundChatRequestTypeID = 0x2ac1bf1c

// Ensuring interfaces in compile-time for RemoveRecentlyFoundChatRequest.
var (
	_ bin.Encoder     = &RemoveRecentlyFoundChatRequest{}
	_ bin.Decoder     = &RemoveRecentlyFoundChatRequest{}
	_ bin.BareEncoder = &RemoveRecentlyFoundChatRequest{}
	_ bin.BareDecoder = &RemoveRecentlyFoundChatRequest{}
)

func (r *RemoveRecentlyFoundChatRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.ChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RemoveRecentlyFoundChatRequest) String() string {
	if r == nil {
		return "RemoveRecentlyFoundChatRequest(nil)"
	}
	type Alias RemoveRecentlyFoundChatRequest
	return fmt.Sprintf("RemoveRecentlyFoundChatRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RemoveRecentlyFoundChatRequest) TypeID() uint32 {
	return RemoveRecentlyFoundChatRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*RemoveRecentlyFoundChatRequest) TypeName() string {
	return "removeRecentlyFoundChat"
}

// TypeInfo returns info about TL type.
func (r *RemoveRecentlyFoundChatRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "removeRecentlyFoundChat",
		ID:   RemoveRecentlyFoundChatRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *RemoveRecentlyFoundChatRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode removeRecentlyFoundChat#2ac1bf1c as nil")
	}
	b.PutID(RemoveRecentlyFoundChatRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RemoveRecentlyFoundChatRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode removeRecentlyFoundChat#2ac1bf1c as nil")
	}
	b.PutInt53(r.ChatID)
	return nil
}

// Decode implements bin.Decoder.
func (r *RemoveRecentlyFoundChatRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode removeRecentlyFoundChat#2ac1bf1c to nil")
	}
	if err := b.ConsumeID(RemoveRecentlyFoundChatRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode removeRecentlyFoundChat#2ac1bf1c: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RemoveRecentlyFoundChatRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode removeRecentlyFoundChat#2ac1bf1c to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode removeRecentlyFoundChat#2ac1bf1c: field chat_id: %w", err)
		}
		r.ChatID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *RemoveRecentlyFoundChatRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode removeRecentlyFoundChat#2ac1bf1c as nil")
	}
	b.ObjStart()
	b.PutID("removeRecentlyFoundChat")
	b.FieldStart("chat_id")
	b.PutInt53(r.ChatID)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *RemoveRecentlyFoundChatRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode removeRecentlyFoundChat#2ac1bf1c to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("removeRecentlyFoundChat"); err != nil {
				return fmt.Errorf("unable to decode removeRecentlyFoundChat#2ac1bf1c: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode removeRecentlyFoundChat#2ac1bf1c: field chat_id: %w", err)
			}
			r.ChatID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (r *RemoveRecentlyFoundChatRequest) GetChatID() (value int64) {
	return r.ChatID
}

// RemoveRecentlyFoundChat invokes method removeRecentlyFoundChat#2ac1bf1c returning error if any.
func (c *Client) RemoveRecentlyFoundChat(ctx context.Context, chatid int64) error {
	var ok Ok

	request := &RemoveRecentlyFoundChatRequest{
		ChatID: chatid,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
