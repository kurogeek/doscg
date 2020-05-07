package main

import (
	"context"

	"github.com/kr/pretty"
	"googlemaps.github.io/maps"
)

func main() {
	gc, err := maps.NewClient(maps.WithAPIKey("AIzaSyBVVZgNZxWGcQqgC4H00a3hDA_-0FMZ2wk"))
	if err != nil {
		panic(err)
	}
	request := maps.DirectionsRequest{
		Origin:      "SCG สำนักงานใหญ่ บางซื่อ 1 Siam Cement Alley, Bang Sue, Bangkok 10800",
		Destination: "centralwOrld, 999/9 Rama I Rd, Pathum Wan, Pathum Wan District, Bangkok 10330",
	}
	route, waypoint, err := gc.Directions(context.Background(), &request)
	if err != nil {
		panic(err)
	}

	pretty.Println(route, waypoint)
}

type XYZ struct {
	X int
	Y int
	Z int
}

// FindXYZ - X, Y, 5, 9, 15, 23, Z  - Please create a new function for finding X, Y, Z value
// X, Y, 5, 9, 15, 23, Z
//  \/ \/ \/ \/  \/  \/
//   0  2  4  6  8   10
// So, Z = 23 + 10 = 33
// Y = 5 - 2 = 3
// X = Y - 0 = 3 - 0 = 3
func FindXYZ() XYZ {
	var xyz XYZ
	xyz.X = findX()
	xyz.Y = findY()
	xyz.Z = findZ()
	return xyz
}

func findX() int {
	return 3
}

func findY() int {
	return 3
}

func findZ() int {
	return 33
}

type BC struct {
	B int
	C int
}

func FindBC(ans1 int, ans2 int) BC {
	var a int = 21
	var bc BC

	bc.B = findNumber(a, ans1)
	bc.B = findNumber(a, ans2)

	return bc
}

// FindB - A = 21
func findNumber(a int, ans int) int {
	return ans - a
}
