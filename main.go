package main

var gWardrobe Wardrobe

func main() {
	tools := LoadConfig("tools.yml")
	gWardrobe := NewWardrobe(tools)

	networkHandler(gWardrobe)
}
