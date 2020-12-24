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

// Invoice represents TL type `invoice#c30aa358`.
// Invoice
//
// See https://core.telegram.org/constructor/invoice for reference.
type Invoice struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Test invoice
	Test bool
	// Set this flag if you require the user's full name to complete the order
	NameRequested bool
	// Set this flag if you require the user's phone number to complete the order
	PhoneRequested bool
	// Set this flag if you require the user's email address to complete the order
	EmailRequested bool
	// Set this flag if you require the user's shipping address to complete the order
	ShippingAddressRequested bool
	// Set this flag if the final price depends on the shipping method
	Flexible bool
	// Set this flag if user's phone number should be sent to provider
	PhoneToProvider bool
	// Set this flag if user's email address should be sent to provider
	EmailToProvider bool
	// Three-letter ISO 4217 currency¹ code
	//
	// Links:
	//  1) https://core.telegram.org/bots/payments#supported-currencies
	Currency string
	// Price breakdown, a list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.)
	Prices []LabeledPrice
}

// InvoiceTypeID is TL type id of Invoice.
const InvoiceTypeID = 0xc30aa358

// String implements fmt.Stringer.
func (i *Invoice) String() string {
	if i == nil {
		return "Invoice(nil)"
	}
	var sb strings.Builder
	sb.WriteString("Invoice")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(i.Flags.String())
	sb.WriteString(",\n")
	sb.WriteString("\tCurrency: ")
	sb.WriteString(fmt.Sprint(i.Currency))
	sb.WriteString(",\n")
	sb.WriteByte('[')
	for _, v := range i.Prices {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (i *Invoice) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode invoice#c30aa358 as nil")
	}
	b.PutID(InvoiceTypeID)
	if err := i.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode invoice#c30aa358: field flags: %w", err)
	}
	b.PutString(i.Currency)
	b.PutVectorHeader(len(i.Prices))
	for idx, v := range i.Prices {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode invoice#c30aa358: field prices element with index %d: %w", idx, err)
		}
	}
	return nil
}

// SetTest sets value of Test conditional field.
func (i *Invoice) SetTest(value bool) {
	if value {
		i.Flags.Set(0)
		i.Test = true
	} else {
		i.Flags.Unset(0)
		i.Test = false
	}
}

// SetNameRequested sets value of NameRequested conditional field.
func (i *Invoice) SetNameRequested(value bool) {
	if value {
		i.Flags.Set(1)
		i.NameRequested = true
	} else {
		i.Flags.Unset(1)
		i.NameRequested = false
	}
}

// SetPhoneRequested sets value of PhoneRequested conditional field.
func (i *Invoice) SetPhoneRequested(value bool) {
	if value {
		i.Flags.Set(2)
		i.PhoneRequested = true
	} else {
		i.Flags.Unset(2)
		i.PhoneRequested = false
	}
}

// SetEmailRequested sets value of EmailRequested conditional field.
func (i *Invoice) SetEmailRequested(value bool) {
	if value {
		i.Flags.Set(3)
		i.EmailRequested = true
	} else {
		i.Flags.Unset(3)
		i.EmailRequested = false
	}
}

// SetShippingAddressRequested sets value of ShippingAddressRequested conditional field.
func (i *Invoice) SetShippingAddressRequested(value bool) {
	if value {
		i.Flags.Set(4)
		i.ShippingAddressRequested = true
	} else {
		i.Flags.Unset(4)
		i.ShippingAddressRequested = false
	}
}

// SetFlexible sets value of Flexible conditional field.
func (i *Invoice) SetFlexible(value bool) {
	if value {
		i.Flags.Set(5)
		i.Flexible = true
	} else {
		i.Flags.Unset(5)
		i.Flexible = false
	}
}

// SetPhoneToProvider sets value of PhoneToProvider conditional field.
func (i *Invoice) SetPhoneToProvider(value bool) {
	if value {
		i.Flags.Set(6)
		i.PhoneToProvider = true
	} else {
		i.Flags.Unset(6)
		i.PhoneToProvider = false
	}
}

// SetEmailToProvider sets value of EmailToProvider conditional field.
func (i *Invoice) SetEmailToProvider(value bool) {
	if value {
		i.Flags.Set(7)
		i.EmailToProvider = true
	} else {
		i.Flags.Unset(7)
		i.EmailToProvider = false
	}
}

// Decode implements bin.Decoder.
func (i *Invoice) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode invoice#c30aa358 to nil")
	}
	if err := b.ConsumeID(InvoiceTypeID); err != nil {
		return fmt.Errorf("unable to decode invoice#c30aa358: %w", err)
	}
	{
		if err := i.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode invoice#c30aa358: field flags: %w", err)
		}
	}
	i.Test = i.Flags.Has(0)
	i.NameRequested = i.Flags.Has(1)
	i.PhoneRequested = i.Flags.Has(2)
	i.EmailRequested = i.Flags.Has(3)
	i.ShippingAddressRequested = i.Flags.Has(4)
	i.Flexible = i.Flags.Has(5)
	i.PhoneToProvider = i.Flags.Has(6)
	i.EmailToProvider = i.Flags.Has(7)
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode invoice#c30aa358: field currency: %w", err)
		}
		i.Currency = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode invoice#c30aa358: field prices: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value LabeledPrice
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode invoice#c30aa358: field prices: %w", err)
			}
			i.Prices = append(i.Prices, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for Invoice.
var (
	_ bin.Encoder = &Invoice{}
	_ bin.Decoder = &Invoice{}
)
