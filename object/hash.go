package object

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

type HashKey struct {
	Type  Type
	Value uint64
}

type HashPair struct {
	Key   Object
	Value Object
}

type Hash struct {
	Pairs  map[HashKey]HashPair
	offset int
}

func (h *Hash) Type() Type { return HASH }
func (h *Hash) String() string {
	var out bytes.Buffer
	var pairs []string
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s",
			pair.Key.String(), pair.Value.String()))
	}
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	return out.String()
}
func (h *Hash) Next() (Object, Object) {
	idx := 0
	hash := make(map[string]HashPair)
	var keys []string
	for _, v := range h.Pairs {
		hash[v.Key.String()] = v
		keys = append(keys, v.Key.String())
	}
	sort.Strings(keys)

	for _, k := range keys {
		if h.offset == idx {
			h.offset += 1
			return hash[k].Key, hash[k].Value
		}
		idx += 1
	}
	return nil, nil
}
func (h *Hash) Reset() {
	h.offset = 0
}
