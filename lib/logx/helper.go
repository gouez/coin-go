package logx

var logger Logger = *NewLogger()

func Info(msg string, params ...any) {
	logger.Info(msg, params...)
}
func Error(msg string, params ...any) {
	logger.Error(msg, params...)
}
func Warn(msg string, params ...any) {
	logger.Warn(msg, params...)
}
func Debug(msg string, params ...any) {
	logger.Debug(msg, params...)
}
