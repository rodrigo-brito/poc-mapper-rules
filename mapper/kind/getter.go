package kind

import (
	"poc-mapper/mapper/model"
)

type Kind interface {
	Name() string
	Key(model.Transaction) string
}

type brand struct{}

func Brand() Kind {
	return brand{}
}

func (o brand) Name() string {
	return "brand"
}

func (o brand) Key(t model.Transaction) string {
	return t.Card.Brand
}

type method struct{}

func Method() Kind {
	return method{}
}

func (o method) Name() string {
	return "method"
}

func (o method) Key(t model.Transaction) string {
	return t.Card.Method
}

type country struct{}

func Country() Kind {
	return country{}
}

func (o country) Name() string {
	return "country"
}

func (o country) Key(t model.Transaction) string {
	return t.Country
}
