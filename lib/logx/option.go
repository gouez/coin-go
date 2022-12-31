package logx

import "io"

type Option func(logger *Logger)

func (o Option) Apply(logger *Logger) {
	o(logger)
}

type options struct{}

func Options() *options {
	return (*options)(nil)
}

func (o *options) WithTimeFormat(format string) Option {
	return func(logger *Logger) {
		logger.timeFormat = format
	}
}

func (o *options) WithWriter(writer io.Writer) Option {
	return func(logger *Logger) {
		logger.Writer = writer
	}
}
