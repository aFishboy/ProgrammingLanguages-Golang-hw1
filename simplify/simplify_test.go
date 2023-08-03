package simplify

import (
	"hw1/expr"
	"testing"
)

//!+Simplify
func TestSimplify(t *testing.T) {
	tests := []struct {
		expr string
		env  expr.Env
		want string
	}{
		{"5", expr.Env{}, "5"},
		{"5 + 2", expr.Env{}, "7"},
		{"X", expr.Env{}, "X"},
		{"X", expr.Env{"X": 2}, "2"},
		{"Y", expr.Env{}, "Y"},
		{"Y", expr.Env{"Y": 2}, "2"},
		{"-X", expr.Env{"X": 2}, "-2"},
		{"+X", expr.Env{"X": 2}, "2"},
		{"-X", expr.Env{}, "(-X)"},
		{"--X", expr.Env{"X": 2}, "2"},
		{"--X", expr.Env{}, "(-(-X))"},
		{"2 + 3", expr.Env{}, "5"},
		{"X + 3", expr.Env{"X": 2}, "5"},
		{"3 + X", expr.Env{"X": 2}, "5"},
		{"X + 3", expr.Env{}, "(X + 3)"},
		{"2 + X", expr.Env{}, "(2 + X)"},
		{"Y + X", expr.Env{}, "(Y + X)"},
		{"X + Y", expr.Env{}, "(X + Y)"},
		{"Y", expr.Env{"X": 2}, "Y"},
		{"Y*0", expr.Env{"X": 2}, "0"},
		{"Y+0", expr.Env{"X": 2}, "Y"},
		{"Y*1", expr.Env{"X": 2}, "Y"},
		{"0*Y", expr.Env{"X": 2}, "0"},
		{"0+Y", expr.Env{"X": 2}, "Y"},
		{"1*Y", expr.Env{"X": 2}, "Y"},
		{"1*Y*1", expr.Env{"X": 2}, "Y"},
		{"X+2", expr.Env{"X": 2}, "4"},
		{"10 / X", expr.Env{"X": 2}, "5"},
		{"Y+2", expr.Env{"X": 2}, "(Y + 2)"},
		{"3 + 5 + X", expr.Env{"X": 2}, "10"},
		{"X + 3 + 5", expr.Env{}, "((X + 3) + 5)"},
		{"2 + X + 3", expr.Env{}, "((2 + X) + 3)"},
		{"X + 2 + 3", expr.Env{}, "((X + 2) + 3)"},
		{"2 + 3 + X", expr.Env{}, "(5 + X)"},
		{"(X + X) - Y", expr.Env{"X": 2}, "(4 - Y)"},
		{"(X + X) - Y", expr.Env{"Y": 8}, "((X + X) - 8)"},
		{"10 - 1 + X - Y", expr.Env{}, "((9 + X) - Y)"},
		{"X + 3 + 5", expr.Env{}, "((X + 3) + 5)"},
		{"-(X + X) - Y", expr.Env{"X": 2}, "(-4 - Y)"},
		{"(X + X) - Y", expr.Env{"Y": 8}, "((X + X) - 8)"},
		{"10 - 1 + X - Y", expr.Env{}, "((9 + X) - Y)"},
		{"X + 3 + 5", expr.Env{}, "((X + 3) + 5)"},
		{"-1", expr.Env{}, "-1"},
		{"-1+2", expr.Env{}, "1"},
		{"(1+2)*3", expr.Env{}, "9"},
		{"X/(2+3)", expr.Env{"X": 10}, "2"},
		{"(X+2)/(Y-3)", expr.Env{"X": 10, "Y": 5}, "6"},
		{"(X+2)*(Y-3)", expr.Env{"X": 10, "Y": 5}, "24"},
		{"((X-2)*(Y-3)+5)/(Y-3)", expr.Env{"X": 10, "Y": 5}, "10.5"},
		{"1/0", expr.Env{}, "+Inf"},
		{"1+2+3", expr.Env{}, "6"},
		{"5*0", expr.Env{}, "0"},
		{"10/2", expr.Env{}, "5"},
		{"(3+4)*2", expr.Env{}, "14"},
		{"(3-4)*2", expr.Env{}, "-2"},
		{"(3-4)/2", expr.Env{}, "-0.5"},
		{"2*(3-4)", expr.Env{}, "-2"},
		{"2*(3+4)", expr.Env{}, "14"},
		{"(3+4)/2", expr.Env{}, "3.5"},
		{"2*(3+4)/2", expr.Env{}, "7"},
		{"(3+4)/2*2", expr.Env{}, "7"},
		{"1.5*2", expr.Env{}, "3"},
		{"5 + 2", expr.Env{}, "7"},
		{"10 - 5", expr.Env{}, "5"},
		{"3 * 4", expr.Env{}, "12"},
		{"15 / 3", expr.Env{}, "5"},
		{"X", expr.Env{}, "X"},
		{"X", expr.Env{"X": 2}, "2"},
		{"X - 4", expr.Env{"X": 7}, "3"},
		{"X + Y", expr.Env{"X": 3, "Y": 5}, "8"},
		{"X * Y", expr.Env{"X": 2, "Y": 4}, "8"},
		{"X / Y", expr.Env{"X": 10, "Y": 2}, "5"},
		{"X + Y - Z", expr.Env{"X": 10, "Y": 5, "Z": 3}, "12"},
		{"X * Y + Z", expr.Env{"X": 2, "Y": 3, "Z": 4}, "10"},
		{"X / Y - Z", expr.Env{"X": 10, "Y": 2, "Z": 1}, "4"},
		{"X * Y / Z", expr.Env{"X": 5, "Y": 10, "Z": 2}, "25"},
		{"X + Y * Z", expr.Env{"X": 2, "Y": 3, "Z": 4}, "14"},
		{"(X + Y) * Z", expr.Env{"X": 2, "Y": 3, "Z": 4}, "20"},
		{"X - Y / Z", expr.Env{"X": 10, "Y": 2, "Z": 4}, "9.5"},
		{"(X + Y) / (Z - 1)", expr.Env{"X": 6, "Y": 4, "Z": 3}, "5"},
		{"X - Y - Z", expr.Env{"X": 10, "Y": 2, "Z": 1}, "7"},
		{"Y", expr.Env{"Y": 2}, "2"},
		{"-Y", expr.Env{"Y": 2}, "-2"},
		{"-Y", expr.Env{}, "(-Y)"},
		{"--Y", expr.Env{"Y": 2}, "2"},
		{"Y*0", expr.Env{"Y": 2}, "0"},
		{"Y+0", expr.Env{"Y": 2}, "2"},
		{"Y*1", expr.Env{"Y": 2}, "2"},
		{"0*Y", expr.Env{"Y": 2}, "0"},
		{"0+Y", expr.Env{"Y": 2}, "2"},
		{"1*Y", expr.Env{"Y": 2}, "2"},
		{"1*Y*1", expr.Env{"Y": 2}, "2"},
		{"Y+2", expr.Env{"Y": 2}, "4"},
		{"10 / Y", expr.Env{"Y": 2}, "5"},
		{"Y+2", expr.Env{}, "(Y + 2)"},
		{"(Y + Y) - Y", expr.Env{"Y": 2}, "2"},
		{"0", expr.Env{}, "0"},
		{"X + 0", expr.Env{"X": 2}, "2"},
		{"X * 0", expr.Env{"X": 2}, "0"},
		{"0 + X", expr.Env{"X": 2}, "2"},
		{"0 * X", expr.Env{"X": 2}, "0"},
		{"X * Y * 0", expr.Env{"X": 2, "Y": 5}, "0"},
		{"(X * Y) + (0 * Z)", expr.Env{"X": 2, "Y": 5, "Z": 7}, "10"},
		{"(X * Y) - (0 + Z)", expr.Env{"X": 2, "Y": 5, "Z": 7}, "3"},
		{"0 - X", expr.Env{"X": 2}, "-2"},
		{"X - 0", expr.Env{"X": 2}, "2"},
		{"0 / X", expr.Env{"X": 2}, "0"},
	}

	for _, test := range tests {
		e, err := expr.Parse(test.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		// Run the method
		result := Simplify(e, test.env)

		// Display the result
		got := expr.Format(result)

		// Check the result
		if got != test.want {
			t.Errorf("Simplify(%s) in %v = %q, want %q\n",
				test.expr, test.env, got, test.want)
		}
	}
}

type Foo int

func (Foo) Eval(env expr.Env) float64 {
	return 0.0
}

func (Foo) Check(vars map[expr.Var]bool) error {
	return nil
}
func TestSimplify_Fail(t *testing.T) {

	func() {
		defer func() {
			if recover() == nil {
				t.Errorf("did not panic, but should\n")
			}
		}()

		var f Foo
		Simplify(f, expr.Env{})
	}()

}

//!-Simplify
