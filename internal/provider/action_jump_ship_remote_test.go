package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestJumpShipAction_Invoke_Happy exercises JumpShipAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestJumpShipAction_Invoke_Happy(t *testing.T) {
	r := &JumpShipAction{client: newMockClientStatus(t, 200, "{}")}
	m := JumpShipActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestJumpShipAction_Invoke_NilClient exercises JumpShipAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestJumpShipAction_Invoke_NilClient(t *testing.T) {
	r := &JumpShipAction{}
	m := JumpShipActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestJumpShipAction_Invoke_BuildError exercises JumpShipAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestJumpShipAction_Invoke_BuildError(t *testing.T) {
	r := &JumpShipAction{client: newMalformedBaseURLClient(t)}
	m := JumpShipActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestJumpShipAction_Invoke_SendError exercises JumpShipAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestJumpShipAction_Invoke_SendError(t *testing.T) {
	r := &JumpShipAction{client: newTransportErrorClient(t)}
	m := JumpShipActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestJumpShipAction_Invoke_APIError exercises JumpShipAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestJumpShipAction_Invoke_APIError(t *testing.T) {
	r := &JumpShipAction{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := JumpShipActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking spacetraders_jump_ship")
}

// TestJumpShipAction_Invoke_APIErrorReadBody exercises JumpShipAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestJumpShipAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &JumpShipAction{client: newMockClientReadErrorBody(t, 500)}
	m := JumpShipActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
