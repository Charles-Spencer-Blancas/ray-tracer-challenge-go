package main

type World struct {
	Objects []Sphere
	Light   PointLight
}

type Computation struct {
	Object  Sphere
	t       float64
	Point   Tuple
	EyeV    Tuple
	NormalV Tuple
}

func defaultWorld() (World, error) {
	l, err := pointLight(point(-10, 10, -10), Color{1, 1, 1})
	if err != nil {
		return World{}, err
	}
	s1 := sphere()
	m := material()
	m.Color = Color{0.8, 1.0, 0.6}
	m.Diffuse = 0.7
	m.Specular = 0.2
	s1.Material = m

	s2 := sphere()
	s2.Transform = scaling(0.5, 0.5, 0.5)

	return World{[]Sphere{s1, s2}, l}, nil
}

func worldRayIntersect(w World, r Ray) ([]Intersection, error) {
	intersections := []Intersection{}
	for _, s := range w.Objects {
		is, err := sphereRayIntersect(s, r)
		if err != nil {
			return []Intersection{}, err
		}
		intersections = append(intersections, is...)
	}

	return sortIntersections(intersections), nil
}

func prepareComputations(i Intersection, r Ray) (Computation, error) {
	p := rayPosition(r, i.t)
	n, err := sphereNormalAt(i.Object, p)
	if err != nil {
		return Computation{}, nil
	}

	return Computation{i.Object, i.t, p, tupleNegate(r.Direction), n}, nil
}
