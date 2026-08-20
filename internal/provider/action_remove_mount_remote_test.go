package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRemoveMountAction_Invoke_Happy exercises RemoveMountAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRemoveMountAction_Invoke_Happy(t *testing.T) {
	r := &RemoveMountAction{client: newMockClientStatus(t, 201, "{}")}
	m := RemoveMountActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRemoveMountAction_Invoke_NilClient exercises RemoveMountAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRemoveMountAction_Invoke_NilClient(t *testing.T) {
	r := &RemoveMountAction{}
	m := RemoveMountActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRemoveMountAction_Invoke_BuildError exercises RemoveMountAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRemoveMountAction_Invoke_BuildError(t *testing.T) {
	r := &RemoveMountAction{client: newMalformedBaseURLClient(t)}
	m := RemoveMountActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRemoveMountAction_Invoke_SendError exercises RemoveMountAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRemoveMountAction_Invoke_SendError(t *testing.T) {
	r := &RemoveMountAction{client: newTransportErrorClient(t)}
	m := RemoveMountActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRemoveMountAction_Invoke_APIError exercises RemoveMountAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRemoveMountAction_Invoke_APIError(t *testing.T) {
	r := &RemoveMountAction{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := RemoveMountActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking spacetraders_remove_mount")
}

// TestRemoveMountAction_Invoke_APIErrorReadBody exercises RemoveMountAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRemoveMountAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RemoveMountAction{client: newMockClientReadErrorBody(t, 500)}
	m := RemoveMountActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
