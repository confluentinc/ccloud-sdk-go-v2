// Code generated by mocker. DO NOT EDIT.
// github.com/travisjeffery/mocker
// Source: api_users_iam_v2.go

package mock

import (
	context "context"
	net_http "net/http"
	sync "sync"

	github_com_confluentinc_ccloud_sdk_go_v2_iam_v2 "github.com/confluentinc/ccloud-sdk-go-v2/iam/v2"
)

// UsersIamV2Api is a mock of UsersIamV2Api interface
type UsersIamV2Api struct {
	lockDeleteIamV2User sync.Mutex
	DeleteIamV2UserFunc func(ctx context.Context, id string) github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiDeleteIamV2UserRequest

	lockDeleteIamV2UserExecute sync.Mutex
	DeleteIamV2UserExecuteFunc func(r github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiDeleteIamV2UserRequest) (*net_http.Response, error)

	lockGetIamV2User sync.Mutex
	GetIamV2UserFunc func(ctx context.Context, id string) github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiGetIamV2UserRequest

	lockGetIamV2UserExecute sync.Mutex
	GetIamV2UserExecuteFunc func(r github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiGetIamV2UserRequest) (github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.IamV2User, *net_http.Response, error)

	lockListIamV2Users sync.Mutex
	ListIamV2UsersFunc func(ctx context.Context) github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiListIamV2UsersRequest

	lockListIamV2UsersExecute sync.Mutex
	ListIamV2UsersExecuteFunc func(r github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiListIamV2UsersRequest) (github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.IamV2UserList, *net_http.Response, error)

	lockUpdateIamV2User sync.Mutex
	UpdateIamV2UserFunc func(ctx context.Context, id string) github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiUpdateIamV2UserRequest

	lockUpdateIamV2UserExecute sync.Mutex
	UpdateIamV2UserExecuteFunc func(r github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiUpdateIamV2UserRequest) (github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.IamV2User, *net_http.Response, error)

	calls struct {
		DeleteIamV2User []struct {
			Ctx context.Context
			Id  string
		}
		DeleteIamV2UserExecute []struct {
			R github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiDeleteIamV2UserRequest
		}
		GetIamV2User []struct {
			Ctx context.Context
			Id  string
		}
		GetIamV2UserExecute []struct {
			R github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiGetIamV2UserRequest
		}
		ListIamV2Users []struct {
			Ctx context.Context
		}
		ListIamV2UsersExecute []struct {
			R github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiListIamV2UsersRequest
		}
		UpdateIamV2User []struct {
			Ctx context.Context
			Id  string
		}
		UpdateIamV2UserExecute []struct {
			R github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiUpdateIamV2UserRequest
		}
	}
}

// DeleteIamV2User mocks base method by wrapping the associated func.
func (m *UsersIamV2Api) DeleteIamV2User(ctx context.Context, id string) github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiDeleteIamV2UserRequest {
	m.lockDeleteIamV2User.Lock()
	defer m.lockDeleteIamV2User.Unlock()

	if m.DeleteIamV2UserFunc == nil {
		panic("mocker: UsersIamV2Api.DeleteIamV2UserFunc is nil but UsersIamV2Api.DeleteIamV2User was called.")
	}

	call := struct {
		Ctx context.Context
		Id  string
	}{
		Ctx: ctx,
		Id:  id,
	}

	m.calls.DeleteIamV2User = append(m.calls.DeleteIamV2User, call)

	return m.DeleteIamV2UserFunc(ctx, id)
}

// DeleteIamV2UserCalled returns true if DeleteIamV2User was called at least once.
func (m *UsersIamV2Api) DeleteIamV2UserCalled() bool {
	m.lockDeleteIamV2User.Lock()
	defer m.lockDeleteIamV2User.Unlock()

	return len(m.calls.DeleteIamV2User) > 0
}

// DeleteIamV2UserCalls returns the calls made to DeleteIamV2User.
func (m *UsersIamV2Api) DeleteIamV2UserCalls() []struct {
	Ctx context.Context
	Id  string
} {
	m.lockDeleteIamV2User.Lock()
	defer m.lockDeleteIamV2User.Unlock()

	return m.calls.DeleteIamV2User
}

// DeleteIamV2UserExecute mocks base method by wrapping the associated func.
func (m *UsersIamV2Api) DeleteIamV2UserExecute(r github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiDeleteIamV2UserRequest) (*net_http.Response, error) {
	m.lockDeleteIamV2UserExecute.Lock()
	defer m.lockDeleteIamV2UserExecute.Unlock()

	if m.DeleteIamV2UserExecuteFunc == nil {
		panic("mocker: UsersIamV2Api.DeleteIamV2UserExecuteFunc is nil but UsersIamV2Api.DeleteIamV2UserExecute was called.")
	}

	call := struct {
		R github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiDeleteIamV2UserRequest
	}{
		R: r,
	}

	m.calls.DeleteIamV2UserExecute = append(m.calls.DeleteIamV2UserExecute, call)

	return m.DeleteIamV2UserExecuteFunc(r)
}

// DeleteIamV2UserExecuteCalled returns true if DeleteIamV2UserExecute was called at least once.
func (m *UsersIamV2Api) DeleteIamV2UserExecuteCalled() bool {
	m.lockDeleteIamV2UserExecute.Lock()
	defer m.lockDeleteIamV2UserExecute.Unlock()

	return len(m.calls.DeleteIamV2UserExecute) > 0
}

// DeleteIamV2UserExecuteCalls returns the calls made to DeleteIamV2UserExecute.
func (m *UsersIamV2Api) DeleteIamV2UserExecuteCalls() []struct {
	R github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiDeleteIamV2UserRequest
} {
	m.lockDeleteIamV2UserExecute.Lock()
	defer m.lockDeleteIamV2UserExecute.Unlock()

	return m.calls.DeleteIamV2UserExecute
}

// GetIamV2User mocks base method by wrapping the associated func.
func (m *UsersIamV2Api) GetIamV2User(ctx context.Context, id string) github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiGetIamV2UserRequest {
	m.lockGetIamV2User.Lock()
	defer m.lockGetIamV2User.Unlock()

	if m.GetIamV2UserFunc == nil {
		panic("mocker: UsersIamV2Api.GetIamV2UserFunc is nil but UsersIamV2Api.GetIamV2User was called.")
	}

	call := struct {
		Ctx context.Context
		Id  string
	}{
		Ctx: ctx,
		Id:  id,
	}

	m.calls.GetIamV2User = append(m.calls.GetIamV2User, call)

	return m.GetIamV2UserFunc(ctx, id)
}

// GetIamV2UserCalled returns true if GetIamV2User was called at least once.
func (m *UsersIamV2Api) GetIamV2UserCalled() bool {
	m.lockGetIamV2User.Lock()
	defer m.lockGetIamV2User.Unlock()

	return len(m.calls.GetIamV2User) > 0
}

// GetIamV2UserCalls returns the calls made to GetIamV2User.
func (m *UsersIamV2Api) GetIamV2UserCalls() []struct {
	Ctx context.Context
	Id  string
} {
	m.lockGetIamV2User.Lock()
	defer m.lockGetIamV2User.Unlock()

	return m.calls.GetIamV2User
}

// GetIamV2UserExecute mocks base method by wrapping the associated func.
func (m *UsersIamV2Api) GetIamV2UserExecute(r github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiGetIamV2UserRequest) (github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.IamV2User, *net_http.Response, error) {
	m.lockGetIamV2UserExecute.Lock()
	defer m.lockGetIamV2UserExecute.Unlock()

	if m.GetIamV2UserExecuteFunc == nil {
		panic("mocker: UsersIamV2Api.GetIamV2UserExecuteFunc is nil but UsersIamV2Api.GetIamV2UserExecute was called.")
	}

	call := struct {
		R github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiGetIamV2UserRequest
	}{
		R: r,
	}

	m.calls.GetIamV2UserExecute = append(m.calls.GetIamV2UserExecute, call)

	return m.GetIamV2UserExecuteFunc(r)
}

// GetIamV2UserExecuteCalled returns true if GetIamV2UserExecute was called at least once.
func (m *UsersIamV2Api) GetIamV2UserExecuteCalled() bool {
	m.lockGetIamV2UserExecute.Lock()
	defer m.lockGetIamV2UserExecute.Unlock()

	return len(m.calls.GetIamV2UserExecute) > 0
}

// GetIamV2UserExecuteCalls returns the calls made to GetIamV2UserExecute.
func (m *UsersIamV2Api) GetIamV2UserExecuteCalls() []struct {
	R github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiGetIamV2UserRequest
} {
	m.lockGetIamV2UserExecute.Lock()
	defer m.lockGetIamV2UserExecute.Unlock()

	return m.calls.GetIamV2UserExecute
}

// ListIamV2Users mocks base method by wrapping the associated func.
func (m *UsersIamV2Api) ListIamV2Users(ctx context.Context) github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiListIamV2UsersRequest {
	m.lockListIamV2Users.Lock()
	defer m.lockListIamV2Users.Unlock()

	if m.ListIamV2UsersFunc == nil {
		panic("mocker: UsersIamV2Api.ListIamV2UsersFunc is nil but UsersIamV2Api.ListIamV2Users was called.")
	}

	call := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}

	m.calls.ListIamV2Users = append(m.calls.ListIamV2Users, call)

	return m.ListIamV2UsersFunc(ctx)
}

// ListIamV2UsersCalled returns true if ListIamV2Users was called at least once.
func (m *UsersIamV2Api) ListIamV2UsersCalled() bool {
	m.lockListIamV2Users.Lock()
	defer m.lockListIamV2Users.Unlock()

	return len(m.calls.ListIamV2Users) > 0
}

// ListIamV2UsersCalls returns the calls made to ListIamV2Users.
func (m *UsersIamV2Api) ListIamV2UsersCalls() []struct {
	Ctx context.Context
} {
	m.lockListIamV2Users.Lock()
	defer m.lockListIamV2Users.Unlock()

	return m.calls.ListIamV2Users
}

// ListIamV2UsersExecute mocks base method by wrapping the associated func.
func (m *UsersIamV2Api) ListIamV2UsersExecute(r github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiListIamV2UsersRequest) (github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.IamV2UserList, *net_http.Response, error) {
	m.lockListIamV2UsersExecute.Lock()
	defer m.lockListIamV2UsersExecute.Unlock()

	if m.ListIamV2UsersExecuteFunc == nil {
		panic("mocker: UsersIamV2Api.ListIamV2UsersExecuteFunc is nil but UsersIamV2Api.ListIamV2UsersExecute was called.")
	}

	call := struct {
		R github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiListIamV2UsersRequest
	}{
		R: r,
	}

	m.calls.ListIamV2UsersExecute = append(m.calls.ListIamV2UsersExecute, call)

	return m.ListIamV2UsersExecuteFunc(r)
}

// ListIamV2UsersExecuteCalled returns true if ListIamV2UsersExecute was called at least once.
func (m *UsersIamV2Api) ListIamV2UsersExecuteCalled() bool {
	m.lockListIamV2UsersExecute.Lock()
	defer m.lockListIamV2UsersExecute.Unlock()

	return len(m.calls.ListIamV2UsersExecute) > 0
}

// ListIamV2UsersExecuteCalls returns the calls made to ListIamV2UsersExecute.
func (m *UsersIamV2Api) ListIamV2UsersExecuteCalls() []struct {
	R github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiListIamV2UsersRequest
} {
	m.lockListIamV2UsersExecute.Lock()
	defer m.lockListIamV2UsersExecute.Unlock()

	return m.calls.ListIamV2UsersExecute
}

// UpdateIamV2User mocks base method by wrapping the associated func.
func (m *UsersIamV2Api) UpdateIamV2User(ctx context.Context, id string) github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiUpdateIamV2UserRequest {
	m.lockUpdateIamV2User.Lock()
	defer m.lockUpdateIamV2User.Unlock()

	if m.UpdateIamV2UserFunc == nil {
		panic("mocker: UsersIamV2Api.UpdateIamV2UserFunc is nil but UsersIamV2Api.UpdateIamV2User was called.")
	}

	call := struct {
		Ctx context.Context
		Id  string
	}{
		Ctx: ctx,
		Id:  id,
	}

	m.calls.UpdateIamV2User = append(m.calls.UpdateIamV2User, call)

	return m.UpdateIamV2UserFunc(ctx, id)
}

// UpdateIamV2UserCalled returns true if UpdateIamV2User was called at least once.
func (m *UsersIamV2Api) UpdateIamV2UserCalled() bool {
	m.lockUpdateIamV2User.Lock()
	defer m.lockUpdateIamV2User.Unlock()

	return len(m.calls.UpdateIamV2User) > 0
}

// UpdateIamV2UserCalls returns the calls made to UpdateIamV2User.
func (m *UsersIamV2Api) UpdateIamV2UserCalls() []struct {
	Ctx context.Context
	Id  string
} {
	m.lockUpdateIamV2User.Lock()
	defer m.lockUpdateIamV2User.Unlock()

	return m.calls.UpdateIamV2User
}

// UpdateIamV2UserExecute mocks base method by wrapping the associated func.
func (m *UsersIamV2Api) UpdateIamV2UserExecute(r github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiUpdateIamV2UserRequest) (github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.IamV2User, *net_http.Response, error) {
	m.lockUpdateIamV2UserExecute.Lock()
	defer m.lockUpdateIamV2UserExecute.Unlock()

	if m.UpdateIamV2UserExecuteFunc == nil {
		panic("mocker: UsersIamV2Api.UpdateIamV2UserExecuteFunc is nil but UsersIamV2Api.UpdateIamV2UserExecute was called.")
	}

	call := struct {
		R github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiUpdateIamV2UserRequest
	}{
		R: r,
	}

	m.calls.UpdateIamV2UserExecute = append(m.calls.UpdateIamV2UserExecute, call)

	return m.UpdateIamV2UserExecuteFunc(r)
}

// UpdateIamV2UserExecuteCalled returns true if UpdateIamV2UserExecute was called at least once.
func (m *UsersIamV2Api) UpdateIamV2UserExecuteCalled() bool {
	m.lockUpdateIamV2UserExecute.Lock()
	defer m.lockUpdateIamV2UserExecute.Unlock()

	return len(m.calls.UpdateIamV2UserExecute) > 0
}

// UpdateIamV2UserExecuteCalls returns the calls made to UpdateIamV2UserExecute.
func (m *UsersIamV2Api) UpdateIamV2UserExecuteCalls() []struct {
	R github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiUpdateIamV2UserRequest
} {
	m.lockUpdateIamV2UserExecute.Lock()
	defer m.lockUpdateIamV2UserExecute.Unlock()

	return m.calls.UpdateIamV2UserExecute
}

// Reset resets the calls made to the mocked methods.
func (m *UsersIamV2Api) Reset() {
	m.lockDeleteIamV2User.Lock()
	m.calls.DeleteIamV2User = nil
	m.lockDeleteIamV2User.Unlock()
	m.lockDeleteIamV2UserExecute.Lock()
	m.calls.DeleteIamV2UserExecute = nil
	m.lockDeleteIamV2UserExecute.Unlock()
	m.lockGetIamV2User.Lock()
	m.calls.GetIamV2User = nil
	m.lockGetIamV2User.Unlock()
	m.lockGetIamV2UserExecute.Lock()
	m.calls.GetIamV2UserExecute = nil
	m.lockGetIamV2UserExecute.Unlock()
	m.lockListIamV2Users.Lock()
	m.calls.ListIamV2Users = nil
	m.lockListIamV2Users.Unlock()
	m.lockListIamV2UsersExecute.Lock()
	m.calls.ListIamV2UsersExecute = nil
	m.lockListIamV2UsersExecute.Unlock()
	m.lockUpdateIamV2User.Lock()
	m.calls.UpdateIamV2User = nil
	m.lockUpdateIamV2User.Unlock()
	m.lockUpdateIamV2UserExecute.Lock()
	m.calls.UpdateIamV2UserExecute = nil
	m.lockUpdateIamV2UserExecute.Unlock()
}
