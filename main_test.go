package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestTranspose(t *testing.T) {
	testCases := []struct {
		In      string
		Query   string
		WantOut string
	}{
		{"a,b,c\nd,e,f", "", ""},
		{"a,b,c\nd,e,f", "{{index . 1}}", "be"},
		{"a,b,c\nd,e,f", "{{index . 1}}\n", "b\ne\n"},
		{"a,b,c\nd,e,f", "Hello {{index . 1}}\n", "Hello b\nHello e\n"},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			out := strings.Builder{}
			in := strings.NewReader(tc.In)
			err := transpose(&out, in, tc.Query)
			t.Logf("got %q", out.String())
			if err != nil {
				t.Fatalf("got error %s", err)
			}
			if g, w := out.String(), tc.WantOut; g != w {
				t.Errorf("got %q, want %q", g, w)
			}
		})
	}
}
