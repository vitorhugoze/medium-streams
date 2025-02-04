package calculations

import (
	"mediumstreams/pkg/budget/item"

	"github.com/jucardi/go-streams/v2/streams"
)

func TotalizeBudgetLevels(budget []*item.BudgetItem) {
	addBudgetItemsValue(budget)
}

func TotalizeBudgetLevelsAsync(budget []*item.BudgetItem) {
	addBudgetItemsValueAsync(budget)
}

func addBudgetItemsValue(budget []*item.BudgetItem) float64 {
	total := 0.0

	streams.FromArray(budget).ForEach(totalizeBudgetItems(&total))
	return total
}

func totalizeBudgetItems(total *float64) func(bi *item.BudgetItem) {
	return func(bi *item.BudgetItem) {
		if len(bi.Children) > 0 {
			levelTotal := addBudgetItemsValue(bi.Children)
			bi.UnitValue = levelTotal

			*total += levelTotal
		} else {
			*total += bi.Quantity * bi.UnitValue
		}
	}
}

func addBudgetItemsValueAsync(budget []*item.BudgetItem, txChan ...chan<- float64) {
	total := 0.0

	rxChan := make(chan float64)
	defer close(rxChan)

	streams.FromArray(budget).ParallelForEach(totalizeBudgetItemsAsync(rxChan), len(budget), true)
	for range len(budget) {
		total += <-rxChan
	}

	if len(txChan) > 0 {
		txChan[0] <- total
	}
}

func totalizeBudgetItemsAsync(totalChan chan float64) func(bi *item.BudgetItem) {
	return func(bi *item.BudgetItem) {
		if len(bi.Children) > 0 {
			go addBudgetItemsValueAsync(bi.Children, totalChan)

			levelTotal := <-totalChan
			bi.UnitValue = levelTotal

			totalChan <- levelTotal
		} else {
			totalChan <- bi.Quantity * bi.UnitValue
		}
	}
}
