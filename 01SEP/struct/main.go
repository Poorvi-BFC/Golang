package main

import "fmt"

// ---- AccountHolder groups the three "personal" objects together ----

type AccountHolder struct {
	personal personalDetails
	fatca    fatcaDetails
	address  addressDetails
}

// ---- Top-level identifier ----

type ClientAccount struct {
	ucc           string
	primaryHolder AccountHolder
	bankDetail    BankDetails
}

type personalDetails struct {
	fname    string
	lname    string
	age      int
	gender   string
	mobileNo string
	pan      string
}

type fatcaDetails struct {
	accountType    string
	maritalStatus  string
	residentStatus string
}

type addressDetails struct {
	homeAddress   string
	officeAddress string
}

type BankDetails struct {
	accountNumber string
	ifscCode      string
	bankName      string
}

func main() {
	primary := AccountHolder{
		personal: personalDetails{
			fname:    "John",
			lname:    "Doe",
			age:      30,
			gender:   "Male",
			mobileNo: "1234567890",
			pan:      "ABCDE1234F",
		},
		fatca: fatcaDetails{
			accountType:    "Savings",
			maritalStatus:  "Single",
			residentStatus: "Resident",
		},
		address: addressDetails{
			homeAddress:   "123 Main St",
			officeAddress: "456 Office Rd",
		},
	}

	// bankDetail := BankDetails{
	// 	accountNumber: "9876543210",
	// 	ifscCode:      "IFSC0001",
	// 	bankName:      "Bank of Go",
	// }

	if primary.personal.pan == "" {
		fmt.Println("PAN is required for the primary account holder.")
	} else {
		fmt.Println("pan is:", primary.personal.pan)
	}
	fmt.Println("Gender:", primary.personal.gender)
}
