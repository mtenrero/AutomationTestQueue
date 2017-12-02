package main

var gWardrobe Wardrobe

func main() {
	tools := LoadConfig("tools.yml")
	gWardrobe := NewWardrobe(tools)

	engine := networkHandler(gWardrobe)
	engine.Run() // listen and serve on 0.0.0.0:8080 by default
}
