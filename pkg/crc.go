package pkg

import (
	"fmt"
)

func CRC(data, key string) string {
	keyLength := len(key)
	var zeros string
	for i := 0; i < keyLength-1; i++ {
		zeros += fmt.Sprint("0")
	}
	d := data + zeros
	dividend := d[0:keyLength]
	for keyLength < len(d) {
		if fmt.Sprintf("%c", dividend[0]) == "1" {
			dividend = xor(key, dividend) + fmt.Sprintf("%c", d[keyLength])
		} else {
			dividend = xor(zeros+"0", dividend) + fmt.Sprintf("%c", d[keyLength])
		}
		keyLength += 1
	}
	if fmt.Sprintf("%c", dividend[0]) == "1" {
		dividend = xor(key, dividend)
	} else {
		dividend = xor(zeros+"0", dividend)
	}
	return data + dividend
}

func CheckSum(frame, key string) string {
	frameLength := len(frame)
	keyLength := len(key)
	var zeros string
	for i := 0; i < keyLength; i++ {
		zeros += fmt.Sprint("0")
	}
	dividend := frame[0:keyLength]
	for keyLength < frameLength {
		if fmt.Sprintf("%c", dividend[0]) == "1" {
			dividend = xor(key, dividend) + fmt.Sprintf("%c", frame[keyLength])
		} else {
			dividend = xor(zeros, dividend) + fmt.Sprintf("%c", frame[keyLength])
		}
		keyLength += 1
	}
	if fmt.Sprintf("%c", dividend[0]) == "1" {
		dividend = xor(key, dividend)
	} else {
		dividend = xor(zeros, dividend)
	}
	return dividend
}

func xor(a string, b string) string {
	bLength := len(b)
	var result string
	for i := 1; i < bLength; i++ {
		if fmt.Sprintf("%c", a[i]) == fmt.Sprintf("%c", b[i]) {
			result += fmt.Sprintf("0")
		} else {
			result += fmt.Sprintf("1")
		}
	}
	return result
}
