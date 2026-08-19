package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestCreateShipSystemScanAction_Invoke_Happy exercises CreateShipSystemScanAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestCreateShipSystemScanAction_Invoke_Happy(t *testing.T) {
	r := &CreateShipSystemScanAction{client: newMockClientStatus(t, 201, "{}")}
	m := CreateShipSystemScanActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestCreateShipSystemScanAction_Invoke_NilClient exercises CreateShipSystemScanAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestCreateShipSystemScanAction_Invoke_NilClient(t *testing.T) {
	r := &CreateShipSystemScanAction{}
	m := CreateShipSystemScanActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestCreateShipSystemScanAction_Invoke_BuildError exercises CreateShipSystemScanAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestCreateShipSystemScanAction_Invoke_BuildError(t *testing.T) {
	r := &CreateShipSystemScanAction{client: newMalformedBaseURLClient(t)}
	m := CreateShipSystemScanActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestCreateShipSystemScanAction_Invoke_SendError exercises CreateShipSystemScanAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestCreateShipSystemScanAction_Invoke_SendError(t *testing.T) {
	r := &CreateShipSystemScanAction{client: newTransportErrorClient(t)}
	m := CreateShipSystemScanActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestCreateShipSystemScanAction_Invoke_APIError exercises CreateShipSystemScanAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestCreateShipSystemScanAction_Invoke_APIError(t *testing.T) {
	r := &CreateShipSystemScanAction{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := CreateShipSystemScanActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking spacetraders_create_ship_system_scan")
}

// TestCreateShipSystemScanAction_Invoke_APIErrorReadBody exercises CreateShipSystemScanAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestCreateShipSystemScanAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &CreateShipSystemScanAction{client: newMockClientReadErrorBody(t, 500)}
	m := CreateShipSystemScanActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
