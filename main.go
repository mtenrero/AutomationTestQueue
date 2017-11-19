package main

var gWardrobe Wardrobe

func main() {
	gWardrobe := NewWardrobe()

	networkHandler(gWardrobe)
}
