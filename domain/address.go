package domain

import "github.com/boreq/errors"

type Address struct {
	address string
}

func NewAddress(address string) (Address, error) {
	if address == "" {
		return Address{}, errors.New("address can not be empty")
	}

	return Address{
		address: address,
	}, nil
}

func MustNewAddress(address string) Address {
	v, err := NewAddress(address)
	if err != nil {
		panic(err)
	}
	return v
}

func (a Address) String() string {
	return a.address
}

func (a Address) IsZero() bool {
	return a == Address{}
}
