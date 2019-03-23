//
// contract/invariant_test.go
// gocontract
//
// Created by steve on 2019-03-04.
// Copyright © 2019 Klassen Software Solutions. All rights reserved.
// Permission is hereby granted for use under the MIT License (https://opensource.org/licenses/MIT).
//

package contract

import "testing"

// Note that the following tests are little more than syntax checks as actually testing a failure
// would cause the program to exit.

func TestInvariantWithoutDefer(t *testing.T) {
	o := MyObject{1, "one"}
	inv := NewInvariant(&o)
	inv.Check()
}

func TestInvariantWithDefer(t *testing.T) {
	o := MyObject{1, "one"}
	inv := NewInvariant(&o)
	defer inv.Check()

	o.IValue++
}

type MyObject struct {
	IValue int
	SValue string
}
