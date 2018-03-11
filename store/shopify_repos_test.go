package store

import (
	"reflect"
	"testing"

	"time"

	"github.com/google/go-github/github"
)

func TestNewRepoStore(t *testing.T) {
	r := &repoStore{
		items:         []*github.Repository{},
		cacheDuration: time.Second * 60,
	}
	tests := []struct {
		name string
		want Repo
	}{
		{
			name: "success",
			want: r,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRepoStore()
			if got.IsExpired() != tt.want.IsExpired() {
				t.Errorf("want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got.GetRepos(), tt.want.GetRepos()) {
				t.Errorf("want %v", got, tt.want)
			}
		})
	}
}

func Test_repoStore_SetRepos(t *testing.T) {
	r := &repoStore{
		items:         []*github.Repository{},
		lastUpdated:   time.Now(),
		cacheDuration: time.Second * 60,
	}
	testID1, testName1, testID2, testName2 := int64(1), "Test 1", int64(2), "Test 2"
	type args struct {
		items []*github.Repository
	}
	tests := []struct {
		name string
		r    *repoStore
		args args
	}{
		{
			name: "success empty",
			r:    r,
			args: args{
				items: []*github.Repository{},
			},
		},
		{
			name: "success with data",
			r:    r,
			args: args{
				items: []*github.Repository{
					{
						ID:   &testID1,
						Name: &testName1,
					},
					{
						ID:   &testID2,
						Name: &testName2,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.SetRepos(tt.args.items)
			if !reflect.DeepEqual(tt.r.GetRepos(), tt.args.items) {
				t.Errorf("want %v", tt.args, tt.r.GetRepos())
			}
		})
	}
}

func Test_repoStore_GetRepos(t *testing.T) {
	testID1, testName1, testID2, testName2 := int64(1), "Test 1", int64(2), "Test 2"
	repos := []*github.Repository{
		&github.Repository{
			ID:   &testID1,
			Name: &testName1,
		},
		&github.Repository{
			ID:   &testID2,
			Name: &testName2,
		},
	}
	r := &repoStore{
		items:         repos,
		lastUpdated:   time.Now(),
		cacheDuration: time.Second * 60,
	}
	if !reflect.DeepEqual(repos, r.GetRepos()) {
		t.Errorf("want %v", repos, r.GetRepos())
	}
}

func Test_repoStore_IsExpired_True(t *testing.T) {
	// Test expired
	r := &repoStore{
		lastUpdated:   time.Now().Add(-time.Second * 61),
		cacheDuration: time.Second * 60,
	}
	if !r.IsExpired() {
		t.Errorf("want %v", true, r.IsExpired())
	}
}

func Test_repoStore_IsExpired_False(t *testing.T) {
	// Test expired
	r := &repoStore{
		lastUpdated:   time.Now(),
		cacheDuration: time.Second * 1,
	}
	if r.IsExpired() {
		t.Errorf("want %v", false, r.IsExpired())
	}
}
