package calculator

import "testing"

func TestDiscountApplied(t *testing.T) {
	calculator := NewDiscountCalculator(100, 20)
	amount := calculator.Calculate(150)

	if amount != 130 {
		t.Fail()
	}
}

func TestDiscountNotApplied(t *testing.T) {
	calculator := NewDiscountCalculator(100, 20)
	amount := calculator.Calculate(50)

	if amount != 50 {
		t.Errorf("expected 50, got %v. failed because the discount was not expected to be applied", amount) // Error = Log + Fail
	}
}
