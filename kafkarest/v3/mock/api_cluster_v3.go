// Code generated by mocker. DO NOT EDIT.
// github.com/travisjeffery/mocker
// Source: api_cluster_v3.go

package mock

import (
	context "context"
	net_http "net/http"
	sync "sync"

	github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3 "github.com/confluentinc/ccloud-sdk-go-v2/kafkarest/v3"
)

// ClusterV3Api is a mock of ClusterV3Api interface
type ClusterV3Api struct {
	lockGetKafkaCluster sync.Mutex
	GetKafkaClusterFunc func(ctx context.Context, clusterId string) github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.ApiGetKafkaClusterRequest

	lockGetKafkaClusterExecute sync.Mutex
	GetKafkaClusterExecuteFunc func(r github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.ApiGetKafkaClusterRequest) (github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.ClusterData, *net_http.Response, error)

	calls struct {
		GetKafkaCluster []struct {
			Ctx       context.Context
			ClusterId string
		}
		GetKafkaClusterExecute []struct {
			R github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.ApiGetKafkaClusterRequest
		}
	}
}

// GetKafkaCluster mocks base method by wrapping the associated func.
func (m *ClusterV3Api) GetKafkaCluster(ctx context.Context, clusterId string) github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.ApiGetKafkaClusterRequest {
	m.lockGetKafkaCluster.Lock()
	defer m.lockGetKafkaCluster.Unlock()

	if m.GetKafkaClusterFunc == nil {
		panic("mocker: ClusterV3Api.GetKafkaClusterFunc is nil but ClusterV3Api.GetKafkaCluster was called.")
	}

	call := struct {
		Ctx       context.Context
		ClusterId string
	}{
		Ctx:       ctx,
		ClusterId: clusterId,
	}

	m.calls.GetKafkaCluster = append(m.calls.GetKafkaCluster, call)

	return m.GetKafkaClusterFunc(ctx, clusterId)
}

// GetKafkaClusterCalled returns true if GetKafkaCluster was called at least once.
func (m *ClusterV3Api) GetKafkaClusterCalled() bool {
	m.lockGetKafkaCluster.Lock()
	defer m.lockGetKafkaCluster.Unlock()

	return len(m.calls.GetKafkaCluster) > 0
}

// GetKafkaClusterCalls returns the calls made to GetKafkaCluster.
func (m *ClusterV3Api) GetKafkaClusterCalls() []struct {
	Ctx       context.Context
	ClusterId string
} {
	m.lockGetKafkaCluster.Lock()
	defer m.lockGetKafkaCluster.Unlock()

	return m.calls.GetKafkaCluster
}

// GetKafkaClusterExecute mocks base method by wrapping the associated func.
func (m *ClusterV3Api) GetKafkaClusterExecute(r github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.ApiGetKafkaClusterRequest) (github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.ClusterData, *net_http.Response, error) {
	m.lockGetKafkaClusterExecute.Lock()
	defer m.lockGetKafkaClusterExecute.Unlock()

	if m.GetKafkaClusterExecuteFunc == nil {
		panic("mocker: ClusterV3Api.GetKafkaClusterExecuteFunc is nil but ClusterV3Api.GetKafkaClusterExecute was called.")
	}

	call := struct {
		R github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.ApiGetKafkaClusterRequest
	}{
		R: r,
	}

	m.calls.GetKafkaClusterExecute = append(m.calls.GetKafkaClusterExecute, call)

	return m.GetKafkaClusterExecuteFunc(r)
}

// GetKafkaClusterExecuteCalled returns true if GetKafkaClusterExecute was called at least once.
func (m *ClusterV3Api) GetKafkaClusterExecuteCalled() bool {
	m.lockGetKafkaClusterExecute.Lock()
	defer m.lockGetKafkaClusterExecute.Unlock()

	return len(m.calls.GetKafkaClusterExecute) > 0
}

// GetKafkaClusterExecuteCalls returns the calls made to GetKafkaClusterExecute.
func (m *ClusterV3Api) GetKafkaClusterExecuteCalls() []struct {
	R github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.ApiGetKafkaClusterRequest
} {
	m.lockGetKafkaClusterExecute.Lock()
	defer m.lockGetKafkaClusterExecute.Unlock()

	return m.calls.GetKafkaClusterExecute
}

// Reset resets the calls made to the mocked methods.
func (m *ClusterV3Api) Reset() {
	m.lockGetKafkaCluster.Lock()
	m.calls.GetKafkaCluster = nil
	m.lockGetKafkaCluster.Unlock()
	m.lockGetKafkaClusterExecute.Lock()
	m.calls.GetKafkaClusterExecute = nil
	m.lockGetKafkaClusterExecute.Unlock()
}
