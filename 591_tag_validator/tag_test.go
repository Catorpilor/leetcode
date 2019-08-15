package tag

import "testing"

func TestIsValid(t *testing.T) {
	st := []struct {
		name string
		code string
		exp  bool
	}{
		{"empty input", "", false},
		{"start tag len eq 0", "<> Hohoho </TAG>", false},
		{"start tag len gt 9", "<THISISATESTFORCODE> HOohoho </THISISATESTFORCODE>", false},
		{"tag is lowercase", "<this> hohoho </this>", false},
		{"unclosed tag", "<THIS> hohoho </THAT>", false},
		{"valid testcase 1", "<DIV>This is the first line <![CDATA[<div>]]></DIV>", true},
		{"valid testcase 2", "<DIV>>>  ![cdata[]] <![CDATA[<div>]>]]>]]>>]</DIV>", true},
		{"failed 1", "<A><![CDATA[</A>]]123></A>", false},
		{"failed 2", "<A></A><B></B>", false},
		{"failed 3", "<DIV>This is the first line <![CDATA[<<<<<<<]]></DIV>", true},
	}

	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := IsValid(tt.code)
			// containTags = false
			if out != tt.exp {
				t.Fatalf("with input code: %s, wanted %t but got %t", tt.code, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
