package cache

import "time"

type Cache struct {
	Key        string `json:"key"`
	Value      string `json:"value"`
	Expiration time.Duration
}
