package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestSupplyConstructionAction_Invoke_Happy exercises SupplyConstructionAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestSupplyConstructionAction_Invoke_Happy(t *testing.T) {
	r := &SupplyConstructionAction{client: newMockClientStatus(t, 201, "{}")}
	m := SupplyConstructionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestSupplyConstructionAction_Invoke_NilClient exercises SupplyConstructionAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSupplyConstructionAction_Invoke_NilClient(t *testing.T) {
	r := &SupplyConstructionAction{}
	m := SupplyConstructionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestSupplyConstructionAction_Invoke_BuildError exercises SupplyConstructionAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestSupplyConstructionAction_Invoke_BuildError(t *testing.T) {
	r := &SupplyConstructionAction{client: newMalformedBaseURLClient(t)}
	m := SupplyConstructionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestSupplyConstructionAction_Invoke_SendError exercises SupplyConstructionAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestSupplyConstructionAction_Invoke_SendError(t *testing.T) {
	r := &SupplyConstructionAction{client: newTransportErrorClient(t)}
	m := SupplyConstructionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestSupplyConstructionAction_Invoke_APIError exercises SupplyConstructionAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestSupplyConstructionAction_Invoke_APIError(t *testing.T) {
	r := &SupplyConstructionAction{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := SupplyConstructionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking spacetraders_supply_construction")
}

// TestSupplyConstructionAction_Invoke_APIErrorReadBody exercises SupplyConstructionAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestSupplyConstructionAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &SupplyConstructionAction{client: newMockClientReadErrorBody(t, 500)}
	m := SupplyConstructionActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
