package tests

import (
	"fmt"
	"testing"
	"xr/xutor/utils"
)

func TestPkByEmail(t *testing.T) {
	emails := []string{
		"nikolaev666@aol.com",
		"nikolaev666@yahoo.com",
		"nnikolaev666@hotmail.com",
		"nnikolay666@gmail.com",
		"nikolaynikolay666@outlook.com",
		"dickson1488@yahoo.com",
		"ddickson1488@gmail.com",
		"nonameson666@yahoo.com",
		"sickdick666@yahoo.com",
		"upalashlyapa@aol.com",
		"dagestan666@yahoo.ca",
		"nagibator05@yahoo.ca",
		"dedpdd@yahoo.ca",
		"niceguy66613@yahoo.com",
		"niceguy66613@aol.com",
		"niceguy66613@hotmail.com",
		"niktester@hotmail.com",
		"nikolay.nikolaev@regium.com",
		"troitskiy.evgeniy@gmail.com",
		"evgeniy.troitskiy@iage.net",
		"kozlov.a.a.62@gmail.com",
		"kenny62rzn17@gmail.com",
		"activate@liveintent.com",
	}

	for _, eml := range emails {
		fmt.Println(eml, "\t\t", utils.GetMurmur3Int64PkStr(eml))
	}

}
