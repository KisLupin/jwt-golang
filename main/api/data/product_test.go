package data

import "testing"

func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name: "lupin",
		Price: 2.00,
		Description: "abc-abc-abc-abc-abc",
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
