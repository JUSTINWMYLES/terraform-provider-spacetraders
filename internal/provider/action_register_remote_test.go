package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRegisterAction_Invoke_Happy exercises RegisterAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRegisterAction_Invoke_Happy(t *testing.T) {
	r := &RegisterAction{client: newMockClientStatus(t, 201, "{}")}
	m := RegisterActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRegisterAction_Invoke_NilClient exercises RegisterAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRegisterAction_Invoke_NilClient(t *testing.T) {
	r := &RegisterAction{}
	m := RegisterActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRegisterAction_Invoke_BuildError exercises RegisterAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRegisterAction_Invoke_BuildError(t *testing.T) {
	r := &RegisterAction{client: newMalformedBaseURLClient(t)}
	m := RegisterActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRegisterAction_Invoke_SendError exercises RegisterAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRegisterAction_Invoke_SendError(t *testing.T) {
	r := &RegisterAction{client: newTransportErrorClient(t)}
	m := RegisterActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRegisterAction_Invoke_APIError exercises RegisterAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRegisterAction_Invoke_APIError(t *testing.T) {
	r := &RegisterAction{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := RegisterActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking spacetraders_register")
}

// TestRegisterAction_Invoke_APIErrorReadBody exercises RegisterAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRegisterAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RegisterAction{client: newMockClientReadErrorBody(t, 500)}
	m := RegisterActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
