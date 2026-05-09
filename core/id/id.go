package id

import (
	"strconv"
	"sync"
)

type counter struct {
	value int
	mx    sync.RWMutex
}

func (c *counter) nextInt() int {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.value++
	return c.value
}

var c = &counter{}

func Generate() string {
	return strconv.Itoa(c.nextInt())
}
