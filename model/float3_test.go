package model

import (
	"testing"

	"gopkg.in/yaml.v2"

	"github.com/stretchr/testify/assert"
)

func TestFloat3_UnmarshalYAML(t *testing.T) {
	testCases := []testCase{
		{
			Input:       []byte("!<Float3> [1.0, 2.0, 3.0]"),
			Expected:    Float3{x: 1, y: 2, z: 3},
			IsErrorCase: false,
		},
		{
			Input:       []byte("!<Float3> [1.0, 2.0, 3.0, 4.0]"),
			IsErrorCase: true,
		},
		{
			Input:       []byte("!<Float3> [\"1\",\"2\",\"3\"]"),
			IsErrorCase: true,
		},
	}
	for _, tc := range testCases {
		f := Float3{}
		err := yaml.Unmarshal(tc.Input, &f)
		if tc.IsErrorCase {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, tc.Expected, f)
		}
	}
}

type testCase struct {
	Input       []byte
	Expected    Float3
	IsErrorCase bool
}

const good = `
header:
  type: !<Enum> GameModes
content:
  duel: !<Mode>
    team_size: !<U32> 1
    duration_s: !<U32> 60
    position: !<Float3> [1.0, 2.0, 3.0]
    requirements:
      - Player: !<Requirement>
          min_battles: !<U32> 4
      - Squad: !<Requirement>
          min_battles: !<U32> 4
    levels:
      - levels/level_1
      - levels/level_2
      - levels/level_3`
