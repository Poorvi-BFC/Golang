package main

import "fmt"

type customer interface {
	GetBillingEmail() string
}
type biller[C customer] interface {
	Charge(customer C, amount float64) bill
	Name() string
}

// TODO: Define the biller[C customer] interface with Charge and Name methods

type user struct {
	UserEmail string
}

func (u user) GetBillingEmail() string {
	return u.UserEmail
}

type bill struct {
	Customer customer
	Amount   float64
}

type userBiller struct {
	planName string
}

func (ub userBiller) Charge(u user, amount float64) bill {
	return bill{
		Customer: u,
		Amount:   amount,
	}
}

func (ub userBiller) Name() string {
	return ub.planName
}

// don't touch below this line

func main() {
	ub := userBiller{planName: "Basic Plan"}
	u := user{UserEmail: "alex@example.com"}

	b := ub.Charge(u, 19.99)
	fmt.Printf("Billed %s ($%.2f) using %s\n", b.Customer.GetBillingEmail(), b.Amount, ub.Name())
}
