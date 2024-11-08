package main

// error-handling-challenge
//
// authors : rm4n0s, rept0id <rad@simplecode.gr>
//
// Notes :
//  - Note that the executionTrace stores the function names in the reverse order.
//  - The final check could had been simplified. rm4n0s provide feedback if possible.
//  - executionTrace update could had been a function. rm4n0s provide feedback if possible.

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"runtime"
)

var ErrBankAccountEmpty = errors.New("account-is-empty")
var ErrInvestmentLost = errors.New("investment-lost")

func f1(_executionTrace *[]string) error {
	// _executionTrace
	// could had been simplified into a function - wanted to follow rules
	defer (func() {
		pc, _, _, _ := runtime.Caller(1)
		fn := runtime.FuncForPC(pc).Name()

		*_executionTrace = append(*_executionTrace, fn)
	})()

	n := rand.IntN(9) + 1

	if n%2 == 0 {
		return ErrBankAccountEmpty
	}

	return ErrInvestmentLost
}

func f2(_executionTrace *[]string) error {
	// _executionTrace
	// could had been simplified into a function - wanted to follow rules
	defer (func() {
		pc, _, _, _ := runtime.Caller(1)
		fn := runtime.FuncForPC(pc).Name()

		*_executionTrace = append(*_executionTrace, fn)
	})()

	return f1(_executionTrace)
}

func f3(_executionTrace *[]string) error {
	// _executionTrace
	// could had been simplified into a function - wanted to follow rules
	defer (func() {
		pc, _, _, _ := runtime.Caller(1)
		fn := runtime.FuncForPC(pc).Name()

		*_executionTrace = append(*_executionTrace, fn)
	})()

	return f1(_executionTrace)
}

func f4(_executionTrace *[]string) error {
	// _executionTrace
	// could had been simplified into a function - wanted to follow rules
	defer (func() {
		pc, _, _, _ := runtime.Caller(1)
		fn := runtime.FuncForPC(pc).Name()

		*_executionTrace = append(*_executionTrace, fn)
	})()

	n := rand.IntN(9) + 1

	if n%2 == 0 {
		return f2(_executionTrace)
	}

	return f3(_executionTrace)
}

func main() {
	var executionTrace []string

	err := f4(&executionTrace)

	/*
	 * print three different messages based on
	 * the execution path of the functions and error:
	 * - for f4()->
	 *			f2()->
	 *				f1()->
	 *					 ErrBankAccountEmpty
	 *					  print "Aand it's gone"
	 * - for f4()->
	 *			f3()->
	 *				f1()->
	 *					ErrInvestmentLost
	 *					 print "The money in your account didn't do well"
	 *
	 * - for the rest of the cases
	 *			print "This line is for bank members only"

	 * also print any type of stack trace for err
	 */
	if err != nil {
		fmt.Println("Print stack trace", err)
		fmt.Println(executionTrace)

		// Only executionTrace[1] changes here but made check strict because footguns are not cool
		/*
		 * Note that the executionTrace stores the function names in the reverse order, e.g.:
		 * 	f4 -> f3 -> f1 => [main.f1 main.f3 main.f4]
		 */
		if len(executionTrace) == 3 {
			if executionTrace[0] == "main.f1" &&
				executionTrace[1] == "main.f3" &&
				executionTrace[2] == "main.f4" {

				fmt.Println("The money in your account didn't do well")
			} else if executionTrace[0] == "main.f1" &&
				executionTrace[1] == "main.f2" &&
				executionTrace[2] == "main.f4" {

				fmt.Println("Aand it's gone")
			} else {
				fmt.Println("This line is for bank members only")
			}
		} else {
			fmt.Println("Unpredicted executionTrace")
		}
	}
}
