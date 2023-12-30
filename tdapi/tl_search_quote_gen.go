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

// SearchQuoteRequest represents TL type `searchQuote#6864011f`.
type SearchQuoteRequest struct {
	// Text in which to search for the quote
	Text FormattedText
	// Quote to search for
	Quote FormattedText
	// Approximate quote position in UTF-16 code units
	QuotePosition int32
}

// SearchQuoteRequestTypeID is TL type id of SearchQuoteRequest.
const SearchQuoteRequestTypeID = 0x6864011f

// Ensuring interfaces in compile-time for SearchQuoteRequest.
var (
	_ bin.Encoder     = &SearchQuoteRequest{}
	_ bin.Decoder     = &SearchQuoteRequest{}
	_ bin.BareEncoder = &SearchQuoteRequest{}
	_ bin.BareDecoder = &SearchQuoteRequest{}
)

func (s *SearchQuoteRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Text.Zero()) {
		return false
	}
	if !(s.Quote.Zero()) {
		return false
	}
	if !(s.QuotePosition == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SearchQuoteRequest) String() string {
	if s == nil {
		return "SearchQuoteRequest(nil)"
	}
	type Alias SearchQuoteRequest
	return fmt.Sprintf("SearchQuoteRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SearchQuoteRequest) TypeID() uint32 {
	return SearchQuoteRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SearchQuoteRequest) TypeName() string {
	return "searchQuote"
}

// TypeInfo returns info about TL type.
func (s *SearchQuoteRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "searchQuote",
		ID:   SearchQuoteRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Text",
			SchemaName: "text",
		},
		{
			Name:       "Quote",
			SchemaName: "quote",
		},
		{
			Name:       "QuotePosition",
			SchemaName: "quote_position",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SearchQuoteRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchQuote#6864011f as nil")
	}
	b.PutID(SearchQuoteRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SearchQuoteRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchQuote#6864011f as nil")
	}
	if err := s.Text.Encode(b); err != nil {
		return fmt.Errorf("unable to encode searchQuote#6864011f: field text: %w", err)
	}
	if err := s.Quote.Encode(b); err != nil {
		return fmt.Errorf("unable to encode searchQuote#6864011f: field quote: %w", err)
	}
	b.PutInt32(s.QuotePosition)
	return nil
}

// Decode implements bin.Decoder.
func (s *SearchQuoteRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchQuote#6864011f to nil")
	}
	if err := b.ConsumeID(SearchQuoteRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode searchQuote#6864011f: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SearchQuoteRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchQuote#6864011f to nil")
	}
	{
		if err := s.Text.Decode(b); err != nil {
			return fmt.Errorf("unable to decode searchQuote#6864011f: field text: %w", err)
		}
	}
	{
		if err := s.Quote.Decode(b); err != nil {
			return fmt.Errorf("unable to decode searchQuote#6864011f: field quote: %w", err)
		}
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode searchQuote#6864011f: field quote_position: %w", err)
		}
		s.QuotePosition = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SearchQuoteRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode searchQuote#6864011f as nil")
	}
	b.ObjStart()
	b.PutID("searchQuote")
	b.Comma()
	b.FieldStart("text")
	if err := s.Text.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode searchQuote#6864011f: field text: %w", err)
	}
	b.Comma()
	b.FieldStart("quote")
	if err := s.Quote.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode searchQuote#6864011f: field quote: %w", err)
	}
	b.Comma()
	b.FieldStart("quote_position")
	b.PutInt32(s.QuotePosition)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SearchQuoteRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode searchQuote#6864011f to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("searchQuote"); err != nil {
				return fmt.Errorf("unable to decode searchQuote#6864011f: %w", err)
			}
		case "text":
			if err := s.Text.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode searchQuote#6864011f: field text: %w", err)
			}
		case "quote":
			if err := s.Quote.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode searchQuote#6864011f: field quote: %w", err)
			}
		case "quote_position":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode searchQuote#6864011f: field quote_position: %w", err)
			}
			s.QuotePosition = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetText returns value of Text field.
func (s *SearchQuoteRequest) GetText() (value FormattedText) {
	if s == nil {
		return
	}
	return s.Text
}

// GetQuote returns value of Quote field.
func (s *SearchQuoteRequest) GetQuote() (value FormattedText) {
	if s == nil {
		return
	}
	return s.Quote
}

// GetQuotePosition returns value of QuotePosition field.
func (s *SearchQuoteRequest) GetQuotePosition() (value int32) {
	if s == nil {
		return
	}
	return s.QuotePosition
}

// SearchQuote invokes method searchQuote#6864011f returning error if any.
func (c *Client) SearchQuote(ctx context.Context, request *SearchQuoteRequest) (*FoundPosition, error) {
	var result FoundPosition

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}