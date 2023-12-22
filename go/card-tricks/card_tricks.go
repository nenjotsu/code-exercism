package cards

var cards = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	var fav []int
	for _, v := range cards {
		if v == 2 || v == 6 || v == 9 {
			fav = append(fav, v)
		}
	}
	return fav
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	for i, v := range slice {
		if i == index {
			return v
		}
	}
	return -1
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	for i := range slice {
		if index == i {
			slice[i] = value
			return slice
		}
	}
	return append(slice, value)
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	if len(values) == 0 {
		return slice
	}
	return append(values, slice...)
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index > len(slice) || index < 0 {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}
