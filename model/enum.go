package model

import "errors"

var ErrValueNotInENUM = errors.New("error value not in enum")

type Enum int

const (
	Undefined Enum = iota
	GameModes
	GameLevels
	GameCharacters
)

func (e Enum) String() string {
	switch e {
	case GameModes:
		return "GameModes"
	case GameLevels:
		return "GameLevels"
	case GameCharacters:
		return "GameCharacters"
	default:
		return "unknown"
	}
}

func (e *Enum) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	err := unmarshal(&s)
	if err != nil {
		return err
	}

	switch s {
	case "GameModes":
		*e = GameModes
	case "GameLevels":
		*e = GameLevels
	case "GameCharacters":
		*e = GameCharacters
	default:
		return ErrValueNotInENUM
	}

	return nil
}
