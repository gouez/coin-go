package logx

type level uint8

const (
	debugLevel level = iota
	infoLevel
	warnLevel
	errorLevel
)

var (
	levels = map[level]string{
		debugLevel: "debug",
		infoLevel:  "info",
		warnLevel:  "warn",
		errorLevel: "error",
	}
)

func (l level) String() string {
	return levels[l]
}
