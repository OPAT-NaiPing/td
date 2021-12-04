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

// BanChatMemberRequest represents TL type `banChatMember#cb107d7c`.
type BanChatMemberRequest struct {
	// Chat identifier
	ChatID int64
	// Member identifier
	MemberID MessageSenderClass
	// Point in time (Unix timestamp) when the user will be unbanned; 0 if never. If the user
	// is banned for more than 366 days or for less than 30 seconds from the current time,
	// the user is considered to be banned forever. Ignored in basic groups
	BannedUntilDate int32
	// Pass true to delete all messages in the chat for the user that is being removed.
	// Always true for supergroups and channels
	RevokeMessages bool
}

// BanChatMemberRequestTypeID is TL type id of BanChatMemberRequest.
const BanChatMemberRequestTypeID = 0xcb107d7c

// Ensuring interfaces in compile-time for BanChatMemberRequest.
var (
	_ bin.Encoder     = &BanChatMemberRequest{}
	_ bin.Decoder     = &BanChatMemberRequest{}
	_ bin.BareEncoder = &BanChatMemberRequest{}
	_ bin.BareDecoder = &BanChatMemberRequest{}
)

func (b *BanChatMemberRequest) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.ChatID == 0) {
		return false
	}
	if !(b.MemberID == nil) {
		return false
	}
	if !(b.BannedUntilDate == 0) {
		return false
	}
	if !(b.RevokeMessages == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BanChatMemberRequest) String() string {
	if b == nil {
		return "BanChatMemberRequest(nil)"
	}
	type Alias BanChatMemberRequest
	return fmt.Sprintf("BanChatMemberRequest%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BanChatMemberRequest) TypeID() uint32 {
	return BanChatMemberRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*BanChatMemberRequest) TypeName() string {
	return "banChatMember"
}

// TypeInfo returns info about TL type.
func (b *BanChatMemberRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "banChatMember",
		ID:   BanChatMemberRequestTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MemberID",
			SchemaName: "member_id",
		},
		{
			Name:       "BannedUntilDate",
			SchemaName: "banned_until_date",
		},
		{
			Name:       "RevokeMessages",
			SchemaName: "revoke_messages",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *BanChatMemberRequest) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode banChatMember#cb107d7c as nil")
	}
	buf.PutID(BanChatMemberRequestTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BanChatMemberRequest) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode banChatMember#cb107d7c as nil")
	}
	buf.PutInt53(b.ChatID)
	if b.MemberID == nil {
		return fmt.Errorf("unable to encode banChatMember#cb107d7c: field member_id is nil")
	}
	if err := b.MemberID.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode banChatMember#cb107d7c: field member_id: %w", err)
	}
	buf.PutInt32(b.BannedUntilDate)
	buf.PutBool(b.RevokeMessages)
	return nil
}

// Decode implements bin.Decoder.
func (b *BanChatMemberRequest) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode banChatMember#cb107d7c to nil")
	}
	if err := buf.ConsumeID(BanChatMemberRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode banChatMember#cb107d7c: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BanChatMemberRequest) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode banChatMember#cb107d7c to nil")
	}
	{
		value, err := buf.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode banChatMember#cb107d7c: field chat_id: %w", err)
		}
		b.ChatID = value
	}
	{
		value, err := DecodeMessageSender(buf)
		if err != nil {
			return fmt.Errorf("unable to decode banChatMember#cb107d7c: field member_id: %w", err)
		}
		b.MemberID = value
	}
	{
		value, err := buf.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode banChatMember#cb107d7c: field banned_until_date: %w", err)
		}
		b.BannedUntilDate = value
	}
	{
		value, err := buf.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode banChatMember#cb107d7c: field revoke_messages: %w", err)
		}
		b.RevokeMessages = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (b *BanChatMemberRequest) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil {
		return fmt.Errorf("can't encode banChatMember#cb107d7c as nil")
	}
	buf.ObjStart()
	buf.PutID("banChatMember")
	buf.FieldStart("chat_id")
	buf.PutInt53(b.ChatID)
	buf.FieldStart("member_id")
	if b.MemberID == nil {
		return fmt.Errorf("unable to encode banChatMember#cb107d7c: field member_id is nil")
	}
	if err := b.MemberID.EncodeTDLibJSON(buf); err != nil {
		return fmt.Errorf("unable to encode banChatMember#cb107d7c: field member_id: %w", err)
	}
	buf.FieldStart("banned_until_date")
	buf.PutInt32(b.BannedUntilDate)
	buf.FieldStart("revoke_messages")
	buf.PutBool(b.RevokeMessages)
	buf.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (b *BanChatMemberRequest) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("can't decode banChatMember#cb107d7c to nil")
	}

	return buf.Obj(func(buf tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := buf.ConsumeID("banChatMember"); err != nil {
				return fmt.Errorf("unable to decode banChatMember#cb107d7c: %w", err)
			}
		case "chat_id":
			value, err := buf.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode banChatMember#cb107d7c: field chat_id: %w", err)
			}
			b.ChatID = value
		case "member_id":
			value, err := DecodeTDLibJSONMessageSender(buf)
			if err != nil {
				return fmt.Errorf("unable to decode banChatMember#cb107d7c: field member_id: %w", err)
			}
			b.MemberID = value
		case "banned_until_date":
			value, err := buf.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode banChatMember#cb107d7c: field banned_until_date: %w", err)
			}
			b.BannedUntilDate = value
		case "revoke_messages":
			value, err := buf.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode banChatMember#cb107d7c: field revoke_messages: %w", err)
			}
			b.RevokeMessages = value
		default:
			return buf.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (b *BanChatMemberRequest) GetChatID() (value int64) {
	return b.ChatID
}

// GetMemberID returns value of MemberID field.
func (b *BanChatMemberRequest) GetMemberID() (value MessageSenderClass) {
	return b.MemberID
}

// GetBannedUntilDate returns value of BannedUntilDate field.
func (b *BanChatMemberRequest) GetBannedUntilDate() (value int32) {
	return b.BannedUntilDate
}

// GetRevokeMessages returns value of RevokeMessages field.
func (b *BanChatMemberRequest) GetRevokeMessages() (value bool) {
	return b.RevokeMessages
}

// BanChatMember invokes method banChatMember#cb107d7c returning error if any.
func (c *Client) BanChatMember(ctx context.Context, request *BanChatMemberRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
