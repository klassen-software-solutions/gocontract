//
// contract/conditions_test.go
// gocontract
//
// Created by steve on 2019-03-02.
// Copyright © 2019 Klassen Software Solutions. All rights reserved.
// Permission is hereby granted for use under the MIT License (https://opensource.org/licenses/MIT).
//

package contract

import "testing"

// Note that the following tests are little more than syntax checks as actually testing a
// failure would cause the program to exit.

func TestPreconditions(t *testing.T) {
	Preconditions(true)

	Preconditions(true, true)

	Preconditions(true,
		myFn(),
		func() bool { return true }(),
		func() bool { return myFn() }())
}

func TestConditions(t *testing.T) {
	Conditions(true)

	Conditions(true, true)

	Conditions(true,
		myFn(),
		func() bool { return true }(),
		func() bool { return myFn() }())
}

func TestPostconditions(t *testing.T) {
	Postconditions(true)

	Postconditions(true, true)

	Postconditions(true,
		myFn(),
		func() bool { return true }(),
		func() bool { return myFn() }())
}

func myFn() bool {
	return true
}
