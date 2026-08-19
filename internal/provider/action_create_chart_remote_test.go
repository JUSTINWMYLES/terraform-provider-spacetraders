package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestCreateChartAction_Invoke_Happy exercises CreateChartAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestCreateChartAction_Invoke_Happy(t *testing.T) {
	r := &CreateChartAction{client: newMockClientStatus(t, 201, "{}")}
	m := CreateChartActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestCreateChartAction_Invoke_NilClient exercises CreateChartAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestCreateChartAction_Invoke_NilClient(t *testing.T) {
	r := &CreateChartAction{}
	m := CreateChartActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestCreateChartAction_Invoke_BuildError exercises CreateChartAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestCreateChartAction_Invoke_BuildError(t *testing.T) {
	r := &CreateChartAction{client: newMalformedBaseURLClient(t)}
	m := CreateChartActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestCreateChartAction_Invoke_SendError exercises CreateChartAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestCreateChartAction_Invoke_SendError(t *testing.T) {
	r := &CreateChartAction{client: newTransportErrorClient(t)}
	m := CreateChartActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestCreateChartAction_Invoke_APIError exercises CreateChartAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestCreateChartAction_Invoke_APIError(t *testing.T) {
	r := &CreateChartAction{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := CreateChartActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking spacetraders_create_chart")
}

// TestCreateChartAction_Invoke_APIErrorReadBody exercises CreateChartAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestCreateChartAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &CreateChartAction{client: newMockClientReadErrorBody(t, 500)}
	m := CreateChartActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
