package nanos

type Message struct {
	ResTo   chan<- Message
	ErrTo   chan<- error
	Content []byte
}
