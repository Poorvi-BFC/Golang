package main

import (
	"errors"
	"fmt"
)

type lineItem interface {
	GetCost() float64
}

type subscription struct {
	monthlyCost float64
}

func (s subscription) GetCost() float64 {
	return s.monthlyCost
}

func chargeForLineItem[T lineItem](newItem T, oldItems []T, balance float64) ([]T, float64, error) {
	// TODO: Calculate newBalance
	// TODO: If newBalance < 0.0, return nil, 0.0, errors.New("insufficient funds")
	// TODO: Append newItem to oldItems and return updated slice, newBalance, nil
	if newBalance := balance - newItem.GetCost(); newBalance < 0.0 {
		return nil, 0.0, errors.New("insufficient funds")
	}
	return append(oldItems, newItem), balance - newItem.GetCost(), nil
}

// don't touch below this line

func main() {
	sub := subscription{monthlyCost: 25.0}
	history := []subscription{}

	updatedHistory, newBal, err := chargeForLineItem(sub, history, 50.0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Success! New Balance: %.2f, Items count: %d\n", newBal, len(updatedHistory))
	}
}
