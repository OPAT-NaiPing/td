// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/gotd/td/bin"
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
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
)

// ChatInviteExported represents TL type `chatInviteExported#6e24fc9d`.
// Exported chat invite
//
// See https://core.telegram.org/constructor/chatInviteExported for reference.
type ChatInviteExported struct {
	// Flags field of ChatInviteExported.
	Flags bin.Fields
	// Revoked field of ChatInviteExported.
	Revoked bool
	// Permanent field of ChatInviteExported.
	Permanent bool
	// Chat invitation link
	Link string
	// AdminID field of ChatInviteExported.
	AdminID int
	// Date field of ChatInviteExported.
	Date int
	// StartDate field of ChatInviteExported.
	//
	// Use SetStartDate and GetStartDate helpers.
	StartDate int
	// ExpireDate field of ChatInviteExported.
	//
	// Use SetExpireDate and GetExpireDate helpers.
	ExpireDate int
	// UsageLimit field of ChatInviteExported.
	//
	// Use SetUsageLimit and GetUsageLimit helpers.
	UsageLimit int
	// Usage field of ChatInviteExported.
	//
	// Use SetUsage and GetUsage helpers.
	Usage int
}

// ChatInviteExportedTypeID is TL type id of ChatInviteExported.
const ChatInviteExportedTypeID = 0x6e24fc9d

func (c *ChatInviteExported) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.Revoked == false) {
		return false
	}
	if !(c.Permanent == false) {
		return false
	}
	if !(c.Link == "") {
		return false
	}
	if !(c.AdminID == 0) {
		return false
	}
	if !(c.Date == 0) {
		return false
	}
	if !(c.StartDate == 0) {
		return false
	}
	if !(c.ExpireDate == 0) {
		return false
	}
	if !(c.UsageLimit == 0) {
		return false
	}
	if !(c.Usage == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatInviteExported) String() string {
	if c == nil {
		return "ChatInviteExported(nil)"
	}
	type Alias ChatInviteExported
	return fmt.Sprintf("ChatInviteExported%+v", Alias(*c))
}

// FillFrom fills ChatInviteExported from given interface.
func (c *ChatInviteExported) FillFrom(from interface {
	GetRevoked() (value bool)
	GetPermanent() (value bool)
	GetLink() (value string)
	GetAdminID() (value int)
	GetDate() (value int)
	GetStartDate() (value int, ok bool)
	GetExpireDate() (value int, ok bool)
	GetUsageLimit() (value int, ok bool)
	GetUsage() (value int, ok bool)
}) {
	c.Revoked = from.GetRevoked()
	c.Permanent = from.GetPermanent()
	c.Link = from.GetLink()
	c.AdminID = from.GetAdminID()
	c.Date = from.GetDate()
	if val, ok := from.GetStartDate(); ok {
		c.StartDate = val
	}

	if val, ok := from.GetExpireDate(); ok {
		c.ExpireDate = val
	}

	if val, ok := from.GetUsageLimit(); ok {
		c.UsageLimit = val
	}

	if val, ok := from.GetUsage(); ok {
		c.Usage = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatInviteExported) TypeID() uint32 {
	return ChatInviteExportedTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatInviteExported) TypeName() string {
	return "chatInviteExported"
}

// TypeInfo returns info about TL type.
func (c *ChatInviteExported) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatInviteExported",
		ID:   ChatInviteExportedTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Revoked",
			SchemaName: "revoked",
			Null:       !c.Flags.Has(0),
		},
		{
			Name:       "Permanent",
			SchemaName: "permanent",
			Null:       !c.Flags.Has(5),
		},
		{
			Name:       "Link",
			SchemaName: "link",
		},
		{
			Name:       "AdminID",
			SchemaName: "admin_id",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
		{
			Name:       "StartDate",
			SchemaName: "start_date",
			Null:       !c.Flags.Has(4),
		},
		{
			Name:       "ExpireDate",
			SchemaName: "expire_date",
			Null:       !c.Flags.Has(1),
		},
		{
			Name:       "UsageLimit",
			SchemaName: "usage_limit",
			Null:       !c.Flags.Has(2),
		},
		{
			Name:       "Usage",
			SchemaName: "usage",
			Null:       !c.Flags.Has(3),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatInviteExported) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatInviteExported#6e24fc9d as nil")
	}
	b.PutID(ChatInviteExportedTypeID)
	if !(c.Revoked == false) {
		c.Flags.Set(0)
	}
	if !(c.Permanent == false) {
		c.Flags.Set(5)
	}
	if !(c.StartDate == 0) {
		c.Flags.Set(4)
	}
	if !(c.ExpireDate == 0) {
		c.Flags.Set(1)
	}
	if !(c.UsageLimit == 0) {
		c.Flags.Set(2)
	}
	if !(c.Usage == 0) {
		c.Flags.Set(3)
	}
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatInviteExported#6e24fc9d: field flags: %w", err)
	}
	b.PutString(c.Link)
	b.PutInt(c.AdminID)
	b.PutInt(c.Date)
	if c.Flags.Has(4) {
		b.PutInt(c.StartDate)
	}
	if c.Flags.Has(1) {
		b.PutInt(c.ExpireDate)
	}
	if c.Flags.Has(2) {
		b.PutInt(c.UsageLimit)
	}
	if c.Flags.Has(3) {
		b.PutInt(c.Usage)
	}
	return nil
}

// SetRevoked sets value of Revoked conditional field.
func (c *ChatInviteExported) SetRevoked(value bool) {
	if value {
		c.Flags.Set(0)
		c.Revoked = true
	} else {
		c.Flags.Unset(0)
		c.Revoked = false
	}
}

// GetRevoked returns value of Revoked conditional field.
func (c *ChatInviteExported) GetRevoked() (value bool) {
	return c.Flags.Has(0)
}

// SetPermanent sets value of Permanent conditional field.
func (c *ChatInviteExported) SetPermanent(value bool) {
	if value {
		c.Flags.Set(5)
		c.Permanent = true
	} else {
		c.Flags.Unset(5)
		c.Permanent = false
	}
}

// GetPermanent returns value of Permanent conditional field.
func (c *ChatInviteExported) GetPermanent() (value bool) {
	return c.Flags.Has(5)
}

// GetLink returns value of Link field.
func (c *ChatInviteExported) GetLink() (value string) {
	return c.Link
}

// GetAdminID returns value of AdminID field.
func (c *ChatInviteExported) GetAdminID() (value int) {
	return c.AdminID
}

// GetDate returns value of Date field.
func (c *ChatInviteExported) GetDate() (value int) {
	return c.Date
}

// SetStartDate sets value of StartDate conditional field.
func (c *ChatInviteExported) SetStartDate(value int) {
	c.Flags.Set(4)
	c.StartDate = value
}

// GetStartDate returns value of StartDate conditional field and
// boolean which is true if field was set.
func (c *ChatInviteExported) GetStartDate() (value int, ok bool) {
	if !c.Flags.Has(4) {
		return value, false
	}
	return c.StartDate, true
}

// SetExpireDate sets value of ExpireDate conditional field.
func (c *ChatInviteExported) SetExpireDate(value int) {
	c.Flags.Set(1)
	c.ExpireDate = value
}

// GetExpireDate returns value of ExpireDate conditional field and
// boolean which is true if field was set.
func (c *ChatInviteExported) GetExpireDate() (value int, ok bool) {
	if !c.Flags.Has(1) {
		return value, false
	}
	return c.ExpireDate, true
}

// SetUsageLimit sets value of UsageLimit conditional field.
func (c *ChatInviteExported) SetUsageLimit(value int) {
	c.Flags.Set(2)
	c.UsageLimit = value
}

// GetUsageLimit returns value of UsageLimit conditional field and
// boolean which is true if field was set.
func (c *ChatInviteExported) GetUsageLimit() (value int, ok bool) {
	if !c.Flags.Has(2) {
		return value, false
	}
	return c.UsageLimit, true
}

// SetUsage sets value of Usage conditional field.
func (c *ChatInviteExported) SetUsage(value int) {
	c.Flags.Set(3)
	c.Usage = value
}

// GetUsage returns value of Usage conditional field and
// boolean which is true if field was set.
func (c *ChatInviteExported) GetUsage() (value int, ok bool) {
	if !c.Flags.Has(3) {
		return value, false
	}
	return c.Usage, true
}

// Decode implements bin.Decoder.
func (c *ChatInviteExported) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatInviteExported#6e24fc9d to nil")
	}
	if err := b.ConsumeID(ChatInviteExportedTypeID); err != nil {
		return fmt.Errorf("unable to decode chatInviteExported#6e24fc9d: %w", err)
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chatInviteExported#6e24fc9d: field flags: %w", err)
		}
	}
	c.Revoked = c.Flags.Has(0)
	c.Permanent = c.Flags.Has(5)
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteExported#6e24fc9d: field link: %w", err)
		}
		c.Link = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteExported#6e24fc9d: field admin_id: %w", err)
		}
		c.AdminID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteExported#6e24fc9d: field date: %w", err)
		}
		c.Date = value
	}
	if c.Flags.Has(4) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteExported#6e24fc9d: field start_date: %w", err)
		}
		c.StartDate = value
	}
	if c.Flags.Has(1) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteExported#6e24fc9d: field expire_date: %w", err)
		}
		c.ExpireDate = value
	}
	if c.Flags.Has(2) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteExported#6e24fc9d: field usage_limit: %w", err)
		}
		c.UsageLimit = value
	}
	if c.Flags.Has(3) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteExported#6e24fc9d: field usage: %w", err)
		}
		c.Usage = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChatInviteExported.
var (
	_ bin.Encoder = &ChatInviteExported{}
	_ bin.Decoder = &ChatInviteExported{}
)