package service

import (
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
