package base64

type Base64Util interface {
	Encode(input []byte) string
	Decode(input string) ([]byte, error)
}
