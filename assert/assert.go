package assert

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

type argumentError struct {
	message string
}

func (err argumentError) Error() string {
	return err.message
}

// IsArgumentError verify if supplied error is an argument validation error.
func IsArgumentError(err error) bool {
	_, ok := err.(argumentError)
	return ok
}

type stateError struct {
	message string
}

func (err stateError) Error() string {
	return err.message
}

// IsStateError verify if the supplied error is a state validation error.
func IsStateError(err error) bool {
	_, ok := err.(stateError)
	return ok
}

// Equaler is an interface that can be implemented by objects that can be compared using specific method.
type Equaler interface {
	// Equal will check if supplied value is equal to receiver, returning true or false.
	Equal(interface{}) bool
}

func equal(obj1, obj2 interface{}) bool {
	if eql, ok := obj1.(Equaler); ok {
		return eql.Equal(obj2)
	}
	return reflect.DeepEqual(obj1, obj2)
}

// Condition will assert that supplied condition is true, returning an argument validation error with
// supplied message if the condition is false. If condition is met, nil is returned.
func Condition(condition bool, message string) error {
	if !condition {
		return argumentError{message}
	}
	return nil
}

// Zeroable is the interface that must be implemented by object that can be seto to zero (e.g. time.Time).
type Zeroable interface {
	IsZero() bool
}

// IsZero verify that supplied Zeroable value is zero (Zeroable.IsZero() must return true) returning
// an error with the format '<argName> should be zero' if condition is not met.
func IsZero(zeroable Zeroable, argName string) error {
	if !zeroable.IsZero() {
		return argumentError{fmt.Sprintf("%s should be zero", argName)}
	}
	return nil
}

// NotZero verify that supplied Zeroable value is not zero (Zeroable.IsZero() must return false), returning
// an error with the format '<argName> should not be zero' if condition is not met.
func NotZero(zeroable Zeroable, argName string) error {
	if zeroable.IsZero() {
		return argumentError{fmt.Sprintf("%s should not be zero", argName)}
	}
	return nil
}

// Validatable is an interface that must be implemented by objects that self validate themselves.
type Validatable interface {
	// Check that the receiver is valid, returning true or false.
	IsValid() bool
}

// IsValid verify that supplied Validatable value is valid (Validatable.IsValid() must return true), returning
// an error with the format '<argName> is not valid' if condition is not met.
func IsValid(validatable Validatable, argName string) error {
	if !validatable.IsValid() {
		return argumentError{fmt.Sprintf("%s is not valid", argName)}
	}
	return nil
}

// Equals verify that two supplied values are equals, returning an error in the format '<argName> must be equal to
// <obj2>' if they are not.
func Equals(obj1, obj2 interface{}, argName string) error {
	if !equal(obj1, obj2) {
		return argumentError{fmt.Sprintf("%s must be equal to %v", argName, obj2)}
	}
	return nil
}

// NotEquals verify that two supplied values are notequals, returning an error in the format '<argName> must not be
// equal to <obj2>' if they are not.
func NotEquals(obj1, obj2 interface{}, argNmae string) error {
	if equal(obj1, obj2) {
		return argumentError{fmt.Sprintf("%s must not be equal to %v", argNmae, obj2)}
	}
	return nil
}

// IsFalse verify that supplied boolean value is false, returning an error in the format '<argName> must be false'
// if it's not
func IsFalse(value bool, argName string) error {
	if value {
		return argumentError{fmt.Sprintf("%s must be false", argName)}
	}
	return nil
}

// IsTrue verify that supplied boolean value is true, returning an error in the format '<argName> must be true'
// if it's not
func IsTrue(value bool, argName string) error {
	if !value {
		return argumentError{fmt.Sprintf("%s must be true", argName)}
	}
	return nil
}

// MinLength verify that supplied string length is not lesser than provided length, returning an error in the format
// '<argName> must be <minLength> characters or less' if it's not.
func MinLength(value string, minLength int, argName string) error {
	length := len(strings.TrimSpace(value))
	if length < minLength {
		return argumentError{fmt.Sprintf("%s must be %d characters or more", argName, minLength)}
	}
	return nil
}

// MaxLength verify that supplied string length is not greater than provided length, returning an error in the format
// '<argName> must be <maxLength> characters or less' if it's not.
func MaxLength(value string, maxLength int, argName string) error {
	length := len(strings.TrimSpace(value))
	if length > maxLength {
		return argumentError{fmt.Sprintf("%s must be %d characters or less", argName, maxLength)}
	}
	return nil
}

// LengthBetween verify that supplied string length is between the provided min and max length, retuning an error in
// the format '<argName> must be between <minLength> and <maxLength> characters' if it's not.
func LengthBetween(value string, minLength, maxLength int, argName string) error {
	length := len(strings.TrimSpace(value))
	if length < minLength || length > maxLength {
		return argumentError{fmt.Sprintf("%s must be between %d and %d characters", argName, minLength, maxLength)}
	}
	return nil
}

// Empty verify that supplied string is empty or blank, returning an error with format '<argName> must be empty'
// if it's not.
func Empty(value, argName string) error {
	length := len(strings.TrimSpace(value))
	if length != 0 {
		return argumentError{fmt.Sprintf("%s must be empty", argName)}
	}
	return nil
}

// NotEmpty verify that supplied string is not empty or blank, returning an error with format '<argName> must not be
// empty' if it's not.
func NotEmpty(value, argName string) error {
	length := len(strings.TrimSpace(value))
	if length == 0 {
		return argumentError{fmt.Sprintf("%s must not be empty", argName)}
	}
	return nil
}

func isNil(a interface{}) bool {
	// See https://stackoverflow.com/questions/13476349/check-for-nil-and-nil-interface-in-go for details
	defer func() { recover() }()
	return a == nil || reflect.ValueOf(a).IsNil()
}

// Nil verify that supplied value is nil, returning an error with format '<argName> musts be nil' if it's not.
func Nil(obj interface{}, argName string) error {
	if !isNil(obj) {
		return argumentError{fmt.Sprintf("%s must be nil", argName)}
	}
	return nil
}

// NotNil verify that supplied value is not nil, returning an error with format '<argName> must not be nil' if it's not.
func NotNil(obj interface{}, argName string) error {
	if isNil(obj) {
		return argumentError{fmt.Sprintf("%s must not be nil", argName)}
	}
	return nil
}

// Matches verify that supplied value matches the supplied regular expression.
func Matches(value string, regexp *regexp.Regexp, argName string) error {
	if !regexp.MatchString(value) {
		return argumentError{fmt.Sprintf("%s does not match the pattern", argName)}
	}
	return nil
}

// NotMatches verify that supplied value does not matches the supplied regular expression.
func NotMatches(value string, regexp *regexp.Regexp, argName string) error {
	if regexp.MatchString(value) {
		return argumentError{fmt.Sprintf("%s does match the pattern", argName)}
	}
	return nil
}

// IntMin verify that supplied integer value is not below the supplied minimum value, returning an error with format
// '<argName> must be greater or equal than <min>' if it's not.
func IntMin(value, min int, argName string) error {
	if value < min {
		return argumentError{fmt.Sprintf("%s must greater or equal than %d", argName, min)}
	}
	return nil
}

// IntMax verify that supplied integer value is not above the supplied maximum value, returning an error with format
// '<argName> must be lower or equal than <min>' if it's not.
func IntMax(value, max int, argName string) error {
	if value > max {
		return argumentError{fmt.Sprintf("%s must lower or equal than %d", argName, max)}
	}
	return nil
}

// IntRange verify that supplied integer value is between minmum and maximum, returning an error with format '<argName>
// must be between <min> and <max>' if it's not.
func IntRange(value, min, max int, argName string) error {
	if value < min || value > max {
		return argumentError{fmt.Sprintf("%s must be between %d and %d", argName, min, max)}
	}
	return nil
}

// State will verify that supplied state is true, returning an error if it's not.
func State(state bool, message string) error {
	if !state {
		return stateError{message}
	}
	return nil
}

// StateNot will verify that supplied state is false, returning an error if it's not.
func StateNot(state bool, message string) error {
	if state {
		return stateError{message}
	}
	return nil
}
