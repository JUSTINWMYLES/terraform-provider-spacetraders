package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestSellCargoAction_Invoke_Happy exercises SellCargoAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestSellCargoAction_Invoke_Happy(t *testing.T) {
	r := &SellCargoAction{client: newMockClientStatus(t, 201, "{}")}
	m := SellCargoActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestSellCargoAction_Invoke_NilClient exercises SellCargoAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestSellCargoAction_Invoke_NilClient(t *testing.T) {
	r := &SellCargoAction{}
	m := SellCargoActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestSellCargoAction_Invoke_BuildError exercises SellCargoAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestSellCargoAction_Invoke_BuildError(t *testing.T) {
	r := &SellCargoAction{client: newMalformedBaseURLClient(t)}
	m := SellCargoActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestSellCargoAction_Invoke_SendError exercises SellCargoAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestSellCargoAction_Invoke_SendError(t *testing.T) {
	r := &SellCargoAction{client: newTransportErrorClient(t)}
	m := SellCargoActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestSellCargoAction_Invoke_APIError exercises SellCargoAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestSellCargoAction_Invoke_APIError(t *testing.T) {
	r := &SellCargoAction{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := SellCargoActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking spacetraders_sell_cargo")
}

// TestSellCargoAction_Invoke_APIErrorReadBody exercises SellCargoAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestSellCargoAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &SellCargoAction{client: newMockClientReadErrorBody(t, 500)}
	m := SellCargoActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
