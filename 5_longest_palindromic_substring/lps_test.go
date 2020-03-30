package lps

import "testing"

func TestLps(t *testing.T) {
	st := []struct {
		name, s, exp string
	}{
		{"empty string", "", ""},
		{"single character", "a", "a"},
		{"two characters", "ab", "a"},
		{"two identical", "aa", "aa"},
		{"test ababb", "ababb", "bab"},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Lps(c.s)
			if ret != c.exp {
				t.Fatalf("expected %s but got %s, with input %s",
					c.exp, ret, c.s)
			}
		})
	}
}

func TestLps2(t *testing.T) {
	st := []struct {
		name, s, exp string
	}{
		{"empty string", "", ""},
		{"single character", "a", "a"},
		{"two characters", "ab", "a"},
		{"two identical", "aa", "aa"},
		{"test abb", "abb", "bb"},
		{"test ccc", "ccc", "ccc"},
		{"test ababb", "ababb", "aba"},
		{"test babadada", "babadada", "adada"},
		{"test abcbcbda", "abcbcbda", "bcbcb"},
		{"test abaxabaxabb", "abaxabaxabb", "baxabaxab"},
		{"test abaxabaxabybaxaby", "abaxabaxabybaxaby", "baxabybaxab"},
		{"failed", "jglknendplocymmvwtoxvebkekzfdhykknufqdkntnqvgfbahsljkobhbxkvyictzkqjqydczuxjkgecdyhixdttxfqmgksrkyvopwprsgoszftuhawflzjyuyrujrxluhzjvbflxgcovilthvuihzttzithnsqbdxtafxrfrblulsakrahulwthhbjcslceewxfxtavljpimaqqlcbrdgtgjryjytgxljxtravwdlnrrauxplempnbfeusgtqzjtzshwieutxdytlrrqvyemlyzolhbkzhyfyttevqnfvmpqjngcnazmaagwihxrhmcibyfkccyrqwnzlzqeuenhwlzhbxqxerfifzncimwqsfatudjihtumrtjtggzleovihifxufvwqeimbxvzlxwcsknksogsbwwdlwulnetdysvsfkonggeedtshxqkgbhoscjgpiel", "sknks"},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Lps2(c.s)
			if ret != c.exp {
				t.Fatalf("expected %s but got %s, with input %s",
					c.exp, ret, c.s)
			}
		})
	}
}

func TestUseBinarySearch(t *testing.T) {
	st := []struct {
		name, s, exp string
	}{
		{"empty string", "", ""},
		{"single character", "a", "a"},
		{"two characters", "ab", "a"},
		{"two identical", "aa", "aa"},
		{"test abb", "abb", "bb"},
		{"test ccc", "ccc", "ccc"},
		{"test ababb", "ababb", "aba"},
		{"test babadada", "babadada", "adada"},
		{"test abcbcbda", "abcbcbda", "bcbcb"},
		{"test abaxabaxabb", "abaxabaxabb", "baxabaxab"},
		{"test abaxabaxabybaxaby", "abaxabaxabybaxaby", "baxabybaxab"},
		{"failed", "jglknendplocymmvwtoxvebkekzfdhykknufqdkntnqvgfbahsljkobhbxkvyictzkqjqydczuxjkgecdyhixdttxfqmgksrkyvopwprsgoszftuhawflzjyuyrujrxluhzjvbflxgcovilthvuihzttzithnsqbdxtafxrfrblulsakrahulwthhbjcslceewxfxtavljpimaqqlcbrdgtgjryjytgxljxtravwdlnrrauxplempnbfeusgtqzjtzshwieutxdytlrrqvyemlyzolhbkzhyfyttevqnfvmpqjngcnazmaagwihxrhmcibyfkccyrqwnzlzqeuenhwlzhbxqxerfifzncimwqsfatudjihtumrtjtggzleovihifxufvwqeimbxvzlxwcsknksogsbwwdlwulnetdysvsfkonggeedtshxqkgbhoscjgpiel", "sknks"},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := useBinarySearch(c.s)
			if ret != c.exp {
				t.Fatalf("expected %s but got %s, with input %s",
					c.exp, ret, c.s)
			}
		})
	}
}

func TestUseCenter(t *testing.T) {
	st := []struct {
		name, s, exp string
	}{
		{"empty string", "", ""},
		{"single character", "a", "a"},
		{"two characters", "ab", "a"},
		{"two identical", "aa", "aa"},
		{"test abb", "abb", "bb"},
		{"test ccc", "ccc", "ccc"},
		{"test ababb", "ababb", "aba"},
		{"test babadada", "babadada", "adada"},
		{"test abcbcbda", "abcbcbda", "bcbcb"},
		{"test abaxabaxabb", "abaxabaxabb", "baxabaxab"},
		{"test abaxabaxabybaxaby", "abaxabaxabybaxaby", "baxabybaxab"},
		{"failed", "jglknendplocymmvwtoxvebkekzfdhykknufqdkntnqvgfbahsljkobhbxkvyictzkqjqydczuxjkgecdyhixdttxfqmgksrkyvopwprsgoszftuhawflzjyuyrujrxluhzjvbflxgcovilthvuihzttzithnsqbdxtafxrfrblulsakrahulwthhbjcslceewxfxtavljpimaqqlcbrdgtgjryjytgxljxtravwdlnrrauxplempnbfeusgtqzjtzshwieutxdytlrrqvyemlyzolhbkzhyfyttevqnfvmpqjngcnazmaagwihxrhmcibyfkccyrqwnzlzqeuenhwlzhbxqxerfifzncimwqsfatudjihtumrtjtggzleovihifxufvwqeimbxvzlxwcsknksogsbwwdlwulnetdysvsfkonggeedtshxqkgbhoscjgpiel", "sknks"},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := useExpendingFromCenter(c.s)
			if ret != c.exp {
				t.Fatalf("expected %s but got %s, with input %s",
					c.exp, ret, c.s)
			}
		})
	}
}
