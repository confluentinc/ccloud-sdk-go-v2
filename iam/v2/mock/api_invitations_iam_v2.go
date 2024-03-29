// Code generated by mocker. DO NOT EDIT.
// github.com/travisjeffery/mocker
// Source: api_invitations_iam_v2.go

package mock

import (
	context "context"
	net_http "net/http"
	sync "sync"

	github_com_confluentinc_ccloud_sdk_go_v2_iam_v2 "github.com/confluentinc/ccloud-sdk-go-v2/iam/v2"
)

// InvitationsIamV2Api is a mock of InvitationsIamV2Api interface
type InvitationsIamV2Api struct {
	lockCreateIamV2Invitation sync.Mutex
	CreateIamV2InvitationFunc func(ctx context.Context) github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiCreateIamV2InvitationRequest

	lockCreateIamV2InvitationExecute sync.Mutex
	CreateIamV2InvitationExecuteFunc func(r github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiCreateIamV2InvitationRequest) (github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.IamV2Invitation, *net_http.Response, error)

	lockDeleteIamV2Invitation sync.Mutex
	DeleteIamV2InvitationFunc func(ctx context.Context, id string) github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiDeleteIamV2InvitationRequest

	lockDeleteIamV2InvitationExecute sync.Mutex
	DeleteIamV2InvitationExecuteFunc func(r github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiDeleteIamV2InvitationRequest) (*net_http.Response, error)

	lockGetIamV2Invitation sync.Mutex
	GetIamV2InvitationFunc func(ctx context.Context, id string) github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiGetIamV2InvitationRequest

	lockGetIamV2InvitationExecute sync.Mutex
	GetIamV2InvitationExecuteFunc func(r github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiGetIamV2InvitationRequest) (github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.IamV2Invitation, *net_http.Response, error)

	lockListIamV2Invitations sync.Mutex
	ListIamV2InvitationsFunc func(ctx context.Context) github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiListIamV2InvitationsRequest

	lockListIamV2InvitationsExecute sync.Mutex
	ListIamV2InvitationsExecuteFunc func(r github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiListIamV2InvitationsRequest) (github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.IamV2InvitationList, *net_http.Response, error)

	calls struct {
		CreateIamV2Invitation []struct {
			Ctx context.Context
		}
		CreateIamV2InvitationExecute []struct {
			R github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiCreateIamV2InvitationRequest
		}
		DeleteIamV2Invitation []struct {
			Ctx context.Context
			Id  string
		}
		DeleteIamV2InvitationExecute []struct {
			R github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiDeleteIamV2InvitationRequest
		}
		GetIamV2Invitation []struct {
			Ctx context.Context
			Id  string
		}
		GetIamV2InvitationExecute []struct {
			R github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiGetIamV2InvitationRequest
		}
		ListIamV2Invitations []struct {
			Ctx context.Context
		}
		ListIamV2InvitationsExecute []struct {
			R github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiListIamV2InvitationsRequest
		}
	}
}

// CreateIamV2Invitation mocks base method by wrapping the associated func.
func (m *InvitationsIamV2Api) CreateIamV2Invitation(ctx context.Context) github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiCreateIamV2InvitationRequest {
	m.lockCreateIamV2Invitation.Lock()
	defer m.lockCreateIamV2Invitation.Unlock()

	if m.CreateIamV2InvitationFunc == nil {
		panic("mocker: InvitationsIamV2Api.CreateIamV2InvitationFunc is nil but InvitationsIamV2Api.CreateIamV2Invitation was called.")
	}

	call := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}

	m.calls.CreateIamV2Invitation = append(m.calls.CreateIamV2Invitation, call)

	return m.CreateIamV2InvitationFunc(ctx)
}

// CreateIamV2InvitationCalled returns true if CreateIamV2Invitation was called at least once.
func (m *InvitationsIamV2Api) CreateIamV2InvitationCalled() bool {
	m.lockCreateIamV2Invitation.Lock()
	defer m.lockCreateIamV2Invitation.Unlock()

	return len(m.calls.CreateIamV2Invitation) > 0
}

// CreateIamV2InvitationCalls returns the calls made to CreateIamV2Invitation.
func (m *InvitationsIamV2Api) CreateIamV2InvitationCalls() []struct {
	Ctx context.Context
} {
	m.lockCreateIamV2Invitation.Lock()
	defer m.lockCreateIamV2Invitation.Unlock()

	return m.calls.CreateIamV2Invitation
}

// CreateIamV2InvitationExecute mocks base method by wrapping the associated func.
func (m *InvitationsIamV2Api) CreateIamV2InvitationExecute(r github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiCreateIamV2InvitationRequest) (github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.IamV2Invitation, *net_http.Response, error) {
	m.lockCreateIamV2InvitationExecute.Lock()
	defer m.lockCreateIamV2InvitationExecute.Unlock()

	if m.CreateIamV2InvitationExecuteFunc == nil {
		panic("mocker: InvitationsIamV2Api.CreateIamV2InvitationExecuteFunc is nil but InvitationsIamV2Api.CreateIamV2InvitationExecute was called.")
	}

	call := struct {
		R github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiCreateIamV2InvitationRequest
	}{
		R: r,
	}

	m.calls.CreateIamV2InvitationExecute = append(m.calls.CreateIamV2InvitationExecute, call)

	return m.CreateIamV2InvitationExecuteFunc(r)
}

// CreateIamV2InvitationExecuteCalled returns true if CreateIamV2InvitationExecute was called at least once.
func (m *InvitationsIamV2Api) CreateIamV2InvitationExecuteCalled() bool {
	m.lockCreateIamV2InvitationExecute.Lock()
	defer m.lockCreateIamV2InvitationExecute.Unlock()

	return len(m.calls.CreateIamV2InvitationExecute) > 0
}

// CreateIamV2InvitationExecuteCalls returns the calls made to CreateIamV2InvitationExecute.
func (m *InvitationsIamV2Api) CreateIamV2InvitationExecuteCalls() []struct {
	R github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiCreateIamV2InvitationRequest
} {
	m.lockCreateIamV2InvitationExecute.Lock()
	defer m.lockCreateIamV2InvitationExecute.Unlock()

	return m.calls.CreateIamV2InvitationExecute
}

// DeleteIamV2Invitation mocks base method by wrapping the associated func.
func (m *InvitationsIamV2Api) DeleteIamV2Invitation(ctx context.Context, id string) github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiDeleteIamV2InvitationRequest {
	m.lockDeleteIamV2Invitation.Lock()
	defer m.lockDeleteIamV2Invitation.Unlock()

	if m.DeleteIamV2InvitationFunc == nil {
		panic("mocker: InvitationsIamV2Api.DeleteIamV2InvitationFunc is nil but InvitationsIamV2Api.DeleteIamV2Invitation was called.")
	}

	call := struct {
		Ctx context.Context
		Id  string
	}{
		Ctx: ctx,
		Id:  id,
	}

	m.calls.DeleteIamV2Invitation = append(m.calls.DeleteIamV2Invitation, call)

	return m.DeleteIamV2InvitationFunc(ctx, id)
}

// DeleteIamV2InvitationCalled returns true if DeleteIamV2Invitation was called at least once.
func (m *InvitationsIamV2Api) DeleteIamV2InvitationCalled() bool {
	m.lockDeleteIamV2Invitation.Lock()
	defer m.lockDeleteIamV2Invitation.Unlock()

	return len(m.calls.DeleteIamV2Invitation) > 0
}

// DeleteIamV2InvitationCalls returns the calls made to DeleteIamV2Invitation.
func (m *InvitationsIamV2Api) DeleteIamV2InvitationCalls() []struct {
	Ctx context.Context
	Id  string
} {
	m.lockDeleteIamV2Invitation.Lock()
	defer m.lockDeleteIamV2Invitation.Unlock()

	return m.calls.DeleteIamV2Invitation
}

// DeleteIamV2InvitationExecute mocks base method by wrapping the associated func.
func (m *InvitationsIamV2Api) DeleteIamV2InvitationExecute(r github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiDeleteIamV2InvitationRequest) (*net_http.Response, error) {
	m.lockDeleteIamV2InvitationExecute.Lock()
	defer m.lockDeleteIamV2InvitationExecute.Unlock()

	if m.DeleteIamV2InvitationExecuteFunc == nil {
		panic("mocker: InvitationsIamV2Api.DeleteIamV2InvitationExecuteFunc is nil but InvitationsIamV2Api.DeleteIamV2InvitationExecute was called.")
	}

	call := struct {
		R github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiDeleteIamV2InvitationRequest
	}{
		R: r,
	}

	m.calls.DeleteIamV2InvitationExecute = append(m.calls.DeleteIamV2InvitationExecute, call)

	return m.DeleteIamV2InvitationExecuteFunc(r)
}

// DeleteIamV2InvitationExecuteCalled returns true if DeleteIamV2InvitationExecute was called at least once.
func (m *InvitationsIamV2Api) DeleteIamV2InvitationExecuteCalled() bool {
	m.lockDeleteIamV2InvitationExecute.Lock()
	defer m.lockDeleteIamV2InvitationExecute.Unlock()

	return len(m.calls.DeleteIamV2InvitationExecute) > 0
}

// DeleteIamV2InvitationExecuteCalls returns the calls made to DeleteIamV2InvitationExecute.
func (m *InvitationsIamV2Api) DeleteIamV2InvitationExecuteCalls() []struct {
	R github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiDeleteIamV2InvitationRequest
} {
	m.lockDeleteIamV2InvitationExecute.Lock()
	defer m.lockDeleteIamV2InvitationExecute.Unlock()

	return m.calls.DeleteIamV2InvitationExecute
}

// GetIamV2Invitation mocks base method by wrapping the associated func.
func (m *InvitationsIamV2Api) GetIamV2Invitation(ctx context.Context, id string) github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiGetIamV2InvitationRequest {
	m.lockGetIamV2Invitation.Lock()
	defer m.lockGetIamV2Invitation.Unlock()

	if m.GetIamV2InvitationFunc == nil {
		panic("mocker: InvitationsIamV2Api.GetIamV2InvitationFunc is nil but InvitationsIamV2Api.GetIamV2Invitation was called.")
	}

	call := struct {
		Ctx context.Context
		Id  string
	}{
		Ctx: ctx,
		Id:  id,
	}

	m.calls.GetIamV2Invitation = append(m.calls.GetIamV2Invitation, call)

	return m.GetIamV2InvitationFunc(ctx, id)
}

// GetIamV2InvitationCalled returns true if GetIamV2Invitation was called at least once.
func (m *InvitationsIamV2Api) GetIamV2InvitationCalled() bool {
	m.lockGetIamV2Invitation.Lock()
	defer m.lockGetIamV2Invitation.Unlock()

	return len(m.calls.GetIamV2Invitation) > 0
}

// GetIamV2InvitationCalls returns the calls made to GetIamV2Invitation.
func (m *InvitationsIamV2Api) GetIamV2InvitationCalls() []struct {
	Ctx context.Context
	Id  string
} {
	m.lockGetIamV2Invitation.Lock()
	defer m.lockGetIamV2Invitation.Unlock()

	return m.calls.GetIamV2Invitation
}

// GetIamV2InvitationExecute mocks base method by wrapping the associated func.
func (m *InvitationsIamV2Api) GetIamV2InvitationExecute(r github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiGetIamV2InvitationRequest) (github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.IamV2Invitation, *net_http.Response, error) {
	m.lockGetIamV2InvitationExecute.Lock()
	defer m.lockGetIamV2InvitationExecute.Unlock()

	if m.GetIamV2InvitationExecuteFunc == nil {
		panic("mocker: InvitationsIamV2Api.GetIamV2InvitationExecuteFunc is nil but InvitationsIamV2Api.GetIamV2InvitationExecute was called.")
	}

	call := struct {
		R github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiGetIamV2InvitationRequest
	}{
		R: r,
	}

	m.calls.GetIamV2InvitationExecute = append(m.calls.GetIamV2InvitationExecute, call)

	return m.GetIamV2InvitationExecuteFunc(r)
}

// GetIamV2InvitationExecuteCalled returns true if GetIamV2InvitationExecute was called at least once.
func (m *InvitationsIamV2Api) GetIamV2InvitationExecuteCalled() bool {
	m.lockGetIamV2InvitationExecute.Lock()
	defer m.lockGetIamV2InvitationExecute.Unlock()

	return len(m.calls.GetIamV2InvitationExecute) > 0
}

// GetIamV2InvitationExecuteCalls returns the calls made to GetIamV2InvitationExecute.
func (m *InvitationsIamV2Api) GetIamV2InvitationExecuteCalls() []struct {
	R github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiGetIamV2InvitationRequest
} {
	m.lockGetIamV2InvitationExecute.Lock()
	defer m.lockGetIamV2InvitationExecute.Unlock()

	return m.calls.GetIamV2InvitationExecute
}

// ListIamV2Invitations mocks base method by wrapping the associated func.
func (m *InvitationsIamV2Api) ListIamV2Invitations(ctx context.Context) github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiListIamV2InvitationsRequest {
	m.lockListIamV2Invitations.Lock()
	defer m.lockListIamV2Invitations.Unlock()

	if m.ListIamV2InvitationsFunc == nil {
		panic("mocker: InvitationsIamV2Api.ListIamV2InvitationsFunc is nil but InvitationsIamV2Api.ListIamV2Invitations was called.")
	}

	call := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}

	m.calls.ListIamV2Invitations = append(m.calls.ListIamV2Invitations, call)

	return m.ListIamV2InvitationsFunc(ctx)
}

// ListIamV2InvitationsCalled returns true if ListIamV2Invitations was called at least once.
func (m *InvitationsIamV2Api) ListIamV2InvitationsCalled() bool {
	m.lockListIamV2Invitations.Lock()
	defer m.lockListIamV2Invitations.Unlock()

	return len(m.calls.ListIamV2Invitations) > 0
}

// ListIamV2InvitationsCalls returns the calls made to ListIamV2Invitations.
func (m *InvitationsIamV2Api) ListIamV2InvitationsCalls() []struct {
	Ctx context.Context
} {
	m.lockListIamV2Invitations.Lock()
	defer m.lockListIamV2Invitations.Unlock()

	return m.calls.ListIamV2Invitations
}

// ListIamV2InvitationsExecute mocks base method by wrapping the associated func.
func (m *InvitationsIamV2Api) ListIamV2InvitationsExecute(r github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiListIamV2InvitationsRequest) (github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.IamV2InvitationList, *net_http.Response, error) {
	m.lockListIamV2InvitationsExecute.Lock()
	defer m.lockListIamV2InvitationsExecute.Unlock()

	if m.ListIamV2InvitationsExecuteFunc == nil {
		panic("mocker: InvitationsIamV2Api.ListIamV2InvitationsExecuteFunc is nil but InvitationsIamV2Api.ListIamV2InvitationsExecute was called.")
	}

	call := struct {
		R github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiListIamV2InvitationsRequest
	}{
		R: r,
	}

	m.calls.ListIamV2InvitationsExecute = append(m.calls.ListIamV2InvitationsExecute, call)

	return m.ListIamV2InvitationsExecuteFunc(r)
}

// ListIamV2InvitationsExecuteCalled returns true if ListIamV2InvitationsExecute was called at least once.
func (m *InvitationsIamV2Api) ListIamV2InvitationsExecuteCalled() bool {
	m.lockListIamV2InvitationsExecute.Lock()
	defer m.lockListIamV2InvitationsExecute.Unlock()

	return len(m.calls.ListIamV2InvitationsExecute) > 0
}

// ListIamV2InvitationsExecuteCalls returns the calls made to ListIamV2InvitationsExecute.
func (m *InvitationsIamV2Api) ListIamV2InvitationsExecuteCalls() []struct {
	R github_com_confluentinc_ccloud_sdk_go_v2_iam_v2.ApiListIamV2InvitationsRequest
} {
	m.lockListIamV2InvitationsExecute.Lock()
	defer m.lockListIamV2InvitationsExecute.Unlock()

	return m.calls.ListIamV2InvitationsExecute
}

// Reset resets the calls made to the mocked methods.
func (m *InvitationsIamV2Api) Reset() {
	m.lockCreateIamV2Invitation.Lock()
	m.calls.CreateIamV2Invitation = nil
	m.lockCreateIamV2Invitation.Unlock()
	m.lockCreateIamV2InvitationExecute.Lock()
	m.calls.CreateIamV2InvitationExecute = nil
	m.lockCreateIamV2InvitationExecute.Unlock()
	m.lockDeleteIamV2Invitation.Lock()
	m.calls.DeleteIamV2Invitation = nil
	m.lockDeleteIamV2Invitation.Unlock()
	m.lockDeleteIamV2InvitationExecute.Lock()
	m.calls.DeleteIamV2InvitationExecute = nil
	m.lockDeleteIamV2InvitationExecute.Unlock()
	m.lockGetIamV2Invitation.Lock()
	m.calls.GetIamV2Invitation = nil
	m.lockGetIamV2Invitation.Unlock()
	m.lockGetIamV2InvitationExecute.Lock()
	m.calls.GetIamV2InvitationExecute = nil
	m.lockGetIamV2InvitationExecute.Unlock()
	m.lockListIamV2Invitations.Lock()
	m.calls.ListIamV2Invitations = nil
	m.lockListIamV2Invitations.Unlock()
	m.lockListIamV2InvitationsExecute.Lock()
	m.calls.ListIamV2InvitationsExecute = nil
	m.lockListIamV2InvitationsExecute.Unlock()
}
