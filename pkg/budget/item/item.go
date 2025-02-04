package item

import (
	"fmt"
	"mediumstreams/pkg/utils"
)

type BudgetItem struct {
	Id        int
	Sequence  int
	Name      string
	Quantity  float64
	UnitValue float64
	Children  []*BudgetItem
}

var id = 0
var possibleNames = []string{"Chair", "Ink", "Notebook", "Monitor", "Headphone", "Coffe Mug", "Pencil"}

func GenerateRandomBudgetItemLevel() *BudgetItem {
	item := BudgetItem{
		Id:        id,
		Sequence:  id,
		Name:      fmt.Sprintf("Level %v", id),
		Quantity:  1,
		UnitValue: 0,
		Children:  []*BudgetItem{},
	}

	id += 1
	return &item
}

func GenerateRandomBudgetItems(items int) []*BudgetItem {
	budgetItems := []*BudgetItem{}
	for i := range items {
		budgetItems = append(budgetItems, generateBudgetItem(i))
	}

	return budgetItems
}

func generateBudgetItem(sequence int) *BudgetItem {
	item := BudgetItem{
		Id:        id,
		Sequence:  sequence,
		Name:      possibleNames[utils.RandomSliceIndex(len(possibleNames))],
		Quantity:  utils.GenerateRandomFloat(100, 2),
		UnitValue: utils.GenerateRandomFloat(100, 2),
		Children:  []*BudgetItem{},
	}

	id += 1
	return &item
}
