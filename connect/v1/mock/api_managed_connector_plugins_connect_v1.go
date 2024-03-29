// Code generated by mocker. DO NOT EDIT.
// github.com/travisjeffery/mocker
// Source: api_managed_connector_plugins_connect_v1.go

package mock

import (
	context "context"
	net_http "net/http"
	sync "sync"

	github_com_confluentinc_ccloud_sdk_go_v2_connect_v1 "github.com/confluentinc/ccloud-sdk-go-v2/connect/v1"
)

// ManagedConnectorPluginsConnectV1Api is a mock of ManagedConnectorPluginsConnectV1Api interface
type ManagedConnectorPluginsConnectV1Api struct {
	lockListConnectv1ConnectorPlugins sync.Mutex
	ListConnectv1ConnectorPluginsFunc func(ctx context.Context, environmentId, kafkaClusterId string) github_com_confluentinc_ccloud_sdk_go_v2_connect_v1.ApiListConnectv1ConnectorPluginsRequest

	lockListConnectv1ConnectorPluginsExecute sync.Mutex
	ListConnectv1ConnectorPluginsExecuteFunc func(r github_com_confluentinc_ccloud_sdk_go_v2_connect_v1.ApiListConnectv1ConnectorPluginsRequest) ([]github_com_confluentinc_ccloud_sdk_go_v2_connect_v1.InlineResponse2002, *net_http.Response, error)

	lockValidateConnectv1ConnectorPlugin sync.Mutex
	ValidateConnectv1ConnectorPluginFunc func(ctx context.Context, pluginName, environmentId, kafkaClusterId string) github_com_confluentinc_ccloud_sdk_go_v2_connect_v1.ApiValidateConnectv1ConnectorPluginRequest

	lockValidateConnectv1ConnectorPluginExecute sync.Mutex
	ValidateConnectv1ConnectorPluginExecuteFunc func(r github_com_confluentinc_ccloud_sdk_go_v2_connect_v1.ApiValidateConnectv1ConnectorPluginRequest) (github_com_confluentinc_ccloud_sdk_go_v2_connect_v1.InlineResponse2003, *net_http.Response, error)

	calls struct {
		ListConnectv1ConnectorPlugins []struct {
			Ctx            context.Context
			EnvironmentId  string
			KafkaClusterId string
		}
		ListConnectv1ConnectorPluginsExecute []struct {
			R github_com_confluentinc_ccloud_sdk_go_v2_connect_v1.ApiListConnectv1ConnectorPluginsRequest
		}
		ValidateConnectv1ConnectorPlugin []struct {
			Ctx            context.Context
			PluginName     string
			EnvironmentId  string
			KafkaClusterId string
		}
		ValidateConnectv1ConnectorPluginExecute []struct {
			R github_com_confluentinc_ccloud_sdk_go_v2_connect_v1.ApiValidateConnectv1ConnectorPluginRequest
		}
	}
}

// ListConnectv1ConnectorPlugins mocks base method by wrapping the associated func.
func (m *ManagedConnectorPluginsConnectV1Api) ListConnectv1ConnectorPlugins(ctx context.Context, environmentId, kafkaClusterId string) github_com_confluentinc_ccloud_sdk_go_v2_connect_v1.ApiListConnectv1ConnectorPluginsRequest {
	m.lockListConnectv1ConnectorPlugins.Lock()
	defer m.lockListConnectv1ConnectorPlugins.Unlock()

	if m.ListConnectv1ConnectorPluginsFunc == nil {
		panic("mocker: ManagedConnectorPluginsConnectV1Api.ListConnectv1ConnectorPluginsFunc is nil but ManagedConnectorPluginsConnectV1Api.ListConnectv1ConnectorPlugins was called.")
	}

	call := struct {
		Ctx            context.Context
		EnvironmentId  string
		KafkaClusterId string
	}{
		Ctx:            ctx,
		EnvironmentId:  environmentId,
		KafkaClusterId: kafkaClusterId,
	}

	m.calls.ListConnectv1ConnectorPlugins = append(m.calls.ListConnectv1ConnectorPlugins, call)

	return m.ListConnectv1ConnectorPluginsFunc(ctx, environmentId, kafkaClusterId)
}

// ListConnectv1ConnectorPluginsCalled returns true if ListConnectv1ConnectorPlugins was called at least once.
func (m *ManagedConnectorPluginsConnectV1Api) ListConnectv1ConnectorPluginsCalled() bool {
	m.lockListConnectv1ConnectorPlugins.Lock()
	defer m.lockListConnectv1ConnectorPlugins.Unlock()

	return len(m.calls.ListConnectv1ConnectorPlugins) > 0
}

// ListConnectv1ConnectorPluginsCalls returns the calls made to ListConnectv1ConnectorPlugins.
func (m *ManagedConnectorPluginsConnectV1Api) ListConnectv1ConnectorPluginsCalls() []struct {
	Ctx            context.Context
	EnvironmentId  string
	KafkaClusterId string
} {
	m.lockListConnectv1ConnectorPlugins.Lock()
	defer m.lockListConnectv1ConnectorPlugins.Unlock()

	return m.calls.ListConnectv1ConnectorPlugins
}

// ListConnectv1ConnectorPluginsExecute mocks base method by wrapping the associated func.
func (m *ManagedConnectorPluginsConnectV1Api) ListConnectv1ConnectorPluginsExecute(r github_com_confluentinc_ccloud_sdk_go_v2_connect_v1.ApiListConnectv1ConnectorPluginsRequest) ([]github_com_confluentinc_ccloud_sdk_go_v2_connect_v1.InlineResponse2002, *net_http.Response, error) {
	m.lockListConnectv1ConnectorPluginsExecute.Lock()
	defer m.lockListConnectv1ConnectorPluginsExecute.Unlock()

	if m.ListConnectv1ConnectorPluginsExecuteFunc == nil {
		panic("mocker: ManagedConnectorPluginsConnectV1Api.ListConnectv1ConnectorPluginsExecuteFunc is nil but ManagedConnectorPluginsConnectV1Api.ListConnectv1ConnectorPluginsExecute was called.")
	}

	call := struct {
		R github_com_confluentinc_ccloud_sdk_go_v2_connect_v1.ApiListConnectv1ConnectorPluginsRequest
	}{
		R: r,
	}

	m.calls.ListConnectv1ConnectorPluginsExecute = append(m.calls.ListConnectv1ConnectorPluginsExecute, call)

	return m.ListConnectv1ConnectorPluginsExecuteFunc(r)
}

// ListConnectv1ConnectorPluginsExecuteCalled returns true if ListConnectv1ConnectorPluginsExecute was called at least once.
func (m *ManagedConnectorPluginsConnectV1Api) ListConnectv1ConnectorPluginsExecuteCalled() bool {
	m.lockListConnectv1ConnectorPluginsExecute.Lock()
	defer m.lockListConnectv1ConnectorPluginsExecute.Unlock()

	return len(m.calls.ListConnectv1ConnectorPluginsExecute) > 0
}

// ListConnectv1ConnectorPluginsExecuteCalls returns the calls made to ListConnectv1ConnectorPluginsExecute.
func (m *ManagedConnectorPluginsConnectV1Api) ListConnectv1ConnectorPluginsExecuteCalls() []struct {
	R github_com_confluentinc_ccloud_sdk_go_v2_connect_v1.ApiListConnectv1ConnectorPluginsRequest
} {
	m.lockListConnectv1ConnectorPluginsExecute.Lock()
	defer m.lockListConnectv1ConnectorPluginsExecute.Unlock()

	return m.calls.ListConnectv1ConnectorPluginsExecute
}

// ValidateConnectv1ConnectorPlugin mocks base method by wrapping the associated func.
func (m *ManagedConnectorPluginsConnectV1Api) ValidateConnectv1ConnectorPlugin(ctx context.Context, pluginName, environmentId, kafkaClusterId string) github_com_confluentinc_ccloud_sdk_go_v2_connect_v1.ApiValidateConnectv1ConnectorPluginRequest {
	m.lockValidateConnectv1ConnectorPlugin.Lock()
	defer m.lockValidateConnectv1ConnectorPlugin.Unlock()

	if m.ValidateConnectv1ConnectorPluginFunc == nil {
		panic("mocker: ManagedConnectorPluginsConnectV1Api.ValidateConnectv1ConnectorPluginFunc is nil but ManagedConnectorPluginsConnectV1Api.ValidateConnectv1ConnectorPlugin was called.")
	}

	call := struct {
		Ctx            context.Context
		PluginName     string
		EnvironmentId  string
		KafkaClusterId string
	}{
		Ctx:            ctx,
		PluginName:     pluginName,
		EnvironmentId:  environmentId,
		KafkaClusterId: kafkaClusterId,
	}

	m.calls.ValidateConnectv1ConnectorPlugin = append(m.calls.ValidateConnectv1ConnectorPlugin, call)

	return m.ValidateConnectv1ConnectorPluginFunc(ctx, pluginName, environmentId, kafkaClusterId)
}

// ValidateConnectv1ConnectorPluginCalled returns true if ValidateConnectv1ConnectorPlugin was called at least once.
func (m *ManagedConnectorPluginsConnectV1Api) ValidateConnectv1ConnectorPluginCalled() bool {
	m.lockValidateConnectv1ConnectorPlugin.Lock()
	defer m.lockValidateConnectv1ConnectorPlugin.Unlock()

	return len(m.calls.ValidateConnectv1ConnectorPlugin) > 0
}

// ValidateConnectv1ConnectorPluginCalls returns the calls made to ValidateConnectv1ConnectorPlugin.
func (m *ManagedConnectorPluginsConnectV1Api) ValidateConnectv1ConnectorPluginCalls() []struct {
	Ctx            context.Context
	PluginName     string
	EnvironmentId  string
	KafkaClusterId string
} {
	m.lockValidateConnectv1ConnectorPlugin.Lock()
	defer m.lockValidateConnectv1ConnectorPlugin.Unlock()

	return m.calls.ValidateConnectv1ConnectorPlugin
}

// ValidateConnectv1ConnectorPluginExecute mocks base method by wrapping the associated func.
func (m *ManagedConnectorPluginsConnectV1Api) ValidateConnectv1ConnectorPluginExecute(r github_com_confluentinc_ccloud_sdk_go_v2_connect_v1.ApiValidateConnectv1ConnectorPluginRequest) (github_com_confluentinc_ccloud_sdk_go_v2_connect_v1.InlineResponse2003, *net_http.Response, error) {
	m.lockValidateConnectv1ConnectorPluginExecute.Lock()
	defer m.lockValidateConnectv1ConnectorPluginExecute.Unlock()

	if m.ValidateConnectv1ConnectorPluginExecuteFunc == nil {
		panic("mocker: ManagedConnectorPluginsConnectV1Api.ValidateConnectv1ConnectorPluginExecuteFunc is nil but ManagedConnectorPluginsConnectV1Api.ValidateConnectv1ConnectorPluginExecute was called.")
	}

	call := struct {
		R github_com_confluentinc_ccloud_sdk_go_v2_connect_v1.ApiValidateConnectv1ConnectorPluginRequest
	}{
		R: r,
	}

	m.calls.ValidateConnectv1ConnectorPluginExecute = append(m.calls.ValidateConnectv1ConnectorPluginExecute, call)

	return m.ValidateConnectv1ConnectorPluginExecuteFunc(r)
}

// ValidateConnectv1ConnectorPluginExecuteCalled returns true if ValidateConnectv1ConnectorPluginExecute was called at least once.
func (m *ManagedConnectorPluginsConnectV1Api) ValidateConnectv1ConnectorPluginExecuteCalled() bool {
	m.lockValidateConnectv1ConnectorPluginExecute.Lock()
	defer m.lockValidateConnectv1ConnectorPluginExecute.Unlock()

	return len(m.calls.ValidateConnectv1ConnectorPluginExecute) > 0
}

// ValidateConnectv1ConnectorPluginExecuteCalls returns the calls made to ValidateConnectv1ConnectorPluginExecute.
func (m *ManagedConnectorPluginsConnectV1Api) ValidateConnectv1ConnectorPluginExecuteCalls() []struct {
	R github_com_confluentinc_ccloud_sdk_go_v2_connect_v1.ApiValidateConnectv1ConnectorPluginRequest
} {
	m.lockValidateConnectv1ConnectorPluginExecute.Lock()
	defer m.lockValidateConnectv1ConnectorPluginExecute.Unlock()

	return m.calls.ValidateConnectv1ConnectorPluginExecute
}

// Reset resets the calls made to the mocked methods.
func (m *ManagedConnectorPluginsConnectV1Api) Reset() {
	m.lockListConnectv1ConnectorPlugins.Lock()
	m.calls.ListConnectv1ConnectorPlugins = nil
	m.lockListConnectv1ConnectorPlugins.Unlock()
	m.lockListConnectv1ConnectorPluginsExecute.Lock()
	m.calls.ListConnectv1ConnectorPluginsExecute = nil
	m.lockListConnectv1ConnectorPluginsExecute.Unlock()
	m.lockValidateConnectv1ConnectorPlugin.Lock()
	m.calls.ValidateConnectv1ConnectorPlugin = nil
	m.lockValidateConnectv1ConnectorPlugin.Unlock()
	m.lockValidateConnectv1ConnectorPluginExecute.Lock()
	m.calls.ValidateConnectv1ConnectorPluginExecute = nil
	m.lockValidateConnectv1ConnectorPluginExecute.Unlock()
}
