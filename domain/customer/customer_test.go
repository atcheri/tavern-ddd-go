package customer_test

import (
	"testing"

	"github.com/atcheri/tavern-ddd-go/domain/customer"
	"github.com/iamkoch/ensure"
)

func TestCustomer_NewCustomer(t *testing.T) {
	type testCase struct {
		test string
		name string
		err  error
	}

	testCases := []testCase{
		{
			test: "Empty name validation",
			name: "",
			err:  customer.ErrInvalidPersonName,
		},
		{
			test: "Valid name",
			name: "Mister Withname",
			err:  nil,
		},
	}

	for _, tc := range testCases {
		var name string
		var err error
		t.Run(tc.test, func(t *testing.T) {
			ensure.That("testing NewCustomer factory function", func(s *ensure.Scenario) {
				s.Given(tc.test, func() {
					name = tc.name
				})

				s.When("Calling the NewCustomer factory function", func() {
					_, err = customer.NewCustomer(name)
				})

				s.Then("NewCustomer returns an error", func() {
					if err != tc.err {
						t.Errorf("expected = %v, but got = %v", err, tc.err)
					}
				})
			}, t)
		})
	}
}
