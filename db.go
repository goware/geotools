package geoy

import (
	"fmt"
	"upper.io/db"
)

// STDistance returns the PostGIS function. See
// http://www.postgis.org/docs/ST_Distance.html.
func STDistance(sourceField string, g Geometry, asField string) db.Raw {
	q := fmt.Sprintf("ST_Distance(%s, '%s') as %s", sourceField, g.WKT(), asField)
	return db.Raw{q}
}

// STDWithin returns the PostGIS function. See
// http://www.postgis.org/docs/ST_DWithin.html.
func STDWithin(sourceField string, g Geometry, distance int) db.Raw {
	q := fmt.Sprintf("ST_DWithin(%s, '%s', %d)", sourceField, g.WKT(), distance)
	return db.Raw{q}
}

// STIntersects returns the PostGIS function. See
// http://www.postgis.org/docs/ST_Intersects.html.
func STIntersects(sourceField string, g Geometry) db.Raw {
	q := fmt.Sprintf("ST_Intersects(%s, '%s')", sourceField, g.WKT())
	return db.Raw{q}
}
