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

// StopBusinessPollRequest represents TL type `stopBusinessPoll#bbeb2160`.
type StopBusinessPollRequest struct {
	// Unique identifier of business connection on behalf of which the message with the poll
	// was sent
	BusinessConnectionID string
	// The chat the message belongs to
	ChatID int64
	// Identifier of the message containing the poll
	MessageID int64
	// The new message reply markup; pass null if none
	ReplyMarkup ReplyMarkupClass
}

// StopBusinessPollRequestTypeID is TL type id of StopBusinessPollRequest.
const StopBusinessPollRequestTypeID = 0xbbeb2160

// Ensuring interfaces in compile-time for StopBusinessPollRequest.
var (
	_ bin.Encoder     = &StopBusinessPollRequest{}
	_ bin.Decoder     = &StopBusinessPollRequest{}
	_ bin.BareEncoder = &StopBusinessPollRequest{}
	_ bin.BareDecoder = &StopBusinessPollRequest{}
)

func (s *StopBusinessPollRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.BusinessConnectionID == "") {
		return false
	}
	if !(s.ChatID == 0) {
		return false
	}
	if !(s.MessageID == 0) {
		return false
	}
	if !(s.ReplyMarkup == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StopBusinessPollRequest) String() string {
	if s == nil {
		return "StopBusinessPollRequest(nil)"
	}
	type Alias StopBusinessPollRequest
	return fmt.Sprintf("StopBusinessPollRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StopBusinessPollRequest) TypeID() uint32 {
	return StopBusinessPollRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*StopBusinessPollRequest) TypeName() string {
	return "stopBusinessPoll"
}

// TypeInfo returns info about TL type.
func (s *StopBusinessPollRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stopBusinessPoll",
		ID:   StopBusinessPollRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "BusinessConnectionID",
			SchemaName: "business_connection_id",
		},
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageID",
			SchemaName: "message_id",
		},
		{
			Name:       "ReplyMarkup",
			SchemaName: "reply_markup",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StopBusinessPollRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode stopBusinessPoll#bbeb2160 as nil")
	}
	b.PutID(StopBusinessPollRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StopBusinessPollRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode stopBusinessPoll#bbeb2160 as nil")
	}
	b.PutString(s.BusinessConnectionID)
	b.PutInt53(s.ChatID)
	b.PutInt53(s.MessageID)
	if s.ReplyMarkup == nil {
		return fmt.Errorf("unable to encode stopBusinessPoll#bbeb2160: field reply_markup is nil")
	}
	if err := s.ReplyMarkup.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stopBusinessPoll#bbeb2160: field reply_markup: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StopBusinessPollRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode stopBusinessPoll#bbeb2160 to nil")
	}
	if err := b.ConsumeID(StopBusinessPollRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stopBusinessPoll#bbeb2160: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StopBusinessPollRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode stopBusinessPoll#bbeb2160 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode stopBusinessPoll#bbeb2160: field business_connection_id: %w", err)
		}
		s.BusinessConnectionID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode stopBusinessPoll#bbeb2160: field chat_id: %w", err)
		}
		s.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode stopBusinessPoll#bbeb2160: field message_id: %w", err)
		}
		s.MessageID = value
	}
	{
		value, err := DecodeReplyMarkup(b)
		if err != nil {
			return fmt.Errorf("unable to decode stopBusinessPoll#bbeb2160: field reply_markup: %w", err)
		}
		s.ReplyMarkup = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *StopBusinessPollRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode stopBusinessPoll#bbeb2160 as nil")
	}
	b.ObjStart()
	b.PutID("stopBusinessPoll")
	b.Comma()
	b.FieldStart("business_connection_id")
	b.PutString(s.BusinessConnectionID)
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(s.ChatID)
	b.Comma()
	b.FieldStart("message_id")
	b.PutInt53(s.MessageID)
	b.Comma()
	b.FieldStart("reply_markup")
	if s.ReplyMarkup == nil {
		return fmt.Errorf("unable to encode stopBusinessPoll#bbeb2160: field reply_markup is nil")
	}
	if err := s.ReplyMarkup.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode stopBusinessPoll#bbeb2160: field reply_markup: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *StopBusinessPollRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode stopBusinessPoll#bbeb2160 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("stopBusinessPoll"); err != nil {
				return fmt.Errorf("unable to decode stopBusinessPoll#bbeb2160: %w", err)
			}
		case "business_connection_id":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode stopBusinessPoll#bbeb2160: field business_connection_id: %w", err)
			}
			s.BusinessConnectionID = value
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode stopBusinessPoll#bbeb2160: field chat_id: %w", err)
			}
			s.ChatID = value
		case "message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode stopBusinessPoll#bbeb2160: field message_id: %w", err)
			}
			s.MessageID = value
		case "reply_markup":
			value, err := DecodeTDLibJSONReplyMarkup(b)
			if err != nil {
				return fmt.Errorf("unable to decode stopBusinessPoll#bbeb2160: field reply_markup: %w", err)
			}
			s.ReplyMarkup = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetBusinessConnectionID returns value of BusinessConnectionID field.
func (s *StopBusinessPollRequest) GetBusinessConnectionID() (value string) {
	if s == nil {
		return
	}
	return s.BusinessConnectionID
}

// GetChatID returns value of ChatID field.
func (s *StopBusinessPollRequest) GetChatID() (value int64) {
	if s == nil {
		return
	}
	return s.ChatID
}

// GetMessageID returns value of MessageID field.
func (s *StopBusinessPollRequest) GetMessageID() (value int64) {
	if s == nil {
		return
	}
	return s.MessageID
}

// GetReplyMarkup returns value of ReplyMarkup field.
func (s *StopBusinessPollRequest) GetReplyMarkup() (value ReplyMarkupClass) {
	if s == nil {
		return
	}
	return s.ReplyMarkup
}

// StopBusinessPoll invokes method stopBusinessPoll#bbeb2160 returning error if any.
func (c *Client) StopBusinessPoll(ctx context.Context, request *StopBusinessPollRequest) (*BusinessMessage, error) {
	var result BusinessMessage

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
