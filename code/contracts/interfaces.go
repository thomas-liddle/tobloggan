package contracts

type Logger interface {
	Printf(format string, args ...interface{})
}
