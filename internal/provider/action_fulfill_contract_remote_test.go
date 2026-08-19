package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestFulfillContractAction_Invoke_Happy exercises FulfillContractAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestFulfillContractAction_Invoke_Happy(t *testing.T) {
	r := &FulfillContractAction{client: newMockClientStatus(t, 200, "{}")}
	m := FulfillContractActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestFulfillContractAction_Invoke_NilClient exercises FulfillContractAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestFulfillContractAction_Invoke_NilClient(t *testing.T) {
	r := &FulfillContractAction{}
	m := FulfillContractActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestFulfillContractAction_Invoke_BuildError exercises FulfillContractAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestFulfillContractAction_Invoke_BuildError(t *testing.T) {
	r := &FulfillContractAction{client: newMalformedBaseURLClient(t)}
	m := FulfillContractActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestFulfillContractAction_Invoke_SendError exercises FulfillContractAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestFulfillContractAction_Invoke_SendError(t *testing.T) {
	r := &FulfillContractAction{client: newTransportErrorClient(t)}
	m := FulfillContractActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestFulfillContractAction_Invoke_APIError exercises FulfillContractAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestFulfillContractAction_Invoke_APIError(t *testing.T) {
	r := &FulfillContractAction{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := FulfillContractActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking spacetraders_fulfill_contract")
}

// TestFulfillContractAction_Invoke_APIErrorReadBody exercises FulfillContractAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestFulfillContractAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &FulfillContractAction{client: newMockClientReadErrorBody(t, 500)}
	m := FulfillContractActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
