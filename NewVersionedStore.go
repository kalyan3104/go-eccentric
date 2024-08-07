package main

import (
	"fmt"
	"sync"
)

// VersionedStore provides a thread-safe key-value store with versioning.
type VersionedStore struct {
	mu        sync.RWMutex
	store     map[string]map[int]string
	version   map[string]int
}

// NewVersionedStore creates a new instance of VersionedStore.
func NewVersionedStore() *VersionedStore {
	return &VersionedStore{
		store:   make(map[string]map[int]string),
		version: make(map[string]int),
	}
}

// Set stores a value with the given key and version.
func (vs *VersionedStore) Set(key string, version int, value string) {
	vs.mu.Lock()
	defer vs.mu.Unlock()

	if _, exists := vs.store[key]; !exists {
		vs.store[key] = make(map[int]string)
	}
	vs.store[key][version] = value
	// Update the latest version for the key
	if version > vs.version[key] {
		vs.version[key] = version
	}
}

// Get retrieves a value by key and version.
func (vs *VersionedStore) Get(key string, version int) (string, bool) {
	vs.mu.RLock()
	defer vs.mu.RUnlock()

	if versions, exists := vs.store[key]; exists {
		value, found := versions[version]
		return value, found
	}
	return "", false
}

// Latest retrieves the latest value for the g
