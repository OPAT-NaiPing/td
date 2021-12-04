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

// RegisterDeviceRequest represents TL type `registerDevice#ff140196`.
type RegisterDeviceRequest struct {
	// Device token
	DeviceToken DeviceTokenClass
	// List of user identifiers of other users currently using the application
	OtherUserIDs []int64
}

// RegisterDeviceRequestTypeID is TL type id of RegisterDeviceRequest.
const RegisterDeviceRequestTypeID = 0xff140196

// Ensuring interfaces in compile-time for RegisterDeviceRequest.
var (
	_ bin.Encoder     = &RegisterDeviceRequest{}
	_ bin.Decoder     = &RegisterDeviceRequest{}
	_ bin.BareEncoder = &RegisterDeviceRequest{}
	_ bin.BareDecoder = &RegisterDeviceRequest{}
)

func (r *RegisterDeviceRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.DeviceToken == nil) {
		return false
	}
	if !(r.OtherUserIDs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RegisterDeviceRequest) String() string {
	if r == nil {
		return "RegisterDeviceRequest(nil)"
	}
	type Alias RegisterDeviceRequest
	return fmt.Sprintf("RegisterDeviceRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RegisterDeviceRequest) TypeID() uint32 {
	return RegisterDeviceRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*RegisterDeviceRequest) TypeName() string {
	return "registerDevice"
}

// TypeInfo returns info about TL type.
func (r *RegisterDeviceRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "registerDevice",
		ID:   RegisterDeviceRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "DeviceToken",
			SchemaName: "device_token",
		},
		{
			Name:       "OtherUserIDs",
			SchemaName: "other_user_ids",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *RegisterDeviceRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode registerDevice#ff140196 as nil")
	}
	b.PutID(RegisterDeviceRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RegisterDeviceRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode registerDevice#ff140196 as nil")
	}
	if r.DeviceToken == nil {
		return fmt.Errorf("unable to encode registerDevice#ff140196: field device_token is nil")
	}
	if err := r.DeviceToken.Encode(b); err != nil {
		return fmt.Errorf("unable to encode registerDevice#ff140196: field device_token: %w", err)
	}
	b.PutInt(len(r.OtherUserIDs))
	for _, v := range r.OtherUserIDs {
		b.PutInt53(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *RegisterDeviceRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode registerDevice#ff140196 to nil")
	}
	if err := b.ConsumeID(RegisterDeviceRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode registerDevice#ff140196: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RegisterDeviceRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode registerDevice#ff140196 to nil")
	}
	{
		value, err := DecodeDeviceToken(b)
		if err != nil {
			return fmt.Errorf("unable to decode registerDevice#ff140196: field device_token: %w", err)
		}
		r.DeviceToken = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode registerDevice#ff140196: field other_user_ids: %w", err)
		}

		if headerLen > 0 {
			r.OtherUserIDs = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode registerDevice#ff140196: field other_user_ids: %w", err)
			}
			r.OtherUserIDs = append(r.OtherUserIDs, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *RegisterDeviceRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode registerDevice#ff140196 as nil")
	}
	b.ObjStart()
	b.PutID("registerDevice")
	b.FieldStart("device_token")
	if r.DeviceToken == nil {
		return fmt.Errorf("unable to encode registerDevice#ff140196: field device_token is nil")
	}
	if err := r.DeviceToken.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode registerDevice#ff140196: field device_token: %w", err)
	}
	b.FieldStart("other_user_ids")
	b.ArrStart()
	for _, v := range r.OtherUserIDs {
		b.PutInt53(v)
	}
	b.ArrEnd()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *RegisterDeviceRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode registerDevice#ff140196 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("registerDevice"); err != nil {
				return fmt.Errorf("unable to decode registerDevice#ff140196: %w", err)
			}
		case "device_token":
			value, err := DecodeTDLibJSONDeviceToken(b)
			if err != nil {
				return fmt.Errorf("unable to decode registerDevice#ff140196: field device_token: %w", err)
			}
			r.DeviceToken = value
		case "other_user_ids":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.Int53()
				if err != nil {
					return fmt.Errorf("unable to decode registerDevice#ff140196: field other_user_ids: %w", err)
				}
				r.OtherUserIDs = append(r.OtherUserIDs, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode registerDevice#ff140196: field other_user_ids: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetDeviceToken returns value of DeviceToken field.
func (r *RegisterDeviceRequest) GetDeviceToken() (value DeviceTokenClass) {
	return r.DeviceToken
}

// GetOtherUserIDs returns value of OtherUserIDs field.
func (r *RegisterDeviceRequest) GetOtherUserIDs() (value []int64) {
	return r.OtherUserIDs
}

// RegisterDevice invokes method registerDevice#ff140196 returning error if any.
func (c *Client) RegisterDevice(ctx context.Context, request *RegisterDeviceRequest) (*PushReceiverID, error) {
	var result PushReceiverID

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
