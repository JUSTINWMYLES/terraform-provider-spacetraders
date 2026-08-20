package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestNavigateShipAction_Invoke_Happy exercises NavigateShipAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestNavigateShipAction_Invoke_Happy(t *testing.T) {
	r := &NavigateShipAction{client: newMockClientStatus(t, 200, "{}")}
	m := NavigateShipActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestNavigateShipAction_Invoke_NilClient exercises NavigateShipAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNavigateShipAction_Invoke_NilClient(t *testing.T) {
	r := &NavigateShipAction{}
	m := NavigateShipActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestNavigateShipAction_Invoke_BuildError exercises NavigateShipAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestNavigateShipAction_Invoke_BuildError(t *testing.T) {
	r := &NavigateShipAction{client: newMalformedBaseURLClient(t)}
	m := NavigateShipActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestNavigateShipAction_Invoke_SendError exercises NavigateShipAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestNavigateShipAction_Invoke_SendError(t *testing.T) {
	r := &NavigateShipAction{client: newTransportErrorClient(t)}
	m := NavigateShipActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestNavigateShipAction_Invoke_APIError exercises NavigateShipAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestNavigateShipAction_Invoke_APIError(t *testing.T) {
	r := &NavigateShipAction{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := NavigateShipActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking spacetraders_navigate_ship")
}

// TestNavigateShipAction_Invoke_APIErrorReadBody exercises NavigateShipAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestNavigateShipAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &NavigateShipAction{client: newMockClientReadErrorBody(t, 500)}
	m := NavigateShipActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
