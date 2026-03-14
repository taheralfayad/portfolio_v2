package data

type Coffee struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	OriginCountry string  `json:"origin_country"`
	Processing    string  `json:"processing"`
	Varietal      string  `json:"varietal"`
	Description   string  `json:"description"`
	Roasts        []Roast `json:"roasts"`
}

type Roast struct {
	ID          int16  `json:"id"`
	RoastLevel  int16  `json:"roast_level"`
	RoasterName string `json:"roaster_name"`
	RoastDate   string `json:"roast_date"`
	Image       string `json:"image"`
	CoffeeID    int16  `json:"coffee_id"`
}

type CoffeeCup struct {
	ID             int    `json:"id"`
	RoastID        int    `json:"roast_id"`
	Temperature    int16  `json:"temperature"`
	DateDrank      string `json:"date_drank"`
	DaysAfterRoast int16  `json:"days_after_roast"`
	Acidity        int16  `json:"acidity"`
	Body           int16  `json:"body"`
	Sweetness      int16  `json:"sweetness"`
	WaterType      string `json:"water_type"`
	GrindSize      int16  `json:"grind_size"`
	Method         string `json:"method"`
	Rating         int16  `json:"rating"`
}
