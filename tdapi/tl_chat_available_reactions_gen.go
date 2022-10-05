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

// ChatAvailableReactionsAll represents TL type `chatAvailableReactionsAll#dff07c4e`.
type ChatAvailableReactionsAll struct {
}

// ChatAvailableReactionsAllTypeID is TL type id of ChatAvailableReactionsAll.
const ChatAvailableReactionsAllTypeID = 0xdff07c4e

// construct implements constructor of ChatAvailableReactionsClass.
func (c ChatAvailableReactionsAll) construct() ChatAvailableReactionsClass { return &c }

// Ensuring interfaces in compile-time for ChatAvailableReactionsAll.
var (
	_ bin.Encoder     = &ChatAvailableReactionsAll{}
	_ bin.Decoder     = &ChatAvailableReactionsAll{}
	_ bin.BareEncoder = &ChatAvailableReactionsAll{}
	_ bin.BareDecoder = &ChatAvailableReactionsAll{}

	_ ChatAvailableReactionsClass = &ChatAvailableReactionsAll{}
)

func (c *ChatAvailableReactionsAll) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatAvailableReactionsAll) String() string {
	if c == nil {
		return "ChatAvailableReactionsAll(nil)"
	}
	type Alias ChatAvailableReactionsAll
	return fmt.Sprintf("ChatAvailableReactionsAll%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatAvailableReactionsAll) TypeID() uint32 {
	return ChatAvailableReactionsAllTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatAvailableReactionsAll) TypeName() string {
	return "chatAvailableReactionsAll"
}

// TypeInfo returns info about TL type.
func (c *ChatAvailableReactionsAll) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatAvailableReactionsAll",
		ID:   ChatAvailableReactionsAllTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatAvailableReactionsAll) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatAvailableReactionsAll#dff07c4e as nil")
	}
	b.PutID(ChatAvailableReactionsAllTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatAvailableReactionsAll) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatAvailableReactionsAll#dff07c4e as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatAvailableReactionsAll) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatAvailableReactionsAll#dff07c4e to nil")
	}
	if err := b.ConsumeID(ChatAvailableReactionsAllTypeID); err != nil {
		return fmt.Errorf("unable to decode chatAvailableReactionsAll#dff07c4e: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatAvailableReactionsAll) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatAvailableReactionsAll#dff07c4e to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatAvailableReactionsAll) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatAvailableReactionsAll#dff07c4e as nil")
	}
	b.ObjStart()
	b.PutID("chatAvailableReactionsAll")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatAvailableReactionsAll) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatAvailableReactionsAll#dff07c4e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatAvailableReactionsAll"); err != nil {
				return fmt.Errorf("unable to decode chatAvailableReactionsAll#dff07c4e: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// ChatAvailableReactionsSome represents TL type `chatAvailableReactionsSome#d3cc0a6b`.
type ChatAvailableReactionsSome struct {
	// The list of reactions
	Reactions []ReactionTypeClass
}

// ChatAvailableReactionsSomeTypeID is TL type id of ChatAvailableReactionsSome.
const ChatAvailableReactionsSomeTypeID = 0xd3cc0a6b

// construct implements constructor of ChatAvailableReactionsClass.
func (c ChatAvailableReactionsSome) construct() ChatAvailableReactionsClass { return &c }

// Ensuring interfaces in compile-time for ChatAvailableReactionsSome.
var (
	_ bin.Encoder     = &ChatAvailableReactionsSome{}
	_ bin.Decoder     = &ChatAvailableReactionsSome{}
	_ bin.BareEncoder = &ChatAvailableReactionsSome{}
	_ bin.BareDecoder = &ChatAvailableReactionsSome{}

	_ ChatAvailableReactionsClass = &ChatAvailableReactionsSome{}
)

func (c *ChatAvailableReactionsSome) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Reactions == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatAvailableReactionsSome) String() string {
	if c == nil {
		return "ChatAvailableReactionsSome(nil)"
	}
	type Alias ChatAvailableReactionsSome
	return fmt.Sprintf("ChatAvailableReactionsSome%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatAvailableReactionsSome) TypeID() uint32 {
	return ChatAvailableReactionsSomeTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatAvailableReactionsSome) TypeName() string {
	return "chatAvailableReactionsSome"
}

// TypeInfo returns info about TL type.
func (c *ChatAvailableReactionsSome) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatAvailableReactionsSome",
		ID:   ChatAvailableReactionsSomeTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Reactions",
			SchemaName: "reactions",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatAvailableReactionsSome) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatAvailableReactionsSome#d3cc0a6b as nil")
	}
	b.PutID(ChatAvailableReactionsSomeTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatAvailableReactionsSome) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatAvailableReactionsSome#d3cc0a6b as nil")
	}
	b.PutInt(len(c.Reactions))
	for idx, v := range c.Reactions {
		if v == nil {
			return fmt.Errorf("unable to encode chatAvailableReactionsSome#d3cc0a6b: field reactions element with index %d is nil", idx)
		}
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare chatAvailableReactionsSome#d3cc0a6b: field reactions element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatAvailableReactionsSome) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatAvailableReactionsSome#d3cc0a6b to nil")
	}
	if err := b.ConsumeID(ChatAvailableReactionsSomeTypeID); err != nil {
		return fmt.Errorf("unable to decode chatAvailableReactionsSome#d3cc0a6b: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatAvailableReactionsSome) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatAvailableReactionsSome#d3cc0a6b to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatAvailableReactionsSome#d3cc0a6b: field reactions: %w", err)
		}

		if headerLen > 0 {
			c.Reactions = make([]ReactionTypeClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeReactionType(b)
			if err != nil {
				return fmt.Errorf("unable to decode chatAvailableReactionsSome#d3cc0a6b: field reactions: %w", err)
			}
			c.Reactions = append(c.Reactions, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatAvailableReactionsSome) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatAvailableReactionsSome#d3cc0a6b as nil")
	}
	b.ObjStart()
	b.PutID("chatAvailableReactionsSome")
	b.Comma()
	b.FieldStart("reactions")
	b.ArrStart()
	for idx, v := range c.Reactions {
		if v == nil {
			return fmt.Errorf("unable to encode chatAvailableReactionsSome#d3cc0a6b: field reactions element with index %d is nil", idx)
		}
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode chatAvailableReactionsSome#d3cc0a6b: field reactions element with index %d: %w", idx, err)
		}
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatAvailableReactionsSome) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatAvailableReactionsSome#d3cc0a6b to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatAvailableReactionsSome"); err != nil {
				return fmt.Errorf("unable to decode chatAvailableReactionsSome#d3cc0a6b: %w", err)
			}
		case "reactions":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := DecodeTDLibJSONReactionType(b)
				if err != nil {
					return fmt.Errorf("unable to decode chatAvailableReactionsSome#d3cc0a6b: field reactions: %w", err)
				}
				c.Reactions = append(c.Reactions, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode chatAvailableReactionsSome#d3cc0a6b: field reactions: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetReactions returns value of Reactions field.
func (c *ChatAvailableReactionsSome) GetReactions() (value []ReactionTypeClass) {
	if c == nil {
		return
	}
	return c.Reactions
}

// ChatAvailableReactionsClassName is schema name of ChatAvailableReactionsClass.
const ChatAvailableReactionsClassName = "ChatAvailableReactions"

// ChatAvailableReactionsClass represents ChatAvailableReactions generic type.
//
// Example:
//
//	g, err := tdapi.DecodeChatAvailableReactions(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.ChatAvailableReactionsAll: // chatAvailableReactionsAll#dff07c4e
//	case *tdapi.ChatAvailableReactionsSome: // chatAvailableReactionsSome#d3cc0a6b
//	default: panic(v)
//	}
type ChatAvailableReactionsClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() ChatAvailableReactionsClass

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

// DecodeChatAvailableReactions implements binary de-serialization for ChatAvailableReactionsClass.
func DecodeChatAvailableReactions(buf *bin.Buffer) (ChatAvailableReactionsClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ChatAvailableReactionsAllTypeID:
		// Decoding chatAvailableReactionsAll#dff07c4e.
		v := ChatAvailableReactionsAll{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatAvailableReactionsClass: %w", err)
		}
		return &v, nil
	case ChatAvailableReactionsSomeTypeID:
		// Decoding chatAvailableReactionsSome#d3cc0a6b.
		v := ChatAvailableReactionsSome{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatAvailableReactionsClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ChatAvailableReactionsClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONChatAvailableReactions implements binary de-serialization for ChatAvailableReactionsClass.
func DecodeTDLibJSONChatAvailableReactions(buf tdjson.Decoder) (ChatAvailableReactionsClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "chatAvailableReactionsAll":
		// Decoding chatAvailableReactionsAll#dff07c4e.
		v := ChatAvailableReactionsAll{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatAvailableReactionsClass: %w", err)
		}
		return &v, nil
	case "chatAvailableReactionsSome":
		// Decoding chatAvailableReactionsSome#d3cc0a6b.
		v := ChatAvailableReactionsSome{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatAvailableReactionsClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ChatAvailableReactionsClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// ChatAvailableReactions boxes the ChatAvailableReactionsClass providing a helper.
type ChatAvailableReactionsBox struct {
	ChatAvailableReactions ChatAvailableReactionsClass
}

// Decode implements bin.Decoder for ChatAvailableReactionsBox.
func (b *ChatAvailableReactionsBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ChatAvailableReactionsBox to nil")
	}
	v, err := DecodeChatAvailableReactions(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ChatAvailableReactions = v
	return nil
}

// Encode implements bin.Encode for ChatAvailableReactionsBox.
func (b *ChatAvailableReactionsBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ChatAvailableReactions == nil {
		return fmt.Errorf("unable to encode ChatAvailableReactionsClass as nil")
	}
	return b.ChatAvailableReactions.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for ChatAvailableReactionsBox.
func (b *ChatAvailableReactionsBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode ChatAvailableReactionsBox to nil")
	}
	v, err := DecodeTDLibJSONChatAvailableReactions(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ChatAvailableReactions = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for ChatAvailableReactionsBox.
func (b *ChatAvailableReactionsBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.ChatAvailableReactions == nil {
		return fmt.Errorf("unable to encode ChatAvailableReactionsClass as nil")
	}
	return b.ChatAvailableReactions.EncodeTDLibJSON(buf)
}