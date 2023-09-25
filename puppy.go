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

func GetVersion1() string {
    return "I'm from v0.0.1"
}

func GetVersion2() string {
    return "I'm from v0.0.2"
}
