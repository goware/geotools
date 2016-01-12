package main

import (
	"fmt"
	"log"

	"upper.io/db"
	"upper.io/db/postgresql"

	"github.com/pressly/geoy"
)

var settings = postgresql.ConnectionURL{
	Address: db.Host("localhost"),
}

type Hub struct {
	Id       int64         `db:"id,omitempty"`
	Name     string        `db:"name"`
	Location geoy.Geometry `db:"location"`
	Radius   int64         `db:"radius"`
}

type HubResult struct {
	Id       int64       `db:"id,omitempty"`
	Name     string      `db:"name"`
	Location *geoy.Point `db:"location"`
	Radius   int64       `db:"radius"`
	Distance float64     `db:"distance,omitempty"`
}

func handleError(err error) {
	if err != nil {
		log.Fatalf("%q\n", err)
	}
}

func main() {
	sess, err := db.Open(postgresql.Adapter, settings)
	handleError(err)
	defer sess.Close()

	hubCollection, err := sess.Collection("hubs")
	handleError(err)

	err = hubCollection.Truncate()
	handleError(err)

	hubs := []Hub{
		Hub{
			Name:     "Hipo",
			Location: geoy.NewPoint(28.986246, 41.05201),
		},
		Hub{
			Name:     "Liberty Village",
			Location: geoy.NewPoint(-79.4211567, 43.6373781),
			Radius:   10 * geoy.KM,
		},
		Hub{
			Name:     "Celbridge",
			Location: geoy.NewPoint(-6.5389921, 53.3392795),
			Radius:   25 * geoy.KM,
		},
	}

	for _, h := range hubs {
		_, err = hubCollection.Append(h)
		handleError(err)
	}

	// Sort hubs by distance from some specified point.
	p := geoy.NewPoint(-7.5, 53.1)
	res := hubCollection.Find()
	res = res.Select("id", "name", "location", geoy.ST_Distance("location", p, "distance")).Sort("distance")
	var results []HubResult
	err = res.All(&results)
	handleError(err)
	fmt.Printf("\nSort hubs by distance from %s\n", p)
	for _, h := range results {
		fmt.Printf("%s (%d) Hub is located at %s which is %0.1fkm from the point.\n", h.Name, h.Id, h.Location, h.Distance/geoy.KM)
	}

	// Sort hubs by distance from some specified point, considering their radius too.
	res = hubCollection.Find()
	res = res.Select("id", "name", "location", "radius", geoy.ST_Distance("st_buffer(location, GREATEST(radius, 1))", p, "distance")).Sort("distance")
	err = res.All(&results)
	handleError(err)
	fmt.Printf("\nSort hubs by distance from %s while considering radius of each hub\n", p)
	for _, h := range results {
		fmt.Printf("%s (%d) Hub is located at %s with radius %dkm which is %0.1fkm from the point.\n", h.Name, h.Id, h.Location, h.Radius/geoy.KM, h.Distance/geoy.KM)
	}

	// Find all hubs inside bbox
	bbox := geoy.NewEnvelope(-100.10, 55.10, 0.10, 0.10)
	res = hubCollection.Find(geoy.ST_Intersects("location", bbox))
	res = res.Select("id", "name", "location")
	err = res.All(&results)
	handleError(err)
	fmt.Printf("\nFind all hubs inside bbox %s\n", bbox)
	for _, h := range results {
		fmt.Printf("%s (%d) Hub is located at %s inside the bbox.\n", h.Name, h.Id, h.Location)
	}

}
