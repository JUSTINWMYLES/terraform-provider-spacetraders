package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRemoveShipModuleAction_Invoke_Happy exercises RemoveShipModuleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRemoveShipModuleAction_Invoke_Happy(t *testing.T) {
	r := &RemoveShipModuleAction{client: newMockClientStatus(t, 201, "{}")}
	m := RemoveShipModuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRemoveShipModuleAction_Invoke_NilClient exercises RemoveShipModuleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRemoveShipModuleAction_Invoke_NilClient(t *testing.T) {
	r := &RemoveShipModuleAction{}
	m := RemoveShipModuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRemoveShipModuleAction_Invoke_BuildError exercises RemoveShipModuleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRemoveShipModuleAction_Invoke_BuildError(t *testing.T) {
	r := &RemoveShipModuleAction{client: newMalformedBaseURLClient(t)}
	m := RemoveShipModuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRemoveShipModuleAction_Invoke_SendError exercises RemoveShipModuleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRemoveShipModuleAction_Invoke_SendError(t *testing.T) {
	r := &RemoveShipModuleAction{client: newTransportErrorClient(t)}
	m := RemoveShipModuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRemoveShipModuleAction_Invoke_APIError exercises RemoveShipModuleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRemoveShipModuleAction_Invoke_APIError(t *testing.T) {
	r := &RemoveShipModuleAction{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := RemoveShipModuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking spacetraders_remove_ship_module")
}

// TestRemoveShipModuleAction_Invoke_APIErrorReadBody exercises RemoveShipModuleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRemoveShipModuleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RemoveShipModuleAction{client: newMockClientReadErrorBody(t, 500)}
	m := RemoveShipModuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
