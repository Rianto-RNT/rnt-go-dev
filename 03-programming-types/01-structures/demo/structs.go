package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {
	ryan := Passenger{"Ryan", 1, false}
	fmt.Println(ryan)

	var (
		morrison = Passenger{Name: "Morrison", TicketNumber: 2}
		saprol   = Passenger{Name: "Saprol", TicketNumber: 3}
	)
	fmt.Println(morrison, saprol)

	var daeng Passenger
	daeng.Name = "Daeng"
	daeng.TicketNumber = 4
	fmt.Println(daeng)

	morrison.Boarded = true
	saprol.Boarded = true

	if morrison.Boarded {
		fmt.Println("Morrison has boarded the bus")
	}

	if saprol.Boarded {
		fmt.Println(saprol.Name, "has boarded the bus")
	}

	daeng.Boarded = true
	bus := Bus{daeng}
	fmt.Println(bus)
	fmt.Println(bus.FrontSeat.Name, "is in the front seat")
}
