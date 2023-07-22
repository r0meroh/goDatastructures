package myStructs

func New_passenger(name string, ticket string, board bool) *Passenger {
	pass := Passenger{name, ticket, board}
	return &pass
}

func Print_name(pass *Passenger) string {
	return pass.Name
}

func Print_ticket(pass *Passenger) string {
	return pass.TicketNumber
}

func Is_boarded(pass *Passenger) bool {
	return pass.Boarded
}
