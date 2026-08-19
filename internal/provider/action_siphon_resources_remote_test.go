package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestSiphonResourcesAction_Invoke_Happy exercises SiphonResourcesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestSiphonResourcesAction_Invoke_Happy(t *testing.T) {
	r := &SiphonResourcesAction{client: newMockClientStatus(t, 201, "{}")}
	m := SiphonResourcesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestSiphonResourcesAction_Invoke_NilClient exercises SiphonResourcesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSiphonResourcesAction_Invoke_NilClient(t *testing.T) {
	r := &SiphonResourcesAction{}
	m := SiphonResourcesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestSiphonResourcesAction_Invoke_BuildError exercises SiphonResourcesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestSiphonResourcesAction_Invoke_BuildError(t *testing.T) {
	r := &SiphonResourcesAction{client: newMalformedBaseURLClient(t)}
	m := SiphonResourcesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestSiphonResourcesAction_Invoke_SendError exercises SiphonResourcesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestSiphonResourcesAction_Invoke_SendError(t *testing.T) {
	r := &SiphonResourcesAction{client: newTransportErrorClient(t)}
	m := SiphonResourcesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestSiphonResourcesAction_Invoke_APIError exercises SiphonResourcesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestSiphonResourcesAction_Invoke_APIError(t *testing.T) {
	r := &SiphonResourcesAction{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := SiphonResourcesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking spacetraders_siphon_resources")
}

// TestSiphonResourcesAction_Invoke_APIErrorReadBody exercises SiphonResourcesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestSiphonResourcesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &SiphonResourcesAction{client: newMockClientReadErrorBody(t, 500)}
	m := SiphonResourcesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
