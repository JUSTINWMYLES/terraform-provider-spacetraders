package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestCreateShipShipScanAction_Invoke_Happy exercises CreateShipShipScanAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestCreateShipShipScanAction_Invoke_Happy(t *testing.T) {
	r := &CreateShipShipScanAction{client: newMockClientStatus(t, 201, "{}")}
	m := CreateShipShipScanActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestCreateShipShipScanAction_Invoke_NilClient exercises CreateShipShipScanAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestCreateShipShipScanAction_Invoke_NilClient(t *testing.T) {
	r := &CreateShipShipScanAction{}
	m := CreateShipShipScanActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestCreateShipShipScanAction_Invoke_BuildError exercises CreateShipShipScanAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestCreateShipShipScanAction_Invoke_BuildError(t *testing.T) {
	r := &CreateShipShipScanAction{client: newMalformedBaseURLClient(t)}
	m := CreateShipShipScanActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestCreateShipShipScanAction_Invoke_SendError exercises CreateShipShipScanAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestCreateShipShipScanAction_Invoke_SendError(t *testing.T) {
	r := &CreateShipShipScanAction{client: newTransportErrorClient(t)}
	m := CreateShipShipScanActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestCreateShipShipScanAction_Invoke_APIError exercises CreateShipShipScanAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestCreateShipShipScanAction_Invoke_APIError(t *testing.T) {
	r := &CreateShipShipScanAction{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := CreateShipShipScanActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking spacetraders_create_ship_ship_scan")
}

// TestCreateShipShipScanAction_Invoke_APIErrorReadBody exercises CreateShipShipScanAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestCreateShipShipScanAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &CreateShipShipScanAction{client: newMockClientReadErrorBody(t, 500)}
	m := CreateShipShipScanActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
