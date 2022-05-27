package main

import "fmt"

// Flight - a struct that
// contains information about flights
type Flight struct {
	Origin      string
	Destination string
	Price       int
}

// SortByPrice sorts flights from highest to lowest
func SortByPrice(flights []Flight) []Flight {
	// implement
	for k := 0; k < len(flights)-1; k++ {
		for key, _ := range flights {
			if key < len(flights)-1 {

				if flights[key].Price < flights[key+1].Price {
					temp := flights[key+1]
					flights[key+1] = flights[key]
					flights[key] = temp
				}
			}
		}
	}

	return flights
}

func printFlights(flights []Flight) {
	for _, flight := range flights {
		fmt.Printf("Origin: %s, Destination: %s, Price: %d", flight.Origin, flight.Destination, flight.Price)
	}
}

func main() {
	// an empty slice of flights
	var flights []Flight
	flights = []Flight{
		{Origin: "Mumbai", Destination: "Delhi", Price: 4000},
		{Origin: "Mumbai", Destination: "Bangluru", Price: 5000},
		{Origin: "Mumbai", Destination: "Chenai", Price: 6000},
	}
	fmt.Println(flights)
	sortedList := SortByPrice(flights)
	printFlights(sortedList)
}

func A() string {
	return "hello"
}

func A() int{
  return 10
}

type user struct{}
func(u user) B()string{
  return "world"
}

type data string
func(d data) B()int{
  return 33
}