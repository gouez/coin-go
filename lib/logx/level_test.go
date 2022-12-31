package logx

import "testing"

func TestLevelString(t *testing.T) {
	levels := map[level]string{
		debugLevel: "debug",
		infoLevel:  "info",
		warnLevel:  "warn",
		errorLevel: "error",
	}

	for lvl, name := range levels {
		if lvl.String() != name {
			t.Errorf("level's name %s is wrong", lvl.String())
		}
	}
}
