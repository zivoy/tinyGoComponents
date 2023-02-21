package rotation

import (
	"github.com/zivoy/tinyGoComponents/utils"
	"tinygo.org/x/drivers/lis3dh"
)

const shrinkFactor int32 = 1_000

type Reference struct {
	reference utils.Plane
}

func NewReference(normal, up utils.Vector3) Reference {
	return Reference{reference: utils.Plane{Normal: normal, Up: up.Normalized()}}
}

func readNormalized(accelerometer lis3dh.Device) (utils.Vector3, error) {
	xi, yi, zi, err := accelerometer.ReadAcceleration()
	if err != nil {
		return utils.Vector3{}, nil
	}

	vector := utils.Vector3{
		X: float64(xi / shrinkFactor),
		Y: float64(yi / shrinkFactor),
		Z: float64(zi / shrinkFactor),
	}
	return vector.Normalized(), nil
}

func (r Reference) Read(accelerometer lis3dh.Device) (angle, magnitude float64, err error) {
	v, err := readNormalized(accelerometer)
	if err != nil {
		return 0, 0, err
	}

	angle, magnitude = r.reference.GetRotation(v)
	return
}
