package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestRepairShipAction_Invoke_Happy exercises RepairShipAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestRepairShipAction_Invoke_Happy(t *testing.T) {
	r := &RepairShipAction{client: newMockClientStatus(t, 200, "{}")}
	m := RepairShipActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestRepairShipAction_Invoke_NilClient exercises RepairShipAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestRepairShipAction_Invoke_NilClient(t *testing.T) {
	r := &RepairShipAction{}
	m := RepairShipActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestRepairShipAction_Invoke_BuildError exercises RepairShipAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestRepairShipAction_Invoke_BuildError(t *testing.T) {
	r := &RepairShipAction{client: newMalformedBaseURLClient(t)}
	m := RepairShipActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestRepairShipAction_Invoke_SendError exercises RepairShipAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestRepairShipAction_Invoke_SendError(t *testing.T) {
	r := &RepairShipAction{client: newTransportErrorClient(t)}
	m := RepairShipActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestRepairShipAction_Invoke_APIError exercises RepairShipAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestRepairShipAction_Invoke_APIError(t *testing.T) {
	r := &RepairShipAction{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := RepairShipActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking spacetraders_repair_ship")
}

// TestRepairShipAction_Invoke_APIErrorReadBody exercises RepairShipAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestRepairShipAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &RepairShipAction{client: newMockClientReadErrorBody(t, 500)}
	m := RepairShipActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
