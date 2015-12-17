package geoy

import (
	"fmt"
	"upper.io/db"
)

func ST_Distance(source_field string, g Geometry, as_field string) db.Raw {
	q := fmt.Sprintf("ST_Distance(%s, '%s') as %s", source_field, g.WKT(), as_field)
	return db.Raw{q}
}

func ST_DWithin(source_field string, g Geometry, distance int) db.Raw {
	q := fmt.Sprintf("ST_DWithin(%s, '%s', %d)", source_field, g.WKT(), distance)
	return db.Raw{q}
}

func ST_Intersects(source_field string, g Geometry) db.Raw {
	q := fmt.Sprintf("ST_Intersects(%s, '%s')", source_field, g.WKT())
	return db.Raw{q}
}
