package main

import(
	"fmt"
)

const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main()  {
	var roles byte = canSeeFinancials | canSeeEurope | isAdmin
	fmt.Printf("%b \n", roles )	
	fmt.Printf("%v \n", isAdmin & roles  == isAdmin )
}