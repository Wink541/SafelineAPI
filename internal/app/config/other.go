package config

import (
	"strings"
)

type KVPair map[string]string

func (kvp *KVPair) Set(str string) {
	kvps := strings.Split(str, ",")
	for _, i := range kvps {
		kv := strings.SplitN(i, "=", 2)
		(*kvp)[kv[0]] = kv[1]
	}
}
