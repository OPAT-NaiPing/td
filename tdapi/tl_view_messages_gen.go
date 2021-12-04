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

// ViewMessagesRequest represents TL type `viewMessages#d6e1005d`.
type ViewMessagesRequest struct {
	// Chat identifier
	ChatID int64
	// If not 0, a message thread identifier in which the messages are being viewed
	MessageThreadID int64
	// The identifiers of the messages being viewed
	MessageIDs []int64
	// True, if messages in closed chats must be marked as read by the request
	ForceRead bool
}

// ViewMessagesRequestTypeID is TL type id of ViewMessagesRequest.
const ViewMessagesRequestTypeID = 0xd6e1005d

// Ensuring interfaces in compile-time for ViewMessagesRequest.
var (
	_ bin.Encoder     = &ViewMessagesRequest{}
	_ bin.Decoder     = &ViewMessagesRequest{}
	_ bin.BareEncoder = &ViewMessagesRequest{}
	_ bin.BareDecoder = &ViewMessagesRequest{}
)

func (v *ViewMessagesRequest) Zero() bool {
	if v == nil {
		return true
	}
	if !(v.ChatID == 0) {
		return false
	}
	if !(v.MessageThreadID == 0) {
		return false
	}
	if !(v.MessageIDs == nil) {
		return false
	}
	if !(v.ForceRead == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (v *ViewMessagesRequest) String() string {
	if v == nil {
		return "ViewMessagesRequest(nil)"
	}
	type Alias ViewMessagesRequest
	return fmt.Sprintf("ViewMessagesRequest%+v", Alias(*v))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ViewMessagesRequest) TypeID() uint32 {
	return ViewMessagesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ViewMessagesRequest) TypeName() string {
	return "viewMessages"
}

// TypeInfo returns info about TL type.
func (v *ViewMessagesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "viewMessages",
		ID:   ViewMessagesRequestTypeID,
	}
	if v == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageThreadID",
			SchemaName: "message_thread_id",
		},
		{
			Name:       "MessageIDs",
			SchemaName: "message_ids",
		},
		{
			Name:       "ForceRead",
			SchemaName: "force_read",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (v *ViewMessagesRequest) Encode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode viewMessages#d6e1005d as nil")
	}
	b.PutID(ViewMessagesRequestTypeID)
	return v.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (v *ViewMessagesRequest) EncodeBare(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode viewMessages#d6e1005d as nil")
	}
	b.PutInt53(v.ChatID)
	b.PutInt53(v.MessageThreadID)
	b.PutInt(len(v.MessageIDs))
	for _, v := range v.MessageIDs {
		b.PutInt53(v)
	}
	b.PutBool(v.ForceRead)
	return nil
}

// Decode implements bin.Decoder.
func (v *ViewMessagesRequest) Decode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode viewMessages#d6e1005d to nil")
	}
	if err := b.ConsumeID(ViewMessagesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode viewMessages#d6e1005d: %w", err)
	}
	return v.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (v *ViewMessagesRequest) DecodeBare(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode viewMessages#d6e1005d to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode viewMessages#d6e1005d: field chat_id: %w", err)
		}
		v.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode viewMessages#d6e1005d: field message_thread_id: %w", err)
		}
		v.MessageThreadID = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode viewMessages#d6e1005d: field message_ids: %w", err)
		}

		if headerLen > 0 {
			v.MessageIDs = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode viewMessages#d6e1005d: field message_ids: %w", err)
			}
			v.MessageIDs = append(v.MessageIDs, value)
		}
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode viewMessages#d6e1005d: field force_read: %w", err)
		}
		v.ForceRead = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (v *ViewMessagesRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if v == nil {
		return fmt.Errorf("can't encode viewMessages#d6e1005d as nil")
	}
	b.ObjStart()
	b.PutID("viewMessages")
	b.FieldStart("chat_id")
	b.PutInt53(v.ChatID)
	b.FieldStart("message_thread_id")
	b.PutInt53(v.MessageThreadID)
	b.FieldStart("message_ids")
	b.ArrStart()
	for _, v := range v.MessageIDs {
		b.PutInt53(v)
	}
	b.ArrEnd()
	b.FieldStart("force_read")
	b.PutBool(v.ForceRead)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (v *ViewMessagesRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if v == nil {
		return fmt.Errorf("can't decode viewMessages#d6e1005d to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("viewMessages"); err != nil {
				return fmt.Errorf("unable to decode viewMessages#d6e1005d: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode viewMessages#d6e1005d: field chat_id: %w", err)
			}
			v.ChatID = value
		case "message_thread_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode viewMessages#d6e1005d: field message_thread_id: %w", err)
			}
			v.MessageThreadID = value
		case "message_ids":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.Int53()
				if err != nil {
					return fmt.Errorf("unable to decode viewMessages#d6e1005d: field message_ids: %w", err)
				}
				v.MessageIDs = append(v.MessageIDs, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode viewMessages#d6e1005d: field message_ids: %w", err)
			}
		case "force_read":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode viewMessages#d6e1005d: field force_read: %w", err)
			}
			v.ForceRead = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (v *ViewMessagesRequest) GetChatID() (value int64) {
	return v.ChatID
}

// GetMessageThreadID returns value of MessageThreadID field.
func (v *ViewMessagesRequest) GetMessageThreadID() (value int64) {
	return v.MessageThreadID
}

// GetMessageIDs returns value of MessageIDs field.
func (v *ViewMessagesRequest) GetMessageIDs() (value []int64) {
	return v.MessageIDs
}

// GetForceRead returns value of ForceRead field.
func (v *ViewMessagesRequest) GetForceRead() (value bool) {
	return v.ForceRead
}

// ViewMessages invokes method viewMessages#d6e1005d returning error if any.
func (c *Client) ViewMessages(ctx context.Context, request *ViewMessagesRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
