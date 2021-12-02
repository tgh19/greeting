package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type GreetTestCase struct {
	Name   string
	Inputs []string
	Exp    string
	Pass   bool
}

func TestGreet(test *testing.T) {
	cases := []GreetTestCase{
		{
			Name:   "R1: Basic single name",
			Inputs: []string{"Trey"},
			Exp:    "Hello, Trey.",
			Pass:   true,
		},
		{
			Name:   "R2: No name",
			Inputs: nil,
			Exp:    "Hello, my friend.",
			Pass:   true,
		},
		{
			Name:   "R3: Shout",
			Inputs: []string{"TREY"},
			Exp:    "HELLO TREY!",
			Pass:   true,
		},
		{
			Name:   "R4: Two names",
			Inputs: []string{"Trey", "Bob"},
			Exp:    "Hello, Trey and Bob.",
			Pass:   true,
		},
		{
			Name:   "R5: N names",
			Inputs: []string{"Trey", "Bob", "Charlie"},
			Exp:    "Hello, Trey, Bob, and Charlie.",
			Pass:   true,
		},
		{
			Name:   "R6: Mix shout with lowercase names",
			Inputs: []string{"Trey", "Bob", "CHARLIE"},
			Exp:    "Hello, Trey and Bob. AND HELLO CHARLIE!",
			Pass:   true,
		},
		{
			Name:   "R7: Handle comma separated strings",
			Inputs: []string{"Trey", "Bob", "CHARLIE,George"},
			Exp:    "Hello, Trey, Bob, and George. AND HELLO CHARLIE!",
			Pass:   true,
		},
		{
			Name:   "R8: Ignore quotes around commas",
			Inputs: []string{"Trey", "Bob", "\"Charlie, George\""},
			Exp:    "Hello, Trey, Bob, and Charlie, George.",
			Pass:   true,
		},
	}
	for _, testCase := range cases {
		test.Run(testCase.Name, func(test *testing.T) {
			if testCase.Pass {
				assert.Equal(test, testCase.Exp, Greet(testCase.Inputs))
			} else {
				assert.NotEqual(test, testCase.Exp, Greet(testCase.Inputs))
			}
		})
	}
}

type IsUpperTestCase struct {
	Name   string
	Inputs string
	Exp    bool
	Pass   bool
}

func TestIsUpper(test *testing.T) {
	cases := []IsUpperTestCase{
		{Name: "Should Pass", Inputs: "BILL", Exp: true, Pass: true},
		{Name: "Non letter", Inputs: "B!LL", Exp: false, Pass: true},
		{Name: "Lowercase", Inputs: "Bill", Exp: false, Pass: true},
	}
	for _, testCase := range cases {
		test.Run(testCase.Name, func(test *testing.T) {
			if testCase.Pass {
				assert.Equal(test, testCase.Exp, IsUpper(testCase.Inputs))
			} else {
				assert.NotEqual(test, testCase.Exp, IsUpper(testCase.Inputs))
			}
		})
	}
}
