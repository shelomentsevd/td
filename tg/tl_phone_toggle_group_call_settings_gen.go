// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}

// PhoneToggleGroupCallSettingsRequest represents TL type `phone.toggleGroupCallSettings#74bbb43d`.
//
// See https://core.telegram.org/method/phone.toggleGroupCallSettings for reference.
type PhoneToggleGroupCallSettingsRequest struct {
	// Flags field of PhoneToggleGroupCallSettingsRequest.
	Flags bin.Fields
	// Call field of PhoneToggleGroupCallSettingsRequest.
	Call InputGroupCall
	// JoinMuted field of PhoneToggleGroupCallSettingsRequest.
	JoinMuted bool
}

// PhoneToggleGroupCallSettingsRequestTypeID is TL type id of PhoneToggleGroupCallSettingsRequest.
const PhoneToggleGroupCallSettingsRequestTypeID = 0x74bbb43d

// String implements fmt.Stringer.
func (t *PhoneToggleGroupCallSettingsRequest) String() string {
	if t == nil {
		return "PhoneToggleGroupCallSettingsRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("PhoneToggleGroupCallSettingsRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(t.Flags.String())
	sb.WriteString(",\n")
	sb.WriteString("\tCall: ")
	sb.WriteString(t.Call.String())
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (t *PhoneToggleGroupCallSettingsRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode phone.toggleGroupCallSettings#74bbb43d as nil")
	}
	b.PutID(PhoneToggleGroupCallSettingsRequestTypeID)
	if err := t.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.toggleGroupCallSettings#74bbb43d: field flags: %w", err)
	}
	if err := t.Call.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.toggleGroupCallSettings#74bbb43d: field call: %w", err)
	}
	return nil
}

// SetJoinMuted sets value of JoinMuted conditional field.
func (t *PhoneToggleGroupCallSettingsRequest) SetJoinMuted(value bool) {
	if value {
		t.Flags.Set(0)
		t.JoinMuted = true
	} else {
		t.Flags.Unset(0)
		t.JoinMuted = false
	}
}

// Decode implements bin.Decoder.
func (t *PhoneToggleGroupCallSettingsRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode phone.toggleGroupCallSettings#74bbb43d to nil")
	}
	if err := b.ConsumeID(PhoneToggleGroupCallSettingsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.toggleGroupCallSettings#74bbb43d: %w", err)
	}
	{
		if err := t.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.toggleGroupCallSettings#74bbb43d: field flags: %w", err)
		}
	}
	{
		if err := t.Call.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.toggleGroupCallSettings#74bbb43d: field call: %w", err)
		}
	}
	t.JoinMuted = t.Flags.Has(0)
	return nil
}

// Ensuring interfaces in compile-time for PhoneToggleGroupCallSettingsRequest.
var (
	_ bin.Encoder = &PhoneToggleGroupCallSettingsRequest{}
	_ bin.Decoder = &PhoneToggleGroupCallSettingsRequest{}
)

// PhoneToggleGroupCallSettings invokes method phone.toggleGroupCallSettings#74bbb43d returning error if any.
//
// See https://core.telegram.org/method/phone.toggleGroupCallSettings for reference.
func (c *Client) PhoneToggleGroupCallSettings(ctx context.Context, request *PhoneToggleGroupCallSettingsRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
