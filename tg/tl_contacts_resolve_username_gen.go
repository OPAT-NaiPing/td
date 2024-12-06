// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// ContactsResolveUsernameRequest represents TL type `contacts.resolveUsername#725afbbc`.
// Resolve a @username to get peer info
//
// See https://core.telegram.org/method/contacts.resolveUsername for reference.
type ContactsResolveUsernameRequest struct {
	// Flags field of ContactsResolveUsernameRequest.
	Flags bin.Fields
	// @username to resolve
	Username string
	// Referer field of ContactsResolveUsernameRequest.
	//
	// Use SetReferer and GetReferer helpers.
	Referer string
}

// ContactsResolveUsernameRequestTypeID is TL type id of ContactsResolveUsernameRequest.
const ContactsResolveUsernameRequestTypeID = 0x725afbbc

// Ensuring interfaces in compile-time for ContactsResolveUsernameRequest.
var (
	_ bin.Encoder     = &ContactsResolveUsernameRequest{}
	_ bin.Decoder     = &ContactsResolveUsernameRequest{}
	_ bin.BareEncoder = &ContactsResolveUsernameRequest{}
	_ bin.BareDecoder = &ContactsResolveUsernameRequest{}
)

func (r *ContactsResolveUsernameRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Flags.Zero()) {
		return false
	}
	if !(r.Username == "") {
		return false
	}
	if !(r.Referer == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ContactsResolveUsernameRequest) String() string {
	if r == nil {
		return "ContactsResolveUsernameRequest(nil)"
	}
	type Alias ContactsResolveUsernameRequest
	return fmt.Sprintf("ContactsResolveUsernameRequest%+v", Alias(*r))
}

// FillFrom fills ContactsResolveUsernameRequest from given interface.
func (r *ContactsResolveUsernameRequest) FillFrom(from interface {
	GetUsername() (value string)
	GetReferer() (value string, ok bool)
}) {
	r.Username = from.GetUsername()
	if val, ok := from.GetReferer(); ok {
		r.Referer = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ContactsResolveUsernameRequest) TypeID() uint32 {
	return ContactsResolveUsernameRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ContactsResolveUsernameRequest) TypeName() string {
	return "contacts.resolveUsername"
}

// TypeInfo returns info about TL type.
func (r *ContactsResolveUsernameRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "contacts.resolveUsername",
		ID:   ContactsResolveUsernameRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Username",
			SchemaName: "username",
		},
		{
			Name:       "Referer",
			SchemaName: "referer",
			Null:       !r.Flags.Has(0),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (r *ContactsResolveUsernameRequest) SetFlags() {
	if !(r.Referer == "") {
		r.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (r *ContactsResolveUsernameRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode contacts.resolveUsername#725afbbc as nil")
	}
	b.PutID(ContactsResolveUsernameRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ContactsResolveUsernameRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode contacts.resolveUsername#725afbbc as nil")
	}
	r.SetFlags()
	if err := r.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode contacts.resolveUsername#725afbbc: field flags: %w", err)
	}
	b.PutString(r.Username)
	if r.Flags.Has(0) {
		b.PutString(r.Referer)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *ContactsResolveUsernameRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode contacts.resolveUsername#725afbbc to nil")
	}
	if err := b.ConsumeID(ContactsResolveUsernameRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.resolveUsername#725afbbc: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ContactsResolveUsernameRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode contacts.resolveUsername#725afbbc to nil")
	}
	{
		if err := r.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode contacts.resolveUsername#725afbbc: field flags: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.resolveUsername#725afbbc: field username: %w", err)
		}
		r.Username = value
	}
	if r.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.resolveUsername#725afbbc: field referer: %w", err)
		}
		r.Referer = value
	}
	return nil
}

// GetUsername returns value of Username field.
func (r *ContactsResolveUsernameRequest) GetUsername() (value string) {
	if r == nil {
		return
	}
	return r.Username
}

// SetReferer sets value of Referer conditional field.
func (r *ContactsResolveUsernameRequest) SetReferer(value string) {
	r.Flags.Set(0)
	r.Referer = value
}

// GetReferer returns value of Referer conditional field and
// boolean which is true if field was set.
func (r *ContactsResolveUsernameRequest) GetReferer() (value string, ok bool) {
	if r == nil {
		return
	}
	if !r.Flags.Has(0) {
		return value, false
	}
	return r.Referer, true
}

// ContactsResolveUsername invokes method contacts.resolveUsername#725afbbc returning error if any.
// Resolve a @username to get peer info
//
// Possible errors:
//
//	400 CONNECTION_LAYER_INVALID: Layer invalid.
//	400 USERNAME_INVALID: The provided username is not valid.
//	400 USERNAME_NOT_OCCUPIED: The provided username is not occupied.
//
// See https://core.telegram.org/method/contacts.resolveUsername for reference.
// Can be used by bots.
func (c *Client) ContactsResolveUsername(ctx context.Context, request *ContactsResolveUsernameRequest) (*ContactsResolvedPeer, error) {
	var result ContactsResolvedPeer

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
