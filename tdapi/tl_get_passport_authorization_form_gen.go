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

// GetPassportAuthorizationFormRequest represents TL type `getPassportAuthorizationForm#61850486`.
type GetPassportAuthorizationFormRequest struct {
	// User identifier of the service's bot
	BotUserID int64
	// Telegram Passport element types requested by the service
	Scope string
	// Service's public key
	PublicKey string
	// Unique request identifier provided by the service
	Nonce string
}

// GetPassportAuthorizationFormRequestTypeID is TL type id of GetPassportAuthorizationFormRequest.
const GetPassportAuthorizationFormRequestTypeID = 0x61850486

// Ensuring interfaces in compile-time for GetPassportAuthorizationFormRequest.
var (
	_ bin.Encoder     = &GetPassportAuthorizationFormRequest{}
	_ bin.Decoder     = &GetPassportAuthorizationFormRequest{}
	_ bin.BareEncoder = &GetPassportAuthorizationFormRequest{}
	_ bin.BareDecoder = &GetPassportAuthorizationFormRequest{}
)

func (g *GetPassportAuthorizationFormRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.BotUserID == 0) {
		return false
	}
	if !(g.Scope == "") {
		return false
	}
	if !(g.PublicKey == "") {
		return false
	}
	if !(g.Nonce == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetPassportAuthorizationFormRequest) String() string {
	if g == nil {
		return "GetPassportAuthorizationFormRequest(nil)"
	}
	type Alias GetPassportAuthorizationFormRequest
	return fmt.Sprintf("GetPassportAuthorizationFormRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetPassportAuthorizationFormRequest) TypeID() uint32 {
	return GetPassportAuthorizationFormRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetPassportAuthorizationFormRequest) TypeName() string {
	return "getPassportAuthorizationForm"
}

// TypeInfo returns info about TL type.
func (g *GetPassportAuthorizationFormRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getPassportAuthorizationForm",
		ID:   GetPassportAuthorizationFormRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "BotUserID",
			SchemaName: "bot_user_id",
		},
		{
			Name:       "Scope",
			SchemaName: "scope",
		},
		{
			Name:       "PublicKey",
			SchemaName: "public_key",
		},
		{
			Name:       "Nonce",
			SchemaName: "nonce",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetPassportAuthorizationFormRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getPassportAuthorizationForm#61850486 as nil")
	}
	b.PutID(GetPassportAuthorizationFormRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetPassportAuthorizationFormRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getPassportAuthorizationForm#61850486 as nil")
	}
	b.PutInt53(g.BotUserID)
	b.PutString(g.Scope)
	b.PutString(g.PublicKey)
	b.PutString(g.Nonce)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetPassportAuthorizationFormRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getPassportAuthorizationForm#61850486 to nil")
	}
	if err := b.ConsumeID(GetPassportAuthorizationFormRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getPassportAuthorizationForm#61850486: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetPassportAuthorizationFormRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getPassportAuthorizationForm#61850486 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getPassportAuthorizationForm#61850486: field bot_user_id: %w", err)
		}
		g.BotUserID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getPassportAuthorizationForm#61850486: field scope: %w", err)
		}
		g.Scope = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getPassportAuthorizationForm#61850486: field public_key: %w", err)
		}
		g.PublicKey = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getPassportAuthorizationForm#61850486: field nonce: %w", err)
		}
		g.Nonce = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetPassportAuthorizationFormRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getPassportAuthorizationForm#61850486 as nil")
	}
	b.ObjStart()
	b.PutID("getPassportAuthorizationForm")
	b.FieldStart("bot_user_id")
	b.PutInt53(g.BotUserID)
	b.FieldStart("scope")
	b.PutString(g.Scope)
	b.FieldStart("public_key")
	b.PutString(g.PublicKey)
	b.FieldStart("nonce")
	b.PutString(g.Nonce)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetPassportAuthorizationFormRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getPassportAuthorizationForm#61850486 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getPassportAuthorizationForm"); err != nil {
				return fmt.Errorf("unable to decode getPassportAuthorizationForm#61850486: %w", err)
			}
		case "bot_user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getPassportAuthorizationForm#61850486: field bot_user_id: %w", err)
			}
			g.BotUserID = value
		case "scope":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getPassportAuthorizationForm#61850486: field scope: %w", err)
			}
			g.Scope = value
		case "public_key":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getPassportAuthorizationForm#61850486: field public_key: %w", err)
			}
			g.PublicKey = value
		case "nonce":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getPassportAuthorizationForm#61850486: field nonce: %w", err)
			}
			g.Nonce = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetBotUserID returns value of BotUserID field.
func (g *GetPassportAuthorizationFormRequest) GetBotUserID() (value int64) {
	return g.BotUserID
}

// GetScope returns value of Scope field.
func (g *GetPassportAuthorizationFormRequest) GetScope() (value string) {
	return g.Scope
}

// GetPublicKey returns value of PublicKey field.
func (g *GetPassportAuthorizationFormRequest) GetPublicKey() (value string) {
	return g.PublicKey
}

// GetNonce returns value of Nonce field.
func (g *GetPassportAuthorizationFormRequest) GetNonce() (value string) {
	return g.Nonce
}

// GetPassportAuthorizationForm invokes method getPassportAuthorizationForm#61850486 returning error if any.
func (c *Client) GetPassportAuthorizationForm(ctx context.Context, request *GetPassportAuthorizationFormRequest) (*PassportAuthorizationForm, error) {
	var result PassportAuthorizationForm

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
