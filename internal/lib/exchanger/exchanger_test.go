package exchanger

import (
	"testing"
)

var PositiveTests = []struct {
	amount    int
	banknotes []int
}{
	{
		0,
		[]int{5000, 2000, 1000, 500, 200, 100, 50},
	},
	{
		50,
		[]int{5000, 2000, 1000, 500, 200, 100, 50},
	},
	{
		100,
		[]int{5000, 2000, 1000, 500, 200, 100, 50},
	},
	{
		200,
		[]int{5000, 2000, 1000, 500, 200, 100, 50},
	},
	{
		300,
		[]int{5000, 2000, 1000, 500, 200, 100, 50},
	},
	{
		400,
		[]int{5000, 2000, 1000, 500, 200, 100, 50},
	},
	{
		2500,
		[]int{5000, 2000, 1000, 500, 200, 100, 50},
	},
	{
		8000,
		[]int{5000, 2000, 1000, 500, 200, 100, 50},
	},
}

var NegativeTests = []struct {
	amount    int
	banknotes []int
}{
	{ // negative amount
		-100,
		[]int{5000, 2000, 1000, 500, 200, 100, 50},
	},
	{ // bad amount
		201,
		[]int{5000, 2000, 1000, 500, 200, 100, 50},
	},
	{ // unsorted list of banknotes
		300,
		[]int{5000, 2000, 1000, 50, 200, 100, 500},
	},
}

func TestExchangeOK(t *testing.T) {
	for _, tc := range PositiveTests {
		items, err := Exchange(tc.amount, tc.banknotes)
		if err != nil {
			t.Errorf("amount: %d, banknotes: %v", tc.amount, tc.banknotes)
		}

		t.Logf("Number of items is %d", len(items))
		if len(items) < 1 {
			t.Fatalf("Number of items is %d (less then 1)", len(items))
		}

		for _, sl := range items {
			if sum := summaOf(sl); sum != tc.amount {
				t.Fatalf("Expected summa: %d, actual summa: %d", tc.amount, sum)
			}
		}
	}
}

func TestExchangeFail(t *testing.T) {
	for _, tc := range NegativeTests {
		_, err := Exchange(tc.amount, tc.banknotes)
		if err == nil {
			t.Errorf("amount: %d, banknotes: %v", tc.amount, tc.banknotes)
		}
	}
}
