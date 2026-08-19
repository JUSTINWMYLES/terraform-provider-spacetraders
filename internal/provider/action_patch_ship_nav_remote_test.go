package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestPatchShipNavAction_Invoke_Happy exercises PatchShipNavAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestPatchShipNavAction_Invoke_Happy(t *testing.T) {
	r := &PatchShipNavAction{client: newMockClientStatus(t, 200, "{}")}
	m := PatchShipNavActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestPatchShipNavAction_Invoke_NilClient exercises PatchShipNavAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPatchShipNavAction_Invoke_NilClient(t *testing.T) {
	r := &PatchShipNavAction{}
	m := PatchShipNavActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestPatchShipNavAction_Invoke_BuildError exercises PatchShipNavAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestPatchShipNavAction_Invoke_BuildError(t *testing.T) {
	r := &PatchShipNavAction{client: newMalformedBaseURLClient(t)}
	m := PatchShipNavActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestPatchShipNavAction_Invoke_SendError exercises PatchShipNavAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestPatchShipNavAction_Invoke_SendError(t *testing.T) {
	r := &PatchShipNavAction{client: newTransportErrorClient(t)}
	m := PatchShipNavActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestPatchShipNavAction_Invoke_APIError exercises PatchShipNavAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestPatchShipNavAction_Invoke_APIError(t *testing.T) {
	r := &PatchShipNavAction{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := PatchShipNavActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking spacetraders_patch_ship_nav")
}

// TestPatchShipNavAction_Invoke_APIErrorReadBody exercises PatchShipNavAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestPatchShipNavAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &PatchShipNavAction{client: newMockClientReadErrorBody(t, 500)}
	m := PatchShipNavActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
