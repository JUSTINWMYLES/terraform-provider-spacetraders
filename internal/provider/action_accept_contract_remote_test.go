package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestAcceptContractAction_Invoke_Happy exercises AcceptContractAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestAcceptContractAction_Invoke_Happy(t *testing.T) {
	r := &AcceptContractAction{client: newMockClientStatus(t, 200, "{}")}
	m := AcceptContractActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestAcceptContractAction_Invoke_NilClient exercises AcceptContractAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestAcceptContractAction_Invoke_NilClient(t *testing.T) {
	r := &AcceptContractAction{}
	m := AcceptContractActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestAcceptContractAction_Invoke_BuildError exercises AcceptContractAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestAcceptContractAction_Invoke_BuildError(t *testing.T) {
	r := &AcceptContractAction{client: newMalformedBaseURLClient(t)}
	m := AcceptContractActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestAcceptContractAction_Invoke_SendError exercises AcceptContractAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestAcceptContractAction_Invoke_SendError(t *testing.T) {
	r := &AcceptContractAction{client: newTransportErrorClient(t)}
	m := AcceptContractActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestAcceptContractAction_Invoke_APIError exercises AcceptContractAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestAcceptContractAction_Invoke_APIError(t *testing.T) {
	r := &AcceptContractAction{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := AcceptContractActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking spacetraders_accept_contract")
}

// TestAcceptContractAction_Invoke_APIErrorReadBody exercises AcceptContractAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestAcceptContractAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &AcceptContractAction{client: newMockClientReadErrorBody(t, 500)}
	m := AcceptContractActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
