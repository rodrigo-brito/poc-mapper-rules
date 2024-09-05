package mapper

import (
	"poc-mapper/mapper/kind"
)

type Key struct {
	Operation string
	Brand     string
	Method    string
	Country   string
}

type KeyRule struct {
	Key  string
	Kind kind.Kind
}

func WithBrand(brand string) KeyRule {
	return KeyRule{
		Key:  brand,
		Kind: kind.Brand(),
	}
}

func WithMethod(method string) KeyRule {
	return KeyRule{
		Key:  method,
		Kind: kind.Method(),
	}
}

func WithCountry(country string) KeyRule {
	return KeyRule{
		Key:  country,
		Kind: kind.Country(),
	}
}
