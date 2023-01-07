package main

import (
	. "patterns/visitor"
)

func main() {
	picture := &Picture{}
	picture.Add(NewLine())
	picture.Add(&Text{Text: "howdy doody\n"})
	picture.Add(NewLine())
	picture.Draw()

	picture.Accept(&SpookySkeletonVisitor{})
	picture.Accept(&SpookyLineVisitor{})
	picture.Draw()
}
