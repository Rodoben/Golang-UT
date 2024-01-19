package main

import (
	"fmt"
	"tdd/05_TestDouble_driving_liscence_generator/drivingliscencegenerator"
)

func main() {
	dl := drivingliscencegenerator.NewDrivingLiscenceNumberGenerator(&drivingliscencegenerator.SpyLogger{}, drivingliscencegenerator.FakeRand{})
	fmt.Println(dl)

	str, err := dl.Generate(drivingliscencegenerator.LiscenceHolderApplicant{})
	fmt.Println(str, err)

	str, err = dl.Generate(drivingliscencegenerator.UnderAgeApplicant{})
	fmt.Println(str, err)

	str, err = dl.Generate(drivingliscencegenerator.ValidApplicant{})
	fmt.Println(str, err)

}
