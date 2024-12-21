package arr

// indexRemove removes the element at the given index
//
// Business logic:
// - if the index out of bounds, will be ignored
//
// Parameters:
// - slice: the slice to remove the element from
// - index: the index of the element to remove
//
// Returns:
// - []T: the new slice
func IndexRemove[T any](slice []T, index int) []T {
	if index < 0 || index >= len(slice) {
		return slice // Nothing to remove or invalid index
	}

	return append(slice[:index], slice[index+1:]...)
}
