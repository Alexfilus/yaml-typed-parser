package model

import "strconv"

type U32 uint32

func (u *U32) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	err := unmarshal(&s)
	if err != nil {
		return err
	}
	val, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return err
	}
	*u = U32(val)

	return nil
}
