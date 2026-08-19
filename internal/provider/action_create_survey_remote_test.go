package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestCreateSurveyAction_Invoke_Happy exercises CreateSurveyAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestCreateSurveyAction_Invoke_Happy(t *testing.T) {
	r := &CreateSurveyAction{client: newMockClientStatus(t, 201, "{}")}
	m := CreateSurveyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestCreateSurveyAction_Invoke_NilClient exercises CreateSurveyAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestCreateSurveyAction_Invoke_NilClient(t *testing.T) {
	r := &CreateSurveyAction{}
	m := CreateSurveyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestCreateSurveyAction_Invoke_BuildError exercises CreateSurveyAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestCreateSurveyAction_Invoke_BuildError(t *testing.T) {
	r := &CreateSurveyAction{client: newMalformedBaseURLClient(t)}
	m := CreateSurveyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestCreateSurveyAction_Invoke_SendError exercises CreateSurveyAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestCreateSurveyAction_Invoke_SendError(t *testing.T) {
	r := &CreateSurveyAction{client: newTransportErrorClient(t)}
	m := CreateSurveyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestCreateSurveyAction_Invoke_APIError exercises CreateSurveyAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestCreateSurveyAction_Invoke_APIError(t *testing.T) {
	r := &CreateSurveyAction{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := CreateSurveyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking spacetraders_create_survey")
}

// TestCreateSurveyAction_Invoke_APIErrorReadBody exercises CreateSurveyAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestCreateSurveyAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &CreateSurveyAction{client: newMockClientReadErrorBody(t, 500)}
	m := CreateSurveyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
