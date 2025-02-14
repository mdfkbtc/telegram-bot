package cmd

import (
	"regexp"
	"strconv"
	"testing"

	"github.com/stretchr/testify/suite"
)

type RunTestSuite struct {
	suite.Suite
}

func (suite *RunTestSuite) SetupTest() {
}

func (suite *RunTestSuite) TestCommandPrice() {
	tests := map[string]struct {
		contains string
		error    bool
	}{
		"btc": {
			contains: "Bitcoin",
			error:    false,
		},
		"Bitcoin": {
			contains: "Bitcoin",
			error:    false,
		},
		"Invalid_coin": {
			contains: "",
			error:    true,
		},
	}

	for s, r := range tests {
		result, err := commandPrice(s)

		if !r.error {
			suite.NoError(err)
		} else {
			suite.Error(err)
		}

		if r.contains != "" {
			suite.Contains(result, r.contains)
		}
	}
}

func (suite *RunTestSuite) TestCommandAthPrice() {
	tests := map[string]struct {
		contains string
		error    bool
	}{
		"vls": {
			contains: "Veles",
			error:    false,
		},
		"Veles": {
			contains: "Veles",
			error:    false,
		},
		"Invalid_coin": {
			contains: "",
			error:    true,
		},
	}

	for s, r := range tests {
		result, err := commandAthPrice(s)

		if !r.error {
			suite.NoError(err)
		} else {
			suite.Error(err)
		}

		if r.contains != "" {
			suite.Contains(result, r.contains)
		}
	}
}


func (suite *RunTestSuite) TestCommandSupply() {
	tests := map[string]struct {
		error bool
	}{
		"btc": {
			error: false,
		},
		"Bitcoin": {
			error: false,
		},
		"Invalid_coin": {
			error: true,
		},
	}

	for s, r := range tests {
		response, err := commandSupply(s)

		if !r.error {
			suite.NoError(err)

			supplyRe := regexp.MustCompile(`(\d+)`)
			supplyStrings := supplyRe.FindStringSubmatch(response)
			suite.NotEmpty(supplyStrings)

			supplyInt, err := strconv.Atoi(supplyStrings[0])
			suite.NoError(err)
			suite.True(supplyInt > 100, "should be more than 100")
		} else {
			suite.Error(err)
		}
	}
}

func (suite *RunTestSuite) TestCommandPriceChange() {
	tests := map[string]struct {
		contains string
		error    bool
	}{
		"btc": {
			contains: "Bitcoin",
			error:    false,
		},
		"Bitcoin": {
			contains: "Bitcoin",
			error:    false,
		},
		"Invalid_coin": {
			contains: "",
			error:    true,
		},
	}

	for s, r := range tests {
		result, err := commandPriceChange(s)

		if !r.error {
			suite.NoError(err)
		} else {
			suite.Error(err)
		}

		if r.contains != "" {
			suite.Contains(result, r.contains)
		}
	}
}
/*
func (suite *RunTestSuite) TestCommandVolume() {
	tests := map[string]struct {
		error bool
	}{
		"btc": {
			error: false,
		},
		"Bitcoin": {
			error: false,
		},
		"Invalid_coin": {
			error: true,
		},
	}

	for s, r := range tests {
		response, err := commandVolume(s)

		if !r.error {
			suite.NoError(err)

			volumeRe := regexp.MustCompile(`[0-9]*[.][0-9]+`)
			volumeStrings := volumeRe.FindStringSubmatch(response)
			suite.NotEmpty(volumeStrings)

			volumeFloat, err := strconv.ParseFloat(volumeStrings[0], 64)
			suite.NoError(err)
			suite.True(volumeFloat > 100, "should be more than 100")
		} else {
			suite.Error(err)
		}
	}
}
*/

func TestRunTestSuite(t *testing.T) {
	suite.Run(t, new(RunTestSuite))
}
