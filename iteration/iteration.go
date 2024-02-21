package iteration

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("could not update definition because the word does not exist")
)

type Dictionary map[string]string
type DictionaryErr string

// cool, error interface for Dictionary
func (e DictionaryErr) Error() string {
	return string(e)
}

// Accepts: word string
// Returns: definition string, err error
// check for word defined in the dicitonary
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// Accepts: word, definition string
// Returns: err error
// Add a word and definition pair to the dictionary
func (d Dictionary) Add(word string, definition string) error {

	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

// Accepts: word, definition string
// Returns: err error
// Update definition paired to word in to the dictionary
func (d Dictionary) Update(word, definition string) error {

	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

// Accepts: word to be deleted
// Delete word definition pair in dicitonary
func (d Dictionary) Delete(word string) {
	delete(d, word)
}

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

// Accepts: dictionary map[string]string
// Accepts: word string
// Returns: definition string
// check for word defined in the dicitonary
func Search(dictionary map[string]string, target string) string {
	return dictionary[target]
}
