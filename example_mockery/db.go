package example

type DB interface {
	Get(val string) string
}
