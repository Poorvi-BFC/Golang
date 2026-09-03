package main

import "fmt"

// ---- AccountHolder groups the three "personal" objects together ----

type AccountHolder struct {
	personal PersonalDetails
	fatca    FatcaDetails
	address  AddressDetails
}

// ---- Top-level identifier ----

type ClientAccount struct {
	ucc           string
	primaryHolder AccountHolder
	bankDetail    BankDetails
}

type PersonalDetails struct {
	fname    string
	lname    string
	age      int
	gender   string
	mobileNo string
	pan      string
}

type FatcaDetails struct {
	accountType    string
	maritalStatus  string
	residentStatus string
}

type AddressDetails struct {
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
		personal: PersonalDetails{
			fname:    "John",
			lname:    "Doe",
			age:      30,
			gender:   "Male",
			mobileNo: "1234567890",
			pan:      "ABCDE1234F",
		},
		fatca: FatcaDetails{
			accountType:    "Savings",
			maritalStatus:  "Single",
			residentStatus: "Resident",
		},
		address: AddressDetails{
			homeAddress:   "123 Main St",
			officeAddress: "456 Office Rd",
		},
	}

	bankDetail := BankDetails{
		accountNumber: "9876543210",
		ifscCode:      "IFSC0001",
		bankName:      "Bank of Go",
	}

	if primary.personal.pan == "" {
		//fmt.Println("PAN is required for the primary account holder.")
	} else {
		//fmt.Println("pan is:", primary.personal.pan)
	}
	if bankDetail.accountNumber == "" {
		fmt.Println("Bank account number is required.")
	} else {
		fmt.Println("Bank account number is:", bankDetail.accountNumber)
	}
	//fmt.Println("Gender:", primary.personal.gender)
}
