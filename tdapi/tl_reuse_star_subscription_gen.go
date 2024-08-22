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

// ReuseStarSubscriptionRequest represents TL type `reuseStarSubscription#2e677441`.
type ReuseStarSubscriptionRequest struct {
	// Identifier of the subscription
	SubscriptionID string
}

// ReuseStarSubscriptionRequestTypeID is TL type id of ReuseStarSubscriptionRequest.
const ReuseStarSubscriptionRequestTypeID = 0x2e677441

// Ensuring interfaces in compile-time for ReuseStarSubscriptionRequest.
var (
	_ bin.Encoder     = &ReuseStarSubscriptionRequest{}
	_ bin.Decoder     = &ReuseStarSubscriptionRequest{}
	_ bin.BareEncoder = &ReuseStarSubscriptionRequest{}
	_ bin.BareDecoder = &ReuseStarSubscriptionRequest{}
)

func (r *ReuseStarSubscriptionRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.SubscriptionID == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReuseStarSubscriptionRequest) String() string {
	if r == nil {
		return "ReuseStarSubscriptionRequest(nil)"
	}
	type Alias ReuseStarSubscriptionRequest
	return fmt.Sprintf("ReuseStarSubscriptionRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReuseStarSubscriptionRequest) TypeID() uint32 {
	return ReuseStarSubscriptionRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ReuseStarSubscriptionRequest) TypeName() string {
	return "reuseStarSubscription"
}

// TypeInfo returns info about TL type.
func (r *ReuseStarSubscriptionRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "reuseStarSubscription",
		ID:   ReuseStarSubscriptionRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SubscriptionID",
			SchemaName: "subscription_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReuseStarSubscriptionRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reuseStarSubscription#2e677441 as nil")
	}
	b.PutID(ReuseStarSubscriptionRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReuseStarSubscriptionRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reuseStarSubscription#2e677441 as nil")
	}
	b.PutString(r.SubscriptionID)
	return nil
}

// Decode implements bin.Decoder.
func (r *ReuseStarSubscriptionRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reuseStarSubscription#2e677441 to nil")
	}
	if err := b.ConsumeID(ReuseStarSubscriptionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode reuseStarSubscription#2e677441: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReuseStarSubscriptionRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reuseStarSubscription#2e677441 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode reuseStarSubscription#2e677441: field subscription_id: %w", err)
		}
		r.SubscriptionID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *ReuseStarSubscriptionRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode reuseStarSubscription#2e677441 as nil")
	}
	b.ObjStart()
	b.PutID("reuseStarSubscription")
	b.Comma()
	b.FieldStart("subscription_id")
	b.PutString(r.SubscriptionID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *ReuseStarSubscriptionRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode reuseStarSubscription#2e677441 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("reuseStarSubscription"); err != nil {
				return fmt.Errorf("unable to decode reuseStarSubscription#2e677441: %w", err)
			}
		case "subscription_id":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode reuseStarSubscription#2e677441: field subscription_id: %w", err)
			}
			r.SubscriptionID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSubscriptionID returns value of SubscriptionID field.
func (r *ReuseStarSubscriptionRequest) GetSubscriptionID() (value string) {
	if r == nil {
		return
	}
	return r.SubscriptionID
}

// ReuseStarSubscription invokes method reuseStarSubscription#2e677441 returning error if any.
func (c *Client) ReuseStarSubscription(ctx context.Context, subscriptionid string) error {
	var ok Ok

	request := &ReuseStarSubscriptionRequest{
		SubscriptionID: subscriptionid,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}