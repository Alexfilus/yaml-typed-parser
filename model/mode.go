package model

type Mode struct {
	TeamSize     U32                      `yaml:"team_size"`
	DurationS    U32                      `yaml:"duration_s"`
	Position     Float3                   `yaml:"position"`
	Requirements []map[string]Requirement `yaml:"requirements"`
	Levels       []string                 `yaml:"levels"`
}
