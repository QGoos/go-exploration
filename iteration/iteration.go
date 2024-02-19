package iteration

// Accepts: character string
// Accepts: count integer
// Returns: string
// generates a repeated character string
func Repeat(character string, count int) string {
	var repeated string

	for i := 0; i < count; i++ {
		repeated += character
	}

	return repeated
}

// Accepts: slice of integers
// Returns: integer
// Sum the integers in a single slice
func SumSlice(nums []int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}

	return sum
}

// Accepts: N slices of integers
// Returns: slice of integers
// sum N individual slices and compile them in another slice
func SumSlices(nums ...[]int) []int {
	var sums []int

	for _, numbers := range nums {
		sums = append(sums, SumSlice(numbers))
	}

	return sums
}
