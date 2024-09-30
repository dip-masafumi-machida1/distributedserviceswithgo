package server

import (
	"fmt"
	"sync"
)

// Log ...
type Log struct {
	mu      sync.Mutex
	records []Record
}

// NewLog ...
func NewLog() *Log {
	return &Log{}
}

// Append ...
func (c *Log) Append(record Record) (uint64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	record.Offset = uint64(len(c.records))
	c.records = append(c.records, record)
	return record.Offset, nil
}

// Read ...
func (c *Log) Read(offset uint64) (Record, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if offset >= uint64(len(c.records)) {
		return Record{}, ErrOffsetNotFound
	}
	return c.records[offset], nil
}

// Record ...
type Record struct {
	Value  []byte `json:"value"`
	Offset uint64 `json:"offset"`
}

// ErrOffsetNotFound is returned when an offset is not found in the log
var ErrOffsetNotFound = fmt.Errorf("offset not found")
