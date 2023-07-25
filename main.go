package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

//arrays in go are of two type one is array and the other is slice.
// array is fixed size and slice is variable size.
//every element in an array is of the same data type in slice.
