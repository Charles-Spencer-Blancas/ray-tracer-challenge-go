package main

import "math"

func translation(x float64, y float64, z float64) Matrix {
	id := matrixConstructIdentity(4)
	id.Values[0][3] = x
	id.Values[1][3] = y
	id.Values[2][3] = z

	return id
}

func scaling(x float64, y float64, z float64) Matrix {
	id := matrixConstructIdentity(4)
	id.Values[0][0] = x
	id.Values[1][1] = y
	id.Values[2][2] = z

	return id
}

func rotation_x(rads float64) Matrix {
	id := matrixConstructIdentity(4)

	id.Values[1][1] = math.Cos(rads)
	id.Values[1][2] = -math.Sin(rads)
	id.Values[2][1] = math.Sin(rads)
	id.Values[2][2] = math.Cos(rads)

	return id
}

func rotation_y(rads float64) Matrix {
	id := matrixConstructIdentity(4)

	id.Values[0][0] = math.Cos(rads)
	id.Values[2][0] = -math.Sin(rads)
	id.Values[0][2] = math.Sin(rads)
	id.Values[2][2] = math.Cos(rads)

	return id
}

func rotation_z(rads float64) Matrix {
	id := matrixConstructIdentity(4)

	id.Values[0][0] = math.Cos(rads)
	id.Values[0][1] = -math.Sin(rads)
	id.Values[1][0] = math.Sin(rads)
	id.Values[1][1] = math.Cos(rads)

	return id
}

func shearing(xy float64, xz float64, yx float64, yz float64, zx float64, zy float64) Matrix {
	m := matrixConstructIdentity(4)
	id := m.Values

	id[0][1] = xy
	id[0][2] = xz
	id[1][0] = yx
	id[1][2] = yz
	id[2][0] = zx
	id[2][1] = zy

	return m
}

func transformation(transforms ...Matrix) (Matrix, error) {
	out := matrixConstructIdentity(4)

	for _, t := range transforms {
		var err error
		out, err = matrix4x4Multiply(t, out)
		if err != nil {
			return Matrix{}, err
		}
	}

	return out, nil
}
