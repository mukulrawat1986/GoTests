package main

// Repeat takes a character and repeats it 5 times
func Repeat(character string, repeatedCount int) string {
	var repeated string
	for i := 0; i < repeatedCount; i++ {
		repeated += character
	}
	return repeated
}
