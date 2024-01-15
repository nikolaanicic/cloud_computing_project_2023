package mapper

type JsonModel interface {
	AsJson() []byte
}
