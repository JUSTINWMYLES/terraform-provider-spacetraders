package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestExtractResourcesWithSurveyAction_Invoke_Happy exercises ExtractResourcesWithSurveyAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestExtractResourcesWithSurveyAction_Invoke_Happy(t *testing.T) {
	r := &ExtractResourcesWithSurveyAction{client: newMockClientStatus(t, 201, "{}")}
	m := ExtractResourcesWithSurveyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestExtractResourcesWithSurveyAction_Invoke_NilClient exercises ExtractResourcesWithSurveyAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestExtractResourcesWithSurveyAction_Invoke_NilClient(t *testing.T) {
	r := &ExtractResourcesWithSurveyAction{}
	m := ExtractResourcesWithSurveyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestExtractResourcesWithSurveyAction_Invoke_BuildError exercises ExtractResourcesWithSurveyAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestExtractResourcesWithSurveyAction_Invoke_BuildError(t *testing.T) {
	r := &ExtractResourcesWithSurveyAction{client: newMalformedBaseURLClient(t)}
	m := ExtractResourcesWithSurveyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestExtractResourcesWithSurveyAction_Invoke_SendError exercises ExtractResourcesWithSurveyAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestExtractResourcesWithSurveyAction_Invoke_SendError(t *testing.T) {
	r := &ExtractResourcesWithSurveyAction{client: newTransportErrorClient(t)}
	m := ExtractResourcesWithSurveyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestExtractResourcesWithSurveyAction_Invoke_APIError exercises ExtractResourcesWithSurveyAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestExtractResourcesWithSurveyAction_Invoke_APIError(t *testing.T) {
	r := &ExtractResourcesWithSurveyAction{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := ExtractResourcesWithSurveyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking spacetraders_extract_resources_with_survey")
}

// TestExtractResourcesWithSurveyAction_Invoke_APIErrorReadBody exercises ExtractResourcesWithSurveyAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestExtractResourcesWithSurveyAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &ExtractResourcesWithSurveyAction{client: newMockClientReadErrorBody(t, 500)}
	m := ExtractResourcesWithSurveyActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
