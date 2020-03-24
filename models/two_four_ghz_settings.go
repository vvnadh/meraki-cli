// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TwoFourGhzSettings TwoFourGhzSettings
//
// Settings related to 2.4Ghz band
//
// swagger:model TwoFourGhzSettings
type TwoFourGhzSettings struct {

	// Determines whether ax radio on 2.4Ghz band is on or off. Can be either true or false. If false, we highly recommend disabling band steering. Defaults to true.
	AxEnabled bool `json:"axEnabled,omitempty"`

	// Sets max power (dBm) of 2.4Ghz band. Can be integer between 5 and 30. Defaults to 30.
	MaxPower int32 `json:"maxPower,omitempty"`

	// Sets min bitrate (Mbps) of 2.4Ghz band. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'. Defaults to 11.
	MinBitrate float64 `json:"minBitrate,omitempty"`

	// Sets min power (dBm) of 2.4Ghz band. Can be integer between 5 and 30. Defaults to 5.
	MinPower int32 `json:"minPower,omitempty"`

	// The RX-SOP level controls the sensitivity of the radio. It is strongly recommended to use RX-SOP only after
	//     consulting a wireless expert. RX-SOP can be configured in the range of -65 to -95 (dBm). A value of null will
	//     reset this to the default.
	Rxsop int32 `json:"rxsop,omitempty"`

	// Sets valid auto channels for 2.4Ghz band. Can be one of '1', '6' or '11'. Defaults to [1, 6, 11].
	ValidAutoChannels []int32 `json:"validAutoChannels"`
}

// Validate validates this two four ghz settings
func (m *TwoFourGhzSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TwoFourGhzSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TwoFourGhzSettings) UnmarshalBinary(b []byte) error {
	var res TwoFourGhzSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}