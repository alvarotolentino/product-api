package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Alvaro",
		Price: 1.00,
		SKU:   "abc-edf-ghi",
	}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
