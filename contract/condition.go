//
// contract/conditions.go
// gocontract
//
// Created by steve on 2019-03-02.
// Copyright © 2019 Klassen Software Solutions. All rights reserved.
// Permission is hereby granted for use under the MIT License (https://opensource.org/licenses/MIT).
//

package contract

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
)

// Preconditions checks the expression(s) given to see that they are all true. If any fail a suitable
// error message is logged and the program is halted if os.Exit(1).
//
// This method should typically be called near the start of a method, possibly right after the parameter
// checking. However, that is not a requirement. (In fact Preconditions, Conditions, and Postconditions
// all do exactly the same thing except for a slight difference in the message that is logged. The
// three methods are provided instead of one simply as a means of documenting their intent.)
func Preconditions(result bool, results ...bool) {
	performConditionChecks("Precondition", result, results)
}

// Conditions checks the expression(s) given to see that they are all true. If any fail a suitable
// error message is logged and the program is halted if os.Exit(1).
//
// This method should typically not be called near the start or end of a method, however that
// is not a requirement. (In fact Preconditions, Conditions, and Postconditions
// all do exactly the same thing except for a slight difference in the message that is logged. The
// three methods are provided instead of one simply as a means of documenting their intent.)
func Conditions(result bool, results ...bool) {
	performConditionChecks("Condition", result, results)
}

// Postconditions checks the expression(s) given to see that they are all true. If any fail a suitable
// error message is logged and the program is halted if os.Exit(1).
//
// This method should typically be called near the end of a method, possibly right after the parameter
// checking. However, that is not a requirement. (In fact Preconditions, Conditions, and Postconditions
// all do exactly the same thing except for a slight difference in the message that is logged. The
// three methods are provided instead of one simply as a means of documenting their intent.)
//
// Typically this is called to check the post-conditions of a successful execution of the function. As
// such it is generally not run when errors cause an early exit. However, if you wish it to run in
// such cases, you can add defer before the call.
func Postconditions(result bool, results ...bool) {
	performConditionChecks("Postcondition", result, results)
}

func performConditionChecks(conditionType string, result bool, results []bool) {
	count := 1 + len(results)

	if !result {
		reportFailureAndExit(conditionType, 1, count)
	}

	for idx := range results {
		if !results[idx] {
			reportFailureAndExit(conditionType, idx+2, count)
		}
	}
}

func generateFailureMessage(conditionType string, expressionNumber, expressionCount int) string {
	pc, _, _, ok := runtime.Caller(4)
	if !ok {
		log.Fatal("runtime.Caller failed, exiting program")
	}

	frames := runtime.CallersFrames([]uintptr{pc})
	frame, _ := frames.Next()

	return fmt.Sprintf("%v failed: expression %v of %v\n   in the method %v in the file %v, line %v",
		conditionType, expressionNumber, expressionCount,
		filepath.Base(frame.Function), filepath.Base(frame.File), frame.Line)
}

func reportFailureAndExit(conditionType string, expressionNumber, expressionCount int) {
	msg := generateFailureMessage(conditionType, expressionNumber, expressionCount)
	log.Fatalf("%v\n   Terminating!", msg)
}
