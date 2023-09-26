package helper

type Resturant struct {
	Name        string
	Rate        string
	ReviewCount string
	Address     string
	Menu        []Menu
}

type Menu struct {
	Category string
	Items    []Item
}

type Item struct {
	Name        string
	Description string
	OldPrice    string
	Price       string
	Currency    string
}
