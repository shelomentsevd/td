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

// GetChatFolderChatCountRequest represents TL type `getChatFolderChatCount#7dd4cbbe`.
type GetChatFolderChatCountRequest struct {
	// The new chat folder
	Folder ChatFolder
}

// GetChatFolderChatCountRequestTypeID is TL type id of GetChatFolderChatCountRequest.
const GetChatFolderChatCountRequestTypeID = 0x7dd4cbbe

// Ensuring interfaces in compile-time for GetChatFolderChatCountRequest.
var (
	_ bin.Encoder     = &GetChatFolderChatCountRequest{}
	_ bin.Decoder     = &GetChatFolderChatCountRequest{}
	_ bin.BareEncoder = &GetChatFolderChatCountRequest{}
	_ bin.BareDecoder = &GetChatFolderChatCountRequest{}
)

func (g *GetChatFolderChatCountRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Folder.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetChatFolderChatCountRequest) String() string {
	if g == nil {
		return "GetChatFolderChatCountRequest(nil)"
	}
	type Alias GetChatFolderChatCountRequest
	return fmt.Sprintf("GetChatFolderChatCountRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetChatFolderChatCountRequest) TypeID() uint32 {
	return GetChatFolderChatCountRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetChatFolderChatCountRequest) TypeName() string {
	return "getChatFolderChatCount"
}

// TypeInfo returns info about TL type.
func (g *GetChatFolderChatCountRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getChatFolderChatCount",
		ID:   GetChatFolderChatCountRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Folder",
			SchemaName: "folder",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetChatFolderChatCountRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatFolderChatCount#7dd4cbbe as nil")
	}
	b.PutID(GetChatFolderChatCountRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetChatFolderChatCountRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatFolderChatCount#7dd4cbbe as nil")
	}
	if err := g.Folder.Encode(b); err != nil {
		return fmt.Errorf("unable to encode getChatFolderChatCount#7dd4cbbe: field folder: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetChatFolderChatCountRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatFolderChatCount#7dd4cbbe to nil")
	}
	if err := b.ConsumeID(GetChatFolderChatCountRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getChatFolderChatCount#7dd4cbbe: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetChatFolderChatCountRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatFolderChatCount#7dd4cbbe to nil")
	}
	{
		if err := g.Folder.Decode(b); err != nil {
			return fmt.Errorf("unable to decode getChatFolderChatCount#7dd4cbbe: field folder: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetChatFolderChatCountRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatFolderChatCount#7dd4cbbe as nil")
	}
	b.ObjStart()
	b.PutID("getChatFolderChatCount")
	b.Comma()
	b.FieldStart("folder")
	if err := g.Folder.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode getChatFolderChatCount#7dd4cbbe: field folder: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetChatFolderChatCountRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatFolderChatCount#7dd4cbbe to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getChatFolderChatCount"); err != nil {
				return fmt.Errorf("unable to decode getChatFolderChatCount#7dd4cbbe: %w", err)
			}
		case "folder":
			if err := g.Folder.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode getChatFolderChatCount#7dd4cbbe: field folder: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetFolder returns value of Folder field.
func (g *GetChatFolderChatCountRequest) GetFolder() (value ChatFolder) {
	if g == nil {
		return
	}
	return g.Folder
}

// GetChatFolderChatCount invokes method getChatFolderChatCount#7dd4cbbe returning error if any.
func (c *Client) GetChatFolderChatCount(ctx context.Context, folder ChatFolder) (*Count, error) {
	var result Count

	request := &GetChatFolderChatCountRequest{
		Folder: folder,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}