package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestWarpShipAction_Invoke_Happy exercises WarpShipAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestWarpShipAction_Invoke_Happy(t *testing.T) {
	r := &WarpShipAction{client: newMockClientStatus(t, 200, "{}")}
	m := WarpShipActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestWarpShipAction_Invoke_NilClient exercises WarpShipAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestWarpShipAction_Invoke_NilClient(t *testing.T) {
	r := &WarpShipAction{}
	m := WarpShipActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestWarpShipAction_Invoke_BuildError exercises WarpShipAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestWarpShipAction_Invoke_BuildError(t *testing.T) {
	r := &WarpShipAction{client: newMalformedBaseURLClient(t)}
	m := WarpShipActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestWarpShipAction_Invoke_SendError exercises WarpShipAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestWarpShipAction_Invoke_SendError(t *testing.T) {
	r := &WarpShipAction{client: newTransportErrorClient(t)}
	m := WarpShipActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestWarpShipAction_Invoke_APIError exercises WarpShipAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestWarpShipAction_Invoke_APIError(t *testing.T) {
	r := &WarpShipAction{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := WarpShipActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking spacetraders_warp_ship")
}

// TestWarpShipAction_Invoke_APIErrorReadBody exercises WarpShipAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestWarpShipAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &WarpShipAction{client: newMockClientReadErrorBody(t, 500)}
	m := WarpShipActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
