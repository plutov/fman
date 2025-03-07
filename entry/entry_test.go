package entry

import "testing"

func TestHighlightSyntax(t *testing.T) {
	testCases := []struct {
		desc     string
		name     string
		preview  string
		expected string
	}{
		{
			desc:     "empty",
			name:     "",
			preview:  "",
			expected: "",
		},
		{
			desc:    "go",
			name:    "go",
			preview: "package main\n\nfunc main()\n{\n}\n",
			expected: `[1m[37mpackage main[0m[1m[37m
[0m[1m[37m
[0m[1m[37mfunc main()[0m[1m[37m
[0m[1m[37m{[0m[1m[37m
[0m[1m[37m}[0m[1m[37m
[0m`,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, _ := HighlightSyntax(tC.name, tC.preview)
			if got != tC.expected {
				t.Errorf("expecting %s, got %v", tC.expected, got)
			}
		})
	}
}
