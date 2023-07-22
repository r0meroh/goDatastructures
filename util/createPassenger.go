package util

import "GOPATH/myStructs"

func Create_passenger(name string, ticketNumber string, boarded bool) *myStructs.Passenger {
	passenger := new(myStructs.Passenger)
	passenger.Name = name
	passenger.TicketNumber = ticketNumber
	// passenger.Boarded is not defined because default value for a bool is false
	return passenger
}
