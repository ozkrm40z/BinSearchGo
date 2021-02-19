package binsearchgo

//BinSearch Model
type BinSearch struct {
	_array    []int
	_arraylen int
}

//Search a number inside an ordered array, returns true if found and position number, false if not found and -1
func (bs *BinSearch) Search(value int) (bool, int) {
	min := 0
	max := bs._arraylen
	for min <= max {
		currentPosition := (min + max) / 2
		currentValue := bs._array[currentPosition]
		if currentValue == value {
			return true, currentPosition
		}

		if currentValue < value {
			min = currentPosition + 1
		} else {
			max = currentPosition - 1
		}
	}

	return false, -1
}

//NewBinSearchService Constructor
func NewBinSearchService(array []int) *BinSearch {
	return &BinSearch{
		_array:    array,
		_arraylen: len(array) - 1,
	}
}
