package main


type ColorGroup struct {
	ID	int
	Name	string
	Colors	[]string
}

group := ColorGroup{
	ID:	1,
	Name: "tanjunchen",
	Colors: []string{"AA","BB","CC","DD"},
}

b,err := json