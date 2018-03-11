package base64

import "testing"

func TestEncodeBase64(t *testing.T) {

	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"ABCD", "QUJDRA=="},
		{"hello", "aGVsbG8="},
		{"base64", "YmFzZTY0"},
		{"hello\nworld\n", "aGVsbG8Kd29ybGQK"},
	}

	for _, test := range tests {
		got := EncodeBase64(test.input)
		if got != test.expected {
			t.Errorf("encode %s, expected %s, but got %s", test.input, test.expected, got)
		}
	}
}
