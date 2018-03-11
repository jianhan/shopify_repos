// Code generated by moq; DO NOT EDIT
// github.com/matryer/moq

package store

import (
	"github.com/google/go-github/github"
	"sync"
)

var (
	lockRepoMockGetRepos  sync.RWMutex
	lockRepoMockIsExpired sync.RWMutex
	lockRepoMockSetRepos  sync.RWMutex
)

// RepoMock is a mock implementation of Repo.
//
//     func TestSomethingThatUsesRepo(t *testing.T) {
//
//         // make and configure a mocked Repo
//         mockedRepo := &RepoMock{
//             GetReposFunc: func() []*github.Repository {
// 	               panic("TODO: mock out the GetRepos method")
//             },
//             IsExpiredFunc: func() bool {
// 	               panic("TODO: mock out the IsExpired method")
//             },
//             SetReposFunc: func(items []*github.Repository)  {
// 	               panic("TODO: mock out the SetRepos method")
//             },
//         }
//
//         // TODO: use mockedRepo in code that requires Repo
//         //       and then make assertions.
//
//     }
type RepoMock struct {
	// GetReposFunc mocks the GetRepos method.
	GetReposFunc func() []*github.Repository

	// IsExpiredFunc mocks the IsExpired method.
	IsExpiredFunc func() bool

	// SetReposFunc mocks the SetRepos method.
	SetReposFunc func(items []*github.Repository)

	// calls tracks calls to the methods.
	calls struct {
		// GetRepos holds details about calls to the GetRepos method.
		GetRepos []struct {
		}
		// IsExpired holds details about calls to the IsExpired method.
		IsExpired []struct {
		}
		// SetRepos holds details about calls to the SetRepos method.
		SetRepos []struct {
			// Items is the items argument value.
			Items []*github.Repository
		}
	}
}

// GetRepos calls GetReposFunc.
func (mock *RepoMock) GetRepos() []*github.Repository {
	if mock.GetReposFunc == nil {
		panic("moq: RepoMock.GetReposFunc is nil but Repo.GetRepos was just called")
	}
	callInfo := struct {
	}{}
	lockRepoMockGetRepos.Lock()
	mock.calls.GetRepos = append(mock.calls.GetRepos, callInfo)
	lockRepoMockGetRepos.Unlock()
	return mock.GetReposFunc()
}

// GetReposCalls gets all the calls that were made to GetRepos.
// Check the length with:
//     len(mockedRepo.GetReposCalls())
func (mock *RepoMock) GetReposCalls() []struct {
} {
	var calls []struct {
	}
	lockRepoMockGetRepos.RLock()
	calls = mock.calls.GetRepos
	lockRepoMockGetRepos.RUnlock()
	return calls
}

// IsExpired calls IsExpiredFunc.
func (mock *RepoMock) IsExpired() bool {
	if mock.IsExpiredFunc == nil {
		panic("moq: RepoMock.IsExpiredFunc is nil but Repo.IsExpired was just called")
	}
	callInfo := struct {
	}{}
	lockRepoMockIsExpired.Lock()
	mock.calls.IsExpired = append(mock.calls.IsExpired, callInfo)
	lockRepoMockIsExpired.Unlock()
	return mock.IsExpiredFunc()
}

// IsExpiredCalls gets all the calls that were made to IsExpired.
// Check the length with:
//     len(mockedRepo.IsExpiredCalls())
func (mock *RepoMock) IsExpiredCalls() []struct {
} {
	var calls []struct {
	}
	lockRepoMockIsExpired.RLock()
	calls = mock.calls.IsExpired
	lockRepoMockIsExpired.RUnlock()
	return calls
}

// SetRepos calls SetReposFunc.
func (mock *RepoMock) SetRepos(items []*github.Repository) {
	if mock.SetReposFunc == nil {
		panic("moq: RepoMock.SetReposFunc is nil but Repo.SetRepos was just called")
	}
	callInfo := struct {
		Items []*github.Repository
	}{
		Items: items,
	}
	lockRepoMockSetRepos.Lock()
	mock.calls.SetRepos = append(mock.calls.SetRepos, callInfo)
	lockRepoMockSetRepos.Unlock()
	mock.SetReposFunc(items)
}

// SetReposCalls gets all the calls that were made to SetRepos.
// Check the length with:
//     len(mockedRepo.SetReposCalls())
func (mock *RepoMock) SetReposCalls() []struct {
	Items []*github.Repository
} {
	var calls []struct {
		Items []*github.Repository
	}
	lockRepoMockSetRepos.RLock()
	calls = mock.calls.SetRepos
	lockRepoMockSetRepos.RUnlock()
	return calls
}
