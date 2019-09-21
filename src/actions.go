package main

type Sphere struct {
	Radius float64
}

func (s *Sphere) SurfaceArea() float64{
	return s.Radius
}

func (s *Sphere) Volumn() float64{
	return s.Radius * s.Radius
}

