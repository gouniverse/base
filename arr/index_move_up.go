package arr

// indexMoveUp moves the element at the given index up
//
// Business logic:
// - if the index is first index, will be ignored
// - if the index out of bounds, will be ignored
//
// Parameters:
// - slice: the slice to move the element from
// - index: the index of the element to move
//
// Returns:
// - []T: the new slice
func IndexMoveUp[T any](slice []T, index int) []T {
	if index <= 0 || index >= len(slice) {
		return slice // Nothing to move or invalid index
	}

	current := slice[index]
	upper := slice[index-1]

	slice[index] = upper
	slice[index-1] = current

	return slice
}
