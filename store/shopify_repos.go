package store

import (
	"sync"
	"time"

	"github.com/google/go-github/github"
)

//go:generate moq -out repo_mock.go . Repo

// Repo defines all exported functions.
type Repo interface {
	SetRepos(items []*github.Repository)
	GetRepos() []*github.Repository
	IsExpired() bool
}

var (
	r    *repoStore
	once sync.Once
)

// NewRepoStore will generate a new instance of repo store followed singleton pattern.
func NewRepoStore() Repo {
	once.Do(func() {
		r = &repoStore{
			items: []*github.Repository{},
			// TODO: cache duration should be in env
			cacheDuration: time.Second * 60,
		}
	})
	return r
}

// repoStore is the implementation of in-memory store, it is concurrent safe and followed
// re-entrant pattern.
type repoStore struct {
	sync.RWMutex
	items         []*github.Repository
	cacheDuration time.Duration
	lastUpdated   time.Time
}

// SetRepos will set all repos within store
func (r *repoStore) SetRepos(items []*github.Repository) {
	r.setRepos(items)
}

func (r *repoStore) setRepos(items []*github.Repository) {
	r.Lock()
	defer r.Unlock()
	r.items = items
	r.lastUpdated = time.Now()
}

// GetRepos retrieve all repos in store.
func (r *repoStore) GetRepos() []*github.Repository {
	return r.getRepos()
}

func (r *repoStore) getRepos() []*github.Repository {
	r.RLock()
	defer r.RUnlock()
	return r.items
}

// IsExpired checks if store is expired.
func (r *repoStore) IsExpired() bool {
	return r.isExpired()
}

func (r *repoStore) isExpired() bool {
	r.RLock()
	defer r.RUnlock()
	return time.Now().After(r.lastUpdated.Add(r.cacheDuration))
}
