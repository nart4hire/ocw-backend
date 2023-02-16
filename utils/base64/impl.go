package base64

import "encoding/base64"

type Base64UtilImpl struct{}

func (Base64UtilImpl) Encode(input []byte) string {
	return base64.StdEncoding.EncodeToString(input)
}

func (Base64UtilImpl) Decode(input string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(input)
}
