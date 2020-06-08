// Package pgstore contains the types for schema 'public'.
package pgstore

import (
	"database/sql/driver"
	"errors"
)

// Code generated by xo. DO NOT EDIT.
// Userstatus is the 'userstatus' enum type from schema 'public'.
type Userstatus uint16

const (
	// UserstatusHealthy is the 'HEALTHY' Userstatus.
	UserstatusHealthy = Userstatus(1)

	// UserstatusPositive is the 'POSITIVE' Userstatus.
	UserstatusPositive = Userstatus(2)

	// UserstatusRecovered is the 'RECOVERED' Userstatus.
	UserstatusRecovered = Userstatus(3)
)

// String returns the string value of the Userstatus.
func (u Userstatus) String() string {
	var enumVal string
	switch u {
	case UserstatusHealthy:
		enumVal = "HEALTHY"

	case UserstatusPositive:
		enumVal = "POSITIVE"

	case UserstatusRecovered:
		enumVal = "RECOVERED"
	}

	return enumVal
}

// MarshalText marshals Userstatus into text.
func (u Userstatus) MarshalText() ([]byte, error) {
	return []byte(u.String()), nil
}

// UnmarshalText unmarshals Userstatus from text.
func (u *Userstatus) UnmarshalText(text []byte) error {
	switch string(text) {
	case "HEALTHY":
		*u = UserstatusHealthy

	case "POSITIVE":
		*u = UserstatusPositive

	case "RECOVERED":
		*u = UserstatusRecovered

	default:
		return errors.New("invalid Userstatus")
	}

	return nil
}

// Value satisfies the sql/driver.Valuer interface for Userstatus.
func (u Userstatus) Value() (driver.Value, error) {
	return u.String(), nil
}

// Scan satisfies the database/sql.Scanner interface for Userstatus.
func (u *Userstatus) Scan(src interface{}) error {
	buf, ok := src.([]byte)
	if !ok {
		return errors.New("invalid Userstatus")
	}

	return u.UnmarshalText(buf)
}