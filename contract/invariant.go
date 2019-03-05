//
// contract/invariant.go
// gocontract
//
// Created by steve on 2019-03-04.
// Copyright © 2019 Klassen Software Solutions. All rights reserved.
// Permission is hereby granted for use under the MIT License (https://opensource.org/licenses/MIT).
//

package contract

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"time"
)

// Invariant marks that a given item must not change within the method.
type Invariant struct {
	obj     interface{}
	nonce   time.Time
	objHash []byte
}

// NewInvariant constructs an Invariant wrapping a given object. Note that obj needs to be a pointer to
// a struct. This code does not work with types such as integers or strings as they end up being copied
// hence any changes are not seen.
func NewInvariant(obj interface{}) Invariant {
	inv := Invariant{obj: obj}
	inv.nonce = time.Now()
	inv.objHash = computeHash(&inv)
	return inv
}

// Checks that the invariant has not changed. If it has changed this will log an error message and halt
// the program via os.Exit(1). (I.e. It acts like a failed condition.)
func (inv *Invariant) Check() {
	h := computeHash(inv)
	performConditionChecks("Invariant", bytes.Equal(inv.objHash, h), []bool{})
}

func computeHash(inv *Invariant) []byte {
	// Note that the use of SHA1 in this case is acceptable as we do not require a cryptographically secure
	// hash, just a reasonably good one.
	h := sha1.New()
	s := fmt.Sprintf("%v", inv.obj)
	h.Write([]byte(inv.nonce.String()))
	h.Write([]byte(s))
	return h.Sum(nil)
}
