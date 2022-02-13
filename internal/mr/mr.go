package mr

import (
	"strings"
)

type Key string

type Value interface{}

// KeyValue
type KeyValue struct {
	Key   Key
	Value Value
}

// KeyValues
func KeyValues(kvs []KeyValue) map[Key][]Value {
	keyValues := make(map[Key][]Value)
	for _, kv := range kvs {
		values, ok := keyValues[kv.Key]
		if !ok {
			values = []Value{}
		}
		keyValues[kv.Key] = append(values, kv.Value)
	}
	return keyValues
}

// Map
func Map(_, content string) []KeyValue {
	kvs := []KeyValue{}
	for _, word := range strings.Fields(content) {
		kv := KeyValue{
			Key:   Key(word),
			Value: 1,
		}
		kvs = append(kvs, kv)
	}
	return kvs
}

// Reduce
func Reduce(_ Key, values []Value) int { return len(values) }
