package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestShipRefineAction_Invoke_Happy exercises ShipRefineAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestShipRefineAction_Invoke_Happy(t *testing.T) {
	r := &ShipRefineAction{client: newMockClientStatus(t, 201, "{}")}
	m := ShipRefineActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestShipRefineAction_Invoke_NilClient exercises ShipRefineAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestShipRefineAction_Invoke_NilClient(t *testing.T) {
	r := &ShipRefineAction{}
	m := ShipRefineActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestShipRefineAction_Invoke_BuildError exercises ShipRefineAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestShipRefineAction_Invoke_BuildError(t *testing.T) {
	r := &ShipRefineAction{client: newMalformedBaseURLClient(t)}
	m := ShipRefineActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestShipRefineAction_Invoke_SendError exercises ShipRefineAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestShipRefineAction_Invoke_SendError(t *testing.T) {
	r := &ShipRefineAction{client: newTransportErrorClient(t)}
	m := ShipRefineActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestShipRefineAction_Invoke_APIError exercises ShipRefineAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestShipRefineAction_Invoke_APIError(t *testing.T) {
	r := &ShipRefineAction{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := ShipRefineActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking spacetraders_ship_refine")
}

// TestShipRefineAction_Invoke_APIErrorReadBody exercises ShipRefineAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestShipRefineAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ShipRefineAction{client: newMockClientReadErrorBody(t, 500)}
	m := ShipRefineActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
