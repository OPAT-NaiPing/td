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

// SetPinnedChatsRequest represents TL type `setPinnedChats#c6c6edf1`.
type SetPinnedChatsRequest struct {
	// Chat list in which to change the order of pinned chats
	ChatList ChatListClass
	// The new list of pinned chats
	ChatIDs []int64
}

// SetPinnedChatsRequestTypeID is TL type id of SetPinnedChatsRequest.
const SetPinnedChatsRequestTypeID = 0xc6c6edf1

// Ensuring interfaces in compile-time for SetPinnedChatsRequest.
var (
	_ bin.Encoder     = &SetPinnedChatsRequest{}
	_ bin.Decoder     = &SetPinnedChatsRequest{}
	_ bin.BareEncoder = &SetPinnedChatsRequest{}
	_ bin.BareDecoder = &SetPinnedChatsRequest{}
)

func (s *SetPinnedChatsRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ChatList == nil) {
		return false
	}
	if !(s.ChatIDs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetPinnedChatsRequest) String() string {
	if s == nil {
		return "SetPinnedChatsRequest(nil)"
	}
	type Alias SetPinnedChatsRequest
	return fmt.Sprintf("SetPinnedChatsRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetPinnedChatsRequest) TypeID() uint32 {
	return SetPinnedChatsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetPinnedChatsRequest) TypeName() string {
	return "setPinnedChats"
}

// TypeInfo returns info about TL type.
func (s *SetPinnedChatsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setPinnedChats",
		ID:   SetPinnedChatsRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatList",
			SchemaName: "chat_list",
		},
		{
			Name:       "ChatIDs",
			SchemaName: "chat_ids",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetPinnedChatsRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setPinnedChats#c6c6edf1 as nil")
	}
	b.PutID(SetPinnedChatsRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetPinnedChatsRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setPinnedChats#c6c6edf1 as nil")
	}
	if s.ChatList == nil {
		return fmt.Errorf("unable to encode setPinnedChats#c6c6edf1: field chat_list is nil")
	}
	if err := s.ChatList.Encode(b); err != nil {
		return fmt.Errorf("unable to encode setPinnedChats#c6c6edf1: field chat_list: %w", err)
	}
	b.PutInt(len(s.ChatIDs))
	for _, v := range s.ChatIDs {
		b.PutInt53(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SetPinnedChatsRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setPinnedChats#c6c6edf1 to nil")
	}
	if err := b.ConsumeID(SetPinnedChatsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setPinnedChats#c6c6edf1: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetPinnedChatsRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setPinnedChats#c6c6edf1 to nil")
	}
	{
		value, err := DecodeChatList(b)
		if err != nil {
			return fmt.Errorf("unable to decode setPinnedChats#c6c6edf1: field chat_list: %w", err)
		}
		s.ChatList = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode setPinnedChats#c6c6edf1: field chat_ids: %w", err)
		}

		if headerLen > 0 {
			s.ChatIDs = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode setPinnedChats#c6c6edf1: field chat_ids: %w", err)
			}
			s.ChatIDs = append(s.ChatIDs, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetPinnedChatsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setPinnedChats#c6c6edf1 as nil")
	}
	b.ObjStart()
	b.PutID("setPinnedChats")
	b.FieldStart("chat_list")
	if s.ChatList == nil {
		return fmt.Errorf("unable to encode setPinnedChats#c6c6edf1: field chat_list is nil")
	}
	if err := s.ChatList.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode setPinnedChats#c6c6edf1: field chat_list: %w", err)
	}
	b.FieldStart("chat_ids")
	b.ArrStart()
	for _, v := range s.ChatIDs {
		b.PutInt53(v)
	}
	b.ArrEnd()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetPinnedChatsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setPinnedChats#c6c6edf1 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setPinnedChats"); err != nil {
				return fmt.Errorf("unable to decode setPinnedChats#c6c6edf1: %w", err)
			}
		case "chat_list":
			value, err := DecodeTDLibJSONChatList(b)
			if err != nil {
				return fmt.Errorf("unable to decode setPinnedChats#c6c6edf1: field chat_list: %w", err)
			}
			s.ChatList = value
		case "chat_ids":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.Int53()
				if err != nil {
					return fmt.Errorf("unable to decode setPinnedChats#c6c6edf1: field chat_ids: %w", err)
				}
				s.ChatIDs = append(s.ChatIDs, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode setPinnedChats#c6c6edf1: field chat_ids: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatList returns value of ChatList field.
func (s *SetPinnedChatsRequest) GetChatList() (value ChatListClass) {
	return s.ChatList
}

// GetChatIDs returns value of ChatIDs field.
func (s *SetPinnedChatsRequest) GetChatIDs() (value []int64) {
	return s.ChatIDs
}

// SetPinnedChats invokes method setPinnedChats#c6c6edf1 returning error if any.
func (c *Client) SetPinnedChats(ctx context.Context, request *SetPinnedChatsRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
