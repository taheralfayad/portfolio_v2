package data

type CoffeeRequest struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	RoastLevel    int16  `json:"roast_level"`
	RoasterName   string `json:"roaster_name`
	OriginCountry string `json:"origin_country"`
	Processing    string `json:"processing"`
	Varietal      string `json:"varietal"`
	Image         string `json:"image"`
	Date          string `json:"date"`
	Description   string `json:"description"`
}

type CoffeeResponse struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	RoastLevel    int16  `json:"roast_level"`
	RoasterName   string `json:"roaster_name`
	OriginCountry string `json:"origin_country"`
	Processing    string `json:"processing"`
	Varietal      string `json:"varietal"`
	Image         string `json:"image"`
	Date          string `json:"date"`
	Description   string `json:"description"`
}

type CoffeeCup struct {
	ID             string `json:"id"`
	Temperature    int16  `json:"temperature"`
	DaysAfterRoast int16  `json:"days_after_roast"`
	Acidity        int16  `json:"acidity"`
	Body           int16  `json:"body"`
	Sweetness      int16  `json:"sweetness"`
	WaterType      string `json:"water_type"`
	GrindSize      int16  `json:"grind_size"`
	Method         string `json:"method"`
	Rating         int16  `json:"rating"`
}
