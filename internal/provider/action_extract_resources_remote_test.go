package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestExtractResourcesAction_Invoke_Happy exercises ExtractResourcesAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestExtractResourcesAction_Invoke_Happy(t *testing.T) {
	r := &ExtractResourcesAction{client: newMockClientStatus(t, 201, "{}")}
	m := ExtractResourcesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestExtractResourcesAction_Invoke_NilClient exercises ExtractResourcesAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestExtractResourcesAction_Invoke_NilClient(t *testing.T) {
	r := &ExtractResourcesAction{}
	m := ExtractResourcesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestExtractResourcesAction_Invoke_BuildError exercises ExtractResourcesAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestExtractResourcesAction_Invoke_BuildError(t *testing.T) {
	r := &ExtractResourcesAction{client: newMalformedBaseURLClient(t)}
	m := ExtractResourcesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestExtractResourcesAction_Invoke_SendError exercises ExtractResourcesAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestExtractResourcesAction_Invoke_SendError(t *testing.T) {
	r := &ExtractResourcesAction{client: newTransportErrorClient(t)}
	m := ExtractResourcesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestExtractResourcesAction_Invoke_APIError exercises ExtractResourcesAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestExtractResourcesAction_Invoke_APIError(t *testing.T) {
	r := &ExtractResourcesAction{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := ExtractResourcesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking spacetraders_extract_resources")
}

// TestExtractResourcesAction_Invoke_APIErrorReadBody exercises ExtractResourcesAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestExtractResourcesAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ExtractResourcesAction{client: newMockClientReadErrorBody(t, 500)}
	m := ExtractResourcesActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
