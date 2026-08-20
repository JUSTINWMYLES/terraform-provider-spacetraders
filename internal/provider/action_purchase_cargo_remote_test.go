package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestPurchaseCargoAction_Invoke_Happy exercises PurchaseCargoAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestPurchaseCargoAction_Invoke_Happy(t *testing.T) {
	r := &PurchaseCargoAction{client: newMockClientStatus(t, 201, "{}")}
	m := PurchaseCargoActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestPurchaseCargoAction_Invoke_NilClient exercises PurchaseCargoAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestPurchaseCargoAction_Invoke_NilClient(t *testing.T) {
	r := &PurchaseCargoAction{}
	m := PurchaseCargoActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestPurchaseCargoAction_Invoke_BuildError exercises PurchaseCargoAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestPurchaseCargoAction_Invoke_BuildError(t *testing.T) {
	r := &PurchaseCargoAction{client: newMalformedBaseURLClient(t)}
	m := PurchaseCargoActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestPurchaseCargoAction_Invoke_SendError exercises PurchaseCargoAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestPurchaseCargoAction_Invoke_SendError(t *testing.T) {
	r := &PurchaseCargoAction{client: newTransportErrorClient(t)}
	m := PurchaseCargoActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestPurchaseCargoAction_Invoke_APIError exercises PurchaseCargoAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestPurchaseCargoAction_Invoke_APIError(t *testing.T) {
	r := &PurchaseCargoAction{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := PurchaseCargoActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking spacetraders_purchase_cargo")
}

// TestPurchaseCargoAction_Invoke_APIErrorReadBody exercises PurchaseCargoAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestPurchaseCargoAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &PurchaseCargoAction{client: newMockClientReadErrorBody(t, 500)}
	m := PurchaseCargoActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
