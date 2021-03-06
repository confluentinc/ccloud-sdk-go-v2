// Code generated by mocker. DO NOT EDIT.
// github.com/travisjeffery/mocker
// Source: api_clusters_cmk_v2.go

package mock

import (
	context "context"
	net_http "net/http"
	sync "sync"

	github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2 "github.com/confluentinc/ccloud-sdk-go-v2/cmk/v2"
)

// ClustersCmkV2Api is a mock of ClustersCmkV2Api interface
type ClustersCmkV2Api struct {
	lockCreateCmkV2Cluster sync.Mutex
	CreateCmkV2ClusterFunc func(ctx context.Context) github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.ApiCreateCmkV2ClusterRequest

	lockCreateCmkV2ClusterExecute sync.Mutex
	CreateCmkV2ClusterExecuteFunc func(r github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.ApiCreateCmkV2ClusterRequest) (github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.CmkV2Cluster, *net_http.Response, error)

	lockDeleteCmkV2Cluster sync.Mutex
	DeleteCmkV2ClusterFunc func(ctx context.Context, id string) github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.ApiDeleteCmkV2ClusterRequest

	lockDeleteCmkV2ClusterExecute sync.Mutex
	DeleteCmkV2ClusterExecuteFunc func(r github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.ApiDeleteCmkV2ClusterRequest) (*net_http.Response, error)

	lockGetCmkV2Cluster sync.Mutex
	GetCmkV2ClusterFunc func(ctx context.Context, id string) github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.ApiGetCmkV2ClusterRequest

	lockGetCmkV2ClusterExecute sync.Mutex
	GetCmkV2ClusterExecuteFunc func(r github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.ApiGetCmkV2ClusterRequest) (github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.CmkV2Cluster, *net_http.Response, error)

	lockListCmkV2Clusters sync.Mutex
	ListCmkV2ClustersFunc func(ctx context.Context) github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.ApiListCmkV2ClustersRequest

	lockListCmkV2ClustersExecute sync.Mutex
	ListCmkV2ClustersExecuteFunc func(r github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.ApiListCmkV2ClustersRequest) (github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.CmkV2ClusterList, *net_http.Response, error)

	lockUpdateCmkV2Cluster sync.Mutex
	UpdateCmkV2ClusterFunc func(ctx context.Context, id string) github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.ApiUpdateCmkV2ClusterRequest

	lockUpdateCmkV2ClusterExecute sync.Mutex
	UpdateCmkV2ClusterExecuteFunc func(r github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.ApiUpdateCmkV2ClusterRequest) (github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.CmkV2Cluster, *net_http.Response, error)

	calls struct {
		CreateCmkV2Cluster []struct {
			Ctx context.Context
		}
		CreateCmkV2ClusterExecute []struct {
			R github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.ApiCreateCmkV2ClusterRequest
		}
		DeleteCmkV2Cluster []struct {
			Ctx context.Context
			Id  string
		}
		DeleteCmkV2ClusterExecute []struct {
			R github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.ApiDeleteCmkV2ClusterRequest
		}
		GetCmkV2Cluster []struct {
			Ctx context.Context
			Id  string
		}
		GetCmkV2ClusterExecute []struct {
			R github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.ApiGetCmkV2ClusterRequest
		}
		ListCmkV2Clusters []struct {
			Ctx context.Context
		}
		ListCmkV2ClustersExecute []struct {
			R github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.ApiListCmkV2ClustersRequest
		}
		UpdateCmkV2Cluster []struct {
			Ctx context.Context
			Id  string
		}
		UpdateCmkV2ClusterExecute []struct {
			R github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.ApiUpdateCmkV2ClusterRequest
		}
	}
}

// CreateCmkV2Cluster mocks base method by wrapping the associated func.
func (m *ClustersCmkV2Api) CreateCmkV2Cluster(ctx context.Context) github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.ApiCreateCmkV2ClusterRequest {
	m.lockCreateCmkV2Cluster.Lock()
	defer m.lockCreateCmkV2Cluster.Unlock()

	if m.CreateCmkV2ClusterFunc == nil {
		panic("mocker: ClustersCmkV2Api.CreateCmkV2ClusterFunc is nil but ClustersCmkV2Api.CreateCmkV2Cluster was called.")
	}

	call := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}

	m.calls.CreateCmkV2Cluster = append(m.calls.CreateCmkV2Cluster, call)

	return m.CreateCmkV2ClusterFunc(ctx)
}

// CreateCmkV2ClusterCalled returns true if CreateCmkV2Cluster was called at least once.
func (m *ClustersCmkV2Api) CreateCmkV2ClusterCalled() bool {
	m.lockCreateCmkV2Cluster.Lock()
	defer m.lockCreateCmkV2Cluster.Unlock()

	return len(m.calls.CreateCmkV2Cluster) > 0
}

// CreateCmkV2ClusterCalls returns the calls made to CreateCmkV2Cluster.
func (m *ClustersCmkV2Api) CreateCmkV2ClusterCalls() []struct {
	Ctx context.Context
} {
	m.lockCreateCmkV2Cluster.Lock()
	defer m.lockCreateCmkV2Cluster.Unlock()

	return m.calls.CreateCmkV2Cluster
}

// CreateCmkV2ClusterExecute mocks base method by wrapping the associated func.
func (m *ClustersCmkV2Api) CreateCmkV2ClusterExecute(r github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.ApiCreateCmkV2ClusterRequest) (github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.CmkV2Cluster, *net_http.Response, error) {
	m.lockCreateCmkV2ClusterExecute.Lock()
	defer m.lockCreateCmkV2ClusterExecute.Unlock()

	if m.CreateCmkV2ClusterExecuteFunc == nil {
		panic("mocker: ClustersCmkV2Api.CreateCmkV2ClusterExecuteFunc is nil but ClustersCmkV2Api.CreateCmkV2ClusterExecute was called.")
	}

	call := struct {
		R github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.ApiCreateCmkV2ClusterRequest
	}{
		R: r,
	}

	m.calls.CreateCmkV2ClusterExecute = append(m.calls.CreateCmkV2ClusterExecute, call)

	return m.CreateCmkV2ClusterExecuteFunc(r)
}

// CreateCmkV2ClusterExecuteCalled returns true if CreateCmkV2ClusterExecute was called at least once.
func (m *ClustersCmkV2Api) CreateCmkV2ClusterExecuteCalled() bool {
	m.lockCreateCmkV2ClusterExecute.Lock()
	defer m.lockCreateCmkV2ClusterExecute.Unlock()

	return len(m.calls.CreateCmkV2ClusterExecute) > 0
}

// CreateCmkV2ClusterExecuteCalls returns the calls made to CreateCmkV2ClusterExecute.
func (m *ClustersCmkV2Api) CreateCmkV2ClusterExecuteCalls() []struct {
	R github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.ApiCreateCmkV2ClusterRequest
} {
	m.lockCreateCmkV2ClusterExecute.Lock()
	defer m.lockCreateCmkV2ClusterExecute.Unlock()

	return m.calls.CreateCmkV2ClusterExecute
}

// DeleteCmkV2Cluster mocks base method by wrapping the associated func.
func (m *ClustersCmkV2Api) DeleteCmkV2Cluster(ctx context.Context, id string) github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.ApiDeleteCmkV2ClusterRequest {
	m.lockDeleteCmkV2Cluster.Lock()
	defer m.lockDeleteCmkV2Cluster.Unlock()

	if m.DeleteCmkV2ClusterFunc == nil {
		panic("mocker: ClustersCmkV2Api.DeleteCmkV2ClusterFunc is nil but ClustersCmkV2Api.DeleteCmkV2Cluster was called.")
	}

	call := struct {
		Ctx context.Context
		Id  string
	}{
		Ctx: ctx,
		Id:  id,
	}

	m.calls.DeleteCmkV2Cluster = append(m.calls.DeleteCmkV2Cluster, call)

	return m.DeleteCmkV2ClusterFunc(ctx, id)
}

// DeleteCmkV2ClusterCalled returns true if DeleteCmkV2Cluster was called at least once.
func (m *ClustersCmkV2Api) DeleteCmkV2ClusterCalled() bool {
	m.lockDeleteCmkV2Cluster.Lock()
	defer m.lockDeleteCmkV2Cluster.Unlock()

	return len(m.calls.DeleteCmkV2Cluster) > 0
}

// DeleteCmkV2ClusterCalls returns the calls made to DeleteCmkV2Cluster.
func (m *ClustersCmkV2Api) DeleteCmkV2ClusterCalls() []struct {
	Ctx context.Context
	Id  string
} {
	m.lockDeleteCmkV2Cluster.Lock()
	defer m.lockDeleteCmkV2Cluster.Unlock()

	return m.calls.DeleteCmkV2Cluster
}

// DeleteCmkV2ClusterExecute mocks base method by wrapping the associated func.
func (m *ClustersCmkV2Api) DeleteCmkV2ClusterExecute(r github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.ApiDeleteCmkV2ClusterRequest) (*net_http.Response, error) {
	m.lockDeleteCmkV2ClusterExecute.Lock()
	defer m.lockDeleteCmkV2ClusterExecute.Unlock()

	if m.DeleteCmkV2ClusterExecuteFunc == nil {
		panic("mocker: ClustersCmkV2Api.DeleteCmkV2ClusterExecuteFunc is nil but ClustersCmkV2Api.DeleteCmkV2ClusterExecute was called.")
	}

	call := struct {
		R github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.ApiDeleteCmkV2ClusterRequest
	}{
		R: r,
	}

	m.calls.DeleteCmkV2ClusterExecute = append(m.calls.DeleteCmkV2ClusterExecute, call)

	return m.DeleteCmkV2ClusterExecuteFunc(r)
}

// DeleteCmkV2ClusterExecuteCalled returns true if DeleteCmkV2ClusterExecute was called at least once.
func (m *ClustersCmkV2Api) DeleteCmkV2ClusterExecuteCalled() bool {
	m.lockDeleteCmkV2ClusterExecute.Lock()
	defer m.lockDeleteCmkV2ClusterExecute.Unlock()

	return len(m.calls.DeleteCmkV2ClusterExecute) > 0
}

// DeleteCmkV2ClusterExecuteCalls returns the calls made to DeleteCmkV2ClusterExecute.
func (m *ClustersCmkV2Api) DeleteCmkV2ClusterExecuteCalls() []struct {
	R github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.ApiDeleteCmkV2ClusterRequest
} {
	m.lockDeleteCmkV2ClusterExecute.Lock()
	defer m.lockDeleteCmkV2ClusterExecute.Unlock()

	return m.calls.DeleteCmkV2ClusterExecute
}

// GetCmkV2Cluster mocks base method by wrapping the associated func.
func (m *ClustersCmkV2Api) GetCmkV2Cluster(ctx context.Context, id string) github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.ApiGetCmkV2ClusterRequest {
	m.lockGetCmkV2Cluster.Lock()
	defer m.lockGetCmkV2Cluster.Unlock()

	if m.GetCmkV2ClusterFunc == nil {
		panic("mocker: ClustersCmkV2Api.GetCmkV2ClusterFunc is nil but ClustersCmkV2Api.GetCmkV2Cluster was called.")
	}

	call := struct {
		Ctx context.Context
		Id  string
	}{
		Ctx: ctx,
		Id:  id,
	}

	m.calls.GetCmkV2Cluster = append(m.calls.GetCmkV2Cluster, call)

	return m.GetCmkV2ClusterFunc(ctx, id)
}

// GetCmkV2ClusterCalled returns true if GetCmkV2Cluster was called at least once.
func (m *ClustersCmkV2Api) GetCmkV2ClusterCalled() bool {
	m.lockGetCmkV2Cluster.Lock()
	defer m.lockGetCmkV2Cluster.Unlock()

	return len(m.calls.GetCmkV2Cluster) > 0
}

// GetCmkV2ClusterCalls returns the calls made to GetCmkV2Cluster.
func (m *ClustersCmkV2Api) GetCmkV2ClusterCalls() []struct {
	Ctx context.Context
	Id  string
} {
	m.lockGetCmkV2Cluster.Lock()
	defer m.lockGetCmkV2Cluster.Unlock()

	return m.calls.GetCmkV2Cluster
}

// GetCmkV2ClusterExecute mocks base method by wrapping the associated func.
func (m *ClustersCmkV2Api) GetCmkV2ClusterExecute(r github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.ApiGetCmkV2ClusterRequest) (github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.CmkV2Cluster, *net_http.Response, error) {
	m.lockGetCmkV2ClusterExecute.Lock()
	defer m.lockGetCmkV2ClusterExecute.Unlock()

	if m.GetCmkV2ClusterExecuteFunc == nil {
		panic("mocker: ClustersCmkV2Api.GetCmkV2ClusterExecuteFunc is nil but ClustersCmkV2Api.GetCmkV2ClusterExecute was called.")
	}

	call := struct {
		R github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.ApiGetCmkV2ClusterRequest
	}{
		R: r,
	}

	m.calls.GetCmkV2ClusterExecute = append(m.calls.GetCmkV2ClusterExecute, call)

	return m.GetCmkV2ClusterExecuteFunc(r)
}

// GetCmkV2ClusterExecuteCalled returns true if GetCmkV2ClusterExecute was called at least once.
func (m *ClustersCmkV2Api) GetCmkV2ClusterExecuteCalled() bool {
	m.lockGetCmkV2ClusterExecute.Lock()
	defer m.lockGetCmkV2ClusterExecute.Unlock()

	return len(m.calls.GetCmkV2ClusterExecute) > 0
}

// GetCmkV2ClusterExecuteCalls returns the calls made to GetCmkV2ClusterExecute.
func (m *ClustersCmkV2Api) GetCmkV2ClusterExecuteCalls() []struct {
	R github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.ApiGetCmkV2ClusterRequest
} {
	m.lockGetCmkV2ClusterExecute.Lock()
	defer m.lockGetCmkV2ClusterExecute.Unlock()

	return m.calls.GetCmkV2ClusterExecute
}

// ListCmkV2Clusters mocks base method by wrapping the associated func.
func (m *ClustersCmkV2Api) ListCmkV2Clusters(ctx context.Context) github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.ApiListCmkV2ClustersRequest {
	m.lockListCmkV2Clusters.Lock()
	defer m.lockListCmkV2Clusters.Unlock()

	if m.ListCmkV2ClustersFunc == nil {
		panic("mocker: ClustersCmkV2Api.ListCmkV2ClustersFunc is nil but ClustersCmkV2Api.ListCmkV2Clusters was called.")
	}

	call := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}

	m.calls.ListCmkV2Clusters = append(m.calls.ListCmkV2Clusters, call)

	return m.ListCmkV2ClustersFunc(ctx)
}

// ListCmkV2ClustersCalled returns true if ListCmkV2Clusters was called at least once.
func (m *ClustersCmkV2Api) ListCmkV2ClustersCalled() bool {
	m.lockListCmkV2Clusters.Lock()
	defer m.lockListCmkV2Clusters.Unlock()

	return len(m.calls.ListCmkV2Clusters) > 0
}

// ListCmkV2ClustersCalls returns the calls made to ListCmkV2Clusters.
func (m *ClustersCmkV2Api) ListCmkV2ClustersCalls() []struct {
	Ctx context.Context
} {
	m.lockListCmkV2Clusters.Lock()
	defer m.lockListCmkV2Clusters.Unlock()

	return m.calls.ListCmkV2Clusters
}

// ListCmkV2ClustersExecute mocks base method by wrapping the associated func.
func (m *ClustersCmkV2Api) ListCmkV2ClustersExecute(r github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.ApiListCmkV2ClustersRequest) (github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.CmkV2ClusterList, *net_http.Response, error) {
	m.lockListCmkV2ClustersExecute.Lock()
	defer m.lockListCmkV2ClustersExecute.Unlock()

	if m.ListCmkV2ClustersExecuteFunc == nil {
		panic("mocker: ClustersCmkV2Api.ListCmkV2ClustersExecuteFunc is nil but ClustersCmkV2Api.ListCmkV2ClustersExecute was called.")
	}

	call := struct {
		R github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.ApiListCmkV2ClustersRequest
	}{
		R: r,
	}

	m.calls.ListCmkV2ClustersExecute = append(m.calls.ListCmkV2ClustersExecute, call)

	return m.ListCmkV2ClustersExecuteFunc(r)
}

// ListCmkV2ClustersExecuteCalled returns true if ListCmkV2ClustersExecute was called at least once.
func (m *ClustersCmkV2Api) ListCmkV2ClustersExecuteCalled() bool {
	m.lockListCmkV2ClustersExecute.Lock()
	defer m.lockListCmkV2ClustersExecute.Unlock()

	return len(m.calls.ListCmkV2ClustersExecute) > 0
}

// ListCmkV2ClustersExecuteCalls returns the calls made to ListCmkV2ClustersExecute.
func (m *ClustersCmkV2Api) ListCmkV2ClustersExecuteCalls() []struct {
	R github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.ApiListCmkV2ClustersRequest
} {
	m.lockListCmkV2ClustersExecute.Lock()
	defer m.lockListCmkV2ClustersExecute.Unlock()

	return m.calls.ListCmkV2ClustersExecute
}

// UpdateCmkV2Cluster mocks base method by wrapping the associated func.
func (m *ClustersCmkV2Api) UpdateCmkV2Cluster(ctx context.Context, id string) github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.ApiUpdateCmkV2ClusterRequest {
	m.lockUpdateCmkV2Cluster.Lock()
	defer m.lockUpdateCmkV2Cluster.Unlock()

	if m.UpdateCmkV2ClusterFunc == nil {
		panic("mocker: ClustersCmkV2Api.UpdateCmkV2ClusterFunc is nil but ClustersCmkV2Api.UpdateCmkV2Cluster was called.")
	}

	call := struct {
		Ctx context.Context
		Id  string
	}{
		Ctx: ctx,
		Id:  id,
	}

	m.calls.UpdateCmkV2Cluster = append(m.calls.UpdateCmkV2Cluster, call)

	return m.UpdateCmkV2ClusterFunc(ctx, id)
}

// UpdateCmkV2ClusterCalled returns true if UpdateCmkV2Cluster was called at least once.
func (m *ClustersCmkV2Api) UpdateCmkV2ClusterCalled() bool {
	m.lockUpdateCmkV2Cluster.Lock()
	defer m.lockUpdateCmkV2Cluster.Unlock()

	return len(m.calls.UpdateCmkV2Cluster) > 0
}

// UpdateCmkV2ClusterCalls returns the calls made to UpdateCmkV2Cluster.
func (m *ClustersCmkV2Api) UpdateCmkV2ClusterCalls() []struct {
	Ctx context.Context
	Id  string
} {
	m.lockUpdateCmkV2Cluster.Lock()
	defer m.lockUpdateCmkV2Cluster.Unlock()

	return m.calls.UpdateCmkV2Cluster
}

// UpdateCmkV2ClusterExecute mocks base method by wrapping the associated func.
func (m *ClustersCmkV2Api) UpdateCmkV2ClusterExecute(r github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.ApiUpdateCmkV2ClusterRequest) (github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.CmkV2Cluster, *net_http.Response, error) {
	m.lockUpdateCmkV2ClusterExecute.Lock()
	defer m.lockUpdateCmkV2ClusterExecute.Unlock()

	if m.UpdateCmkV2ClusterExecuteFunc == nil {
		panic("mocker: ClustersCmkV2Api.UpdateCmkV2ClusterExecuteFunc is nil but ClustersCmkV2Api.UpdateCmkV2ClusterExecute was called.")
	}

	call := struct {
		R github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.ApiUpdateCmkV2ClusterRequest
	}{
		R: r,
	}

	m.calls.UpdateCmkV2ClusterExecute = append(m.calls.UpdateCmkV2ClusterExecute, call)

	return m.UpdateCmkV2ClusterExecuteFunc(r)
}

// UpdateCmkV2ClusterExecuteCalled returns true if UpdateCmkV2ClusterExecute was called at least once.
func (m *ClustersCmkV2Api) UpdateCmkV2ClusterExecuteCalled() bool {
	m.lockUpdateCmkV2ClusterExecute.Lock()
	defer m.lockUpdateCmkV2ClusterExecute.Unlock()

	return len(m.calls.UpdateCmkV2ClusterExecute) > 0
}

// UpdateCmkV2ClusterExecuteCalls returns the calls made to UpdateCmkV2ClusterExecute.
func (m *ClustersCmkV2Api) UpdateCmkV2ClusterExecuteCalls() []struct {
	R github_com_confluentinc_ccloud_sdk_go_v2_cmk_v2.ApiUpdateCmkV2ClusterRequest
} {
	m.lockUpdateCmkV2ClusterExecute.Lock()
	defer m.lockUpdateCmkV2ClusterExecute.Unlock()

	return m.calls.UpdateCmkV2ClusterExecute
}

// Reset resets the calls made to the mocked methods.
func (m *ClustersCmkV2Api) Reset() {
	m.lockCreateCmkV2Cluster.Lock()
	m.calls.CreateCmkV2Cluster = nil
	m.lockCreateCmkV2Cluster.Unlock()
	m.lockCreateCmkV2ClusterExecute.Lock()
	m.calls.CreateCmkV2ClusterExecute = nil
	m.lockCreateCmkV2ClusterExecute.Unlock()
	m.lockDeleteCmkV2Cluster.Lock()
	m.calls.DeleteCmkV2Cluster = nil
	m.lockDeleteCmkV2Cluster.Unlock()
	m.lockDeleteCmkV2ClusterExecute.Lock()
	m.calls.DeleteCmkV2ClusterExecute = nil
	m.lockDeleteCmkV2ClusterExecute.Unlock()
	m.lockGetCmkV2Cluster.Lock()
	m.calls.GetCmkV2Cluster = nil
	m.lockGetCmkV2Cluster.Unlock()
	m.lockGetCmkV2ClusterExecute.Lock()
	m.calls.GetCmkV2ClusterExecute = nil
	m.lockGetCmkV2ClusterExecute.Unlock()
	m.lockListCmkV2Clusters.Lock()
	m.calls.ListCmkV2Clusters = nil
	m.lockListCmkV2Clusters.Unlock()
	m.lockListCmkV2ClustersExecute.Lock()
	m.calls.ListCmkV2ClustersExecute = nil
	m.lockListCmkV2ClustersExecute.Unlock()
	m.lockUpdateCmkV2Cluster.Lock()
	m.calls.UpdateCmkV2Cluster = nil
	m.lockUpdateCmkV2Cluster.Unlock()
	m.lockUpdateCmkV2ClusterExecute.Lock()
	m.calls.UpdateCmkV2ClusterExecute = nil
	m.lockUpdateCmkV2ClusterExecute.Unlock()
}
