package tag

import (
	"fmt"
	"strings"

	"github.com/catorpilor/LeetCode/utils"
)

type sb struct {
	containTags bool
}

func IsValid(code string) bool {
	n := len(code)
	fmt.Println(n)
	if n < 7 {
		return false
	}
	if code[0] != '<' || code[n-1] != '>' {
		return false
	}
	cts := sb{}
	st := utils.NewStack()
	for i := 0; i < n; i++ {
		ended := false
		cl := 0
		if st.IsEmpty() && cts.containTags {
			// <A></A><B></B>
			return false
		}
		if code[i] == '<' {
			if !st.IsEmpty() && code[i+1] == '!' {
				cl = strings.Index(code[i+1:], "]]>")
				if cl > 0 {
					cl += i + 1
				}
				fmt.Printf("cdata section cl: %d, i:%d\n", cl, i)
				if cl < 0 || !isValidCDATA(code[i+2:cl]) {
					fmt.Printf("failed here with cl: %d, and cdata start index: %d\n", cl, i+2)
					return false
				}
			} else {
				if code[i+1] == '/' {
					ended = true
					i++
				}
				cl = strings.Index(code[i:], ">")
				if cl > 0 {
					cl += i
				}
				fmt.Printf("cdata section cl: %d, i:%d\n", cl, i)
				if cl < 0 || !validTag(code[i+1:cl], st, ended, &cts) {
					fmt.Printf("failed here with cl: %d, and tag's start index: %d\n", cl, i+1)
					return false
				}
			}
			i = cl
			fmt.Printf("setting i to %d\n", cl)
		}
	}
	fmt.Printf("st: %t and containtags: %t\n", st.IsEmpty(), cts.containTags)
	return st.IsEmpty() && cts.containTags
}

func validTag(code string, st *utils.Stack, ending bool, cts *sb) (b bool) {
	// fmt.Printf("tag is %s\n", code)
	n := len(code)
	if n < 1 || n > 9 {
		return false
	}
	for i := 0; i < n; i++ {
		if !isUpperCase(code[i]) {
			return
		}
	}
	if ending {
		// closed tag
		if !st.IsEmpty() && st.Top().(string) == code {
			st.Pop()
		} else {
			return false
		}
	} else {
		st.Push(code)
		cts.containTags = true
	}
	return true
}

func isUpperCase(c byte) bool {
	if c < 65 || c > 90 {
		return false
	}
	return true
}

func isValidCDATA(code string) bool {
	return strings.Index(code, "[CDATA[") == 0
}
