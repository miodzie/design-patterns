package main

import . "patterns/composite"

func main() {
	picture := &Picture{}
	picture.Add(&Line{})
	picture.Add(&Line{})
	picture.Add(&Rectangle{})
	picture.Add(&Line{})
	picture.Add(&Text{Text: "howdy doody\n"})
	picture.Add(&Line{})

	pic2 := &Picture{}
	pic2.Add(&Rectangle{})
	pic2.Add(&Text{Text: "\tI AM THE BOX GHOST OOOO!!"})
	pic2.Add(&Rectangle{})

	picture.Add(pic2)

	picture.Draw()
}
