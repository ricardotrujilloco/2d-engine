package main

type boundingCircleScaler struct {
	boundingCircles    []*boundingCircle
	maxRadius          float64
	isMaxRadiusReached bool
}

func newBoundingCircleScaler(
	boundingCircles []*boundingCircle,
	maxRadius float64,
) *boundingCircleScaler {
	var scaler boundingCircleScaler
	scaler.boundingCircles = boundingCircles
	scaler.maxRadius = maxRadius
	return &scaler
}

func (scaler *boundingCircleScaler) onUpdate(parameters updateParameters) error {
	if !scaler.isMaxRadiusReached {
		for _, circle := range scaler.boundingCircles {
			if circle.radius < scaler.maxRadius {
				circle.radius += explosionSpeed * parameters.elapsed
			} else {
				scaler.isMaxRadiusReached = true
			}
		}
	}
	return nil
}
