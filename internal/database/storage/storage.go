package storage

import (
	"sync"

	"go.uber.org/zap"
)

type Storage struct {
	m      map[string]string
	logger *zap.Logger
	mu     sync.RWMutex
}

func New(logger *zap.Logger) *Storage {
	return &Storage{
		m:      make(map[string]string),
		logger: logger,
		mu:     sync.RWMutex{},
	}
}

func (s *Storage) Save(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = value
}

func (s *Storage) Delete(key string) bool {
	s.mu.RLock()
	_, ok := s.m[key]
	s.mu.RUnlock()
	if !ok {
		return false
	}
	s.mu.Lock()
	delete(s.m, key)
	s.mu.Unlock()
	return true
}

func (s *Storage) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	v, ok := s.m[key]
	return v, ok
}
