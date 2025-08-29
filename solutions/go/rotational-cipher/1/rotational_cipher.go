package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
	result := make([]byte, len(plain))

	for i, c := range []byte(plain) {
		switch {
		case c >= 'a' && c <= 'z':
			result[i] = byte((int(c-'a')+shiftKey)%26) + 'a'
		case c >= 'A' && c <= 'Z':
			result[i] = byte((int(c-'A')+shiftKey)%26) + 'A'
		default:
			result[i] = c
		}
	}

	return string(result)
}
