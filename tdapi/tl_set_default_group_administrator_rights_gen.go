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

// SetDefaultGroupAdministratorRightsRequest represents TL type `setDefaultGroupAdministratorRights#65577768`.
type SetDefaultGroupAdministratorRightsRequest struct {
	// Default administrator rights for adding the bot to basic group and supergroup chats;
	// may be null
	DefaultGroupAdministratorRights ChatAdministratorRights
}

// SetDefaultGroupAdministratorRightsRequestTypeID is TL type id of SetDefaultGroupAdministratorRightsRequest.
const SetDefaultGroupAdministratorRightsRequestTypeID = 0x65577768

// Ensuring interfaces in compile-time for SetDefaultGroupAdministratorRightsRequest.
var (
	_ bin.Encoder     = &SetDefaultGroupAdministratorRightsRequest{}
	_ bin.Decoder     = &SetDefaultGroupAdministratorRightsRequest{}
	_ bin.BareEncoder = &SetDefaultGroupAdministratorRightsRequest{}
	_ bin.BareDecoder = &SetDefaultGroupAdministratorRightsRequest{}
)

func (s *SetDefaultGroupAdministratorRightsRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.DefaultGroupAdministratorRights.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetDefaultGroupAdministratorRightsRequest) String() string {
	if s == nil {
		return "SetDefaultGroupAdministratorRightsRequest(nil)"
	}
	type Alias SetDefaultGroupAdministratorRightsRequest
	return fmt.Sprintf("SetDefaultGroupAdministratorRightsRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetDefaultGroupAdministratorRightsRequest) TypeID() uint32 {
	return SetDefaultGroupAdministratorRightsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetDefaultGroupAdministratorRightsRequest) TypeName() string {
	return "setDefaultGroupAdministratorRights"
}

// TypeInfo returns info about TL type.
func (s *SetDefaultGroupAdministratorRightsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setDefaultGroupAdministratorRights",
		ID:   SetDefaultGroupAdministratorRightsRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "DefaultGroupAdministratorRights",
			SchemaName: "default_group_administrator_rights",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetDefaultGroupAdministratorRightsRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setDefaultGroupAdministratorRights#65577768 as nil")
	}
	b.PutID(SetDefaultGroupAdministratorRightsRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetDefaultGroupAdministratorRightsRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setDefaultGroupAdministratorRights#65577768 as nil")
	}
	if err := s.DefaultGroupAdministratorRights.Encode(b); err != nil {
		return fmt.Errorf("unable to encode setDefaultGroupAdministratorRights#65577768: field default_group_administrator_rights: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SetDefaultGroupAdministratorRightsRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setDefaultGroupAdministratorRights#65577768 to nil")
	}
	if err := b.ConsumeID(SetDefaultGroupAdministratorRightsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setDefaultGroupAdministratorRights#65577768: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetDefaultGroupAdministratorRightsRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setDefaultGroupAdministratorRights#65577768 to nil")
	}
	{
		if err := s.DefaultGroupAdministratorRights.Decode(b); err != nil {
			return fmt.Errorf("unable to decode setDefaultGroupAdministratorRights#65577768: field default_group_administrator_rights: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetDefaultGroupAdministratorRightsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setDefaultGroupAdministratorRights#65577768 as nil")
	}
	b.ObjStart()
	b.PutID("setDefaultGroupAdministratorRights")
	b.Comma()
	b.FieldStart("default_group_administrator_rights")
	if err := s.DefaultGroupAdministratorRights.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode setDefaultGroupAdministratorRights#65577768: field default_group_administrator_rights: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetDefaultGroupAdministratorRightsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setDefaultGroupAdministratorRights#65577768 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setDefaultGroupAdministratorRights"); err != nil {
				return fmt.Errorf("unable to decode setDefaultGroupAdministratorRights#65577768: %w", err)
			}
		case "default_group_administrator_rights":
			if err := s.DefaultGroupAdministratorRights.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode setDefaultGroupAdministratorRights#65577768: field default_group_administrator_rights: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetDefaultGroupAdministratorRights returns value of DefaultGroupAdministratorRights field.
func (s *SetDefaultGroupAdministratorRightsRequest) GetDefaultGroupAdministratorRights() (value ChatAdministratorRights) {
	if s == nil {
		return
	}
	return s.DefaultGroupAdministratorRights
}

// SetDefaultGroupAdministratorRights invokes method setDefaultGroupAdministratorRights#65577768 returning error if any.
func (c *Client) SetDefaultGroupAdministratorRights(ctx context.Context, defaultgroupadministratorrights ChatAdministratorRights) error {
	var ok Ok

	request := &SetDefaultGroupAdministratorRightsRequest{
		DefaultGroupAdministratorRights: defaultgroupadministratorrights,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}