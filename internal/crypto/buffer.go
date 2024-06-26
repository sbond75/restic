package crypto

// NewBlobBuffer returns a buffer that is large enough to hold a blob of size
// plaintext bytes, including the crypto overhead.
func NewBlobBuffer(size int) []byte {
	return make([]byte, size, size+Extension)
}

// PlaintextLength returns the plaintext length of a blob with ciphertextSize
// bytes.
func PlaintextLength(ciphertextSize int, encrypt bool) int {
	if !encrypt {
		return ciphertextSize
	}
	return ciphertextSize - Extension
}

// CiphertextLength returns the encrypted length of a blob with plaintextSize
// bytes.
func CiphertextLength(plaintextSize int, encrypt bool) int {
	if !encrypt {
		return plaintextSize
	}
	return plaintextSize + Extension
}
