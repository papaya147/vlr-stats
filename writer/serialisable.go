package writer

type Serialisable interface {
	GetHeaders() []string
	GetRecords() [][]string
}
