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

// BotsUpdateUserEmojiStatusRequest represents TL type `bots.updateUserEmojiStatus#ed9f30c5`.
//
// See https://core.telegram.org/method/bots.updateUserEmojiStatus for reference.
type BotsUpdateUserEmojiStatusRequest struct {
	// UserID field of BotsUpdateUserEmojiStatusRequest.
	UserID InputUserClass
	// EmojiStatus field of BotsUpdateUserEmojiStatusRequest.
	EmojiStatus EmojiStatusClass
}

// BotsUpdateUserEmojiStatusRequestTypeID is TL type id of BotsUpdateUserEmojiStatusRequest.
const BotsUpdateUserEmojiStatusRequestTypeID = 0xed9f30c5

// Ensuring interfaces in compile-time for BotsUpdateUserEmojiStatusRequest.
var (
	_ bin.Encoder     = &BotsUpdateUserEmojiStatusRequest{}
	_ bin.Decoder     = &BotsUpdateUserEmojiStatusRequest{}
	_ bin.BareEncoder = &BotsUpdateUserEmojiStatusRequest{}
	_ bin.BareDecoder = &BotsUpdateUserEmojiStatusRequest{}
)

func (u *BotsUpdateUserEmojiStatusRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.UserID == nil) {
		return false
	}
	if !(u.EmojiStatus == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *BotsUpdateUserEmojiStatusRequest) String() string {
	if u == nil {
		return "BotsUpdateUserEmojiStatusRequest(nil)"
	}
	type Alias BotsUpdateUserEmojiStatusRequest
	return fmt.Sprintf("BotsUpdateUserEmojiStatusRequest%+v", Alias(*u))
}

// FillFrom fills BotsUpdateUserEmojiStatusRequest from given interface.
func (u *BotsUpdateUserEmojiStatusRequest) FillFrom(from interface {
	GetUserID() (value InputUserClass)
	GetEmojiStatus() (value EmojiStatusClass)
}) {
	u.UserID = from.GetUserID()
	u.EmojiStatus = from.GetEmojiStatus()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotsUpdateUserEmojiStatusRequest) TypeID() uint32 {
	return BotsUpdateUserEmojiStatusRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*BotsUpdateUserEmojiStatusRequest) TypeName() string {
	return "bots.updateUserEmojiStatus"
}

// TypeInfo returns info about TL type.
func (u *BotsUpdateUserEmojiStatusRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "bots.updateUserEmojiStatus",
		ID:   BotsUpdateUserEmojiStatusRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "EmojiStatus",
			SchemaName: "emoji_status",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *BotsUpdateUserEmojiStatusRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode bots.updateUserEmojiStatus#ed9f30c5 as nil")
	}
	b.PutID(BotsUpdateUserEmojiStatusRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *BotsUpdateUserEmojiStatusRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode bots.updateUserEmojiStatus#ed9f30c5 as nil")
	}
	if u.UserID == nil {
		return fmt.Errorf("unable to encode bots.updateUserEmojiStatus#ed9f30c5: field user_id is nil")
	}
	if err := u.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode bots.updateUserEmojiStatus#ed9f30c5: field user_id: %w", err)
	}
	if u.EmojiStatus == nil {
		return fmt.Errorf("unable to encode bots.updateUserEmojiStatus#ed9f30c5: field emoji_status is nil")
	}
	if err := u.EmojiStatus.Encode(b); err != nil {
		return fmt.Errorf("unable to encode bots.updateUserEmojiStatus#ed9f30c5: field emoji_status: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *BotsUpdateUserEmojiStatusRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode bots.updateUserEmojiStatus#ed9f30c5 to nil")
	}
	if err := b.ConsumeID(BotsUpdateUserEmojiStatusRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode bots.updateUserEmojiStatus#ed9f30c5: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *BotsUpdateUserEmojiStatusRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode bots.updateUserEmojiStatus#ed9f30c5 to nil")
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode bots.updateUserEmojiStatus#ed9f30c5: field user_id: %w", err)
		}
		u.UserID = value
	}
	{
		value, err := DecodeEmojiStatus(b)
		if err != nil {
			return fmt.Errorf("unable to decode bots.updateUserEmojiStatus#ed9f30c5: field emoji_status: %w", err)
		}
		u.EmojiStatus = value
	}
	return nil
}

// GetUserID returns value of UserID field.
func (u *BotsUpdateUserEmojiStatusRequest) GetUserID() (value InputUserClass) {
	if u == nil {
		return
	}
	return u.UserID
}

// GetEmojiStatus returns value of EmojiStatus field.
func (u *BotsUpdateUserEmojiStatusRequest) GetEmojiStatus() (value EmojiStatusClass) {
	if u == nil {
		return
	}
	return u.EmojiStatus
}

// GetEmojiStatusAsNotEmpty returns mapped value of EmojiStatus field.
func (u *BotsUpdateUserEmojiStatusRequest) GetEmojiStatusAsNotEmpty() (NotEmptyEmojiStatus, bool) {
	return u.EmojiStatus.AsNotEmpty()
}

// BotsUpdateUserEmojiStatus invokes method bots.updateUserEmojiStatus#ed9f30c5 returning error if any.
//
// See https://core.telegram.org/method/bots.updateUserEmojiStatus for reference.
func (c *Client) BotsUpdateUserEmojiStatus(ctx context.Context, request *BotsUpdateUserEmojiStatusRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}