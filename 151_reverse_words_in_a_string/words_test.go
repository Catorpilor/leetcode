package words


import "testing"

func TestReverseWords(t *testing.T) {
	st := []struct{
		name string
		s string
		exp string
	}{
		{"empty string", "", ""},
		{"single words", "123", "123"},
		{"testcase1", "hello world!", "world! hello"},
		{"testcase2", "  hi my   name", "name my hi"},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T){
			out := reverseWords(tt.s)
			if out != tt.exp {
				t.Fatalf("with input s:%s wanted %s but got %s", tt.s, tt.exp, out)
			}
			t.Log("pass")
		})
	}

}
