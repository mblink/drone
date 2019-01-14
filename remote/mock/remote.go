// Copyright 2018 Drone.IO Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mock

import (
	"net/http"

	"github.com/mblink/drone/model"
	"github.com/stretchr/testify/mock"
)

// Remote is an autogenerated mock type for the Remote type
type Remote struct {
	mock.Mock
}

// Activate provides a mock function with given fields: u, r, link
func (_m *Remote) Activate(u *model.User, r *model.Repo, link string) error {
	ret := _m.Called(u, r, link)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.User, *model.Repo, string) error); ok {
		r0 = rf(u, r, link)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Auth provides a mock function with given fields: token, secret
func (_m *Remote) Auth(token string, secret string) (string, error) {
	ret := _m.Called(token, secret)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(token, secret)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(token, secret)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Deactivate provides a mock function with given fields: u, r, link
func (_m *Remote) Deactivate(u *model.User, r *model.Repo, link string) error {
	ret := _m.Called(u, r, link)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.User, *model.Repo, string) error); ok {
		r0 = rf(u, r, link)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// File provides a mock function with given fields: u, r, b, f
func (_m *Remote) File(u *model.User, r *model.Repo, b *model.Build, f string) ([]byte, error) {
	ret := _m.Called(u, r, b, f)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(*model.User, *model.Repo, *model.Build, string) []byte); ok {
		r0 = rf(u, r, b, f)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.User, *model.Repo, *model.Build, string) error); ok {
		r1 = rf(u, r, b, f)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FileRef provides a mock function with given fields: u, r, ref, f
func (_m *Remote) FileRef(u *model.User, r *model.Repo, ref string, f string) ([]byte, error) {
	ret := _m.Called(u, r, ref, f)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(*model.User, *model.Repo, string, string) []byte); ok {
		r0 = rf(u, r, ref, f)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.User, *model.Repo, string, string) error); ok {
		r1 = rf(u, r, ref, f)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Hook provides a mock function with given fields: r
func (_m *Remote) Hook(r *http.Request) (*model.Repo, *model.Build, error) {
	ret := _m.Called(r)

	var r0 *model.Repo
	if rf, ok := ret.Get(0).(func(*http.Request) *model.Repo); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Repo)
		}
	}

	var r1 *model.Build
	if rf, ok := ret.Get(1).(func(*http.Request) *model.Build); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.Build)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*http.Request) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Login provides a mock function with given fields: w, r
func (_m *Remote) Login(w http.ResponseWriter, r *http.Request) (*model.User, error) {
	ret := _m.Called(w, r)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(http.ResponseWriter, *http.Request) *model.User); ok {
		r0 = rf(w, r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(http.ResponseWriter, *http.Request) error); ok {
		r1 = rf(w, r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Netrc provides a mock function with given fields: u, r
func (_m *Remote) Netrc(u *model.User, r *model.Repo) (*model.Netrc, error) {
	ret := _m.Called(u, r)

	var r0 *model.Netrc
	if rf, ok := ret.Get(0).(func(*model.User, *model.Repo) *model.Netrc); ok {
		r0 = rf(u, r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Netrc)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.User, *model.Repo) error); ok {
		r1 = rf(u, r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Perm provides a mock function with given fields: u, owner, repo
func (_m *Remote) Perm(u *model.User, owner string, repo string) (*model.Perm, error) {
	ret := _m.Called(u, owner, repo)

	var r0 *model.Perm
	if rf, ok := ret.Get(0).(func(*model.User, string, string) *model.Perm); ok {
		r0 = rf(u, owner, repo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Perm)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.User, string, string) error); ok {
		r1 = rf(u, owner, repo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repo provides a mock function with given fields: u, owner, repo
func (_m *Remote) Repo(u *model.User, owner string, repo string) (*model.Repo, error) {
	ret := _m.Called(u, owner, repo)

	var r0 *model.Repo
	if rf, ok := ret.Get(0).(func(*model.User, string, string) *model.Repo); ok {
		r0 = rf(u, owner, repo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Repo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.User, string, string) error); ok {
		r1 = rf(u, owner, repo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repos provides a mock function with given fields: u
func (_m *Remote) Repos(u *model.User) ([]*model.RepoLite, error) {
	ret := _m.Called(u)

	var r0 []*model.RepoLite
	if rf, ok := ret.Get(0).(func(*model.User) []*model.RepoLite); ok {
		r0 = rf(u)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.RepoLite)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.User) error); ok {
		r1 = rf(u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Status provides a mock function with given fields: u, r, b, link
func (_m *Remote) Status(u *model.User, r *model.Repo, b *model.Build, link string) error {
	ret := _m.Called(u, r, b, link)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.User, *model.Repo, *model.Build, string) error); ok {
		r0 = rf(u, r, b, link)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TeamPerm provides a mock function with given fields: u, org
func (_m *Remote) TeamPerm(u *model.User, org string) (*model.Perm, error) {
	ret := _m.Called(u, org)

	var r0 *model.Perm
	if rf, ok := ret.Get(0).(func(*model.User, string) *model.Perm); ok {
		r0 = rf(u, org)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Perm)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.User, string) error); ok {
		r1 = rf(u, org)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Teams provides a mock function with given fields: u
func (_m *Remote) Teams(u *model.User) ([]*model.Team, error) {
	ret := _m.Called(u)

	var r0 []*model.Team
	if rf, ok := ret.Get(0).(func(*model.User) []*model.Team); ok {
		r0 = rf(u)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Team)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.User) error); ok {
		r1 = rf(u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
