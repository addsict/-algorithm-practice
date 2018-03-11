package base64

import "testing"

func TestBase64(t *testing.T) {
	tests := []struct {
		decoded string
		encoded string
	}{
		{"", ""},
		{"ABCD", "QUJDRA=="},
		{"hello", "aGVsbG8="},
		{"base64", "YmFzZTY0"},
		{"hello\nworld\n", "aGVsbG8Kd29ybGQK"},
	}

	t.Run("EncodeBase64", func(t *testing.T) {
		for _, test := range tests {
			got := EncodeBase64(test.decoded)
			if got != test.encoded {
				t.Errorf("encode %s, expected %s, but got %s", test.decoded, test.encoded, got)
			}
		}
	})

	t.Run("DecodeBase64", func(t *testing.T) {
		for _, test := range tests {
			got := DecodeBase64(test.encoded)
			if got != test.decoded {
				t.Errorf("decode %s, expected %s, but got %s", test.encoded, test.decoded, got)
			}
		}
	})
}
