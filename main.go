package main

import "fmt"

type Projectile struct {
	Position Tuple // point
	Velocity Tuple // vector
}

type Environment struct {
	gravity Tuple // vector
	wind    Tuple // vector
}

func tick(e Environment, p Projectile) Projectile {
	pos := tupleAdd(p.Position, p.Velocity)
	vel := tupleAdd(tupleAdd(p.Velocity, e.gravity), e.wind)

	return Projectile{pos, vel}
}

func main() {
	p := Projectile{point(0, 1, 0), tupleMultiply(vector(1, 1, 0), 0.2)}
	e := Environment{vector(0, -0.1, 0), vector(-0.01, 0, 0)}
	t := 0

	for p.Position.Y > 0.0 {
		t += 1
		fmt.Printf("tick: %d, pos: %v\n", t, p.Position)
		p = tick(e, p)
	}
}
