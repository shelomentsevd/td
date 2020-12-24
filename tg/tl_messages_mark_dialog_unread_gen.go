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

// MessagesMarkDialogUnreadRequest represents TL type `messages.markDialogUnread#c286d98f`.
// Manually mark dialog as unread
//
// See https://core.telegram.org/method/messages.markDialogUnread for reference.
type MessagesMarkDialogUnreadRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Mark as unread/read
	Unread bool
	// Dialog
	Peer InputDialogPeerClass
}

// MessagesMarkDialogUnreadRequestTypeID is TL type id of MessagesMarkDialogUnreadRequest.
const MessagesMarkDialogUnreadRequestTypeID = 0xc286d98f

// String implements fmt.Stringer.
func (m *MessagesMarkDialogUnreadRequest) String() string {
	if m == nil {
		return "MessagesMarkDialogUnreadRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesMarkDialogUnreadRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(m.Flags.String())
	sb.WriteString(",\n")
	sb.WriteString("\tPeer: ")
	sb.WriteString(m.Peer.String())
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (m *MessagesMarkDialogUnreadRequest) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messages.markDialogUnread#c286d98f as nil")
	}
	b.PutID(MessagesMarkDialogUnreadRequestTypeID)
	if err := m.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.markDialogUnread#c286d98f: field flags: %w", err)
	}
	if m.Peer == nil {
		return fmt.Errorf("unable to encode messages.markDialogUnread#c286d98f: field peer is nil")
	}
	if err := m.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.markDialogUnread#c286d98f: field peer: %w", err)
	}
	return nil
}

// SetUnread sets value of Unread conditional field.
func (m *MessagesMarkDialogUnreadRequest) SetUnread(value bool) {
	if value {
		m.Flags.Set(0)
		m.Unread = true
	} else {
		m.Flags.Unset(0)
		m.Unread = false
	}
}

// Decode implements bin.Decoder.
func (m *MessagesMarkDialogUnreadRequest) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messages.markDialogUnread#c286d98f to nil")
	}
	if err := b.ConsumeID(MessagesMarkDialogUnreadRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.markDialogUnread#c286d98f: %w", err)
	}
	{
		if err := m.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.markDialogUnread#c286d98f: field flags: %w", err)
		}
	}
	m.Unread = m.Flags.Has(0)
	{
		value, err := DecodeInputDialogPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.markDialogUnread#c286d98f: field peer: %w", err)
		}
		m.Peer = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesMarkDialogUnreadRequest.
var (
	_ bin.Encoder = &MessagesMarkDialogUnreadRequest{}
	_ bin.Decoder = &MessagesMarkDialogUnreadRequest{}
)

// MessagesMarkDialogUnread invokes method messages.markDialogUnread#c286d98f returning error if any.
// Manually mark dialog as unread
//
// See https://core.telegram.org/method/messages.markDialogUnread for reference.
func (c *Client) MessagesMarkDialogUnread(ctx context.Context, request *MessagesMarkDialogUnreadRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
