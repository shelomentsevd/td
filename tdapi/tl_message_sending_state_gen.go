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

// MessageSendingStatePending represents TL type `messageSendingStatePending#f32b63b4`.
type MessageSendingStatePending struct {
	// Non-persistent message sending identifier, specified by the application
	SendingID int32
}

// MessageSendingStatePendingTypeID is TL type id of MessageSendingStatePending.
const MessageSendingStatePendingTypeID = 0xf32b63b4

// construct implements constructor of MessageSendingStateClass.
func (m MessageSendingStatePending) construct() MessageSendingStateClass { return &m }

// Ensuring interfaces in compile-time for MessageSendingStatePending.
var (
	_ bin.Encoder     = &MessageSendingStatePending{}
	_ bin.Decoder     = &MessageSendingStatePending{}
	_ bin.BareEncoder = &MessageSendingStatePending{}
	_ bin.BareDecoder = &MessageSendingStatePending{}

	_ MessageSendingStateClass = &MessageSendingStatePending{}
)

func (m *MessageSendingStatePending) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.SendingID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageSendingStatePending) String() string {
	if m == nil {
		return "MessageSendingStatePending(nil)"
	}
	type Alias MessageSendingStatePending
	return fmt.Sprintf("MessageSendingStatePending%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageSendingStatePending) TypeID() uint32 {
	return MessageSendingStatePendingTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageSendingStatePending) TypeName() string {
	return "messageSendingStatePending"
}

// TypeInfo returns info about TL type.
func (m *MessageSendingStatePending) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messageSendingStatePending",
		ID:   MessageSendingStatePendingTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SendingID",
			SchemaName: "sending_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessageSendingStatePending) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageSendingStatePending#f32b63b4 as nil")
	}
	b.PutID(MessageSendingStatePendingTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessageSendingStatePending) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageSendingStatePending#f32b63b4 as nil")
	}
	b.PutInt32(m.SendingID)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageSendingStatePending) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageSendingStatePending#f32b63b4 to nil")
	}
	if err := b.ConsumeID(MessageSendingStatePendingTypeID); err != nil {
		return fmt.Errorf("unable to decode messageSendingStatePending#f32b63b4: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessageSendingStatePending) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageSendingStatePending#f32b63b4 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode messageSendingStatePending#f32b63b4: field sending_id: %w", err)
		}
		m.SendingID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (m *MessageSendingStatePending) EncodeTDLibJSON(b tdjson.Encoder) error {
	if m == nil {
		return fmt.Errorf("can't encode messageSendingStatePending#f32b63b4 as nil")
	}
	b.ObjStart()
	b.PutID("messageSendingStatePending")
	b.Comma()
	b.FieldStart("sending_id")
	b.PutInt32(m.SendingID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (m *MessageSendingStatePending) DecodeTDLibJSON(b tdjson.Decoder) error {
	if m == nil {
		return fmt.Errorf("can't decode messageSendingStatePending#f32b63b4 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("messageSendingStatePending"); err != nil {
				return fmt.Errorf("unable to decode messageSendingStatePending#f32b63b4: %w", err)
			}
		case "sending_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode messageSendingStatePending#f32b63b4: field sending_id: %w", err)
			}
			m.SendingID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSendingID returns value of SendingID field.
func (m *MessageSendingStatePending) GetSendingID() (value int32) {
	if m == nil {
		return
	}
	return m.SendingID
}

// MessageSendingStateFailed represents TL type `messageSendingStateFailed#ac81ee5e`.
type MessageSendingStateFailed struct {
	// The cause of the message sending failure
	Error Error
	// True, if the message can be re-sent using resendMessages or
	// readdQuickReplyShortcutMessages
	CanRetry bool
	// True, if the message can be re-sent only on behalf of a different sender
	NeedAnotherSender bool
	// True, if the message can be re-sent only if another quote is chosen in the message
	// that is replied by the given message
	NeedAnotherReplyQuote bool
	// True, if the message can be re-sent only if the message to be replied is removed. This
	// will be done automatically by resendMessages
	NeedDropReply bool
	// Time left before the message can be re-sent, in seconds. No update is sent when this
	// field changes
	RetryAfter float64
}

// MessageSendingStateFailedTypeID is TL type id of MessageSendingStateFailed.
const MessageSendingStateFailedTypeID = 0xac81ee5e

// construct implements constructor of MessageSendingStateClass.
func (m MessageSendingStateFailed) construct() MessageSendingStateClass { return &m }

// Ensuring interfaces in compile-time for MessageSendingStateFailed.
var (
	_ bin.Encoder     = &MessageSendingStateFailed{}
	_ bin.Decoder     = &MessageSendingStateFailed{}
	_ bin.BareEncoder = &MessageSendingStateFailed{}
	_ bin.BareDecoder = &MessageSendingStateFailed{}

	_ MessageSendingStateClass = &MessageSendingStateFailed{}
)

func (m *MessageSendingStateFailed) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.Error.Zero()) {
		return false
	}
	if !(m.CanRetry == false) {
		return false
	}
	if !(m.NeedAnotherSender == false) {
		return false
	}
	if !(m.NeedAnotherReplyQuote == false) {
		return false
	}
	if !(m.NeedDropReply == false) {
		return false
	}
	if !(m.RetryAfter == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageSendingStateFailed) String() string {
	if m == nil {
		return "MessageSendingStateFailed(nil)"
	}
	type Alias MessageSendingStateFailed
	return fmt.Sprintf("MessageSendingStateFailed%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageSendingStateFailed) TypeID() uint32 {
	return MessageSendingStateFailedTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageSendingStateFailed) TypeName() string {
	return "messageSendingStateFailed"
}

// TypeInfo returns info about TL type.
func (m *MessageSendingStateFailed) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messageSendingStateFailed",
		ID:   MessageSendingStateFailedTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Error",
			SchemaName: "error",
		},
		{
			Name:       "CanRetry",
			SchemaName: "can_retry",
		},
		{
			Name:       "NeedAnotherSender",
			SchemaName: "need_another_sender",
		},
		{
			Name:       "NeedAnotherReplyQuote",
			SchemaName: "need_another_reply_quote",
		},
		{
			Name:       "NeedDropReply",
			SchemaName: "need_drop_reply",
		},
		{
			Name:       "RetryAfter",
			SchemaName: "retry_after",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessageSendingStateFailed) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageSendingStateFailed#ac81ee5e as nil")
	}
	b.PutID(MessageSendingStateFailedTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessageSendingStateFailed) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageSendingStateFailed#ac81ee5e as nil")
	}
	if err := m.Error.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messageSendingStateFailed#ac81ee5e: field error: %w", err)
	}
	b.PutBool(m.CanRetry)
	b.PutBool(m.NeedAnotherSender)
	b.PutBool(m.NeedAnotherReplyQuote)
	b.PutBool(m.NeedDropReply)
	b.PutDouble(m.RetryAfter)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageSendingStateFailed) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageSendingStateFailed#ac81ee5e to nil")
	}
	if err := b.ConsumeID(MessageSendingStateFailedTypeID); err != nil {
		return fmt.Errorf("unable to decode messageSendingStateFailed#ac81ee5e: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessageSendingStateFailed) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageSendingStateFailed#ac81ee5e to nil")
	}
	{
		if err := m.Error.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messageSendingStateFailed#ac81ee5e: field error: %w", err)
		}
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode messageSendingStateFailed#ac81ee5e: field can_retry: %w", err)
		}
		m.CanRetry = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode messageSendingStateFailed#ac81ee5e: field need_another_sender: %w", err)
		}
		m.NeedAnotherSender = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode messageSendingStateFailed#ac81ee5e: field need_another_reply_quote: %w", err)
		}
		m.NeedAnotherReplyQuote = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode messageSendingStateFailed#ac81ee5e: field need_drop_reply: %w", err)
		}
		m.NeedDropReply = value
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode messageSendingStateFailed#ac81ee5e: field retry_after: %w", err)
		}
		m.RetryAfter = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (m *MessageSendingStateFailed) EncodeTDLibJSON(b tdjson.Encoder) error {
	if m == nil {
		return fmt.Errorf("can't encode messageSendingStateFailed#ac81ee5e as nil")
	}
	b.ObjStart()
	b.PutID("messageSendingStateFailed")
	b.Comma()
	b.FieldStart("error")
	if err := m.Error.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode messageSendingStateFailed#ac81ee5e: field error: %w", err)
	}
	b.Comma()
	b.FieldStart("can_retry")
	b.PutBool(m.CanRetry)
	b.Comma()
	b.FieldStart("need_another_sender")
	b.PutBool(m.NeedAnotherSender)
	b.Comma()
	b.FieldStart("need_another_reply_quote")
	b.PutBool(m.NeedAnotherReplyQuote)
	b.Comma()
	b.FieldStart("need_drop_reply")
	b.PutBool(m.NeedDropReply)
	b.Comma()
	b.FieldStart("retry_after")
	b.PutDouble(m.RetryAfter)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (m *MessageSendingStateFailed) DecodeTDLibJSON(b tdjson.Decoder) error {
	if m == nil {
		return fmt.Errorf("can't decode messageSendingStateFailed#ac81ee5e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("messageSendingStateFailed"); err != nil {
				return fmt.Errorf("unable to decode messageSendingStateFailed#ac81ee5e: %w", err)
			}
		case "error":
			if err := m.Error.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode messageSendingStateFailed#ac81ee5e: field error: %w", err)
			}
		case "can_retry":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode messageSendingStateFailed#ac81ee5e: field can_retry: %w", err)
			}
			m.CanRetry = value
		case "need_another_sender":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode messageSendingStateFailed#ac81ee5e: field need_another_sender: %w", err)
			}
			m.NeedAnotherSender = value
		case "need_another_reply_quote":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode messageSendingStateFailed#ac81ee5e: field need_another_reply_quote: %w", err)
			}
			m.NeedAnotherReplyQuote = value
		case "need_drop_reply":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode messageSendingStateFailed#ac81ee5e: field need_drop_reply: %w", err)
			}
			m.NeedDropReply = value
		case "retry_after":
			value, err := b.Double()
			if err != nil {
				return fmt.Errorf("unable to decode messageSendingStateFailed#ac81ee5e: field retry_after: %w", err)
			}
			m.RetryAfter = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetError returns value of Error field.
func (m *MessageSendingStateFailed) GetError() (value Error) {
	if m == nil {
		return
	}
	return m.Error
}

// GetCanRetry returns value of CanRetry field.
func (m *MessageSendingStateFailed) GetCanRetry() (value bool) {
	if m == nil {
		return
	}
	return m.CanRetry
}

// GetNeedAnotherSender returns value of NeedAnotherSender field.
func (m *MessageSendingStateFailed) GetNeedAnotherSender() (value bool) {
	if m == nil {
		return
	}
	return m.NeedAnotherSender
}

// GetNeedAnotherReplyQuote returns value of NeedAnotherReplyQuote field.
func (m *MessageSendingStateFailed) GetNeedAnotherReplyQuote() (value bool) {
	if m == nil {
		return
	}
	return m.NeedAnotherReplyQuote
}

// GetNeedDropReply returns value of NeedDropReply field.
func (m *MessageSendingStateFailed) GetNeedDropReply() (value bool) {
	if m == nil {
		return
	}
	return m.NeedDropReply
}

// GetRetryAfter returns value of RetryAfter field.
func (m *MessageSendingStateFailed) GetRetryAfter() (value float64) {
	if m == nil {
		return
	}
	return m.RetryAfter
}

// MessageSendingStateClassName is schema name of MessageSendingStateClass.
const MessageSendingStateClassName = "MessageSendingState"

// MessageSendingStateClass represents MessageSendingState generic type.
//
// Example:
//
//	g, err := tdapi.DecodeMessageSendingState(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.MessageSendingStatePending: // messageSendingStatePending#f32b63b4
//	case *tdapi.MessageSendingStateFailed: // messageSendingStateFailed#ac81ee5e
//	default: panic(v)
//	}
type MessageSendingStateClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() MessageSendingStateClass

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

	EncodeTDLibJSON(b tdjson.Encoder) error
	DecodeTDLibJSON(b tdjson.Decoder) error
}

// DecodeMessageSendingState implements binary de-serialization for MessageSendingStateClass.
func DecodeMessageSendingState(buf *bin.Buffer) (MessageSendingStateClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessageSendingStatePendingTypeID:
		// Decoding messageSendingStatePending#f32b63b4.
		v := MessageSendingStatePending{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageSendingStateClass: %w", err)
		}
		return &v, nil
	case MessageSendingStateFailedTypeID:
		// Decoding messageSendingStateFailed#ac81ee5e.
		v := MessageSendingStateFailed{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageSendingStateClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessageSendingStateClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONMessageSendingState implements binary de-serialization for MessageSendingStateClass.
func DecodeTDLibJSONMessageSendingState(buf tdjson.Decoder) (MessageSendingStateClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "messageSendingStatePending":
		// Decoding messageSendingStatePending#f32b63b4.
		v := MessageSendingStatePending{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageSendingStateClass: %w", err)
		}
		return &v, nil
	case "messageSendingStateFailed":
		// Decoding messageSendingStateFailed#ac81ee5e.
		v := MessageSendingStateFailed{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageSendingStateClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessageSendingStateClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// MessageSendingState boxes the MessageSendingStateClass providing a helper.
type MessageSendingStateBox struct {
	MessageSendingState MessageSendingStateClass
}

// Decode implements bin.Decoder for MessageSendingStateBox.
func (b *MessageSendingStateBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessageSendingStateBox to nil")
	}
	v, err := DecodeMessageSendingState(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.MessageSendingState = v
	return nil
}

// Encode implements bin.Encode for MessageSendingStateBox.
func (b *MessageSendingStateBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.MessageSendingState == nil {
		return fmt.Errorf("unable to encode MessageSendingStateClass as nil")
	}
	return b.MessageSendingState.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for MessageSendingStateBox.
func (b *MessageSendingStateBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessageSendingStateBox to nil")
	}
	v, err := DecodeTDLibJSONMessageSendingState(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.MessageSendingState = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for MessageSendingStateBox.
func (b *MessageSendingStateBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.MessageSendingState == nil {
		return fmt.Errorf("unable to encode MessageSendingStateClass as nil")
	}
	return b.MessageSendingState.EncodeTDLibJSON(buf)
}
