package utils

import (
	"testing"
)

func TestGetFibonacciNumberWithPositionBinet(t *testing.T) {
	cases := []struct {
		input  int
		expect int
	}{
		{
			input:  0,
			expect: 0,
		},
		{
			input:  1,
			expect: 1,
		},
		{
			input:  2,
			expect: 1,
		},
		{
			input:  3,
			expect: 2,
		},
		{
			input:  4,
			expect: 3,
		},
		{
			input:  5,
			expect: 5,
		},
		{
			input:  6,
			expect: 8,
		},
		{
			input:  7,
			expect: 13,
		},
		{
			input:  8,
			expect: 21,
		},
		{
			input:  9,
			expect: 34,
		},
		{
			input:  10,
			expect: 55,
		},
		{
			input:  15,
			expect: 610,
		},
		{
			input:  19,
			expect: 4181,
		},
	}

	for _, cs := range cases {
		actual, err := GetFibonacciNumberWithPositionBinet(cs.input)
		if err != nil {
			t.Errorf("Invalid input: %d", cs.input)
			continue
		}
		if actual != cs.expect {
			t.Errorf("%d doesn't equal %d", actual, cs.expect)
			continue
		}
	}
}

func TestGetFibonacciNumberWithPosition(t *testing.T) {
	cases := []struct {
		input  int
		expect int
	}{
		{
			input:  0,
			expect: 0,
		},
		{
			input:  1,
			expect: 1,
		},
		{
			input:  2,
			expect: 1,
		},
		{
			input:  3,
			expect: 2,
		},
		{
			input:  4,
			expect: 3,
		},
		{
			input:  5,
			expect: 5,
		},
		{
			input:  6,
			expect: 8,
		},
		{
			input:  7,
			expect: 13,
		},
		{
			input:  8,
			expect: 21,
		},
		{
			input:  9,
			expect: 34,
		},
		{
			input:  10,
			expect: 55,
		},
		{
			input:  15,
			expect: 610,
		},
		{
			input:  19,
			expect: 4181,
		},
	}

	for _, cs := range cases {
		actual, err := GetFibonacciNumberWithPosition(cs.input)
		if err != nil {
			t.Errorf("Invalid input: %d", cs.input)
			continue
		}
		if actual != cs.expect {
			t.Errorf("%d doesn't equal %d", actual, cs.expect)
			continue
		}
	}
}
