package store

import (
	"sync"
	"time"

	"github.com/google/go-github/github"
)

type Repo interface {
	SetRepos(items []*github.Repository)
	GetRepos() []*github.Repository
	IsExpired() bool
}

var (
	r    *repoStore
	once sync.Once
)

func NewRepoStore() Repo {
	once.Do(func() {
		r = &repoStore{
			items:         []*github.Repository{},
			lastUpdated:   time.Now(),
			cacheDuration: time.Second * 2,
		}
	})
	return r
}

type repoStore struct {
	sync.RWMutex
	items         []*github.Repository
	cacheDuration time.Duration
	lastUpdated   time.Time
}

func (r *repoStore) SetRepos(items []*github.Repository) {
	r.setRepos(items)
}

func (r *repoStore) setRepos(items []*github.Repository) {
	r.Lock()
	defer r.Unlock()
	r.items = items
	r.lastUpdated = time.Now()
}

func (r *repoStore) GetRepos() []*github.Repository {
	return r.getRepos()
}

func (r *repoStore) getRepos() []*github.Repository {
	r.RLock()
	defer r.RUnlock()
	return r.items
}

func (r *repoStore) IsExpired() bool {
	return r.isExpired()
}

func (r *repoStore) isExpired() bool {
	r.RLock()
	defer r.RUnlock()
	return time.Now().After(r.lastUpdated.Add(r.cacheDuration))
}
