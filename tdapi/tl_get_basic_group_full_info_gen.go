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

// GetBasicGroupFullInfoRequest represents TL type `getBasicGroupFullInfo#9365e32b`.
type GetBasicGroupFullInfoRequest struct {
	// Basic group identifier
	BasicGroupID int64
}

// GetBasicGroupFullInfoRequestTypeID is TL type id of GetBasicGroupFullInfoRequest.
const GetBasicGroupFullInfoRequestTypeID = 0x9365e32b

// Ensuring interfaces in compile-time for GetBasicGroupFullInfoRequest.
var (
	_ bin.Encoder     = &GetBasicGroupFullInfoRequest{}
	_ bin.Decoder     = &GetBasicGroupFullInfoRequest{}
	_ bin.BareEncoder = &GetBasicGroupFullInfoRequest{}
	_ bin.BareDecoder = &GetBasicGroupFullInfoRequest{}
)

func (g *GetBasicGroupFullInfoRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.BasicGroupID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetBasicGroupFullInfoRequest) String() string {
	if g == nil {
		return "GetBasicGroupFullInfoRequest(nil)"
	}
	type Alias GetBasicGroupFullInfoRequest
	return fmt.Sprintf("GetBasicGroupFullInfoRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetBasicGroupFullInfoRequest) TypeID() uint32 {
	return GetBasicGroupFullInfoRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetBasicGroupFullInfoRequest) TypeName() string {
	return "getBasicGroupFullInfo"
}

// TypeInfo returns info about TL type.
func (g *GetBasicGroupFullInfoRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getBasicGroupFullInfo",
		ID:   GetBasicGroupFullInfoRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "BasicGroupID",
			SchemaName: "basic_group_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetBasicGroupFullInfoRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getBasicGroupFullInfo#9365e32b as nil")
	}
	b.PutID(GetBasicGroupFullInfoRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetBasicGroupFullInfoRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getBasicGroupFullInfo#9365e32b as nil")
	}
	b.PutInt53(g.BasicGroupID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetBasicGroupFullInfoRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getBasicGroupFullInfo#9365e32b to nil")
	}
	if err := b.ConsumeID(GetBasicGroupFullInfoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getBasicGroupFullInfo#9365e32b: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetBasicGroupFullInfoRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getBasicGroupFullInfo#9365e32b to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getBasicGroupFullInfo#9365e32b: field basic_group_id: %w", err)
		}
		g.BasicGroupID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetBasicGroupFullInfoRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getBasicGroupFullInfo#9365e32b as nil")
	}
	b.ObjStart()
	b.PutID("getBasicGroupFullInfo")
	b.FieldStart("basic_group_id")
	b.PutInt53(g.BasicGroupID)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetBasicGroupFullInfoRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getBasicGroupFullInfo#9365e32b to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getBasicGroupFullInfo"); err != nil {
				return fmt.Errorf("unable to decode getBasicGroupFullInfo#9365e32b: %w", err)
			}
		case "basic_group_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getBasicGroupFullInfo#9365e32b: field basic_group_id: %w", err)
			}
			g.BasicGroupID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetBasicGroupID returns value of BasicGroupID field.
func (g *GetBasicGroupFullInfoRequest) GetBasicGroupID() (value int64) {
	return g.BasicGroupID
}

// GetBasicGroupFullInfo invokes method getBasicGroupFullInfo#9365e32b returning error if any.
func (c *Client) GetBasicGroupFullInfo(ctx context.Context, basicgroupid int64) (*BasicGroupFullInfo, error) {
	var result BasicGroupFullInfo

	request := &GetBasicGroupFullInfoRequest{
		BasicGroupID: basicgroupid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
