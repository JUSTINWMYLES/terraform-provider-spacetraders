package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestTransferCargoAction_Invoke_Happy exercises TransferCargoAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestTransferCargoAction_Invoke_Happy(t *testing.T) {
	r := &TransferCargoAction{client: newMockClientStatus(t, 200, "{}")}
	m := TransferCargoActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestTransferCargoAction_Invoke_NilClient exercises TransferCargoAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestTransferCargoAction_Invoke_NilClient(t *testing.T) {
	r := &TransferCargoAction{}
	m := TransferCargoActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestTransferCargoAction_Invoke_BuildError exercises TransferCargoAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestTransferCargoAction_Invoke_BuildError(t *testing.T) {
	r := &TransferCargoAction{client: newMalformedBaseURLClient(t)}
	m := TransferCargoActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestTransferCargoAction_Invoke_SendError exercises TransferCargoAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestTransferCargoAction_Invoke_SendError(t *testing.T) {
	r := &TransferCargoAction{client: newTransportErrorClient(t)}
	m := TransferCargoActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestTransferCargoAction_Invoke_APIError exercises TransferCargoAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestTransferCargoAction_Invoke_APIError(t *testing.T) {
	r := &TransferCargoAction{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := TransferCargoActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking spacetraders_transfer_cargo")
}

// TestTransferCargoAction_Invoke_APIErrorReadBody exercises TransferCargoAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestTransferCargoAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &TransferCargoAction{client: newMockClientReadErrorBody(t, 500)}
	m := TransferCargoActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
