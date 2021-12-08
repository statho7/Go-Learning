package main

// import(
// 	"fmt"
// )

// func main()  {
// 	// statePopulations := make(map[string]int)

// 	// statePopulations = map[string]int{
// 	// 	"California":		39250017,
// 	// 	"Texas":			27862596,
// 	// 	"Florida":			20612439,
// 	// 	"New York":			19745286,
// 	// 	"Pennsylvania":		12802503,
// 	// 	"Illinois":			12801539,
// 	// 	"Ohio":				11614373,
// 	// }

// 	// m := map[[3]int]string{}
// 	// fmt.Println(statePopulations, m)
// 	// fmt.Println(statePopulations)
// 	// fmt.Println(statePopulations["Ohio"])
// 	// statePopulations["Georgia"] = 10310371
// 	// delete(statePopulations, "Georgia")
// 	// fmt.Println(statePopulations)
// 	// fmt.Println(statePopulations["Georgia"])

// 	// pop, ok := statePopulations["Oho"]
// 	// fmt.Println(pop, ok)
// 	// pop, ok = statePopulations["Ohio"]
// 	// fmt.Println(pop, ok)

// 	// fmt.Println(len(statePopulations))

// 	// sp := statePopulations
// 	// delete(sp, "Ohio")
// 	// fmt.Println(sp)
// 	// fmt.Println(statePopulations)

//     // m := make(map[string]int)

//     // m["k1"] = 7
//     // m["k2"] = 13

//     // fmt.Println("map:", m)

//     // v1 := m["k1"]
//     // fmt.Println("v1: ", v1)

//     // fmt.Println("len:", len(m))

//     // delete(m, "k2")
//     // fmt.Println("map:", m)

//     // _, prs := m["k2"]
//     // fmt.Println("prs:", prs)

//     // n := map[string]int{"foo": 1, "bar": 2}
//     // fmt.Println("map:", n)
// }

// Structs
// // type Doctor struct{
// // 	Number int
// // 	ActorName string
// // 	Companions []string
// // }
// type Doctor struct{
// 	number int
// 	actorName string
// 	companions []string
// }

// func main(){
// 	// aDoctor := Doctor{
// 	// 	number: 3,
// 	// 	actorName: "Jon Pertwee",
// 	// 	companions: []string {
// 	// 		"Liz Shaw",
// 	// 		"Jo Grant",
// 	// 		"Sarah Jane Smith",
// 	// 	},
// 	// }
// 	// // aDoctor := Doctor{
// 	// // 	3,
// 	// // 	"Jon Pertwee",
// 	// // 	[]string {
// 	// // 		"Liz Shaw",
// 	// // 		"Jo Grant",
// 	// // 		"Sarah Jane Smith",
// 	// // 	},
// 	// // }

// 	// fmt.Println(aDoctor)
// 	// fmt.Println(aDoctor.actorName)
// 	// fmt.Println(aDoctor.companions)
// 	// fmt.Println(aDoctor.companions[1])

// }

// type person struct {
//     name string
//     age  int
// }

// func main()  {
// 	//anonymous struct
// 	aDoctor := struct{name string}{name: "Jon Pertwee"}
// 	// anotherDoctor := aDoctor
// 	anotherDoctor := &aDoctor
// 	anotherDoctor.name = "Tom Baker"
// 	fmt.Println(aDoctor)
// 	fmt.Println(anotherDoctor)
// }

// func newPerson(name string) *person {

//     p := person{name: name}
//     p.age = 42
//     return &p
// }

// func main() {

//     fmt.Println(person{"Bob", 20})

//     fmt.Println(person{name: "Alice", age: 30})

//     fmt.Println(person{name: "Fred"})

//     fmt.Println(&person{name: "Ann", age: 40})

//     fmt.Println(newPerson("Jon"))

//     s := person{name: "Sean", age: 50}
//     fmt.Println(s.name)

//     sp := &s
//     fmt.Println(sp.age)

//     sp.age = 51
//     fmt.Println(sp.age)
// }


// Composition
// type Animal struct {
// 	Name string
// 	Origin string
// }

// type Bird struct {
// 	Animal
// 	SpeedKPH float32
// 	CanFly bool
// }

// func main()  {
// 	// b := Bird{}
// 	// b.Name = "Emu"
// 	// b.Origin = "Australia"
// 	// b.SpeedKPH = 48
// 	// b.CanFly = false

// 	b := Bird{
// 		Animal: Animal{Name: "Emu", Origin: "Australia"},
// 		SpeedKPH: 48,
// 		CanFly: false,
// 	}

// 	fmt.Println(b)
// 	fmt.Println(b.Name)
// 	fmt.Println(b.Origin)
// 	fmt.Println(b.Origin)
// 	fmt.Println(b.CanFly)
// }


import(
	"fmt"
	"reflect"
)

type Animal struct {
	Name string `required max:"100"`
	Origin string
}

func main()  {

	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName(("Name"))
	fmt.Println(field.Tag)
}