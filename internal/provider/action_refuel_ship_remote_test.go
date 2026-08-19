package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRefuelShipAction_Invoke_Happy exercises RefuelShipAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRefuelShipAction_Invoke_Happy(t *testing.T) {
	r := &RefuelShipAction{client: newMockClientStatus(t, 200, "{}")}
	m := RefuelShipActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRefuelShipAction_Invoke_NilClient exercises RefuelShipAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRefuelShipAction_Invoke_NilClient(t *testing.T) {
	r := &RefuelShipAction{}
	m := RefuelShipActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRefuelShipAction_Invoke_BuildError exercises RefuelShipAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRefuelShipAction_Invoke_BuildError(t *testing.T) {
	r := &RefuelShipAction{client: newMalformedBaseURLClient(t)}
	m := RefuelShipActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRefuelShipAction_Invoke_SendError exercises RefuelShipAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRefuelShipAction_Invoke_SendError(t *testing.T) {
	r := &RefuelShipAction{client: newTransportErrorClient(t)}
	m := RefuelShipActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRefuelShipAction_Invoke_APIError exercises RefuelShipAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRefuelShipAction_Invoke_APIError(t *testing.T) {
	r := &RefuelShipAction{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := RefuelShipActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking spacetraders_refuel_ship")
}

// TestRefuelShipAction_Invoke_APIErrorReadBody exercises RefuelShipAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRefuelShipAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RefuelShipAction{client: newMockClientReadErrorBody(t, 500)}
	m := RefuelShipActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
