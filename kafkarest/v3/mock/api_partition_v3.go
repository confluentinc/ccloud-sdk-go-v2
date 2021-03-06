// Code generated by mocker. DO NOT EDIT.
// github.com/travisjeffery/mocker
// Source: api_partition_v3.go

package mock

import (
	context "context"
	net_http "net/http"
	sync "sync"

	github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3 "github.com/confluentinc/ccloud-sdk-go-v2/kafkarest/v3"
)

// PartitionV3Api is a mock of PartitionV3Api interface
type PartitionV3Api struct {
	lockGetKafkaConsumerLag sync.Mutex
	GetKafkaConsumerLagFunc func(ctx context.Context, clusterId, consumerGroupId, topicName string, partitionId int32) github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.ApiGetKafkaConsumerLagRequest

	lockGetKafkaConsumerLagExecute sync.Mutex
	GetKafkaConsumerLagExecuteFunc func(r github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.ApiGetKafkaConsumerLagRequest) (github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.ConsumerLagData, *net_http.Response, error)

	lockGetKafkaPartition sync.Mutex
	GetKafkaPartitionFunc func(ctx context.Context, clusterId, topicName string, partitionId int32) github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.ApiGetKafkaPartitionRequest

	lockGetKafkaPartitionExecute sync.Mutex
	GetKafkaPartitionExecuteFunc func(r github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.ApiGetKafkaPartitionRequest) (github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.PartitionData, *net_http.Response, error)

	lockListKafkaPartitions sync.Mutex
	ListKafkaPartitionsFunc func(ctx context.Context, clusterId, topicName string) github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.ApiListKafkaPartitionsRequest

	lockListKafkaPartitionsExecute sync.Mutex
	ListKafkaPartitionsExecuteFunc func(r github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.ApiListKafkaPartitionsRequest) (github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.PartitionDataList, *net_http.Response, error)

	calls struct {
		GetKafkaConsumerLag []struct {
			Ctx             context.Context
			ClusterId       string
			ConsumerGroupId string
			TopicName       string
			PartitionId     int32
		}
		GetKafkaConsumerLagExecute []struct {
			R github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.ApiGetKafkaConsumerLagRequest
		}
		GetKafkaPartition []struct {
			Ctx         context.Context
			ClusterId   string
			TopicName   string
			PartitionId int32
		}
		GetKafkaPartitionExecute []struct {
			R github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.ApiGetKafkaPartitionRequest
		}
		ListKafkaPartitions []struct {
			Ctx       context.Context
			ClusterId string
			TopicName string
		}
		ListKafkaPartitionsExecute []struct {
			R github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.ApiListKafkaPartitionsRequest
		}
	}
}

// GetKafkaConsumerLag mocks base method by wrapping the associated func.
func (m *PartitionV3Api) GetKafkaConsumerLag(ctx context.Context, clusterId, consumerGroupId, topicName string, partitionId int32) github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.ApiGetKafkaConsumerLagRequest {
	m.lockGetKafkaConsumerLag.Lock()
	defer m.lockGetKafkaConsumerLag.Unlock()

	if m.GetKafkaConsumerLagFunc == nil {
		panic("mocker: PartitionV3Api.GetKafkaConsumerLagFunc is nil but PartitionV3Api.GetKafkaConsumerLag was called.")
	}

	call := struct {
		Ctx             context.Context
		ClusterId       string
		ConsumerGroupId string
		TopicName       string
		PartitionId     int32
	}{
		Ctx:             ctx,
		ClusterId:       clusterId,
		ConsumerGroupId: consumerGroupId,
		TopicName:       topicName,
		PartitionId:     partitionId,
	}

	m.calls.GetKafkaConsumerLag = append(m.calls.GetKafkaConsumerLag, call)

	return m.GetKafkaConsumerLagFunc(ctx, clusterId, consumerGroupId, topicName, partitionId)
}

// GetKafkaConsumerLagCalled returns true if GetKafkaConsumerLag was called at least once.
func (m *PartitionV3Api) GetKafkaConsumerLagCalled() bool {
	m.lockGetKafkaConsumerLag.Lock()
	defer m.lockGetKafkaConsumerLag.Unlock()

	return len(m.calls.GetKafkaConsumerLag) > 0
}

// GetKafkaConsumerLagCalls returns the calls made to GetKafkaConsumerLag.
func (m *PartitionV3Api) GetKafkaConsumerLagCalls() []struct {
	Ctx             context.Context
	ClusterId       string
	ConsumerGroupId string
	TopicName       string
	PartitionId     int32
} {
	m.lockGetKafkaConsumerLag.Lock()
	defer m.lockGetKafkaConsumerLag.Unlock()

	return m.calls.GetKafkaConsumerLag
}

// GetKafkaConsumerLagExecute mocks base method by wrapping the associated func.
func (m *PartitionV3Api) GetKafkaConsumerLagExecute(r github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.ApiGetKafkaConsumerLagRequest) (github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.ConsumerLagData, *net_http.Response, error) {
	m.lockGetKafkaConsumerLagExecute.Lock()
	defer m.lockGetKafkaConsumerLagExecute.Unlock()

	if m.GetKafkaConsumerLagExecuteFunc == nil {
		panic("mocker: PartitionV3Api.GetKafkaConsumerLagExecuteFunc is nil but PartitionV3Api.GetKafkaConsumerLagExecute was called.")
	}

	call := struct {
		R github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.ApiGetKafkaConsumerLagRequest
	}{
		R: r,
	}

	m.calls.GetKafkaConsumerLagExecute = append(m.calls.GetKafkaConsumerLagExecute, call)

	return m.GetKafkaConsumerLagExecuteFunc(r)
}

// GetKafkaConsumerLagExecuteCalled returns true if GetKafkaConsumerLagExecute was called at least once.
func (m *PartitionV3Api) GetKafkaConsumerLagExecuteCalled() bool {
	m.lockGetKafkaConsumerLagExecute.Lock()
	defer m.lockGetKafkaConsumerLagExecute.Unlock()

	return len(m.calls.GetKafkaConsumerLagExecute) > 0
}

// GetKafkaConsumerLagExecuteCalls returns the calls made to GetKafkaConsumerLagExecute.
func (m *PartitionV3Api) GetKafkaConsumerLagExecuteCalls() []struct {
	R github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.ApiGetKafkaConsumerLagRequest
} {
	m.lockGetKafkaConsumerLagExecute.Lock()
	defer m.lockGetKafkaConsumerLagExecute.Unlock()

	return m.calls.GetKafkaConsumerLagExecute
}

// GetKafkaPartition mocks base method by wrapping the associated func.
func (m *PartitionV3Api) GetKafkaPartition(ctx context.Context, clusterId, topicName string, partitionId int32) github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.ApiGetKafkaPartitionRequest {
	m.lockGetKafkaPartition.Lock()
	defer m.lockGetKafkaPartition.Unlock()

	if m.GetKafkaPartitionFunc == nil {
		panic("mocker: PartitionV3Api.GetKafkaPartitionFunc is nil but PartitionV3Api.GetKafkaPartition was called.")
	}

	call := struct {
		Ctx         context.Context
		ClusterId   string
		TopicName   string
		PartitionId int32
	}{
		Ctx:         ctx,
		ClusterId:   clusterId,
		TopicName:   topicName,
		PartitionId: partitionId,
	}

	m.calls.GetKafkaPartition = append(m.calls.GetKafkaPartition, call)

	return m.GetKafkaPartitionFunc(ctx, clusterId, topicName, partitionId)
}

// GetKafkaPartitionCalled returns true if GetKafkaPartition was called at least once.
func (m *PartitionV3Api) GetKafkaPartitionCalled() bool {
	m.lockGetKafkaPartition.Lock()
	defer m.lockGetKafkaPartition.Unlock()

	return len(m.calls.GetKafkaPartition) > 0
}

// GetKafkaPartitionCalls returns the calls made to GetKafkaPartition.
func (m *PartitionV3Api) GetKafkaPartitionCalls() []struct {
	Ctx         context.Context
	ClusterId   string
	TopicName   string
	PartitionId int32
} {
	m.lockGetKafkaPartition.Lock()
	defer m.lockGetKafkaPartition.Unlock()

	return m.calls.GetKafkaPartition
}

// GetKafkaPartitionExecute mocks base method by wrapping the associated func.
func (m *PartitionV3Api) GetKafkaPartitionExecute(r github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.ApiGetKafkaPartitionRequest) (github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.PartitionData, *net_http.Response, error) {
	m.lockGetKafkaPartitionExecute.Lock()
	defer m.lockGetKafkaPartitionExecute.Unlock()

	if m.GetKafkaPartitionExecuteFunc == nil {
		panic("mocker: PartitionV3Api.GetKafkaPartitionExecuteFunc is nil but PartitionV3Api.GetKafkaPartitionExecute was called.")
	}

	call := struct {
		R github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.ApiGetKafkaPartitionRequest
	}{
		R: r,
	}

	m.calls.GetKafkaPartitionExecute = append(m.calls.GetKafkaPartitionExecute, call)

	return m.GetKafkaPartitionExecuteFunc(r)
}

// GetKafkaPartitionExecuteCalled returns true if GetKafkaPartitionExecute was called at least once.
func (m *PartitionV3Api) GetKafkaPartitionExecuteCalled() bool {
	m.lockGetKafkaPartitionExecute.Lock()
	defer m.lockGetKafkaPartitionExecute.Unlock()

	return len(m.calls.GetKafkaPartitionExecute) > 0
}

// GetKafkaPartitionExecuteCalls returns the calls made to GetKafkaPartitionExecute.
func (m *PartitionV3Api) GetKafkaPartitionExecuteCalls() []struct {
	R github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.ApiGetKafkaPartitionRequest
} {
	m.lockGetKafkaPartitionExecute.Lock()
	defer m.lockGetKafkaPartitionExecute.Unlock()

	return m.calls.GetKafkaPartitionExecute
}

// ListKafkaPartitions mocks base method by wrapping the associated func.
func (m *PartitionV3Api) ListKafkaPartitions(ctx context.Context, clusterId, topicName string) github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.ApiListKafkaPartitionsRequest {
	m.lockListKafkaPartitions.Lock()
	defer m.lockListKafkaPartitions.Unlock()

	if m.ListKafkaPartitionsFunc == nil {
		panic("mocker: PartitionV3Api.ListKafkaPartitionsFunc is nil but PartitionV3Api.ListKafkaPartitions was called.")
	}

	call := struct {
		Ctx       context.Context
		ClusterId string
		TopicName string
	}{
		Ctx:       ctx,
		ClusterId: clusterId,
		TopicName: topicName,
	}

	m.calls.ListKafkaPartitions = append(m.calls.ListKafkaPartitions, call)

	return m.ListKafkaPartitionsFunc(ctx, clusterId, topicName)
}

// ListKafkaPartitionsCalled returns true if ListKafkaPartitions was called at least once.
func (m *PartitionV3Api) ListKafkaPartitionsCalled() bool {
	m.lockListKafkaPartitions.Lock()
	defer m.lockListKafkaPartitions.Unlock()

	return len(m.calls.ListKafkaPartitions) > 0
}

// ListKafkaPartitionsCalls returns the calls made to ListKafkaPartitions.
func (m *PartitionV3Api) ListKafkaPartitionsCalls() []struct {
	Ctx       context.Context
	ClusterId string
	TopicName string
} {
	m.lockListKafkaPartitions.Lock()
	defer m.lockListKafkaPartitions.Unlock()

	return m.calls.ListKafkaPartitions
}

// ListKafkaPartitionsExecute mocks base method by wrapping the associated func.
func (m *PartitionV3Api) ListKafkaPartitionsExecute(r github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.ApiListKafkaPartitionsRequest) (github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.PartitionDataList, *net_http.Response, error) {
	m.lockListKafkaPartitionsExecute.Lock()
	defer m.lockListKafkaPartitionsExecute.Unlock()

	if m.ListKafkaPartitionsExecuteFunc == nil {
		panic("mocker: PartitionV3Api.ListKafkaPartitionsExecuteFunc is nil but PartitionV3Api.ListKafkaPartitionsExecute was called.")
	}

	call := struct {
		R github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.ApiListKafkaPartitionsRequest
	}{
		R: r,
	}

	m.calls.ListKafkaPartitionsExecute = append(m.calls.ListKafkaPartitionsExecute, call)

	return m.ListKafkaPartitionsExecuteFunc(r)
}

// ListKafkaPartitionsExecuteCalled returns true if ListKafkaPartitionsExecute was called at least once.
func (m *PartitionV3Api) ListKafkaPartitionsExecuteCalled() bool {
	m.lockListKafkaPartitionsExecute.Lock()
	defer m.lockListKafkaPartitionsExecute.Unlock()

	return len(m.calls.ListKafkaPartitionsExecute) > 0
}

// ListKafkaPartitionsExecuteCalls returns the calls made to ListKafkaPartitionsExecute.
func (m *PartitionV3Api) ListKafkaPartitionsExecuteCalls() []struct {
	R github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.ApiListKafkaPartitionsRequest
} {
	m.lockListKafkaPartitionsExecute.Lock()
	defer m.lockListKafkaPartitionsExecute.Unlock()

	return m.calls.ListKafkaPartitionsExecute
}

// Reset resets the calls made to the mocked methods.
func (m *PartitionV3Api) Reset() {
	m.lockGetKafkaConsumerLag.Lock()
	m.calls.GetKafkaConsumerLag = nil
	m.lockGetKafkaConsumerLag.Unlock()
	m.lockGetKafkaConsumerLagExecute.Lock()
	m.calls.GetKafkaConsumerLagExecute = nil
	m.lockGetKafkaConsumerLagExecute.Unlock()
	m.lockGetKafkaPartition.Lock()
	m.calls.GetKafkaPartition = nil
	m.lockGetKafkaPartition.Unlock()
	m.lockGetKafkaPartitionExecute.Lock()
	m.calls.GetKafkaPartitionExecute = nil
	m.lockGetKafkaPartitionExecute.Unlock()
	m.lockListKafkaPartitions.Lock()
	m.calls.ListKafkaPartitions = nil
	m.lockListKafkaPartitions.Unlock()
	m.lockListKafkaPartitionsExecute.Lock()
	m.calls.ListKafkaPartitionsExecute = nil
	m.lockListKafkaPartitionsExecute.Unlock()
}
