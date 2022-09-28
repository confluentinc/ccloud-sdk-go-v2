// Code generated by mocker. DO NOT EDIT.
// github.com/travisjeffery/mocker
// Source: api_api_keys_iam_v2.go

package mock

import (
	context "context"
	net_http "net/http"
	sync "sync"

	github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2 "github.com/confluentinc/ccloud-sdk-go-v2/apikeys/v2"
)

// APIKeysIamV2Api is a mock of APIKeysIamV2Api interface
type APIKeysIamV2Api struct {
	lockCreateIamV2ApiKey sync.Mutex
	CreateIamV2ApiKeyFunc func(ctx context.Context) github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.ApiCreateIamV2ApiKeyRequest

	lockCreateIamV2ApiKeyExecute sync.Mutex
	CreateIamV2ApiKeyExecuteFunc func(r github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.ApiCreateIamV2ApiKeyRequest) (github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.IamV2ApiKey, *net_http.Response, error)

	lockDeleteIamV2ApiKey sync.Mutex
	DeleteIamV2ApiKeyFunc func(ctx context.Context, id string) github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.ApiDeleteIamV2ApiKeyRequest

	lockDeleteIamV2ApiKeyExecute sync.Mutex
	DeleteIamV2ApiKeyExecuteFunc func(r github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.ApiDeleteIamV2ApiKeyRequest) (*net_http.Response, error)

	lockGetIamV2ApiKey sync.Mutex
	GetIamV2ApiKeyFunc func(ctx context.Context, id string) github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.ApiGetIamV2ApiKeyRequest

	lockGetIamV2ApiKeyExecute sync.Mutex
	GetIamV2ApiKeyExecuteFunc func(r github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.ApiGetIamV2ApiKeyRequest) (github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.IamV2ApiKey, *net_http.Response, error)

	lockListIamV2ApiKeys sync.Mutex
	ListIamV2ApiKeysFunc func(ctx context.Context) github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.ApiListIamV2ApiKeysRequest

	lockListIamV2ApiKeysExecute sync.Mutex
	ListIamV2ApiKeysExecuteFunc func(r github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.ApiListIamV2ApiKeysRequest) (github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.IamV2ApiKeyList, *net_http.Response, error)

	lockUpdateIamV2ApiKey sync.Mutex
	UpdateIamV2ApiKeyFunc func(ctx context.Context, id string) github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.ApiUpdateIamV2ApiKeyRequest

	lockUpdateIamV2ApiKeyExecute sync.Mutex
	UpdateIamV2ApiKeyExecuteFunc func(r github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.ApiUpdateIamV2ApiKeyRequest) (github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.IamV2ApiKey, *net_http.Response, error)

	calls struct {
		CreateIamV2ApiKey []struct {
			Ctx context.Context
		}
		CreateIamV2ApiKeyExecute []struct {
			R github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.ApiCreateIamV2ApiKeyRequest
		}
		DeleteIamV2ApiKey []struct {
			Ctx context.Context
			Id  string
		}
		DeleteIamV2ApiKeyExecute []struct {
			R github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.ApiDeleteIamV2ApiKeyRequest
		}
		GetIamV2ApiKey []struct {
			Ctx context.Context
			Id  string
		}
		GetIamV2ApiKeyExecute []struct {
			R github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.ApiGetIamV2ApiKeyRequest
		}
		ListIamV2ApiKeys []struct {
			Ctx context.Context
		}
		ListIamV2ApiKeysExecute []struct {
			R github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.ApiListIamV2ApiKeysRequest
		}
		UpdateIamV2ApiKey []struct {
			Ctx context.Context
			Id  string
		}
		UpdateIamV2ApiKeyExecute []struct {
			R github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.ApiUpdateIamV2ApiKeyRequest
		}
	}
}

// CreateIamV2ApiKey mocks base method by wrapping the associated func.
func (m *APIKeysIamV2Api) CreateIamV2ApiKey(ctx context.Context) github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.ApiCreateIamV2ApiKeyRequest {
	m.lockCreateIamV2ApiKey.Lock()
	defer m.lockCreateIamV2ApiKey.Unlock()

	if m.CreateIamV2ApiKeyFunc == nil {
		panic("mocker: APIKeysIamV2Api.CreateIamV2ApiKeyFunc is nil but APIKeysIamV2Api.CreateIamV2ApiKey was called.")
	}

	call := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}

	m.calls.CreateIamV2ApiKey = append(m.calls.CreateIamV2ApiKey, call)

	return m.CreateIamV2ApiKeyFunc(ctx)
}

// CreateIamV2ApiKeyCalled returns true if CreateIamV2ApiKey was called at least once.
func (m *APIKeysIamV2Api) CreateIamV2ApiKeyCalled() bool {
	m.lockCreateIamV2ApiKey.Lock()
	defer m.lockCreateIamV2ApiKey.Unlock()

	return len(m.calls.CreateIamV2ApiKey) > 0
}

// CreateIamV2ApiKeyCalls returns the calls made to CreateIamV2ApiKey.
func (m *APIKeysIamV2Api) CreateIamV2ApiKeyCalls() []struct {
	Ctx context.Context
} {
	m.lockCreateIamV2ApiKey.Lock()
	defer m.lockCreateIamV2ApiKey.Unlock()

	return m.calls.CreateIamV2ApiKey
}

// CreateIamV2ApiKeyExecute mocks base method by wrapping the associated func.
func (m *APIKeysIamV2Api) CreateIamV2ApiKeyExecute(r github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.ApiCreateIamV2ApiKeyRequest) (github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.IamV2ApiKey, *net_http.Response, error) {
	m.lockCreateIamV2ApiKeyExecute.Lock()
	defer m.lockCreateIamV2ApiKeyExecute.Unlock()

	if m.CreateIamV2ApiKeyExecuteFunc == nil {
		panic("mocker: APIKeysIamV2Api.CreateIamV2ApiKeyExecuteFunc is nil but APIKeysIamV2Api.CreateIamV2ApiKeyExecute was called.")
	}

	call := struct {
		R github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.ApiCreateIamV2ApiKeyRequest
	}{
		R: r,
	}

	m.calls.CreateIamV2ApiKeyExecute = append(m.calls.CreateIamV2ApiKeyExecute, call)

	return m.CreateIamV2ApiKeyExecuteFunc(r)
}

// CreateIamV2ApiKeyExecuteCalled returns true if CreateIamV2ApiKeyExecute was called at least once.
func (m *APIKeysIamV2Api) CreateIamV2ApiKeyExecuteCalled() bool {
	m.lockCreateIamV2ApiKeyExecute.Lock()
	defer m.lockCreateIamV2ApiKeyExecute.Unlock()

	return len(m.calls.CreateIamV2ApiKeyExecute) > 0
}

// CreateIamV2ApiKeyExecuteCalls returns the calls made to CreateIamV2ApiKeyExecute.
func (m *APIKeysIamV2Api) CreateIamV2ApiKeyExecuteCalls() []struct {
	R github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.ApiCreateIamV2ApiKeyRequest
} {
	m.lockCreateIamV2ApiKeyExecute.Lock()
	defer m.lockCreateIamV2ApiKeyExecute.Unlock()

	return m.calls.CreateIamV2ApiKeyExecute
}

// DeleteIamV2ApiKey mocks base method by wrapping the associated func.
func (m *APIKeysIamV2Api) DeleteIamV2ApiKey(ctx context.Context, id string) github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.ApiDeleteIamV2ApiKeyRequest {
	m.lockDeleteIamV2ApiKey.Lock()
	defer m.lockDeleteIamV2ApiKey.Unlock()

	if m.DeleteIamV2ApiKeyFunc == nil {
		panic("mocker: APIKeysIamV2Api.DeleteIamV2ApiKeyFunc is nil but APIKeysIamV2Api.DeleteIamV2ApiKey was called.")
	}

	call := struct {
		Ctx context.Context
		Id  string
	}{
		Ctx: ctx,
		Id:  id,
	}

	m.calls.DeleteIamV2ApiKey = append(m.calls.DeleteIamV2ApiKey, call)

	return m.DeleteIamV2ApiKeyFunc(ctx, id)
}

// DeleteIamV2ApiKeyCalled returns true if DeleteIamV2ApiKey was called at least once.
func (m *APIKeysIamV2Api) DeleteIamV2ApiKeyCalled() bool {
	m.lockDeleteIamV2ApiKey.Lock()
	defer m.lockDeleteIamV2ApiKey.Unlock()

	return len(m.calls.DeleteIamV2ApiKey) > 0
}

// DeleteIamV2ApiKeyCalls returns the calls made to DeleteIamV2ApiKey.
func (m *APIKeysIamV2Api) DeleteIamV2ApiKeyCalls() []struct {
	Ctx context.Context
	Id  string
} {
	m.lockDeleteIamV2ApiKey.Lock()
	defer m.lockDeleteIamV2ApiKey.Unlock()

	return m.calls.DeleteIamV2ApiKey
}

// DeleteIamV2ApiKeyExecute mocks base method by wrapping the associated func.
func (m *APIKeysIamV2Api) DeleteIamV2ApiKeyExecute(r github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.ApiDeleteIamV2ApiKeyRequest) (*net_http.Response, error) {
	m.lockDeleteIamV2ApiKeyExecute.Lock()
	defer m.lockDeleteIamV2ApiKeyExecute.Unlock()

	if m.DeleteIamV2ApiKeyExecuteFunc == nil {
		panic("mocker: APIKeysIamV2Api.DeleteIamV2ApiKeyExecuteFunc is nil but APIKeysIamV2Api.DeleteIamV2ApiKeyExecute was called.")
	}

	call := struct {
		R github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.ApiDeleteIamV2ApiKeyRequest
	}{
		R: r,
	}

	m.calls.DeleteIamV2ApiKeyExecute = append(m.calls.DeleteIamV2ApiKeyExecute, call)

	return m.DeleteIamV2ApiKeyExecuteFunc(r)
}

// DeleteIamV2ApiKeyExecuteCalled returns true if DeleteIamV2ApiKeyExecute was called at least once.
func (m *APIKeysIamV2Api) DeleteIamV2ApiKeyExecuteCalled() bool {
	m.lockDeleteIamV2ApiKeyExecute.Lock()
	defer m.lockDeleteIamV2ApiKeyExecute.Unlock()

	return len(m.calls.DeleteIamV2ApiKeyExecute) > 0
}

// DeleteIamV2ApiKeyExecuteCalls returns the calls made to DeleteIamV2ApiKeyExecute.
func (m *APIKeysIamV2Api) DeleteIamV2ApiKeyExecuteCalls() []struct {
	R github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.ApiDeleteIamV2ApiKeyRequest
} {
	m.lockDeleteIamV2ApiKeyExecute.Lock()
	defer m.lockDeleteIamV2ApiKeyExecute.Unlock()

	return m.calls.DeleteIamV2ApiKeyExecute
}

// GetIamV2ApiKey mocks base method by wrapping the associated func.
func (m *APIKeysIamV2Api) GetIamV2ApiKey(ctx context.Context, id string) github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.ApiGetIamV2ApiKeyRequest {
	m.lockGetIamV2ApiKey.Lock()
	defer m.lockGetIamV2ApiKey.Unlock()

	if m.GetIamV2ApiKeyFunc == nil {
		panic("mocker: APIKeysIamV2Api.GetIamV2ApiKeyFunc is nil but APIKeysIamV2Api.GetIamV2ApiKey was called.")
	}

	call := struct {
		Ctx context.Context
		Id  string
	}{
		Ctx: ctx,
		Id:  id,
	}

	m.calls.GetIamV2ApiKey = append(m.calls.GetIamV2ApiKey, call)

	return m.GetIamV2ApiKeyFunc(ctx, id)
}

// GetIamV2ApiKeyCalled returns true if GetIamV2ApiKey was called at least once.
func (m *APIKeysIamV2Api) GetIamV2ApiKeyCalled() bool {
	m.lockGetIamV2ApiKey.Lock()
	defer m.lockGetIamV2ApiKey.Unlock()

	return len(m.calls.GetIamV2ApiKey) > 0
}

// GetIamV2ApiKeyCalls returns the calls made to GetIamV2ApiKey.
func (m *APIKeysIamV2Api) GetIamV2ApiKeyCalls() []struct {
	Ctx context.Context
	Id  string
} {
	m.lockGetIamV2ApiKey.Lock()
	defer m.lockGetIamV2ApiKey.Unlock()

	return m.calls.GetIamV2ApiKey
}

// GetIamV2ApiKeyExecute mocks base method by wrapping the associated func.
func (m *APIKeysIamV2Api) GetIamV2ApiKeyExecute(r github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.ApiGetIamV2ApiKeyRequest) (github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.IamV2ApiKey, *net_http.Response, error) {
	m.lockGetIamV2ApiKeyExecute.Lock()
	defer m.lockGetIamV2ApiKeyExecute.Unlock()

	if m.GetIamV2ApiKeyExecuteFunc == nil {
		panic("mocker: APIKeysIamV2Api.GetIamV2ApiKeyExecuteFunc is nil but APIKeysIamV2Api.GetIamV2ApiKeyExecute was called.")
	}

	call := struct {
		R github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.ApiGetIamV2ApiKeyRequest
	}{
		R: r,
	}

	m.calls.GetIamV2ApiKeyExecute = append(m.calls.GetIamV2ApiKeyExecute, call)

	return m.GetIamV2ApiKeyExecuteFunc(r)
}

// GetIamV2ApiKeyExecuteCalled returns true if GetIamV2ApiKeyExecute was called at least once.
func (m *APIKeysIamV2Api) GetIamV2ApiKeyExecuteCalled() bool {
	m.lockGetIamV2ApiKeyExecute.Lock()
	defer m.lockGetIamV2ApiKeyExecute.Unlock()

	return len(m.calls.GetIamV2ApiKeyExecute) > 0
}

// GetIamV2ApiKeyExecuteCalls returns the calls made to GetIamV2ApiKeyExecute.
func (m *APIKeysIamV2Api) GetIamV2ApiKeyExecuteCalls() []struct {
	R github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.ApiGetIamV2ApiKeyRequest
} {
	m.lockGetIamV2ApiKeyExecute.Lock()
	defer m.lockGetIamV2ApiKeyExecute.Unlock()

	return m.calls.GetIamV2ApiKeyExecute
}

// ListIamV2ApiKeys mocks base method by wrapping the associated func.
func (m *APIKeysIamV2Api) ListIamV2ApiKeys(ctx context.Context) github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.ApiListIamV2ApiKeysRequest {
	m.lockListIamV2ApiKeys.Lock()
	defer m.lockListIamV2ApiKeys.Unlock()

	if m.ListIamV2ApiKeysFunc == nil {
		panic("mocker: APIKeysIamV2Api.ListIamV2ApiKeysFunc is nil but APIKeysIamV2Api.ListIamV2ApiKeys was called.")
	}

	call := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}

	m.calls.ListIamV2ApiKeys = append(m.calls.ListIamV2ApiKeys, call)

	return m.ListIamV2ApiKeysFunc(ctx)
}

// ListIamV2ApiKeysCalled returns true if ListIamV2ApiKeys was called at least once.
func (m *APIKeysIamV2Api) ListIamV2ApiKeysCalled() bool {
	m.lockListIamV2ApiKeys.Lock()
	defer m.lockListIamV2ApiKeys.Unlock()

	return len(m.calls.ListIamV2ApiKeys) > 0
}

// ListIamV2ApiKeysCalls returns the calls made to ListIamV2ApiKeys.
func (m *APIKeysIamV2Api) ListIamV2ApiKeysCalls() []struct {
	Ctx context.Context
} {
	m.lockListIamV2ApiKeys.Lock()
	defer m.lockListIamV2ApiKeys.Unlock()

	return m.calls.ListIamV2ApiKeys
}

// ListIamV2ApiKeysExecute mocks base method by wrapping the associated func.
func (m *APIKeysIamV2Api) ListIamV2ApiKeysExecute(r github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.ApiListIamV2ApiKeysRequest) (github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.IamV2ApiKeyList, *net_http.Response, error) {
	m.lockListIamV2ApiKeysExecute.Lock()
	defer m.lockListIamV2ApiKeysExecute.Unlock()

	if m.ListIamV2ApiKeysExecuteFunc == nil {
		panic("mocker: APIKeysIamV2Api.ListIamV2ApiKeysExecuteFunc is nil but APIKeysIamV2Api.ListIamV2ApiKeysExecute was called.")
	}

	call := struct {
		R github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.ApiListIamV2ApiKeysRequest
	}{
		R: r,
	}

	m.calls.ListIamV2ApiKeysExecute = append(m.calls.ListIamV2ApiKeysExecute, call)

	return m.ListIamV2ApiKeysExecuteFunc(r)
}

// ListIamV2ApiKeysExecuteCalled returns true if ListIamV2ApiKeysExecute was called at least once.
func (m *APIKeysIamV2Api) ListIamV2ApiKeysExecuteCalled() bool {
	m.lockListIamV2ApiKeysExecute.Lock()
	defer m.lockListIamV2ApiKeysExecute.Unlock()

	return len(m.calls.ListIamV2ApiKeysExecute) > 0
}

// ListIamV2ApiKeysExecuteCalls returns the calls made to ListIamV2ApiKeysExecute.
func (m *APIKeysIamV2Api) ListIamV2ApiKeysExecuteCalls() []struct {
	R github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.ApiListIamV2ApiKeysRequest
} {
	m.lockListIamV2ApiKeysExecute.Lock()
	defer m.lockListIamV2ApiKeysExecute.Unlock()

	return m.calls.ListIamV2ApiKeysExecute
}

// UpdateIamV2ApiKey mocks base method by wrapping the associated func.
func (m *APIKeysIamV2Api) UpdateIamV2ApiKey(ctx context.Context, id string) github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.ApiUpdateIamV2ApiKeyRequest {
	m.lockUpdateIamV2ApiKey.Lock()
	defer m.lockUpdateIamV2ApiKey.Unlock()

	if m.UpdateIamV2ApiKeyFunc == nil {
		panic("mocker: APIKeysIamV2Api.UpdateIamV2ApiKeyFunc is nil but APIKeysIamV2Api.UpdateIamV2ApiKey was called.")
	}

	call := struct {
		Ctx context.Context
		Id  string
	}{
		Ctx: ctx,
		Id:  id,
	}

	m.calls.UpdateIamV2ApiKey = append(m.calls.UpdateIamV2ApiKey, call)

	return m.UpdateIamV2ApiKeyFunc(ctx, id)
}

// UpdateIamV2ApiKeyCalled returns true if UpdateIamV2ApiKey was called at least once.
func (m *APIKeysIamV2Api) UpdateIamV2ApiKeyCalled() bool {
	m.lockUpdateIamV2ApiKey.Lock()
	defer m.lockUpdateIamV2ApiKey.Unlock()

	return len(m.calls.UpdateIamV2ApiKey) > 0
}

// UpdateIamV2ApiKeyCalls returns the calls made to UpdateIamV2ApiKey.
func (m *APIKeysIamV2Api) UpdateIamV2ApiKeyCalls() []struct {
	Ctx context.Context
	Id  string
} {
	m.lockUpdateIamV2ApiKey.Lock()
	defer m.lockUpdateIamV2ApiKey.Unlock()

	return m.calls.UpdateIamV2ApiKey
}

// UpdateIamV2ApiKeyExecute mocks base method by wrapping the associated func.
func (m *APIKeysIamV2Api) UpdateIamV2ApiKeyExecute(r github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.ApiUpdateIamV2ApiKeyRequest) (github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.IamV2ApiKey, *net_http.Response, error) {
	m.lockUpdateIamV2ApiKeyExecute.Lock()
	defer m.lockUpdateIamV2ApiKeyExecute.Unlock()

	if m.UpdateIamV2ApiKeyExecuteFunc == nil {
		panic("mocker: APIKeysIamV2Api.UpdateIamV2ApiKeyExecuteFunc is nil but APIKeysIamV2Api.UpdateIamV2ApiKeyExecute was called.")
	}

	call := struct {
		R github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.ApiUpdateIamV2ApiKeyRequest
	}{
		R: r,
	}

	m.calls.UpdateIamV2ApiKeyExecute = append(m.calls.UpdateIamV2ApiKeyExecute, call)

	return m.UpdateIamV2ApiKeyExecuteFunc(r)
}

// UpdateIamV2ApiKeyExecuteCalled returns true if UpdateIamV2ApiKeyExecute was called at least once.
func (m *APIKeysIamV2Api) UpdateIamV2ApiKeyExecuteCalled() bool {
	m.lockUpdateIamV2ApiKeyExecute.Lock()
	defer m.lockUpdateIamV2ApiKeyExecute.Unlock()

	return len(m.calls.UpdateIamV2ApiKeyExecute) > 0
}

// UpdateIamV2ApiKeyExecuteCalls returns the calls made to UpdateIamV2ApiKeyExecute.
func (m *APIKeysIamV2Api) UpdateIamV2ApiKeyExecuteCalls() []struct {
	R github_com_confluentinc_ccloud_sdk_go_v2_apikeys_v2.ApiUpdateIamV2ApiKeyRequest
} {
	m.lockUpdateIamV2ApiKeyExecute.Lock()
	defer m.lockUpdateIamV2ApiKeyExecute.Unlock()

	return m.calls.UpdateIamV2ApiKeyExecute
}

// Reset resets the calls made to the mocked methods.
func (m *APIKeysIamV2Api) Reset() {
	m.lockCreateIamV2ApiKey.Lock()
	m.calls.CreateIamV2ApiKey = nil
	m.lockCreateIamV2ApiKey.Unlock()
	m.lockCreateIamV2ApiKeyExecute.Lock()
	m.calls.CreateIamV2ApiKeyExecute = nil
	m.lockCreateIamV2ApiKeyExecute.Unlock()
	m.lockDeleteIamV2ApiKey.Lock()
	m.calls.DeleteIamV2ApiKey = nil
	m.lockDeleteIamV2ApiKey.Unlock()
	m.lockDeleteIamV2ApiKeyExecute.Lock()
	m.calls.DeleteIamV2ApiKeyExecute = nil
	m.lockDeleteIamV2ApiKeyExecute.Unlock()
	m.lockGetIamV2ApiKey.Lock()
	m.calls.GetIamV2ApiKey = nil
	m.lockGetIamV2ApiKey.Unlock()
	m.lockGetIamV2ApiKeyExecute.Lock()
	m.calls.GetIamV2ApiKeyExecute = nil
	m.lockGetIamV2ApiKeyExecute.Unlock()
	m.lockListIamV2ApiKeys.Lock()
	m.calls.ListIamV2ApiKeys = nil
	m.lockListIamV2ApiKeys.Unlock()
	m.lockListIamV2ApiKeysExecute.Lock()
	m.calls.ListIamV2ApiKeysExecute = nil
	m.lockListIamV2ApiKeysExecute.Unlock()
	m.lockUpdateIamV2ApiKey.Lock()
	m.calls.UpdateIamV2ApiKey = nil
	m.lockUpdateIamV2ApiKey.Unlock()
	m.lockUpdateIamV2ApiKeyExecute.Lock()
	m.calls.UpdateIamV2ApiKeyExecute = nil
	m.lockUpdateIamV2ApiKeyExecute.Unlock()
}