package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestNegotiateContractAction_Invoke_Happy exercises NegotiateContractAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestNegotiateContractAction_Invoke_Happy(t *testing.T) {
	r := &NegotiateContractAction{client: newMockClientStatus(t, 201, "{}")}
	m := NegotiateContractActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestNegotiateContractAction_Invoke_NilClient exercises NegotiateContractAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestNegotiateContractAction_Invoke_NilClient(t *testing.T) {
	r := &NegotiateContractAction{}
	m := NegotiateContractActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestNegotiateContractAction_Invoke_BuildError exercises NegotiateContractAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestNegotiateContractAction_Invoke_BuildError(t *testing.T) {
	r := &NegotiateContractAction{client: newMalformedBaseURLClient(t)}
	m := NegotiateContractActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestNegotiateContractAction_Invoke_SendError exercises NegotiateContractAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestNegotiateContractAction_Invoke_SendError(t *testing.T) {
	r := &NegotiateContractAction{client: newTransportErrorClient(t)}
	m := NegotiateContractActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestNegotiateContractAction_Invoke_APIError exercises NegotiateContractAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestNegotiateContractAction_Invoke_APIError(t *testing.T) {
	r := &NegotiateContractAction{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := NegotiateContractActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking spacetraders_negotiate_contract")
}

// TestNegotiateContractAction_Invoke_APIErrorReadBody exercises NegotiateContractAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestNegotiateContractAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &NegotiateContractAction{client: newMockClientReadErrorBody(t, 500)}
	m := NegotiateContractActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
