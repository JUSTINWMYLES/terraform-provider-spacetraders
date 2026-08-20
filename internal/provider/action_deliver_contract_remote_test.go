package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestDeliverContractAction_Invoke_Happy exercises DeliverContractAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestDeliverContractAction_Invoke_Happy(t *testing.T) {
	r := &DeliverContractAction{client: newMockClientStatus(t, 200, "{}")}
	m := DeliverContractActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestDeliverContractAction_Invoke_NilClient exercises DeliverContractAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestDeliverContractAction_Invoke_NilClient(t *testing.T) {
	r := &DeliverContractAction{}
	m := DeliverContractActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestDeliverContractAction_Invoke_BuildError exercises DeliverContractAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestDeliverContractAction_Invoke_BuildError(t *testing.T) {
	r := &DeliverContractAction{client: newMalformedBaseURLClient(t)}
	m := DeliverContractActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestDeliverContractAction_Invoke_SendError exercises DeliverContractAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestDeliverContractAction_Invoke_SendError(t *testing.T) {
	r := &DeliverContractAction{client: newTransportErrorClient(t)}
	m := DeliverContractActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestDeliverContractAction_Invoke_APIError exercises DeliverContractAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestDeliverContractAction_Invoke_APIError(t *testing.T) {
	r := &DeliverContractAction{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := DeliverContractActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking spacetraders_deliver_contract")
}

// TestDeliverContractAction_Invoke_APIErrorReadBody exercises DeliverContractAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestDeliverContractAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &DeliverContractAction{client: newMockClientReadErrorBody(t, 500)}
	m := DeliverContractActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
