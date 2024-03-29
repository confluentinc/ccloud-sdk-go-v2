// Code generated by mocker. DO NOT EDIT.
// github.com/travisjeffery/mocker
// Source: api_records_v3.go

package mock

import (
	context "context"
	net_http "net/http"
	sync "sync"

	github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3 "github.com/confluentinc/ccloud-sdk-go-v2/kafkarest/v3"
)

// RecordsV3Api is a mock of RecordsV3Api interface
type RecordsV3Api struct {
	lockProduceRecord sync.Mutex
	ProduceRecordFunc func(ctx context.Context, clusterId, topicName string) github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.ApiProduceRecordRequest

	lockProduceRecordExecute sync.Mutex
	ProduceRecordExecuteFunc func(r github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.ApiProduceRecordRequest) (github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.ProduceResponse, *net_http.Response, error)

	calls struct {
		ProduceRecord []struct {
			Ctx       context.Context
			ClusterId string
			TopicName string
		}
		ProduceRecordExecute []struct {
			R github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.ApiProduceRecordRequest
		}
	}
}

// ProduceRecord mocks base method by wrapping the associated func.
func (m *RecordsV3Api) ProduceRecord(ctx context.Context, clusterId, topicName string) github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.ApiProduceRecordRequest {
	m.lockProduceRecord.Lock()
	defer m.lockProduceRecord.Unlock()

	if m.ProduceRecordFunc == nil {
		panic("mocker: RecordsV3Api.ProduceRecordFunc is nil but RecordsV3Api.ProduceRecord was called.")
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

	m.calls.ProduceRecord = append(m.calls.ProduceRecord, call)

	return m.ProduceRecordFunc(ctx, clusterId, topicName)
}

// ProduceRecordCalled returns true if ProduceRecord was called at least once.
func (m *RecordsV3Api) ProduceRecordCalled() bool {
	m.lockProduceRecord.Lock()
	defer m.lockProduceRecord.Unlock()

	return len(m.calls.ProduceRecord) > 0
}

// ProduceRecordCalls returns the calls made to ProduceRecord.
func (m *RecordsV3Api) ProduceRecordCalls() []struct {
	Ctx       context.Context
	ClusterId string
	TopicName string
} {
	m.lockProduceRecord.Lock()
	defer m.lockProduceRecord.Unlock()

	return m.calls.ProduceRecord
}

// ProduceRecordExecute mocks base method by wrapping the associated func.
func (m *RecordsV3Api) ProduceRecordExecute(r github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.ApiProduceRecordRequest) (github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.ProduceResponse, *net_http.Response, error) {
	m.lockProduceRecordExecute.Lock()
	defer m.lockProduceRecordExecute.Unlock()

	if m.ProduceRecordExecuteFunc == nil {
		panic("mocker: RecordsV3Api.ProduceRecordExecuteFunc is nil but RecordsV3Api.ProduceRecordExecute was called.")
	}

	call := struct {
		R github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.ApiProduceRecordRequest
	}{
		R: r,
	}

	m.calls.ProduceRecordExecute = append(m.calls.ProduceRecordExecute, call)

	return m.ProduceRecordExecuteFunc(r)
}

// ProduceRecordExecuteCalled returns true if ProduceRecordExecute was called at least once.
func (m *RecordsV3Api) ProduceRecordExecuteCalled() bool {
	m.lockProduceRecordExecute.Lock()
	defer m.lockProduceRecordExecute.Unlock()

	return len(m.calls.ProduceRecordExecute) > 0
}

// ProduceRecordExecuteCalls returns the calls made to ProduceRecordExecute.
func (m *RecordsV3Api) ProduceRecordExecuteCalls() []struct {
	R github_com_confluentinc_ccloud_sdk_go_v2_kafkarest_v3.ApiProduceRecordRequest
} {
	m.lockProduceRecordExecute.Lock()
	defer m.lockProduceRecordExecute.Unlock()

	return m.calls.ProduceRecordExecute
}

// Reset resets the calls made to the mocked methods.
func (m *RecordsV3Api) Reset() {
	m.lockProduceRecord.Lock()
	m.calls.ProduceRecord = nil
	m.lockProduceRecord.Unlock()
	m.lockProduceRecordExecute.Lock()
	m.calls.ProduceRecordExecute = nil
	m.lockProduceRecordExecute.Unlock()
}
