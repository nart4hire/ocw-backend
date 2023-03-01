package cache

import "fmt"

type String struct {
	Key             Key
	Value           string
	ExpiryInMinutes int
}

type Hash struct {
	Key             Key
	Values          []Value
	ExpiryInMinutes int
}

type Key struct {
	Hash string
	Id   string
}

type Value struct {
	Field string
	Store string
}

func (c *Hash) AppendValue(value Value) *Hash {
	c.Values = append(c.Values, value)
	return c
}

func NewKey(hash string, id string) *Key {
	return &Key{hash, id}
}

func (k Key) String() string {
	return fmt.Sprintf("%s:%s", k.Hash, k.Id)
}

func NewValue(field string, store string) *Value {
	return &Value{field, store}
}

func NewString(key Key, value string, expiryInMinutes int) *String {
	return &String{key, value, expiryInMinutes}
}

func NewHash(key Key, initValue Value, expiryInMinutes int) *Hash {
	return &Hash{key, []Value{initValue}, expiryInMinutes}
}

func (c *Hash) Slice() []interface{} {
	slice := make([]interface{}, len(c.Values)*2+1)
	slice[0] = c.Key.String()
	for i := range make([]int, len(c.Values)) {
		slice[i*2+1] = c.Values[i].Field
		slice[i*2+2] = c.Values[i].Store
	}
	return slice
}
