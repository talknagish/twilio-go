package twilio

import (
	"context"
	"net/url"
)

const addressPathPart = "Addresses"

type AddressService struct {
	client *Client
}

type Address struct {
	AccountSid       string `json:"account_sid"`
	City             string `json:"city"`
	CustomerName     string `json:"customer_name"`
	DateCreated      string `json:"date_created"`
	DateUpdated      string `json:"date_updated"`
	EmergencyEnabled bool   `json:"emergency_enabled"`
	FriendlyName     string `json:"friendly_name"`
	IsoCountry       string `json:"iso_country"`
	PostalCode       string `json:"postal_code"`
	Region           string `json:"region"`
	Sid              string `json:"sid"`
	Street           string `json:"street"`
	Validated        bool   `json:"validated"`
	Verified         bool   `json:"verified"`
	Uri              string `json:"uri"`
}

// Create a new address
func (as *AddressService) Create(ctx context.Context, data url.Values) (*Address, error) {
	address := new(Address)
	pathPart := addressPathPart

	err := as.client.CreateResource(ctx, pathPart, data, address)

	return address, err
}

// Update an existing address
func (as *AddressService) Update(ctx context.Context, sid string, data url.Values) (*Address, error) {
	address := new(Address)
	pathPart := addressPathPart

	err := as.client.UpdateResource(ctx, pathPart, sid, data, address)

	return address, err
}

// Delete an address
func (as *AddressService) Delete(ctx context.Context, sid string) error {
	return as.client.DeleteResource(ctx, addressPathPart, sid)
}
