package budget

import (
	"mediumstreams/pkg/budget/item"
)

func GenerateBudget(levels, itemsPerLevel, nestingDepth int) []*item.BudgetItem {
	budget := []*item.BudgetItem{}

	for range levels {
		budgetLevel := item.GenerateRandomBudgetItemLevel()

		if nestingDepth > 1 {
			budgetLevel.Children = GenerateBudget(levels, itemsPerLevel, nestingDepth-1)
		}
		if nestingDepth == 1 {
			budgetLevel.Children = item.GenerateRandomBudgetItems(itemsPerLevel)
		}

		budget = append(budget, budgetLevel)
	}

	return budget
}
