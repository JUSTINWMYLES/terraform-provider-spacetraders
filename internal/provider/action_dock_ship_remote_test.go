package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDockShipAction_Invoke_Happy exercises DockShipAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDockShipAction_Invoke_Happy(t *testing.T) {
	r := &DockShipAction{client: newMockClientStatus(t, 200, "{}")}
	m := DockShipActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDockShipAction_Invoke_NilClient exercises DockShipAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDockShipAction_Invoke_NilClient(t *testing.T) {
	r := &DockShipAction{}
	m := DockShipActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDockShipAction_Invoke_BuildError exercises DockShipAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDockShipAction_Invoke_BuildError(t *testing.T) {
	r := &DockShipAction{client: newMalformedBaseURLClient(t)}
	m := DockShipActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDockShipAction_Invoke_SendError exercises DockShipAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDockShipAction_Invoke_SendError(t *testing.T) {
	r := &DockShipAction{client: newTransportErrorClient(t)}
	m := DockShipActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDockShipAction_Invoke_APIError exercises DockShipAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDockShipAction_Invoke_APIError(t *testing.T) {
	r := &DockShipAction{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := DockShipActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking spacetraders_dock_ship")
}

// TestDockShipAction_Invoke_APIErrorReadBody exercises DockShipAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDockShipAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DockShipAction{client: newMockClientReadErrorBody(t, 500)}
	m := DockShipActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
