package main

import "math"

type boundingCircle struct {
	center vector
	radius float64
}

func collides(c1, c2 boundingCircle) bool {
	dist := math.Sqrt(math.Pow(c2.center.x-c1.center.x, 2) +
		math.Pow(c2.center.y-c1.center.y, 2))

	return dist <= c1.radius+c2.radius
}

func checkCollisions() error {
	for i := 0; i < len(elements)-1; i++ {
		collidableElement := elements[i]
		for j := i + 1; j < len(elements); j++ {
			movingElement := elements[j]
			boundingCircle1 := collidableElement.getBoundingCircle()
			boundingCircle2 := movingElement.getBoundingCircle()
			if collides(boundingCircle1, boundingCircle2) && *collidableElement.isActive() && *movingElement.isActive() {
				err := collidableElement.onCollision(movingElement)
				if err != nil {
					return err
				}
				err = movingElement.onCollision(collidableElement)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
