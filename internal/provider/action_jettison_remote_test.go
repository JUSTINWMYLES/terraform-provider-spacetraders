package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestJettisonAction_Invoke_Happy exercises JettisonAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestJettisonAction_Invoke_Happy(t *testing.T) {
	r := &JettisonAction{client: newMockClientStatus(t, 200, "{}")}
	m := JettisonActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestJettisonAction_Invoke_NilClient exercises JettisonAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestJettisonAction_Invoke_NilClient(t *testing.T) {
	r := &JettisonAction{}
	m := JettisonActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestJettisonAction_Invoke_BuildError exercises JettisonAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestJettisonAction_Invoke_BuildError(t *testing.T) {
	r := &JettisonAction{client: newMalformedBaseURLClient(t)}
	m := JettisonActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestJettisonAction_Invoke_SendError exercises JettisonAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestJettisonAction_Invoke_SendError(t *testing.T) {
	r := &JettisonAction{client: newTransportErrorClient(t)}
	m := JettisonActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestJettisonAction_Invoke_APIError exercises JettisonAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestJettisonAction_Invoke_APIError(t *testing.T) {
	r := &JettisonAction{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := JettisonActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking spacetraders_jettison")
}

// TestJettisonAction_Invoke_APIErrorReadBody exercises JettisonAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestJettisonAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &JettisonAction{client: newMockClientReadErrorBody(t, 500)}
	m := JettisonActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
