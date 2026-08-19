package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestOrbitShipAction_Invoke_Happy exercises OrbitShipAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestOrbitShipAction_Invoke_Happy(t *testing.T) {
	r := &OrbitShipAction{client: newMockClientStatus(t, 200, "{}")}
	m := OrbitShipActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestOrbitShipAction_Invoke_NilClient exercises OrbitShipAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestOrbitShipAction_Invoke_NilClient(t *testing.T) {
	r := &OrbitShipAction{}
	m := OrbitShipActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestOrbitShipAction_Invoke_BuildError exercises OrbitShipAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestOrbitShipAction_Invoke_BuildError(t *testing.T) {
	r := &OrbitShipAction{client: newMalformedBaseURLClient(t)}
	m := OrbitShipActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestOrbitShipAction_Invoke_SendError exercises OrbitShipAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestOrbitShipAction_Invoke_SendError(t *testing.T) {
	r := &OrbitShipAction{client: newTransportErrorClient(t)}
	m := OrbitShipActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestOrbitShipAction_Invoke_APIError exercises OrbitShipAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestOrbitShipAction_Invoke_APIError(t *testing.T) {
	r := &OrbitShipAction{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := OrbitShipActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking spacetraders_orbit_ship")
}

// TestOrbitShipAction_Invoke_APIErrorReadBody exercises OrbitShipAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestOrbitShipAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &OrbitShipAction{client: newMockClientReadErrorBody(t, 500)}
	m := OrbitShipActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
