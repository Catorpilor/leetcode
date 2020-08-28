package character

// StreamChecker
type StreamChecker struct{}

// COnstructor create a StreamChecker based on words
func Constructor(words []string) *StreamChecker {
	return &StreamChecker{}
}

// Query query a letter
func (sc *StreamChecker) Query(letter byte) bool {
	return false
}
