package puppy

import "github.com/akerage/dog"

func Bark() string {
	return "Woof"
}

func Barks() string {
	return "WoofWoofWoof"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}
