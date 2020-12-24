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

// PhoneCallProtocol represents TL type `phoneCallProtocol#fc878fc8`.
// Protocol info for libtgvoip
//
// See https://core.telegram.org/constructor/phoneCallProtocol for reference.
type PhoneCallProtocol struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether to allow P2P connection to the other participant
	UDPP2P bool
	// Whether to allow connection to the other participants through the reflector servers
	UDPReflector bool
	// Minimum layer for remote libtgvoip
	MinLayer int
	// Maximum layer for remote libtgvoip
	MaxLayer int
	// When using phone.requestCall¹ and phone.acceptCall², specify all library versions supported by the client. The server will merge and choose the best library version supported by both peers, returning only the best value in the result of the callee's phone.acceptCall³ and in the phoneCallAccepted⁴ update received by the caller.
	//
	// Links:
	//  1) https://core.telegram.org/method/phone.requestCall
	//  2) https://core.telegram.org/method/phone.acceptCall
	//  3) https://core.telegram.org/method/phone.acceptCall
	//  4) https://core.telegram.org/constructor/phoneCallAccepted
	LibraryVersions []string
}

// PhoneCallProtocolTypeID is TL type id of PhoneCallProtocol.
const PhoneCallProtocolTypeID = 0xfc878fc8

// String implements fmt.Stringer.
func (p *PhoneCallProtocol) String() string {
	if p == nil {
		return "PhoneCallProtocol(nil)"
	}
	var sb strings.Builder
	sb.WriteString("PhoneCallProtocol")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(p.Flags.String())
	sb.WriteString(",\n")
	sb.WriteString("\tMinLayer: ")
	sb.WriteString(fmt.Sprint(p.MinLayer))
	sb.WriteString(",\n")
	sb.WriteString("\tMaxLayer: ")
	sb.WriteString(fmt.Sprint(p.MaxLayer))
	sb.WriteString(",\n")
	sb.WriteByte('[')
	for _, v := range p.LibraryVersions {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (p *PhoneCallProtocol) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode phoneCallProtocol#fc878fc8 as nil")
	}
	b.PutID(PhoneCallProtocolTypeID)
	if err := p.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phoneCallProtocol#fc878fc8: field flags: %w", err)
	}
	b.PutInt(p.MinLayer)
	b.PutInt(p.MaxLayer)
	b.PutVectorHeader(len(p.LibraryVersions))
	for _, v := range p.LibraryVersions {
		b.PutString(v)
	}
	return nil
}

// SetUDPP2P sets value of UDPP2P conditional field.
func (p *PhoneCallProtocol) SetUDPP2P(value bool) {
	if value {
		p.Flags.Set(0)
		p.UDPP2P = true
	} else {
		p.Flags.Unset(0)
		p.UDPP2P = false
	}
}

// SetUDPReflector sets value of UDPReflector conditional field.
func (p *PhoneCallProtocol) SetUDPReflector(value bool) {
	if value {
		p.Flags.Set(1)
		p.UDPReflector = true
	} else {
		p.Flags.Unset(1)
		p.UDPReflector = false
	}
}

// Decode implements bin.Decoder.
func (p *PhoneCallProtocol) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode phoneCallProtocol#fc878fc8 to nil")
	}
	if err := b.ConsumeID(PhoneCallProtocolTypeID); err != nil {
		return fmt.Errorf("unable to decode phoneCallProtocol#fc878fc8: %w", err)
	}
	{
		if err := p.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phoneCallProtocol#fc878fc8: field flags: %w", err)
		}
	}
	p.UDPP2P = p.Flags.Has(0)
	p.UDPReflector = p.Flags.Has(1)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phoneCallProtocol#fc878fc8: field min_layer: %w", err)
		}
		p.MinLayer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phoneCallProtocol#fc878fc8: field max_layer: %w", err)
		}
		p.MaxLayer = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode phoneCallProtocol#fc878fc8: field library_versions: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode phoneCallProtocol#fc878fc8: field library_versions: %w", err)
			}
			p.LibraryVersions = append(p.LibraryVersions, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for PhoneCallProtocol.
var (
	_ bin.Encoder = &PhoneCallProtocol{}
	_ bin.Decoder = &PhoneCallProtocol{}
)
