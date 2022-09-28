package work_test

import (
	"myBigint/work"
	"testing"
)

func TestAdd(t *testing.T) {

	tests := []struct {
		num1     work.Bigint
		num2     work.Bigint
		expected string
	}{
		{
			num1:     work.Bigint{Value: "11"},
			num2:     work.Bigint{Value: "89"},
			expected: "100",
		},
		{
			num1:     work.Bigint{Value: "1"},
			num2:     work.Bigint{Value: "123456789"},
			expected: "123456790",
		},
		{
			num1:     work.Bigint{Value: "999"},
			num2:     work.Bigint{Value: "1"},
			expected: "1000",
		},
	}

	for i := 0; i < len(tests); i++ {
		total := work.Add(tests[i].num1, tests[i].num2)

		if total.Value != tests[i].expected {
			t.Errorf("Wrong ansver: %v, Correct answer expected: %v", total.Value, tests[i].expected)
			t.Errorf("Input values: %v, %v", tests[i].num1.Value, tests[i].num2.Value)

		}
	}
}
//*=====================================================================================

func TestSub(t *testing.T) {
	a, b := work.Bigint{Value: "111"}, work.Bigint{Value: "11"}
	total := work.Sub(a, b)
	expected := work.Bigint{Value: "100"}

	if total != expected {
		t.Errorf("Wrong ansver: %v, Correct answer expected: %v", total, expected)
	}
}
//*=====================================================================================

func TestMultiples(t *testing.T) {
	a, b := work.Bigint{Value: "11"}, work.Bigint{Value: "11"}
	total := work.Multiples(a, b)
	expected := work.Bigint{Value: "121"}

	if total != expected {
		t.Errorf("Wrong ansver: %v, Correct answer expected: %v", total, expected)
	}
}

//*=====================================================================================

func TestDivision(t *testing.T) {
	a, b := work.Bigint{Value: "999"}, work.Bigint{Value: "9"}
	total := work.Division(a, b)
	expected := work.Bigint{Value: "111"}

	if total != expected {
		t.Errorf("Wrong ansver: %v, Correct answer expected: %v", total, expected)
	}
}

//*=====================================================================================

func TestMod(t *testing.T) {
	a, b := work.Bigint{Value: "1234"}, work.Bigint{Value: "10"}
	total := work.Mod(a, b)
	expected := work.Bigint{Value: "4"}

	if total != expected {
		t.Errorf("Wrong ansver: %v, Correct answer expected: %v", total, expected)
	}
}

//*=====================================================================================

func TestAbs(t *testing.T) {
	a := work.Bigint{Value: "-123"}
	box := work.Abs(a)
	expected := work.Bigint{Value: "123"}

	if box != expected {
		t.Errorf("Wrong ansver: %v, Correct answer expected: %v", box, expected)
	}
}
