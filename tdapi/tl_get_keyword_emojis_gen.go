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

// GetKeywordEmojisRequest represents TL type `getKeywordEmojis#7af81263`.
type GetKeywordEmojisRequest struct {
	// Text to search for
	Text string
	// List of possible IETF language tags of the user's input language; may be empty if
	// unknown
	InputLanguageCodes []string
}

// GetKeywordEmojisRequestTypeID is TL type id of GetKeywordEmojisRequest.
const GetKeywordEmojisRequestTypeID = 0x7af81263

// Ensuring interfaces in compile-time for GetKeywordEmojisRequest.
var (
	_ bin.Encoder     = &GetKeywordEmojisRequest{}
	_ bin.Decoder     = &GetKeywordEmojisRequest{}
	_ bin.BareEncoder = &GetKeywordEmojisRequest{}
	_ bin.BareDecoder = &GetKeywordEmojisRequest{}
)

func (g *GetKeywordEmojisRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Text == "") {
		return false
	}
	if !(g.InputLanguageCodes == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetKeywordEmojisRequest) String() string {
	if g == nil {
		return "GetKeywordEmojisRequest(nil)"
	}
	type Alias GetKeywordEmojisRequest
	return fmt.Sprintf("GetKeywordEmojisRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetKeywordEmojisRequest) TypeID() uint32 {
	return GetKeywordEmojisRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetKeywordEmojisRequest) TypeName() string {
	return "getKeywordEmojis"
}

// TypeInfo returns info about TL type.
func (g *GetKeywordEmojisRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getKeywordEmojis",
		ID:   GetKeywordEmojisRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Text",
			SchemaName: "text",
		},
		{
			Name:       "InputLanguageCodes",
			SchemaName: "input_language_codes",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetKeywordEmojisRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getKeywordEmojis#7af81263 as nil")
	}
	b.PutID(GetKeywordEmojisRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetKeywordEmojisRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getKeywordEmojis#7af81263 as nil")
	}
	b.PutString(g.Text)
	b.PutInt(len(g.InputLanguageCodes))
	for _, v := range g.InputLanguageCodes {
		b.PutString(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetKeywordEmojisRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getKeywordEmojis#7af81263 to nil")
	}
	if err := b.ConsumeID(GetKeywordEmojisRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getKeywordEmojis#7af81263: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetKeywordEmojisRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getKeywordEmojis#7af81263 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getKeywordEmojis#7af81263: field text: %w", err)
		}
		g.Text = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode getKeywordEmojis#7af81263: field input_language_codes: %w", err)
		}

		if headerLen > 0 {
			g.InputLanguageCodes = make([]string, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getKeywordEmojis#7af81263: field input_language_codes: %w", err)
			}
			g.InputLanguageCodes = append(g.InputLanguageCodes, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetKeywordEmojisRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getKeywordEmojis#7af81263 as nil")
	}
	b.ObjStart()
	b.PutID("getKeywordEmojis")
	b.Comma()
	b.FieldStart("text")
	b.PutString(g.Text)
	b.Comma()
	b.FieldStart("input_language_codes")
	b.ArrStart()
	for _, v := range g.InputLanguageCodes {
		b.PutString(v)
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
func (g *GetKeywordEmojisRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getKeywordEmojis#7af81263 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getKeywordEmojis"); err != nil {
				return fmt.Errorf("unable to decode getKeywordEmojis#7af81263: %w", err)
			}
		case "text":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getKeywordEmojis#7af81263: field text: %w", err)
			}
			g.Text = value
		case "input_language_codes":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.String()
				if err != nil {
					return fmt.Errorf("unable to decode getKeywordEmojis#7af81263: field input_language_codes: %w", err)
				}
				g.InputLanguageCodes = append(g.InputLanguageCodes, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode getKeywordEmojis#7af81263: field input_language_codes: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetText returns value of Text field.
func (g *GetKeywordEmojisRequest) GetText() (value string) {
	if g == nil {
		return
	}
	return g.Text
}

// GetInputLanguageCodes returns value of InputLanguageCodes field.
func (g *GetKeywordEmojisRequest) GetInputLanguageCodes() (value []string) {
	if g == nil {
		return
	}
	return g.InputLanguageCodes
}

// GetKeywordEmojis invokes method getKeywordEmojis#7af81263 returning error if any.
func (c *Client) GetKeywordEmojis(ctx context.Context, request *GetKeywordEmojisRequest) (*Emojis, error) {
	var result Emojis

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}