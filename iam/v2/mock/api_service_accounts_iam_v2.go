// Code generated by mocker. DO NOT EDIT.
// github.com/travisjeffery/mocker
// Source: api_service_accounts_iam_v2.go

package mock

import (
	context "context"
	net_http "net/http"
	sync "sync"

	github_com_confluentinc_ccloud_sdk_go_v2_iam_v2 "github.com/confluentinc/ccloud-sdk-go-v2/iam/v2"
)

// ServiceAccountsIamV2Api is a mock of ServiceAccountsIamV2Api interface
type ServiceAccountsIamV2Api struct {
	lockCreateIamV2ServiceAccount sync.Mutex
	CreateIamV2ServiceAccountFunc func(ctx context.Context) github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiCreateIamV2ServiceAccountRequest

	lockCreateIamV2ServiceAccountExecute sync.Mutex
	CreateIamV2ServiceAccountExecuteFunc func(r github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiCreateIamV2ServiceAccountRequest) (github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.IamV2ServiceAccount, *net_http.Response, error)

	lockDeleteIamV2ServiceAccount sync.Mutex
	DeleteIamV2ServiceAccountFunc func(ctx context.Context, id string) github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiDeleteIamV2ServiceAccountRequest

	lockDeleteIamV2ServiceAccountExecute sync.Mutex
	DeleteIamV2ServiceAccountExecuteFunc func(r github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiDeleteIamV2ServiceAccountRequest) (*net_http.Response, error)

	lockGetIamV2ServiceAccount sync.Mutex
	GetIamV2ServiceAccountFunc func(ctx context.Context, id string) github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiGetIamV2ServiceAccountRequest

	lockGetIamV2ServiceAccountExecute sync.Mutex
	GetIamV2ServiceAccountExecuteFunc func(r github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiGetIamV2ServiceAccountRequest) (github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.IamV2ServiceAccount, *net_http.Response, error)

	lockListIamV2ServiceAccounts sync.Mutex
	ListIamV2ServiceAccountsFunc func(ctx context.Context) github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiListIamV2ServiceAccountsRequest

	lockListIamV2ServiceAccountsExecute sync.Mutex
	ListIamV2ServiceAccountsExecuteFunc func(r github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiListIamV2ServiceAccountsRequest) (github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.IamV2ServiceAccountList, *net_http.Response, error)

	lockUpdateIamV2ServiceAccount sync.Mutex
	UpdateIamV2ServiceAccountFunc func(ctx context.Context, id string) github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiUpdateIamV2ServiceAccountRequest

	lockUpdateIamV2ServiceAccountExecute sync.Mutex
	UpdateIamV2ServiceAccountExecuteFunc func(r github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiUpdateIamV2ServiceAccountRequest) (github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.IamV2ServiceAccount, *net_http.Response, error)

	calls struct {
		CreateIamV2ServiceAccount []struct {
			Ctx context.Context
		}
		CreateIamV2ServiceAccountExecute []struct {
			R github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiCreateIamV2ServiceAccountRequest
		}
		DeleteIamV2ServiceAccount []struct {
			Ctx context.Context
			Id  string
		}
		DeleteIamV2ServiceAccountExecute []struct {
			R github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiDeleteIamV2ServiceAccountRequest
		}
		GetIamV2ServiceAccount []struct {
			Ctx context.Context
			Id  string
		}
		GetIamV2ServiceAccountExecute []struct {
			R github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiGetIamV2ServiceAccountRequest
		}
		ListIamV2ServiceAccounts []struct {
			Ctx context.Context
		}
		ListIamV2ServiceAccountsExecute []struct {
			R github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiListIamV2ServiceAccountsRequest
		}
		UpdateIamV2ServiceAccount []struct {
			Ctx context.Context
			Id  string
		}
		UpdateIamV2ServiceAccountExecute []struct {
			R github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiUpdateIamV2ServiceAccountRequest
		}
	}
}

// CreateIamV2ServiceAccount mocks base method by wrapping the associated func.
func (m *ServiceAccountsIamV2Api) CreateIamV2ServiceAccount(ctx context.Context) github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiCreateIamV2ServiceAccountRequest {
	m.lockCreateIamV2ServiceAccount.Lock()
	defer m.lockCreateIamV2ServiceAccount.Unlock()

	if m.CreateIamV2ServiceAccountFunc == nil {
		panic("mocker: ServiceAccountsIamV2Api.CreateIamV2ServiceAccountFunc is nil but ServiceAccountsIamV2Api.CreateIamV2ServiceAccount was called.")
	}

	call := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}

	m.calls.CreateIamV2ServiceAccount = append(m.calls.CreateIamV2ServiceAccount, call)

	return m.CreateIamV2ServiceAccountFunc(ctx)
}

// CreateIamV2ServiceAccountCalled returns true if CreateIamV2ServiceAccount was called at least once.
func (m *ServiceAccountsIamV2Api) CreateIamV2ServiceAccountCalled() bool {
	m.lockCreateIamV2ServiceAccount.Lock()
	defer m.lockCreateIamV2ServiceAccount.Unlock()

	return len(m.calls.CreateIamV2ServiceAccount) > 0
}

// CreateIamV2ServiceAccountCalls returns the calls made to CreateIamV2ServiceAccount.
func (m *ServiceAccountsIamV2Api) CreateIamV2ServiceAccountCalls() []struct {
	Ctx context.Context
} {
	m.lockCreateIamV2ServiceAccount.Lock()
	defer m.lockCreateIamV2ServiceAccount.Unlock()

	return m.calls.CreateIamV2ServiceAccount
}

// CreateIamV2ServiceAccountExecute mocks base method by wrapping the associated func.
func (m *ServiceAccountsIamV2Api) CreateIamV2ServiceAccountExecute(r github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiCreateIamV2ServiceAccountRequest) (github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.IamV2ServiceAccount, *net_http.Response, error) {
	m.lockCreateIamV2ServiceAccountExecute.Lock()
	defer m.lockCreateIamV2ServiceAccountExecute.Unlock()

	if m.CreateIamV2ServiceAccountExecuteFunc == nil {
		panic("mocker: ServiceAccountsIamV2Api.CreateIamV2ServiceAccountExecuteFunc is nil but ServiceAccountsIamV2Api.CreateIamV2ServiceAccountExecute was called.")
	}

	call := struct {
		R github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiCreateIamV2ServiceAccountRequest
	}{
		R: r,
	}

	m.calls.CreateIamV2ServiceAccountExecute = append(m.calls.CreateIamV2ServiceAccountExecute, call)

	return m.CreateIamV2ServiceAccountExecuteFunc(r)
}

// CreateIamV2ServiceAccountExecuteCalled returns true if CreateIamV2ServiceAccountExecute was called at least once.
func (m *ServiceAccountsIamV2Api) CreateIamV2ServiceAccountExecuteCalled() bool {
	m.lockCreateIamV2ServiceAccountExecute.Lock()
	defer m.lockCreateIamV2ServiceAccountExecute.Unlock()

	return len(m.calls.CreateIamV2ServiceAccountExecute) > 0
}

// CreateIamV2ServiceAccountExecuteCalls returns the calls made to CreateIamV2ServiceAccountExecute.
func (m *ServiceAccountsIamV2Api) CreateIamV2ServiceAccountExecuteCalls() []struct {
	R github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiCreateIamV2ServiceAccountRequest
} {
	m.lockCreateIamV2ServiceAccountExecute.Lock()
	defer m.lockCreateIamV2ServiceAccountExecute.Unlock()

	return m.calls.CreateIamV2ServiceAccountExecute
}

// DeleteIamV2ServiceAccount mocks base method by wrapping the associated func.
func (m *ServiceAccountsIamV2Api) DeleteIamV2ServiceAccount(ctx context.Context, id string) github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiDeleteIamV2ServiceAccountRequest {
	m.lockDeleteIamV2ServiceAccount.Lock()
	defer m.lockDeleteIamV2ServiceAccount.Unlock()

	if m.DeleteIamV2ServiceAccountFunc == nil {
		panic("mocker: ServiceAccountsIamV2Api.DeleteIamV2ServiceAccountFunc is nil but ServiceAccountsIamV2Api.DeleteIamV2ServiceAccount was called.")
	}

	call := struct {
		Ctx context.Context
		Id  string
	}{
		Ctx: ctx,
		Id:  id,
	}

	m.calls.DeleteIamV2ServiceAccount = append(m.calls.DeleteIamV2ServiceAccount, call)

	return m.DeleteIamV2ServiceAccountFunc(ctx, id)
}

// DeleteIamV2ServiceAccountCalled returns true if DeleteIamV2ServiceAccount was called at least once.
func (m *ServiceAccountsIamV2Api) DeleteIamV2ServiceAccountCalled() bool {
	m.lockDeleteIamV2ServiceAccount.Lock()
	defer m.lockDeleteIamV2ServiceAccount.Unlock()

	return len(m.calls.DeleteIamV2ServiceAccount) > 0
}

// DeleteIamV2ServiceAccountCalls returns the calls made to DeleteIamV2ServiceAccount.
func (m *ServiceAccountsIamV2Api) DeleteIamV2ServiceAccountCalls() []struct {
	Ctx context.Context
	Id  string
} {
	m.lockDeleteIamV2ServiceAccount.Lock()
	defer m.lockDeleteIamV2ServiceAccount.Unlock()

	return m.calls.DeleteIamV2ServiceAccount
}

// DeleteIamV2ServiceAccountExecute mocks base method by wrapping the associated func.
func (m *ServiceAccountsIamV2Api) DeleteIamV2ServiceAccountExecute(r github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiDeleteIamV2ServiceAccountRequest) (*net_http.Response, error) {
	m.lockDeleteIamV2ServiceAccountExecute.Lock()
	defer m.lockDeleteIamV2ServiceAccountExecute.Unlock()

	if m.DeleteIamV2ServiceAccountExecuteFunc == nil {
		panic("mocker: ServiceAccountsIamV2Api.DeleteIamV2ServiceAccountExecuteFunc is nil but ServiceAccountsIamV2Api.DeleteIamV2ServiceAccountExecute was called.")
	}

	call := struct {
		R github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiDeleteIamV2ServiceAccountRequest
	}{
		R: r,
	}

	m.calls.DeleteIamV2ServiceAccountExecute = append(m.calls.DeleteIamV2ServiceAccountExecute, call)

	return m.DeleteIamV2ServiceAccountExecuteFunc(r)
}

// DeleteIamV2ServiceAccountExecuteCalled returns true if DeleteIamV2ServiceAccountExecute was called at least once.
func (m *ServiceAccountsIamV2Api) DeleteIamV2ServiceAccountExecuteCalled() bool {
	m.lockDeleteIamV2ServiceAccountExecute.Lock()
	defer m.lockDeleteIamV2ServiceAccountExecute.Unlock()

	return len(m.calls.DeleteIamV2ServiceAccountExecute) > 0
}

// DeleteIamV2ServiceAccountExecuteCalls returns the calls made to DeleteIamV2ServiceAccountExecute.
func (m *ServiceAccountsIamV2Api) DeleteIamV2ServiceAccountExecuteCalls() []struct {
	R github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiDeleteIamV2ServiceAccountRequest
} {
	m.lockDeleteIamV2ServiceAccountExecute.Lock()
	defer m.lockDeleteIamV2ServiceAccountExecute.Unlock()

	return m.calls.DeleteIamV2ServiceAccountExecute
}

// GetIamV2ServiceAccount mocks base method by wrapping the associated func.
func (m *ServiceAccountsIamV2Api) GetIamV2ServiceAccount(ctx context.Context, id string) github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiGetIamV2ServiceAccountRequest {
	m.lockGetIamV2ServiceAccount.Lock()
	defer m.lockGetIamV2ServiceAccount.Unlock()

	if m.GetIamV2ServiceAccountFunc == nil {
		panic("mocker: ServiceAccountsIamV2Api.GetIamV2ServiceAccountFunc is nil but ServiceAccountsIamV2Api.GetIamV2ServiceAccount was called.")
	}

	call := struct {
		Ctx context.Context
		Id  string
	}{
		Ctx: ctx,
		Id:  id,
	}

	m.calls.GetIamV2ServiceAccount = append(m.calls.GetIamV2ServiceAccount, call)

	return m.GetIamV2ServiceAccountFunc(ctx, id)
}

// GetIamV2ServiceAccountCalled returns true if GetIamV2ServiceAccount was called at least once.
func (m *ServiceAccountsIamV2Api) GetIamV2ServiceAccountCalled() bool {
	m.lockGetIamV2ServiceAccount.Lock()
	defer m.lockGetIamV2ServiceAccount.Unlock()

	return len(m.calls.GetIamV2ServiceAccount) > 0
}

// GetIamV2ServiceAccountCalls returns the calls made to GetIamV2ServiceAccount.
func (m *ServiceAccountsIamV2Api) GetIamV2ServiceAccountCalls() []struct {
	Ctx context.Context
	Id  string
} {
	m.lockGetIamV2ServiceAccount.Lock()
	defer m.lockGetIamV2ServiceAccount.Unlock()

	return m.calls.GetIamV2ServiceAccount
}

// GetIamV2ServiceAccountExecute mocks base method by wrapping the associated func.
func (m *ServiceAccountsIamV2Api) GetIamV2ServiceAccountExecute(r github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiGetIamV2ServiceAccountRequest) (github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.IamV2ServiceAccount, *net_http.Response, error) {
	m.lockGetIamV2ServiceAccountExecute.Lock()
	defer m.lockGetIamV2ServiceAccountExecute.Unlock()

	if m.GetIamV2ServiceAccountExecuteFunc == nil {
		panic("mocker: ServiceAccountsIamV2Api.GetIamV2ServiceAccountExecuteFunc is nil but ServiceAccountsIamV2Api.GetIamV2ServiceAccountExecute was called.")
	}

	call := struct {
		R github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiGetIamV2ServiceAccountRequest
	}{
		R: r,
	}

	m.calls.GetIamV2ServiceAccountExecute = append(m.calls.GetIamV2ServiceAccountExecute, call)

	return m.GetIamV2ServiceAccountExecuteFunc(r)
}

// GetIamV2ServiceAccountExecuteCalled returns true if GetIamV2ServiceAccountExecute was called at least once.
func (m *ServiceAccountsIamV2Api) GetIamV2ServiceAccountExecuteCalled() bool {
	m.lockGetIamV2ServiceAccountExecute.Lock()
	defer m.lockGetIamV2ServiceAccountExecute.Unlock()

	return len(m.calls.GetIamV2ServiceAccountExecute) > 0
}

// GetIamV2ServiceAccountExecuteCalls returns the calls made to GetIamV2ServiceAccountExecute.
func (m *ServiceAccountsIamV2Api) GetIamV2ServiceAccountExecuteCalls() []struct {
	R github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiGetIamV2ServiceAccountRequest
} {
	m.lockGetIamV2ServiceAccountExecute.Lock()
	defer m.lockGetIamV2ServiceAccountExecute.Unlock()

	return m.calls.GetIamV2ServiceAccountExecute
}

// ListIamV2ServiceAccounts mocks base method by wrapping the associated func.
func (m *ServiceAccountsIamV2Api) ListIamV2ServiceAccounts(ctx context.Context) github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiListIamV2ServiceAccountsRequest {
	m.lockListIamV2ServiceAccounts.Lock()
	defer m.lockListIamV2ServiceAccounts.Unlock()

	if m.ListIamV2ServiceAccountsFunc == nil {
		panic("mocker: ServiceAccountsIamV2Api.ListIamV2ServiceAccountsFunc is nil but ServiceAccountsIamV2Api.ListIamV2ServiceAccounts was called.")
	}

	call := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}

	m.calls.ListIamV2ServiceAccounts = append(m.calls.ListIamV2ServiceAccounts, call)

	return m.ListIamV2ServiceAccountsFunc(ctx)
}

// ListIamV2ServiceAccountsCalled returns true if ListIamV2ServiceAccounts was called at least once.
func (m *ServiceAccountsIamV2Api) ListIamV2ServiceAccountsCalled() bool {
	m.lockListIamV2ServiceAccounts.Lock()
	defer m.lockListIamV2ServiceAccounts.Unlock()

	return len(m.calls.ListIamV2ServiceAccounts) > 0
}

// ListIamV2ServiceAccountsCalls returns the calls made to ListIamV2ServiceAccounts.
func (m *ServiceAccountsIamV2Api) ListIamV2ServiceAccountsCalls() []struct {
	Ctx context.Context
} {
	m.lockListIamV2ServiceAccounts.Lock()
	defer m.lockListIamV2ServiceAccounts.Unlock()

	return m.calls.ListIamV2ServiceAccounts
}

// ListIamV2ServiceAccountsExecute mocks base method by wrapping the associated func.
func (m *ServiceAccountsIamV2Api) ListIamV2ServiceAccountsExecute(r github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiListIamV2ServiceAccountsRequest) (github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.IamV2ServiceAccountList, *net_http.Response, error) {
	m.lockListIamV2ServiceAccountsExecute.Lock()
	defer m.lockListIamV2ServiceAccountsExecute.Unlock()

	if m.ListIamV2ServiceAccountsExecuteFunc == nil {
		panic("mocker: ServiceAccountsIamV2Api.ListIamV2ServiceAccountsExecuteFunc is nil but ServiceAccountsIamV2Api.ListIamV2ServiceAccountsExecute was called.")
	}

	call := struct {
		R github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiListIamV2ServiceAccountsRequest
	}{
		R: r,
	}

	m.calls.ListIamV2ServiceAccountsExecute = append(m.calls.ListIamV2ServiceAccountsExecute, call)

	return m.ListIamV2ServiceAccountsExecuteFunc(r)
}

// ListIamV2ServiceAccountsExecuteCalled returns true if ListIamV2ServiceAccountsExecute was called at least once.
func (m *ServiceAccountsIamV2Api) ListIamV2ServiceAccountsExecuteCalled() bool {
	m.lockListIamV2ServiceAccountsExecute.Lock()
	defer m.lockListIamV2ServiceAccountsExecute.Unlock()

	return len(m.calls.ListIamV2ServiceAccountsExecute) > 0
}

// ListIamV2ServiceAccountsExecuteCalls returns the calls made to ListIamV2ServiceAccountsExecute.
func (m *ServiceAccountsIamV2Api) ListIamV2ServiceAccountsExecuteCalls() []struct {
	R github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiListIamV2ServiceAccountsRequest
} {
	m.lockListIamV2ServiceAccountsExecute.Lock()
	defer m.lockListIamV2ServiceAccountsExecute.Unlock()

	return m.calls.ListIamV2ServiceAccountsExecute
}

// UpdateIamV2ServiceAccount mocks base method by wrapping the associated func.
func (m *ServiceAccountsIamV2Api) UpdateIamV2ServiceAccount(ctx context.Context, id string) github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiUpdateIamV2ServiceAccountRequest {
	m.lockUpdateIamV2ServiceAccount.Lock()
	defer m.lockUpdateIamV2ServiceAccount.Unlock()

	if m.UpdateIamV2ServiceAccountFunc == nil {
		panic("mocker: ServiceAccountsIamV2Api.UpdateIamV2ServiceAccountFunc is nil but ServiceAccountsIamV2Api.UpdateIamV2ServiceAccount was called.")
	}

	call := struct {
		Ctx context.Context
		Id  string
	}{
		Ctx: ctx,
		Id:  id,
	}

	m.calls.UpdateIamV2ServiceAccount = append(m.calls.UpdateIamV2ServiceAccount, call)

	return m.UpdateIamV2ServiceAccountFunc(ctx, id)
}

// UpdateIamV2ServiceAccountCalled returns true if UpdateIamV2ServiceAccount was called at least once.
func (m *ServiceAccountsIamV2Api) UpdateIamV2ServiceAccountCalled() bool {
	m.lockUpdateIamV2ServiceAccount.Lock()
	defer m.lockUpdateIamV2ServiceAccount.Unlock()

	return len(m.calls.UpdateIamV2ServiceAccount) > 0
}

// UpdateIamV2ServiceAccountCalls returns the calls made to UpdateIamV2ServiceAccount.
func (m *ServiceAccountsIamV2Api) UpdateIamV2ServiceAccountCalls() []struct {
	Ctx context.Context
	Id  string
} {
	m.lockUpdateIamV2ServiceAccount.Lock()
	defer m.lockUpdateIamV2ServiceAccount.Unlock()

	return m.calls.UpdateIamV2ServiceAccount
}

// UpdateIamV2ServiceAccountExecute mocks base method by wrapping the associated func.
func (m *ServiceAccountsIamV2Api) UpdateIamV2ServiceAccountExecute(r github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiUpdateIamV2ServiceAccountRequest) (github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.IamV2ServiceAccount, *net_http.Response, error) {
	m.lockUpdateIamV2ServiceAccountExecute.Lock()
	defer m.lockUpdateIamV2ServiceAccountExecute.Unlock()

	if m.UpdateIamV2ServiceAccountExecuteFunc == nil {
		panic("mocker: ServiceAccountsIamV2Api.UpdateIamV2ServiceAccountExecuteFunc is nil but ServiceAccountsIamV2Api.UpdateIamV2ServiceAccountExecute was called.")
	}

	call := struct {
		R github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiUpdateIamV2ServiceAccountRequest
	}{
		R: r,
	}

	m.calls.UpdateIamV2ServiceAccountExecute = append(m.calls.UpdateIamV2ServiceAccountExecute, call)

	return m.UpdateIamV2ServiceAccountExecuteFunc(r)
}

// UpdateIamV2ServiceAccountExecuteCalled returns true if UpdateIamV2ServiceAccountExecute was called at least once.
func (m *ServiceAccountsIamV2Api) UpdateIamV2ServiceAccountExecuteCalled() bool {
	m.lockUpdateIamV2ServiceAccountExecute.Lock()
	defer m.lockUpdateIamV2ServiceAccountExecute.Unlock()

	return len(m.calls.UpdateIamV2ServiceAccountExecute) > 0
}

// UpdateIamV2ServiceAccountExecuteCalls returns the calls made to UpdateIamV2ServiceAccountExecute.
func (m *ServiceAccountsIamV2Api) UpdateIamV2ServiceAccountExecuteCalls() []struct {
	R github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiUpdateIamV2ServiceAccountRequest
} {
	m.lockUpdateIamV2ServiceAccountExecute.Lock()
	defer m.lockUpdateIamV2ServiceAccountExecute.Unlock()

	return m.calls.UpdateIamV2ServiceAccountExecute
}

// Reset resets the calls made to the mocked methods.
func (m *ServiceAccountsIamV2Api) Reset() {
	m.lockCreateIamV2ServiceAccount.Lock()
	m.calls.CreateIamV2ServiceAccount = nil
	m.lockCreateIamV2ServiceAccount.Unlock()
	m.lockCreateIamV2ServiceAccountExecute.Lock()
	m.calls.CreateIamV2ServiceAccountExecute = nil
	m.lockCreateIamV2ServiceAccountExecute.Unlock()
	m.lockDeleteIamV2ServiceAccount.Lock()
	m.calls.DeleteIamV2ServiceAccount = nil
	m.lockDeleteIamV2ServiceAccount.Unlock()
	m.lockDeleteIamV2ServiceAccountExecute.Lock()
	m.calls.DeleteIamV2ServiceAccountExecute = nil
	m.lockDeleteIamV2ServiceAccountExecute.Unlock()
	m.lockGetIamV2ServiceAccount.Lock()
	m.calls.GetIamV2ServiceAccount = nil
	m.lockGetIamV2ServiceAccount.Unlock()
	m.lockGetIamV2ServiceAccountExecute.Lock()
	m.calls.GetIamV2ServiceAccountExecute = nil
	m.lockGetIamV2ServiceAccountExecute.Unlock()
	m.lockListIamV2ServiceAccounts.Lock()
	m.calls.ListIamV2ServiceAccounts = nil
	m.lockListIamV2ServiceAccounts.Unlock()
	m.lockListIamV2ServiceAccountsExecute.Lock()
	m.calls.ListIamV2ServiceAccountsExecute = nil
	m.lockListIamV2ServiceAccountsExecute.Unlock()
	m.lockUpdateIamV2ServiceAccount.Lock()
	m.calls.UpdateIamV2ServiceAccount = nil
	m.lockUpdateIamV2ServiceAccount.Unlock()
	m.lockUpdateIamV2ServiceAccountExecute.Lock()
	m.calls.UpdateIamV2ServiceAccountExecute = nil
	m.lockUpdateIamV2ServiceAccountExecute.Unlock()
}