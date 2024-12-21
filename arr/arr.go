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

// indexMoveDown moves the element at the given index down
//
// Business logic:
// - if the index is last index, will be ignored
// - if the index out of bounds, will be ignored
//
// Parameters:
// - slice: the slice to move the element from
// - index: the index of the element to move
//
// Returns:
// - []T: the new slice
func IndexMoveDown[T any](slice []T, index int) []T {
	if index < 0 || index >= len(slice)-1 {
		return slice // Nothing to move or invalid index
	}

	current := slice[index]
	lower := slice[index+1]

	slice[index] = lower
	slice[index+1] = current

	return slice
}

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
func indexMoveUp[T any](slice []T, index int) []T {
	if index <= 0 || index >= len(slice) {
		return slice // Nothing to move or invalid index
	}

	current := slice[index]
	upper := slice[index-1]

	slice[index] = upper
	slice[index-1] = current

	return slice
}
