package common

import "testing"

// Given a string that contains brackets (), [] and {}
// validate if brackets are valid and balanced
func isValidBrackets(str string) bool {
	if len(str) == 0 {
		return true
	}

	var stk []byte

	for i := 0; i < len(str); i++ {
		if str[i] == '(' || str[i] == '[' || str[i] == '{' {
			stk = append(stk, str[i])
			continue
		}

		if len(stk) == 0 {
			return false
		}
		c := stk[len(stk)-1] // peek. top element

		if str[i] == ')' && c != '(' {
			return false
		} else if str[i] == ']' && c != '[' {
			return false
		} else if str[i] == '}' && c != '{' {
			return false
		}

		stk = stk[:len(stk)-1] // pop, because of correct combo
	}
	return len(stk) == 0
}

func TestIsValidBrackets(t *testing.T) {
	testCases := []struct {
		str string
		exp bool
	}{
		{
			str: "()[]{}",
			exp: true,
		},
		{
			str: "([])",
			exp: true,
		},
		{
			str: "([)]{}",
			exp: false,
		},
		{
			str: "([][}])",
			exp: false,
		},
		{
			str: "()[]{}",
			exp: true,
		},
		{
			str: ")))",
			exp: false,
		},
		{
			str: "{{{",
			exp: false,
		},
	}
	for i, tc := range testCases {
		got := isValidBrackets(tc.str)
		if tc.exp != got {
			t.Errorf("tc %d: exp %v, got %v", i, tc.exp, got)
		}
	}
}
