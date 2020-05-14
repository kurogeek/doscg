package entity

type XYZ struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

type BC struct {
	B int `json:"b"`
	C int `json:"c"`
}

type LatLng struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Place struct {
	Name     string `json:"name"`
	Location LatLng `json:"location"`
}

type BestRoute struct {
	OriginLocation      Place    `json:"originLocation"`
	DestinationLocation Place    `json:"destinationLocation"`
	Polyline            []LatLng `json:"polyline"`
}

type BotMessage struct {
	ReplyToken string
	Text       string
}
