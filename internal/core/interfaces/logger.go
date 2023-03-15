package interfaces

type Logger interface {
	Error(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
}
