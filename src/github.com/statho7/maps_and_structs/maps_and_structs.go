package main

import(
	"fmt"
)

func main()  {
	// statePopulations := make(map[string]int)

	// statePopulations = map[string]int{
	// 	"California":		39250017,
	// 	"Texas":			27862596,
	// 	"Florida":			20612439,
	// 	"New York":			19745286,
	// 	"Pennsylvania":		12802503,
	// 	"Illinois":			12801539,
	// 	"Ohio":				11614373,
	// }
	
	// m := map[[3]int]string{}
	// fmt.Println(statePopulations, m)
	// fmt.Println(statePopulations)
	// fmt.Println(statePopulations["Ohio"])
	// statePopulations["Georgia"] = 10310371
	// delete(statePopulations, "Georgia")
	// fmt.Println(statePopulations)
	// fmt.Println(statePopulations["Georgia"])

	// pop, ok := statePopulations["Oho"]
	// fmt.Println(pop, ok)
	// pop, ok = statePopulations["Ohio"]
	// fmt.Println(pop, ok)

	// fmt.Println(len(statePopulations))

	// sp := statePopulations
	// delete(sp, "Ohio")
	// fmt.Println(sp)
	// fmt.Println(statePopulations)

    m := make(map[string]int)

    m["k1"] = 7
    m["k2"] = 13

    fmt.Println("map:", m)

    v1 := m["k1"]
    fmt.Println("v1: ", v1)

    fmt.Println("len:", len(m))

    delete(m, "k2")
    fmt.Println("map:", m)

    _, prs := m["k2"]
    fmt.Println("prs:", prs)

    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)
}