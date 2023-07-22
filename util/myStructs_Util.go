package util

import (
	"GOPATH/myStructs"
	"fmt"
)

func Run_myStructs() {
	Hugo := myStructs.New_passenger("Hugo", "xbn1234", false)
	fmt.Println(myStructs.Print_name(Hugo))
	fmt.Println(myStructs.Print_ticket(Hugo))
	fmt.Println(myStructs.Is_boarded(Hugo))
}
