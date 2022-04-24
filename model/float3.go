package model

type Float3 struct {
	x, y, z float64
}

func (f *Float3) UnmarshalYAML(unmarshal func(interface{}) error) error {
	vals := [3]float64{}
	err := unmarshal(&vals)
	if err != nil {
		return err
	}
	f.x = vals[0]
	f.y = vals[1]
	f.z = vals[2]

	return nil
}
