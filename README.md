# gocontract
Go tools to assist with "programming by contract"

## Prerequisites

All prerequisites will be handled automatically provide this library is checked out using the command:

    go get github.com/klassen-software-solutions/gocontract.git

## Description

This library contains items useful in coding based on 
[Design by Contract](https://en.wikipedia.org/wiki/Design_by_contract). Specifically it provides the ability
to add checks on preconditions, postconditions, middle conditions, and invariants in an expressive
manner.

[API Documentation](https://godoc.org/github.com/klassen-software-solutions/gocontract) 

### To terminate or not to terminate?

A common debate in programming by contract is what action should be taken when a contract
condition fails. We have taken the opinion that condition failures will write an error message
to the standard error device, and then exit the program via `os.Exit(1)`. (This is largely
following the advice of Ferguson, Schneier, and Kohno in their book _Cryptography Engineering_.
It was that book as well as Schneier's book _Click Here to Kill Everybody_ that were the primary
influencers in our decision to create this library.)

In addition we chose that the condition checks would always take place regardless of whether
your code is compiled in debug mode or production mode.

### Customizing the actions

At present we provide no facility for customizing what happens on a failed condition check. However,
if you really want to do it differently, feel free to take this code and modify it for your own 
purposes.

The key method that you need to change is `performConditionChecks` and is found in `condition.go`.
It's use is common to all the conditions so modifying that one method will affect them all.

### Using this library

#### Condition checks

All the condition types, `Preconditions`, `Conditions`,  and `Postconditions` work the same way. In fact
all three do the same thing except for slight differences in the message that is reported. The are 
presented as three methods only to help document the intent.

Specifically, the intent is that preconditions are checked at or near the start of a method, postconditions
at or near the end of a method, and conditions anywhere in between. Each of them takes any number of
expressions returning boolean values. If any of the expressions returns false, a suitable message
is produced on the standard error device, and the program is halted.

For example, the following could be used to check assumptions at the start of a given method.

    type MyStats struct {
        MinValueSeen, MaxValueSeen, AverageValueSeen int
        Count uint
    }
    
    func (s *MyStats) AddAValue(value int) {
        oldCount := s.Count
        contract.Preconditions(
            s.MinValueSeen <= s.AverageValueSeen,
            s.AverageValueSeen <= s.MaxValueSeen
        )
        
        // ... update the stats ...
        
        contract.Postconditions(
            s.MinValueSeen <= value,
            value <= s.MaxValueSeen,
            s.Count = oldCount + 1
        )
    }

You could add a `defer` in front of the postconditions call, but typically we don't on the assumption
that the postconditions are only true if no error occurs. Of course you could have two sets of 
post-conditions, one for a valid exit and one for an error exit, but we have never bothered to do
that in our own code.

An invariant is similar except that you don't give it an expression, you simply call the `Check` method
to see if the invariant has changed. _Note: Our implementation of an invariant does not actually check
for inequality. Rather its computes a hash based on the `%v` representation of the object and check if
that hash has changed. This is not technically a full check of the invariant, but it is generally good
enough to see if something has changed._

Very often an invariant is used to ensure that a method that should not change an object has indeed
not changed the object. This could look like the following:

    func (s *MyStats) PrintTheStats() {
        inv := contract.NewInvariant(s)
        defer inv.Check()
        
        // ... print out the representation of the stats ...
    }
   
Note the user of `defer` to ensure the variant is always checked. You could leave this out if you so
desire, but generally an invariant should be invariant regardless of how the method exits.

## Contributing

If you wish to make changes to this library that you believe will be useful to others, you can
contribute to the project. If you do, there are a number of policies you should follow:

* Check the issues to see if there are already similar bug reports or enhancements.
* Feel free to add bug reports and enhancements at any time.
* Don't work on things that are not part of an approved project.
* Don't create projects - they are only created the by owner.
* Projects are created based on conversations on the wiki.
* Feel free to initiate or join conversations on the wiki.
* Follow our [Go Coding Standards](https://www.kss.cc/standards/go.html).
