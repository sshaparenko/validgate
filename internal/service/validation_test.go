package service

import (
	"log"
	"testing"
)

func TestValidateNumber(t *testing.T) {
	type test struct {
		input []string
		want  bool
	}

	tests := map[string]test{
		"valid_visa": {[]string{
			"4861686289870368",
			"4716443548206662",
			"4485994849642627787",
		}, true},
		"valid_visa_whitespace": {[]string{
			"4111 1111 1111 1111",
			"4012 8888 8888 1881",
			"4222 2222 2222 2",
		}, true},
		"invalid_visa": {[]string{
			"4111111111111110",
			"5111111111111111",
			"4222 2222 2222 222",
			"",
		}, false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			for _, n := range tc.input {
				got, _ := VailedateNumber(n)
				if got != tc.want {
					t.Fatalf("%s: expected: \"%v\", got \"%v\"\ntested card number: %s", name, tc.want, got, n)
				}
			}
		})
	}
}

func TestCheckExpDate(t *testing.T) {

	type Input struct {
		ExpMonth int
		ExpYear  int
	}

	type test struct {
		input []Input
		want  bool
	}

	tests := map[string]test{
		"valid_dates": {[]Input{{
			ExpMonth: 11,
			ExpYear:  2024}, {
			ExpMonth: 10,
			ExpYear:  2025,
		}}, true},
		"invalid_dates": {[]Input{{
			ExpMonth: 6,
			ExpYear:  2024}, {
			ExpMonth: 10,
			ExpYear:  2024}, {
			ExpMonth: 10,
			ExpYear:  2023}, {
			ExpMonth: 0,
			ExpYear:  0,
		}}, false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			for _, d := range tc.input {
				got := CheckExpDate(d.ExpMonth, d.ExpYear)
				if got != tc.want {
					log.Println(d.ExpMonth)
					log.Println(d.ExpYear)
					t.Fatalf("%s: expected: \"%v\", got \"%v\"\n", name, tc.want, got)
				}
			}
		})
	}
}
