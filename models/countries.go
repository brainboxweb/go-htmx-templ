package models

type Country struct {
	Name        string
	Capital     string
	Population  int
	Continent   string
	Description string
}

var Countries = []Country{
	{
		Name:        "Japan",
		Capital:     "Tokyo",
		Population:  127500000,
		Continent:   "Asia",
		Description: "An island country known for its unique culture and technology.",
	},
	{
		Name:        "Brazil",
		Capital:     "Brasilia",
		Population:  214300000,
		Continent:   "South America",
		Description: "The largest country in South America.",
	},
	{
		Name:        "France",
		Capital:     "Paris",
		Population:  673900000,
		Continent:   "Europe",
		Description: "Known for its art, culture and cuisine.",
	},
}
