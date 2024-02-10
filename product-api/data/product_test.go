package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "aj",
		Price: 2.10,
		SKU:   "a-b-c",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
