package entity

import "errors"

var (
	NoRouteError = errors.New("mapservice: no route found")
	NoLegsFound = errors.New("mapservice: no legs found")
)