package memory

import (
	"testing"

	"github.com/atcheri/tavern-ddd-go/domain/customer"
	"github.com/google/uuid"
	"github.com/iamkoch/ensure"
)

func TestMemory_GetCustomer(t *testing.T) {
	type testCase struct {
		name string
		id   uuid.UUID
		err  error
	}

	// Create a fake customer to add to repository
	cust, err := customer.NewCustomer("Percy")
	if err != nil {
		t.Fatal(err)
	}
	id := cust.GetID()

	// Create the repo to use, and add some test Data to it for testing
	repo := MemoryCustomerRepository{
		customers: map[uuid.UUID]customer.Customer{
			id: cust,
		},
	}

	testCases := []testCase{
		{
			name: "No Customer By ID",
			id:   uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			err:  customer.ErrCustomerNotFound,
		}, {
			name: "Customer By ID",
			id:   id,
			err:  nil,
		},
	}

	for _, tc := range testCases {
		var err error
		var id uuid.UUID
		t.Run("Create a new Customer and get the ID", func(t *testing.T) {
			ensure.That("testing the GetID method", func(s *ensure.Scenario) {
				s.Given(tc.name, func() {
					id = tc.id
				})

				s.When("Calling the GetID method", func() {
					_, err = repo.Get(id)
				})

				s.Then("it returns the corresponding error", func() {
					if err != tc.err {
						t.Errorf("expected = %v, but got = %v", err, tc.err)
					}
				})
			}, t)
		})
	}
}

func TestMemory_AddCustomer(t *testing.T) {
	type testCase struct {
		name string
		cust string
		err  error
	}

	testCases := []testCase{
		{
			name: "A customer with a valid name",
			cust: "Percy",
			err:  nil,
		},
	}

	for _, tc := range testCases {
		repo := MemoryCustomerRepository{
			customers: map[uuid.UUID]customer.Customer{},
		}

		t.Run("Add a new customer in the repository", func(t *testing.T) {
			var err error
			var cust customer.Customer
			ensure.That("testing the Add method", func(s *ensure.Scenario) {
				s.Given(tc.name, func() {
					c, e := customer.NewCustomer(tc.cust)
					if e != nil {
						t.Fatal(e)
					}
					cust = c
				})
				s.When("Calling the Add method", func() {
					err = repo.Add(cust)
				})
				s.Then("it returns the corresponding error", func() {
					if err != tc.err {
						t.Errorf("expected = %v, but got = %v", err, tc.err)
					}
				})
			}, t)
		})
	}
}
