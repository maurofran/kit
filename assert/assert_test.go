package assert_test

import (
	. "github.com/maurofran/assert"

	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/types"
	"time"
)

var _ = Describe("Assert", func() {
	Describe("Condition", func() {
		It("Should return error if condition is false", func() {
			Expect(Condition(false, "message")).To(BeArgumentError())
		})
		It("Should return nil if condition is true", func() {
			Expect(Condition(true, "message")).To(BeNil())
		})
	})

	Describe("IsZero", func() {
		It("Should return an error if supplied value is not zero", func() {
			Expect(IsZero(time.Now(), "argument")).To(BeArgumentError())
		})
		It("Should return nil if supplied value is zero", func() {
			Expect(IsZero(time.Time{}, "argument")).To(BeNil())
		})
	})

	Describe("NotZero", func() {
		It("Should return an error if supplied value is zero", func() {
			Expect(NotZero(time.Time{}, "argument")).To(BeArgumentError())
		})
		It("Should return nil if supplied value is not zero", func() {
			Expect(NotZero(time.Now(), "argument")).To(BeNil())
		})
	})

	Describe("IsValid", func() {
		It("Should return an error if supplied object is not valid", func() {
			Expect(IsValid(validatableMock{false}, "argument")).To(BeArgumentError())
		})
		It("Should return nil if supplied object is valid", func() {
			Expect(IsValid(validatableMock{true}, "argument")).To(BeNil())
		})
	})

	Describe("Equals", func() {
		It("Should return an error if supplied objects are not equals", func() {
			Expect(Equals("Foo", "Baz", "argument")).To(BeArgumentError())
		})
		It("Should return nil if supplied objects are equals", func() {
			Expect(Equals("Foo", "Foo", "argument")).To(BeNil())
		})
	})

	Describe("NotEquals", func() {
		It("Should return an error if supplied objects are t equals", func() {
			Expect(NotEquals("Foo", "Foo", "argument")).To(BeArgumentError())
		})
		It("Should return nil if supplied objects are not equals", func() {
			Expect(NotEquals("Foo", "Baz", "argument")).To(BeNil())
		})
	})

	Describe("IsFalse", func() {
		It("Should return an error if supplied value is true", func() {
			Expect(IsFalse(true, "argument")).To(BeArgumentError())
		})
		It("Should return nil if supplied value is false", func() {
			Expect(IsFalse(false, "argument")).To(BeNil())
		})
	})

	Describe("IsTrue", func() {
		It("Should return an error if supplied value is false", func() {
			Expect(IsTrue(false, "argument")).To(BeArgumentError())
		})
		It("Should return nil if supplied value is true", func() {
			Expect(IsTrue(true, "argument")).To(BeNil())
		})
	})

	Describe("MinLength", func() {
		It("Should return an error if supplied value shorter than minimum length", func() {
			Expect(MinLength("foo", 4, "argument")).To(BeArgumentError())
		})
		It("Should return nil if supplied value is equal to minimum length", func() {
			Expect(MinLength("foob", 4, "argument")).To(BeNil())
		})
		It("Should return nil if supplied value is longer than minimum length", func() {
			Expect(MinLength("foobaz", 4, "argument")).To(BeNil())
		})
	})

	Describe("MaxLength", func() {
		It("Should return an error if supplied value longer than maximum length", func() {
			Expect(MaxLength("foobaz", 4, "argument")).To(BeArgumentError())
		})
		It("Should return nil if supplied value is equal to maximum length", func() {
			Expect(MaxLength("foob", 4, "argument")).To(BeNil())
		})
		It("Should return nil if supplied value is shorter than maximum length", func() {
			Expect(MaxLength("foo", 4, "argument")).To(BeNil())
		})
	})

	Describe("LengthBetween", func() {
		It("Should return an error if supplied value longer than maximum length", func() {
			Expect(LengthBetween("foobaz", 2, 4, "argument")).To(BeArgumentError())
		})
		It("Should return an error if supplied value shorter than minium length", func() {
			Expect(LengthBetween("f", 2, 4, "argument")).To(BeArgumentError())
		})
		It("Should return nil if supplied value is equal to maximum length", func() {
			Expect(LengthBetween("fo", 2, 4, "argument")).To(BeNil())
		})
		It("Should return nil if supplied value is equal to minium length", func() {
			Expect(LengthBetween("foob", 2, 4, "argument")).To(BeNil())
		})
		It("Should return nil if supplied value is between minium and maximum length", func() {
			Expect(LengthBetween("foo", 2, 4, "argument")).To(BeNil())
		})
	})

	Describe("Empty", func() {
		It("Should return an error if supplied value is not empty", func() {
			Expect(Empty("foo", "argument")).To(BeArgumentError())
		})
		It("Should return nil if supplied value is empty", func() {
			Expect(Empty("", "argument")).To(BeNil())
		})
		It("Should return nil if supplied value is blank", func() {
			Expect(Empty("   ", "argument")).To(BeNil())
		})
	})

	Describe("NotEmpty", func() {
		It("Should return an error if supplied value is empty", func() {
			Expect(NotEmpty("", "argument")).To(BeArgumentError())
		})
		It("Should return an error if supplied value is blank", func() {
			Expect(NotEmpty("  ", "argument")).To(BeArgumentError())
		})
		It("Should return nil if supplied value is not empty", func() {
			Expect(NotEmpty("foo", "argument")).To(BeNil())
		})
	})

	Describe("Nil", func() {
		It("Should return an error if supplied value is not nil", func() {
			obj := ""
			Expect(Nil(obj, "argument")).To(BeArgumentError())
		})
		It("Should return an error if supplied value is not nil", func() {
			var obj interface{}
			obj = "foo"
			Expect(Nil(obj, "argument")).To(BeArgumentError())
		})
		It("Should return nil if supplied value is nil", func() {
			Expect(Nil(nil, "argument")).To(BeNil())
		})
		It("Should return nil if supplied value is not empty", func() {
			var obj interface{}
			Expect(Nil(obj, "argument")).To(BeNil())
		})
	})

	Describe("NotNil", func() {
		It("Should return an error if supplied value is nil", func() {
			Expect(NotNil(nil, "argument")).To(BeArgumentError())
		})
		It("Should return an error if supplied value is not empty", func() {
			var obj interface{}
			Expect(NotNil(obj, "argument")).To(BeArgumentError())
		})
		It("Should return nil if supplied value is not nil", func() {
			obj := ""
			Expect(NotNil(obj, "argument")).To(BeNil())
		})
		It("Should return nil if supplied value is not nil", func() {
			var obj interface{}
			obj = time.Now() //"foo"
			Expect(NotNil(&obj, "argument")).To(BeNil())
		})
	})

	Describe("IntMin", func() {
		It("Should return an error if supplied value is lesser than minimum", func() {
			Expect(IntMin(3, 4, "argument")).To(BeArgumentError())
		})
		It("Should return nil if supplied value is equal to minimum", func() {
			Expect(IntMin(4, 4, "argument")).To(BeNil())
		})
		It("Should return nil if supplied value is greater than minimum", func() {
			Expect(IntMin(7, 4, "argument")).To(BeNil())
		})
	})

	Describe("IntMax", func() {
		It("Should return an error if supplied value is lesser than minimum", func() {
			Expect(IntMin(3, 4, "argument")).To(BeArgumentError())
		})
		It("Should return nil if supplied value is equal to minimum", func() {
			Expect(IntMin(4, 4, "argument")).To(BeNil())
		})
		It("Should return nil if supplied value is greater than minimum", func() {
			Expect(IntMin(7, 4, "argument")).To(BeNil())
		})
	})

	Describe("IntRange", func() {
		It("Should return an error if supplied value is greater than maximum", func() {
			Expect(IntRange(5, 2, 4, "argument")).To(BeArgumentError())
		})
		It("Should return an error if supplied value is lesser than minium", func() {
			Expect(IntRange(1, 2, 4, "argument")).To(BeArgumentError())
		})
		It("Should return nil if supplied value is equal to maximum", func() {
			Expect(IntRange(4, 2, 4, "argument")).To(BeNil())
		})
		It("Should return nil if supplied value is equal to minium", func() {
			Expect(IntRange(2, 2, 4, "argument")).To(BeNil())
		})
		It("Should return nil if supplied value is between minium and maximum", func() {
			Expect(IntRange(3, 2, 4, "argument")).To(BeNil())
		})
	})

	Describe("State", func() {
		It("Should return an error if supplied value is false", func() {
			Expect(State(false, "message")).To(BeStateError())
		})
		It("Should return nil if supplied value is false", func() {
			Expect(State(true, "message")).To(BeNil())
		})
	})

	Describe("StateNot", func() {
		It("Should return an error if supplied value is false", func() {
			Expect(StateNot(true, "message")).To(BeStateError())
		})
		It("Should return nil if supplied value is true", func() {
			Expect(StateNot(false, "message")).To(BeNil())
		})
	})
})

// Mock types
type validatableMock struct {
	valid bool
}

func (v validatableMock) IsValid() bool {
	return v.valid
}

// Custom matchers.

func BeArgumentError() GomegaMatcher {
	return &beArgumentErrorMatcher{}
}

type beArgumentErrorMatcher struct {
}

func (m *beArgumentErrorMatcher) Match(actual interface{}) (bool, error) {
	err, ok := actual.(error)
	return ok && IsArgumentError(err), nil
}

func (m *beArgumentErrorMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n\t%#v\nto be an argumentError", actual)
}

func (m *beArgumentErrorMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n\t%#v\nto not be an argumentError", actual)
}

func BeStateError() GomegaMatcher {
	return &beStateErrorMatcher{}
}

type beStateErrorMatcher struct {
}

func (m *beStateErrorMatcher) Match(actual interface{}) (bool, error) {
	err, ok := actual.(error)
	return ok && IsStateError(err), nil
}

func (m *beStateErrorMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n\t%#v\nto be an stateError", actual)
}

func (m *beStateErrorMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n\t%#v\nto not be an stateError", actual)
}
