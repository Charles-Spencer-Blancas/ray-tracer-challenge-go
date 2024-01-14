package main

func sphereNormalAt(s Sphere, p Tuple) (Tuple, error) {
	inv, err := matrixInverse(s.Transform)
	if err != nil {
		return Tuple{}, err
	}
	objectPoint, err := matrix4x4TupleMultiply(inv, p)
	if err != nil {
		return Tuple{}, err
	}
	// Get the normal in object space
	objectNormal := tupleSubtract(objectPoint, point(0, 0, 0))
	// Convert the normal from object to world space
	worldNormal, err := matrix4x4TupleMultiply(matrixTranspose(inv), objectNormal)
	if err != nil {
		return Tuple{}, err
	}
	worldNormal.W = 0
	return vectorNormalize(worldNormal), nil
}
