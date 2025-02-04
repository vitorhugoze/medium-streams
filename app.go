package main

import (
	"mediumstreams/pkg/budget"
	"mediumstreams/pkg/budget/calculations"
)

func app() {
	b := budget.GenerateBudget(5, 5, 2)

	calculations.TotalizeBudgetLevels(b)
}
