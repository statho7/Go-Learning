package main

import(
	"fmt"
)

func main()  {
	statePopulations := make(map[string]int)

	statePopulations = map[string]int{
		"California":		39250017,
		"Texas":			27862596,
		"Florida":			20612439,
		"New York":			19745286,
		"Pennsylvania":		12802503,
		"Illinois":			12801539,
		"Ohio":				11614373,
	}
	// m := map[[3]int]string{}
	// fmt.Println(statePopulations, m)
	// fmt.Println(statePopulations)
	fmt.Println(statePopulations["Ohio"])
	statePopulations["Georgia"] = 10310371
	delete(statePopulations, "Georgia")
	fmt.Println(statePopulations)
}