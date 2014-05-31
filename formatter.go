package gomaplog

type Formatter interface {
	Format(event LogEvent) ([]byte, error)
}
