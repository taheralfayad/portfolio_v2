package data

type BookMeta struct {
	Title            string  `json:"title"`
	Authors          string  `json:"authors"`
	Status           string  `json:"status,omitempty"`
	LastReadTime     string  `json:"last_read_time"`
	PercentFinished  float64 `json:"percent_finished,omitempty"`
	CurrentlyReading bool
}

type DeviceEntry struct {
	File          string `json:"file"`
	Text          string `json:"text"`
	Dim           bool   `json:"dim"`
	MandatoryTime string `json:"mandatory"`
	SelectEnabled bool   `json:"select_enabled"`
	Time          int64  `json:"time"`
}

type PostBooksPayload struct {
	CurrentlyReading DeviceEntry         `json:"currentlyReading"`
	Library          map[string]BookMeta `json:"library"`
}
