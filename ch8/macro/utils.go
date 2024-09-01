package macro

// isalphanum -- checks if ch is [0-9A-Za-z]
func isalphanum(ch uint8) bool {
	return isletter(ch) || ('0' <= ch && ch <= '9')
}

// isletter -- true if c is a letter of either case
func isletter(ch uint8) bool {
	return ('a' <= ch && ch <= 'z') ||
		('A' <= ch && ch <= 'Z')
}

// cscopy -- copy ch[i] - array of chars - to string
func cscopy(arr []byte, i int) string {
	const ENDSTR byte = '\n'
	start := i
	for arr[i] != ENDSTR {
		i++
	}
	end := i

	// Strings in Go are immutable, let's make a string from a slice of the array
	return string(arr[start:end])
}

// sccopy -- copy string s to cb[i] - array of chars
func sccopy(arr []byte, temp string, i int) {
	const ENDSTR byte = '\n'
	for j := 0; j < len(temp) && temp[j] != ENDSTR; {
		arr[i] = temp[j]
		i++
		j++
	}
	arr[i] = ENDSTR
}
