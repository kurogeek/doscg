package entity

type XYZ struct {
	X int
	Y int
	Z int
}

type BC struct {
	B int
	C int
}

type LatLng struct {
	Lat float64
	Lng float64
}

type Place struct {
	Name     string
	Location LatLng
}

type BestRoute struct {
	OriginLocation      Place
	DestinationLocation Place
	Polyline            []LatLng
}

type BotMessage struct {
	ReplyToken string
	Text       string
}
