package geos

import (
	"github.com/spatial-go/geos/geo"
)

type GEOSAlgorithm struct{}

// Returns the area of a polygonal geometry
func (G GEOSAlgorithm) Area(g Geometry) (float64, error) {
	s := MarshalString(g)
	return geo.Area(s)
}

// Returns the closure of the combinatorial boundary of this Geometry
func (G GEOSAlgorithm) Boundary(g Geometry) (Geometry, error) {
	s := MarshalString(g)
	wkt, e := geo.Boundary(s)
	if e != nil {
		return nil, e
	}
	geometry, e := UnmarshalString(wkt)
	if e != nil {
		return nil, e
	}
	return geometry, nil
}
// Computes the geometric center of a geometry, or equivalently, the center of mass of the geometry as a POINT.
// For [MULTI]POINTs, this is computed as the arithmetic mean of the input coordinates.
// For [MULTI]LINESTRINGs, this is computed as the weighted length of each line segment.
// For [MULTI]POLYGONs, "weight" is thought in terms of area.
// If an empty geometry is supplied, an empty GEOMETRYCOLLECTION is returned.
// If NULL is supplied, NULL is returned.
// If CIRCULARSTRING or COMPOUNDCURVE are supplied, they are converted to linestring wtih CurveToLine first, then same than for LINESTRING
func (G GEOSAlgorithm) Centroid(g Geometry) (Geometry, error) {
	s := MarshalString(g)
	centroid, e := geo.Centroid(s)
	if e != nil {
		return nil, e
	}
	geometry, e := UnmarshalString(centroid)
	if e != nil {
		return nil, e
	}
	return geometry, nil
}
// Returns true if this Geometry has no anomalous geometric points, such as self intersection or self tangency.
func (G GEOSAlgorithm) IsSimple(g Geometry) (bool, error) {
	s := MarshalString(g)
	return geo.IsSimple(s)
}
// returns the 2D Cartesian length of the geometry if it is a LineString, MultiLineString
func (G GEOSAlgorithm) Length(g Geometry) (float64, error) {
	s := MarshalString(g)
	return geo.Length(s)
}
// returns the minimum 2D Cartesian (planar) distance between two geometries, in projected units (spatial ref units).
func (G GEOSAlgorithm) Distance(g1 Geometry, g2 Geometry) (float64, error) {
	geom1 := MarshalString(g1)
	geom2 := MarshalString(g2)
	return geo.Distance(geom1, geom2)
}

// Returns the Hausdorff distance between two geometries, a measure of how similar or dissimilar 2 geometries are.
// Implements algorithm for computing a distance metric which can be thought of as the "Discrete Hausdorff Distance".
// This is the Hausdorff distance restricted to discrete points for one of the geometries
func (G GEOSAlgorithm) HausdorffDistance(g1 Geometry, g2 Geometry) (float64, error) {
	geom1 := MarshalString(g1)
	geom2 := MarshalString(g2)
	return geo.HausdorffDistance(geom1, geom2)
}

// Returns true if this Geometry is an empty geometry.
// If true, then this Geometry represents an empty geometry collection, polygon, point etc.
func (G GEOSAlgorithm) IsEmpty(g Geometry) (bool, error) {
	wkt := MarshalString(g)
	return geo.IsEmpty(wkt)
}

// Returns the  minimum bounding box for the supplied geometry, as a geometry.
// The polygon is defined by the corner points of the bounding box ((MINX, MINY), (MINX, MAXY), (MAXX, MAXY), (MAXX, MINY), (MINX, MINY)).
func (G GEOSAlgorithm) Envelope() (*Geometry, error) {
	panic("implement me")
}

//Computes the convex hull of a geometry. The convex hull is the smallest convex geometry that encloses all geometries in the input.
//In the general case the convex hull is a Polygon. The convex hull of two or more collinear points is a two-point LineString. The convex hull of one or more identical points is a Point.
func (G GEOSAlgorithm) ConvexHull() (*Geometry, error) {
	panic("implement me")
}

//UnaryUnion does dissolve boundaries between components of a multipolygon (invalid) and does perform union between the components of a geometrycollection
func (G GEOSAlgorithm) UnaryUnion() (*Geometry, error) {
	panic("implement me")
}

// Returns a POINT guaranteed to intersect a surface.
func (G GEOSAlgorithm) PointOnSurface() (*Geometry, error) {
	panic("implement me")
}

// Returns a (set of) LineString(s) formed by sewing together the constituent line work of a MULTILINESTRING.
func (G GEOSAlgorithm) LineMerge() (*Geometry, error) {
	panic("implement me")
}

// Returns a "simplified" version of the given geometry using the Douglas-Peucker algorithm,May not preserve topology
func (G GEOSAlgorithm) Simplify(tolerance float64) (*Geometry, error) {
	panic("implement me")
}

// SimplifyP returns a geometry simplified by amount given by tolerance.
// Unlike Simplify, SimplifyP guarantees it will preserve topology.
func (G GEOSAlgorithm) SimplifyP(tolerance float64) (*Geometry, error) {
	panic("implement me")
}

// Returns a geometry that represents the point set intersection of the Geometries.
func (G GEOSAlgorithm) Intersection(other *Geometry) (*Geometry, error) {
	panic("implement me")
}

// Returns a geometry that represents that part of geometry A that does not intersect with geometry B.
// One can think of this as GeometryA - Intersection(A,B).
// If A is completely contained in B then an empty geometry collection is returned.
func (G GEOSAlgorithm) Difference(other *Geometry) (*Geometry, error) {
	panic("implement me")
}

// Returns a geometry that represents the portions of A and B that do not intersect.
// It is called a symmetric difference because SymDifference(A,B) = SymDifference(B,A).
// One can think of this as Union(geomA,geomB) - Intersection(A,B).
func (G GEOSAlgorithm) SymDifference(other *Geometry) (*Geometry, error) {
	panic("implement me")
}

// Union returns a new geometry representing all points in this geometry and the other.
func (G GEOSAlgorithm) Union(other *Geometry) (*Geometry, error) {
	panic("implement me")
}

//Overlaps, Touches, Within all imply geometries are not spatially disjoint.
// If any of the aforementioned returns true, then the geometries are not spatially disjoint.
// Disjoint implies false for spatial intersection.
func (G GEOSAlgorithm) Disjoint(other *Geometry) (bool, error) {
	panic("implement me")
}

//Returns TRUE if the only points in common between g1 and g2 lie in the union of the boundaries of g1 and g2.
// The ouches relation applies to all Area/Area, Line/Line, Line/Area, Point/Area and Point/Line pairs of relationships, but not to the Point/Point pair.
func (G GEOSAlgorithm) Touches(other *Geometry) (bool, error) {
	panic("implement me")
}

// If a geometry  shares any portion of space then they intersect
func (G GEOSAlgorithm) Intersects(other *Geometry) (bool, error) {
	panic("implement me")
}

// Returns TRUE if the Geometries "spatially overlap". By that we mean they intersect, but one does not completely contain another.
func (G GEOSAlgorithm) Overlaps(other *Geometry) (bool, error) {
	panic("implement me")
}
// Returns TRUE if the given Geometries are "spatially equal".
func (G GEOSAlgorithm) Equals(other *Geometry) (bool, error) {
	panic("implement me")
}
// Returns TRUE if no point in Geometry B is outside Geometry A
func (G GEOSAlgorithm) Covers(other *Geometry) (bool, error) {
	panic("implement me")
}
// Returns TRUE if no point in Geometry A is outside Geometry B
func (G GEOSAlgorithm) CoveredBy(other *Geometry) (bool, error) {
	panic("implement me")
}
// IsRing returns true if the lineal geometry has the ring property.
func (G GEOSAlgorithm) IsRing() (bool, error) {
	panic("implement me")
}

func (G GEOSAlgorithm) HasZ() (bool, error) {
	panic("implement me")
}
// Returns TRUE if the LINESTRING's start and end points are coincident.
// For Polyhedral Surfaces, reports if the surface is areal (open) or volumetric (closed).
func (G GEOSAlgorithm) IsClosed() (bool, error) {
	panic("implement me")
}

func (G GEOSAlgorithm) SRID() (int, error) {
	panic("implement me")
}

func (G GEOSAlgorithm) SetSRID(srid int) {
	panic("implement me")
}

func (G GEOSAlgorithm) NGeometry() (int, error) {
	panic("implement me")
}

// Returns a geometry that represents all points whose distance from this Geometry is less than or equal to distance.
func (G GEOSAlgorithm) Buffer(g Geometry, width float64, quadsegs int32) Geometry {
	panic("implement me")
}

func (G GEOSAlgorithm) EqualsExact(g1 Geometry, g2 Geometry, tolerance float64) bool {
	panic("implement me")
}

func (G GEOSAlgorithm) HausdorffDistanceDensify(s Geometry, d Geometry, densifyFrac float64) (float64, error) {
	panic("implement me")
}
// Relate computes the intersection matrix (Dimensionally Extended
// Nine-Intersection Model (DE-9IM) matrix) for the spatial relationship between
// the two geometries.
func (G GEOSAlgorithm) Relate(s Geometry, d Geometry, ) {
	panic("implement me")
}


// Crosses takes two geometry objects and returns TRUE if their intersection "spatially cross",
// that is, the geometries have some, but not all interior points in common.
// The intersection of the interiors of the geometries must not be the empty set and must have a dimensionality less than the maximum dimension of the two input geometries.
// Additionally, the intersection of the two geometries must not equal either of the source geometries. Otherwise, it returns FALSE.
func (G GEOSAlgorithm) Crosses(g1 Geometry, g2 Geometry) (bool, error) {
	geom1 := MarshalString(g1)
	geom2 := MarshalString(g2)
	return geo.Crosses(geom1, geom2)
}

//Returns TRUE if geometry A is completely inside geometry B.
// For this function to make sense, the source geometries must both be of the same coordinate projection,
// having the same SRID.
func (G GEOSAlgorithm) Within(g1 Geometry, g2 Geometry) (bool, error) {
	geom1 := MarshalString(g1)
	geom2 := MarshalString(g2)
	return geo.Within(geom1, geom2)
}

// Geometry A contains Geometry B if and only if no points of B lie in the exterior of A,
// and at least one point of the interior of B lies in the interior of A.
// An important subtlety of this definition is that A does not contain its boundary, but A does contain itself.
//Returns TRUE if geometry B is completely inside geometry A.
// For this function to make sense, the source geometries must both be of the same coordinate projection, having the same SRID.
func (G GEOSAlgorithm) Contains(g1 Geometry, g2 Geometry) (bool, error) {
	geom1 := MarshalString(g1)
	geom2 := MarshalString(g2)
	return geo.Contains(geom1, geom2)
}

// UniquePoints return all distinct vertices of input geometry as a MultiPoint.
func (G GEOSAlgorithm) UniquePoints(g Geometry) (Geometry, error) {
	geom := MarshalString(g)
	wkt, e := geo.UniquePoints(geom)
	if e != nil {
		return nil, e
	}
	geometry, e := UnmarshalString(wkt)
	if e != nil {
		return nil, e
	}
	return geometry, nil

}

// Returns a collection containing paths shared by the two input geometries.
// Those going in the same direction are in the first element of the collection, those going in the opposite direction are in the second element.
// The paths themselves are given in the direction of the first geometry.
func (G GEOSAlgorithm) SharedPaths(g1 Geometry, g2 Geometry) (string, error) {
	geom1 := MarshalString(g1)
	geom2 := MarshalString(g2)
	s, e := geo.SharedPaths(geom1, geom2)
	if e != nil {
		return "", e
	}
	return s, nil
}

// Snaps the vertices and segments of a geometry to another Geometry's vertices.
// A snap distance tolerance is used to control where snapping is performed.
// The result geometry is the input geometry with the vertices snapped.
// If no snapping occurs then the input geometry is returned unchanged.
func (G GEOSAlgorithm) Snap(input Geometry, reference Geometry, tolerance float64) (Geometry, error) {
	inGeom := MarshalString(input)
	refGeom := MarshalString(reference)
	s, e := geo.Snap(inGeom, refGeom, tolerance)
	if e != nil {
		return nil, e
	}
	geometry, e := UnmarshalString(s)
	if e != nil {
		return nil, e
	}
	return geometry, nil
}
