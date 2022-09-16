package main

import "fmt"

// /*
// Input = [
//   {"id":1, "item": 2, "city": "Bengaluru"},
//   {"id":2, "item": 4, "city": "Mysore"},
//   {"id":1, "item": 5, "city": "Unnao"},
//   {"id":2, "item": 6, "city": "Agra"},
//   {"id":1, "item": 7, "city": "Kanpur"},
//   {"id":3, "item": 8, "city": "Ooty"}
// ]

// Output = [
//   {"id":3, "item": 8, "city": "Ooty"},
//   {"id":2, "item": 6, "city": "Agra"},
//   {"id":1, "item": 7, "city": "Kanpur"}
// ]
// zfw-jgee-abk
// */

type input struct {
	id   int
	item int
	city string
}
type inputs []input

func main() {

	in := inputs{{id: 1, item: 2, city: "mumbai"}, {id: 2, item: 5, city: "delhi"}, {id: 2, item: 6, city: "chenai"}}
	fmt.Println("input :-> ", in)
	out := inputs{}
	temp := map[int]input{}
	//result := inputs{}
	for _, value := range in {
		if temp[value.id].id == 0 {
			fmt.Println("in")
			item := input{id: value.id, item: value.item, city: value.city}
			temp[value.id] = item
			out = append(out, item)
		} else {
			item := input{id: value.id, item: value.item, city: value.city}
			//temp[value.id] = item
			out = append(out, item)

		}
	}
	//append(result)
	fmt.Println("Output = ", out)

}
