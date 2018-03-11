package base64

import (
	"encoding/binary"
)

var convertMap map[uint8]rune = map[uint8]rune{
	0x00: 'A',
	0x01: 'B',
	0x02: 'C',
	0x03: 'D',
	0x04: 'E',
	0x05: 'F',
	0x06: 'G',
	0x07: 'H',
	0x08: 'I',
	0x09: 'J',
	0x0a: 'K',
	0x0b: 'L',
	0x0c: 'M',
	0x0d: 'N',
	0x0e: 'O',
	0x0f: 'P',
	0x10: 'Q',
	0x11: 'R',
	0x12: 'S',
	0x13: 'T',
	0x14: 'U',
	0x15: 'V',
	0x16: 'W',
	0x17: 'X',
	0x18: 'Y',
	0x19: 'Z',
	0x1a: 'a',
	0x1b: 'b',
	0x1c: 'c',
	0x1d: 'd',
	0x1e: 'e',
	0x1f: 'f',
	0x20: 'g',
	0x21: 'h',
	0x22: 'i',
	0x23: 'j',
	0x24: 'k',
	0x25: 'l',
	0x26: 'm',
	0x27: 'n',
	0x28: 'o',
	0x29: 'p',
	0x2a: 'q',
	0x2b: 'r',
	0x2c: 's',
	0x2d: 't',
	0x2e: 'u',
	0x2f: 'v',
	0x30: 'w',
	0x31: 'x',
	0x32: 'y',
	0x33: 'z',
	0x34: '0',
	0x35: '1',
	0x36: '2',
	0x37: '3',
	0x38: '4',
	0x39: '5',
	0x3a: '6',
	0x3b: '7',
	0x3c: '8',
	0x3d: '9',
	0x3e: '+',
	0x3f: '/',
}

func EncodeBase64(text string) string {
	byteText := []byte(text)
	var encoded []rune

	var i uint32
	for i = 0; i < uint32(len(byteText)/3); i++ {
		b := binary.BigEndian.Uint32([]byte{0x00, byteText[i*3+0], byteText[i*3+1], byteText[i*3+2]})
		encoded = append(encoded, convertMap[0x3f&uint8(b>>18)])
		encoded = append(encoded, convertMap[0x3f&uint8(b>>12)])
		encoded = append(encoded, convertMap[0x3f&uint8(b>>6)])
		encoded = append(encoded, convertMap[0x3f&uint8(b>>0)])
	}

	if mod := len(byteText) % 3; mod != 0 {
		if mod == 1 {
			b := binary.BigEndian.Uint32([]byte{0x00, 0x00, byteText[i*3+0], 0x00})
			encoded = append(encoded, convertMap[0x3f&uint8(b>>10)])
			encoded = append(encoded, convertMap[0x3f&uint8(b>>4)])
			encoded = append(encoded, '=')
			encoded = append(encoded, '=')
		} else if mod == 2 {
			b := binary.BigEndian.Uint32([]byte{0x00, byteText[i*3+0], byteText[i*3+1], 0x00})
			encoded = append(encoded, convertMap[0x3f&uint8(b>>18)])
			encoded = append(encoded, convertMap[0x3f&uint8(b>>12)])
			encoded = append(encoded, convertMap[0x3f&uint8(b>>6)])
			encoded = append(encoded, '=')
		}
	}

	return string(encoded)
}

func DecodeBase64(text string) string {
	return ""
}
