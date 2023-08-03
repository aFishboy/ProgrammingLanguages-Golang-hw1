package triangle

import "testing"

func TestGetTriangleType(t *testing.T) {
	type Test struct {
		a, b, c  int
		expected triangleType
	}

	var tests = []Test{
		{30001, 6, 2, UnknownTriangle},
		{5, 20001, 5, UnknownTriangle},
		{5, 5, 10001, UnknownTriangle},
		{-1, 5, 5, UnknownTriangle},
		{5, -1, 5, UnknownTriangle},
		{5, 5, -1, UnknownTriangle},
		{5, 5, 15, InvalidTriangle},
		{5, 15, 5, InvalidTriangle},
		{15, 5, 5, InvalidTriangle},
		{5, 4, 3, RightTriangle},
		{5, 4, 4, AcuteTriangle},
		{6, 4, 3, ObtuseTriangle},
	}

	for _, test := range tests {
		actual := getTriangleType(test.a, test.b, test.c)
		if actual != test.expected {
			t.Errorf("getTriangleType(%d, %d, %d)=%v; want %v", test.a, test.b, test.c, actual, test.expected)
		}
	}
}
