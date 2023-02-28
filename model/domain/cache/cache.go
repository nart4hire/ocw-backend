package cache

import "fmt"

type Cache struct {
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

func (c *Cache) AppendValue(value Value) *Cache {
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

func NewCache(key Key, initValue Value, expiryInMinutes int) *Cache {
	return &Cache{key, []Value{initValue}, expiryInMinutes}
}

func (c *Cache) Slice() ([]interface{}) {
	slice := make([]interface{}, len(c.Values) * 2 + 1)
	slice[0] = c.Key.String()
	for i := range make([]int, len(c.Values)) {
		slice[i * 2 + 1] = c.Values[i].Field
		slice[i * 2 + 2] = c.Values[i].Store
	}
	return slice
}
