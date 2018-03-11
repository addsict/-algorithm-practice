package base64

import (
	"strings"
)

var convertMap map[uint8]rune = map[uint8]rune{
	0x00: 'A', 0x01: 'B', 0x02: 'C', 0x03: 'D', 0x04: 'E', 0x05: 'F', 0x06: 'G', 0x07: 'H',
	0x08: 'I', 0x09: 'J', 0x0a: 'K', 0x0b: 'L', 0x0c: 'M', 0x0d: 'N', 0x0e: 'O', 0x0f: 'P',
	0x10: 'Q', 0x11: 'R', 0x12: 'S', 0x13: 'T', 0x14: 'U', 0x15: 'V', 0x16: 'W', 0x17: 'X',
	0x18: 'Y', 0x19: 'Z', 0x1a: 'a', 0x1b: 'b', 0x1c: 'c', 0x1d: 'd', 0x1e: 'e', 0x1f: 'f',
	0x20: 'g', 0x21: 'h', 0x22: 'i', 0x23: 'j', 0x24: 'k', 0x25: 'l', 0x26: 'm', 0x27: 'n',
	0x28: 'o', 0x29: 'p', 0x2a: 'q', 0x2b: 'r', 0x2c: 's', 0x2d: 't', 0x2e: 'u', 0x2f: 'v',
	0x30: 'w', 0x31: 'x', 0x32: 'y', 0x33: 'z', 0x34: '0', 0x35: '1', 0x36: '2', 0x37: '3',
	0x38: '4', 0x39: '5', 0x3a: '6', 0x3b: '7', 0x3c: '8', 0x3d: '9', 0x3e: '+', 0x3f: '/',
}

var rConvertMap map[rune]uint8 = map[rune]uint8{
	'A': 0x00, 'B': 0x01, 'C': 0x02, 'D': 0x03, 'E': 0x04, 'F': 0x05, 'G': 0x06, 'H': 0x07,
	'I': 0x08, 'J': 0x09, 'K': 0x0a, 'L': 0x0b, 'M': 0x0c, 'N': 0x0d, 'O': 0x0e, 'P': 0x0f,
	'Q': 0x10, 'R': 0x11, 'S': 0x12, 'T': 0x13, 'U': 0x14, 'V': 0x15, 'W': 0x16, 'X': 0x17,
	'Y': 0x18, 'Z': 0x19, 'a': 0x1a, 'b': 0x1b, 'c': 0x1c, 'd': 0x1d, 'e': 0x1e, 'f': 0x1f,
	'g': 0x20, 'h': 0x21, 'i': 0x22, 'j': 0x23, 'k': 0x24, 'l': 0x25, 'm': 0x26, 'n': 0x27,
	'o': 0x28, 'p': 0x29, 'q': 0x2a, 'r': 0x2b, 's': 0x2c, 't': 0x2d, 'u': 0x2e, 'v': 0x2f,
	'w': 0x30, 'x': 0x31, 'y': 0x32, 'z': 0x33, '0': 0x34, '1': 0x35, '2': 0x36, '3': 0x37,
	'4': 0x38, '5': 0x39, '6': 0x3a, '7': 0x3b, '8': 0x3c, '9': 0x3d, '+': 0x3e, '/': 0x3f,
}

func EncodeBase64(text string) string {
	byteText := []byte(text)
	var encoded []rune

	var i uint32
	for i = 0; i < uint32(len(byteText)/3); i++ {
		b := uint32(byteText[i*3+0])<<16 | uint32(byteText[i*3+1])<<8 | uint32(byteText[i*3+2])<<0
		encoded = append(encoded, convertMap[0x3f&uint8(b>>18)])
		encoded = append(encoded, convertMap[0x3f&uint8(b>>12)])
		encoded = append(encoded, convertMap[0x3f&uint8(b>>6)])
		encoded = append(encoded, convertMap[0x3f&uint8(b>>0)])
	}

	if mod := len(byteText) % 3; mod != 0 {
		if mod == 1 {
			b := uint32(byteText[i*3+0]) << 8
			encoded = append(encoded, convertMap[0x3f&uint8(b>>10)])
			encoded = append(encoded, convertMap[0x3f&uint8(b>>4)])
			encoded = append(encoded, '=')
			encoded = append(encoded, '=')
		} else if mod == 2 {
			b := uint32(byteText[i*3+0])<<16 | uint32(byteText[i*3+1])<<8
			encoded = append(encoded, convertMap[0x3f&uint8(b>>18)])
			encoded = append(encoded, convertMap[0x3f&uint8(b>>12)])
			encoded = append(encoded, convertMap[0x3f&uint8(b>>6)])
			encoded = append(encoded, '=')
		}
	}

	return string(encoded)
}

func DecodeBase64(text string) string {
	byteText := []byte(strings.TrimRight(text, "="))
	var decoded []rune

	var i uint32
	for i = 0; i < uint32(len(byteText)/4); i++ {
		tmp := uint32(rConvertMap[rune(byteText[i*4+0])])<<24 |
			uint32(rConvertMap[rune(byteText[i*4+1])])<<16 |
			uint32(rConvertMap[rune(byteText[i*4+2])])<<8 |
			uint32(rConvertMap[rune(byteText[i*4+3])])<<0
		b := uint32(uint8(tmp>>24)<<2|(0x03&(uint8(tmp>>16)>>4)))<<16 |
			uint32(uint8(tmp>>16)<<4|(0x0f&(uint8(tmp>>8)>>2)))<<8 |
			uint32(uint8(tmp>>8)<<6|(0x3f&(uint8(tmp>>0)>>0)))<<0

		decoded = append(decoded, rune(uint8(b>>16)))
		decoded = append(decoded, rune(uint8(b>>8)))
		decoded = append(decoded, rune(uint8(b>>0)))
	}

	if mod := len(byteText) % 4; mod != 0 {
		if mod == 2 {
			tmp := uint32(rConvertMap[rune(byteText[i*4+0])])<<8 |
				uint32(rConvertMap[rune(byteText[i*4+1])])<<0
			b := uint32(uint8(tmp>>8)<<2 | (0x03 & (uint8(tmp>>0) >> 4)))
			decoded = append(decoded, rune(uint8(b>>0)))
		} else if mod == 3 {
			tmp := uint32(rConvertMap[rune(byteText[i*4+0])])<<16 |
				uint32(rConvertMap[rune(byteText[i*4+1])])<<8 |
				uint32(rConvertMap[rune(byteText[i*4+2])])<<0
			b := uint32(uint8(tmp>>16)<<2|(0x03&(uint8(tmp>>8)>>4)))<<8 |
				uint32(uint8(tmp>>8)<<4|(0x0f&(uint8(tmp>>0)>>2)))<<0
			decoded = append(decoded, rune(uint8(b>>8)))
			decoded = append(decoded, rune(uint8(b>>0)))
		}
	}

	return string(decoded)
}
