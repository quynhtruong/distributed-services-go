package server

import (
	"fmt"
	"sync"
)

type Log struct {
	mu      sync.Mutex
	records []Record
}

type Record struct {
	Value  []byte `json:"value"`
	Offset int64  `json:"offset"`
}

var ErrorOffsetNotFound = fmt.Errorf("Offset not found")

func NewLog() *Log {
	return &Log{}
}

func (c *Log) Append(record Record) (int64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	record.Offset = int64(len(c.records))
	c.records = append(c.records, record)
	return record.Offset, nil
}

func (c *Log) Read(offset int64) (Record, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if offset > int64(len(c.records)) {
		return Record{}, ErrorOffsetNotFound
	}
	return c.records[offset], nil
}
