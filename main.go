package main

import (
	"poc-mapper/mapper"
	"poc-mapper/mapper/model"
)

func main() {
	trxMaster := model.Transaction{
		Operation: "purchase",
		Card: model.Card{
			Brand:  "master",
			Method: "credit",
		},
		Country: "BRA",
		Amount:  100,
	}

	trxVisa := model.Transaction{
		Operation: "purchase",
		Card: model.Card{
			Brand:  "visa",
			Method: "credit",
		},
		Country: "BRA",
		Amount:  100,
	}

	statusMapper := mapper.NewMapper()

	// CUSTOMIZATION FOR - VISA CREDIT IN BRASIL (Priority 1)
	rules := []mapper.KeyRule{
		mapper.WithBrand("visa"),
		mapper.WithMethod("credit"),
		mapper.WithCountry("BRA"),
	}
	statusMapper.Add(map[string]mapper.Result{
		"B1": mapper.CallForAuth, // NOT SUPPORTED IN BRASIL
	}, rules...)

	// CUSTOMIZATION FOR - BRASIL (Priority 2)
	rules = []mapper.KeyRule{
		mapper.WithCountry("BRA"),
	}
	statusMapper.Add(map[string]mapper.Result{
		"B1": mapper.Approved,
	}, rules...)

	// CUSTOMIZATION FOR - MASTER CREDIT (Priority 3)
	rules = []mapper.KeyRule{
		mapper.WithBrand("master"),
		mapper.WithMethod("credit"),
	}
	statusMapper.Add(map[string]mapper.Result{
		"M1": mapper.Approved,
	}, rules...)

	// default mapper (fallback)
	statusMapper.Add(map[string]mapper.Result{
		"00": mapper.Approved,
		"01": mapper.Rejected,
		"02": mapper.RejectedOtherReason,
		"03": mapper.RejectedByProvider,
		"04": mapper.Contingency,
	})

	println("DEFAULT = ", statusMapper.Get(trxMaster, "00"))          // EXPECT APPROVED by default mapper (Priority 4)
	println("MASTER CREDIT = ", statusMapper.Get(trxMaster, "M1"))    // EXPECT APPROVED by MASTER CREDIT (Priority 3)
	println("BRASIL = ", statusMapper.Get(trxMaster, "B1"))           // EXPECT APPROVED by BRASIL (Priority 2)
	println("BRASIL VISA CREDIT = ", statusMapper.Get(trxVisa, "B1")) // EXPECT CALL_FOR_AUTH by VISA CREDIT IN BRASIL (Priority 1), same code before, buy with high priority
	println("INVALID = ", statusMapper.Get(trxVisa, "404"))           // EXPECT NOT_MAPPED
}
