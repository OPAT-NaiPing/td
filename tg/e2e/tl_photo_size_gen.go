// Code generated by gotdgen, DO NOT EDIT.

package e2e

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

// PhotoSizeEmpty represents TL type `photoSizeEmpty#e17e23c`.
// Empty constructor. Image with this thumbnail is unavailable.
//
// See https://core.telegram.org/constructor/photoSizeEmpty for reference.
type PhotoSizeEmpty struct {
	// Thumbnail type »¹
	//
	// Links:
	//  1) https://core.telegram.org/api/files#image-thumbnail-types
	Type string
}

// PhotoSizeEmptyTypeID is TL type id of PhotoSizeEmpty.
const PhotoSizeEmptyTypeID = 0xe17e23c

// construct implements constructor of PhotoSizeClass.
func (p PhotoSizeEmpty) construct() PhotoSizeClass { return &p }

// Ensuring interfaces in compile-time for PhotoSizeEmpty.
var (
	_ bin.Encoder     = &PhotoSizeEmpty{}
	_ bin.Decoder     = &PhotoSizeEmpty{}
	_ bin.BareEncoder = &PhotoSizeEmpty{}
	_ bin.BareDecoder = &PhotoSizeEmpty{}

	_ PhotoSizeClass = &PhotoSizeEmpty{}
)

func (p *PhotoSizeEmpty) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Type == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PhotoSizeEmpty) String() string {
	if p == nil {
		return "PhotoSizeEmpty(nil)"
	}
	type Alias PhotoSizeEmpty
	return fmt.Sprintf("PhotoSizeEmpty%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhotoSizeEmpty) TypeID() uint32 {
	return PhotoSizeEmptyTypeID
}

// TypeName returns name of type in TL schema.
func (*PhotoSizeEmpty) TypeName() string {
	return "photoSizeEmpty"
}

// TypeInfo returns info about TL type.
func (p *PhotoSizeEmpty) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "photoSizeEmpty",
		ID:   PhotoSizeEmptyTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Type",
			SchemaName: "type",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PhotoSizeEmpty) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode photoSizeEmpty#e17e23c as nil")
	}
	b.PutID(PhotoSizeEmptyTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PhotoSizeEmpty) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode photoSizeEmpty#e17e23c as nil")
	}
	b.PutString(p.Type)
	return nil
}

// Decode implements bin.Decoder.
func (p *PhotoSizeEmpty) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode photoSizeEmpty#e17e23c to nil")
	}
	if err := b.ConsumeID(PhotoSizeEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode photoSizeEmpty#e17e23c: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PhotoSizeEmpty) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode photoSizeEmpty#e17e23c to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode photoSizeEmpty#e17e23c: field type: %w", err)
		}
		p.Type = value
	}
	return nil
}

// GetType returns value of Type field.
func (p *PhotoSizeEmpty) GetType() (value string) {
	if p == nil {
		return
	}
	return p.Type
}

// PhotoSize represents TL type `photoSize#77bfb61b`.
// Image description.
//
// See https://core.telegram.org/constructor/photoSize for reference.
type PhotoSize struct {
	// Thumbnail type »¹
	//
	// Links:
	//  1) https://core.telegram.org/api/files#image-thumbnail-types
	Type string
	// Location field of PhotoSize.
	Location FileLocationClass
	// Image width
	W int
	// Image height
	H int
	// File size
	Size int
}

// PhotoSizeTypeID is TL type id of PhotoSize.
const PhotoSizeTypeID = 0x77bfb61b

// construct implements constructor of PhotoSizeClass.
func (p PhotoSize) construct() PhotoSizeClass { return &p }

// Ensuring interfaces in compile-time for PhotoSize.
var (
	_ bin.Encoder     = &PhotoSize{}
	_ bin.Decoder     = &PhotoSize{}
	_ bin.BareEncoder = &PhotoSize{}
	_ bin.BareDecoder = &PhotoSize{}

	_ PhotoSizeClass = &PhotoSize{}
)

func (p *PhotoSize) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Type == "") {
		return false
	}
	if !(p.Location == nil) {
		return false
	}
	if !(p.W == 0) {
		return false
	}
	if !(p.H == 0) {
		return false
	}
	if !(p.Size == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PhotoSize) String() string {
	if p == nil {
		return "PhotoSize(nil)"
	}
	type Alias PhotoSize
	return fmt.Sprintf("PhotoSize%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhotoSize) TypeID() uint32 {
	return PhotoSizeTypeID
}

// TypeName returns name of type in TL schema.
func (*PhotoSize) TypeName() string {
	return "photoSize"
}

// TypeInfo returns info about TL type.
func (p *PhotoSize) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "photoSize",
		ID:   PhotoSizeTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Type",
			SchemaName: "type",
		},
		{
			Name:       "Location",
			SchemaName: "location",
		},
		{
			Name:       "W",
			SchemaName: "w",
		},
		{
			Name:       "H",
			SchemaName: "h",
		},
		{
			Name:       "Size",
			SchemaName: "size",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PhotoSize) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode photoSize#77bfb61b as nil")
	}
	b.PutID(PhotoSizeTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PhotoSize) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode photoSize#77bfb61b as nil")
	}
	b.PutString(p.Type)
	if p.Location == nil {
		return fmt.Errorf("unable to encode photoSize#77bfb61b: field location is nil")
	}
	if err := p.Location.Encode(b); err != nil {
		return fmt.Errorf("unable to encode photoSize#77bfb61b: field location: %w", err)
	}
	b.PutInt(p.W)
	b.PutInt(p.H)
	b.PutInt(p.Size)
	return nil
}

// Decode implements bin.Decoder.
func (p *PhotoSize) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode photoSize#77bfb61b to nil")
	}
	if err := b.ConsumeID(PhotoSizeTypeID); err != nil {
		return fmt.Errorf("unable to decode photoSize#77bfb61b: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PhotoSize) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode photoSize#77bfb61b to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode photoSize#77bfb61b: field type: %w", err)
		}
		p.Type = value
	}
	{
		value, err := DecodeFileLocation(b)
		if err != nil {
			return fmt.Errorf("unable to decode photoSize#77bfb61b: field location: %w", err)
		}
		p.Location = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode photoSize#77bfb61b: field w: %w", err)
		}
		p.W = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode photoSize#77bfb61b: field h: %w", err)
		}
		p.H = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode photoSize#77bfb61b: field size: %w", err)
		}
		p.Size = value
	}
	return nil
}

// GetType returns value of Type field.
func (p *PhotoSize) GetType() (value string) {
	if p == nil {
		return
	}
	return p.Type
}

// GetLocation returns value of Location field.
func (p *PhotoSize) GetLocation() (value FileLocationClass) {
	if p == nil {
		return
	}
	return p.Location
}

// GetW returns value of W field.
func (p *PhotoSize) GetW() (value int) {
	if p == nil {
		return
	}
	return p.W
}

// GetH returns value of H field.
func (p *PhotoSize) GetH() (value int) {
	if p == nil {
		return
	}
	return p.H
}

// GetSize returns value of Size field.
func (p *PhotoSize) GetSize() (value int) {
	if p == nil {
		return
	}
	return p.Size
}

// PhotoCachedSize represents TL type `photoCachedSize#e9a734fa`.
// Description of an image and its content.
//
// See https://core.telegram.org/constructor/photoCachedSize for reference.
type PhotoCachedSize struct {
	// Thumbnail type
	Type string
	// Location field of PhotoCachedSize.
	Location FileLocationClass
	// Image width
	W int
	// Image height
	H int
	// Binary data, file content
	Bytes []byte
}

// PhotoCachedSizeTypeID is TL type id of PhotoCachedSize.
const PhotoCachedSizeTypeID = 0xe9a734fa

// construct implements constructor of PhotoSizeClass.
func (p PhotoCachedSize) construct() PhotoSizeClass { return &p }

// Ensuring interfaces in compile-time for PhotoCachedSize.
var (
	_ bin.Encoder     = &PhotoCachedSize{}
	_ bin.Decoder     = &PhotoCachedSize{}
	_ bin.BareEncoder = &PhotoCachedSize{}
	_ bin.BareDecoder = &PhotoCachedSize{}

	_ PhotoSizeClass = &PhotoCachedSize{}
)

func (p *PhotoCachedSize) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Type == "") {
		return false
	}
	if !(p.Location == nil) {
		return false
	}
	if !(p.W == 0) {
		return false
	}
	if !(p.H == 0) {
		return false
	}
	if !(p.Bytes == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PhotoCachedSize) String() string {
	if p == nil {
		return "PhotoCachedSize(nil)"
	}
	type Alias PhotoCachedSize
	return fmt.Sprintf("PhotoCachedSize%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhotoCachedSize) TypeID() uint32 {
	return PhotoCachedSizeTypeID
}

// TypeName returns name of type in TL schema.
func (*PhotoCachedSize) TypeName() string {
	return "photoCachedSize"
}

// TypeInfo returns info about TL type.
func (p *PhotoCachedSize) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "photoCachedSize",
		ID:   PhotoCachedSizeTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Type",
			SchemaName: "type",
		},
		{
			Name:       "Location",
			SchemaName: "location",
		},
		{
			Name:       "W",
			SchemaName: "w",
		},
		{
			Name:       "H",
			SchemaName: "h",
		},
		{
			Name:       "Bytes",
			SchemaName: "bytes",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PhotoCachedSize) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode photoCachedSize#e9a734fa as nil")
	}
	b.PutID(PhotoCachedSizeTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PhotoCachedSize) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode photoCachedSize#e9a734fa as nil")
	}
	b.PutString(p.Type)
	if p.Location == nil {
		return fmt.Errorf("unable to encode photoCachedSize#e9a734fa: field location is nil")
	}
	if err := p.Location.Encode(b); err != nil {
		return fmt.Errorf("unable to encode photoCachedSize#e9a734fa: field location: %w", err)
	}
	b.PutInt(p.W)
	b.PutInt(p.H)
	b.PutBytes(p.Bytes)
	return nil
}

// Decode implements bin.Decoder.
func (p *PhotoCachedSize) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode photoCachedSize#e9a734fa to nil")
	}
	if err := b.ConsumeID(PhotoCachedSizeTypeID); err != nil {
		return fmt.Errorf("unable to decode photoCachedSize#e9a734fa: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PhotoCachedSize) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode photoCachedSize#e9a734fa to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode photoCachedSize#e9a734fa: field type: %w", err)
		}
		p.Type = value
	}
	{
		value, err := DecodeFileLocation(b)
		if err != nil {
			return fmt.Errorf("unable to decode photoCachedSize#e9a734fa: field location: %w", err)
		}
		p.Location = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode photoCachedSize#e9a734fa: field w: %w", err)
		}
		p.W = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode photoCachedSize#e9a734fa: field h: %w", err)
		}
		p.H = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode photoCachedSize#e9a734fa: field bytes: %w", err)
		}
		p.Bytes = value
	}
	return nil
}

// GetType returns value of Type field.
func (p *PhotoCachedSize) GetType() (value string) {
	if p == nil {
		return
	}
	return p.Type
}

// GetLocation returns value of Location field.
func (p *PhotoCachedSize) GetLocation() (value FileLocationClass) {
	if p == nil {
		return
	}
	return p.Location
}

// GetW returns value of W field.
func (p *PhotoCachedSize) GetW() (value int) {
	if p == nil {
		return
	}
	return p.W
}

// GetH returns value of H field.
func (p *PhotoCachedSize) GetH() (value int) {
	if p == nil {
		return
	}
	return p.H
}

// GetBytes returns value of Bytes field.
func (p *PhotoCachedSize) GetBytes() (value []byte) {
	if p == nil {
		return
	}
	return p.Bytes
}

// PhotoSizeClassName is schema name of PhotoSizeClass.
const PhotoSizeClassName = "PhotoSize"

// PhotoSizeClass represents PhotoSize generic type.
//
// See https://core.telegram.org/type/PhotoSize for reference.
//
// Example:
//
//	g, err := e2e.DecodePhotoSize(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *e2e.PhotoSizeEmpty: // photoSizeEmpty#e17e23c
//	case *e2e.PhotoSize: // photoSize#77bfb61b
//	case *e2e.PhotoCachedSize: // photoCachedSize#e9a734fa
//	default: panic(v)
//	}
type PhotoSizeClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() PhotoSizeClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	// Thumbnail type »¹
	//
	// Links:
	//  1) https://core.telegram.org/api/files#image-thumbnail-types
	GetType() (value string)
}

// DecodePhotoSize implements binary de-serialization for PhotoSizeClass.
func DecodePhotoSize(buf *bin.Buffer) (PhotoSizeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case PhotoSizeEmptyTypeID:
		// Decoding photoSizeEmpty#e17e23c.
		v := PhotoSizeEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PhotoSizeClass: %w", err)
		}
		return &v, nil
	case PhotoSizeTypeID:
		// Decoding photoSize#77bfb61b.
		v := PhotoSize{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PhotoSizeClass: %w", err)
		}
		return &v, nil
	case PhotoCachedSizeTypeID:
		// Decoding photoCachedSize#e9a734fa.
		v := PhotoCachedSize{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PhotoSizeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode PhotoSizeClass: %w", bin.NewUnexpectedID(id))
	}
}

// PhotoSize boxes the PhotoSizeClass providing a helper.
type PhotoSizeBox struct {
	PhotoSize PhotoSizeClass
}

// Decode implements bin.Decoder for PhotoSizeBox.
func (b *PhotoSizeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode PhotoSizeBox to nil")
	}
	v, err := DecodePhotoSize(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.PhotoSize = v
	return nil
}

// Encode implements bin.Encode for PhotoSizeBox.
func (b *PhotoSizeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.PhotoSize == nil {
		return fmt.Errorf("unable to encode PhotoSizeClass as nil")
	}
	return b.PhotoSize.Encode(buf)
}
