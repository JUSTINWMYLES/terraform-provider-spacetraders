package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestCreateShipWaypointScanAction_Invoke_Happy exercises CreateShipWaypointScanAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestCreateShipWaypointScanAction_Invoke_Happy(t *testing.T) {
	r := &CreateShipWaypointScanAction{client: newMockClientStatus(t, 201, "{}")}
	m := CreateShipWaypointScanActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestCreateShipWaypointScanAction_Invoke_NilClient exercises CreateShipWaypointScanAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestCreateShipWaypointScanAction_Invoke_NilClient(t *testing.T) {
	r := &CreateShipWaypointScanAction{}
	m := CreateShipWaypointScanActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestCreateShipWaypointScanAction_Invoke_BuildError exercises CreateShipWaypointScanAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestCreateShipWaypointScanAction_Invoke_BuildError(t *testing.T) {
	r := &CreateShipWaypointScanAction{client: newMalformedBaseURLClient(t)}
	m := CreateShipWaypointScanActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestCreateShipWaypointScanAction_Invoke_SendError exercises CreateShipWaypointScanAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestCreateShipWaypointScanAction_Invoke_SendError(t *testing.T) {
	r := &CreateShipWaypointScanAction{client: newTransportErrorClient(t)}
	m := CreateShipWaypointScanActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestCreateShipWaypointScanAction_Invoke_APIError exercises CreateShipWaypointScanAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestCreateShipWaypointScanAction_Invoke_APIError(t *testing.T) {
	r := &CreateShipWaypointScanAction{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := CreateShipWaypointScanActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking spacetraders_create_ship_waypoint_scan")
}

// TestCreateShipWaypointScanAction_Invoke_APIErrorReadBody exercises CreateShipWaypointScanAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestCreateShipWaypointScanAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &CreateShipWaypointScanAction{client: newMockClientReadErrorBody(t, 500)}
	m := CreateShipWaypointScanActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
