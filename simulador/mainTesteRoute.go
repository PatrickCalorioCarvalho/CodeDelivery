package main
import (
		"fmt"
		route2 "github.com/patrickcaloriocarvalho/GeoCarMaps/application/route"
)

func main() {
	route := route2.Route{
		ID: "1",
		ClientID: "1",
	}
	route.LoadPositions()
	stringJson, _ := route.ExportJsonPositions()

	for _, v := range stringJson {
		fmt.Println(v)
	}
}