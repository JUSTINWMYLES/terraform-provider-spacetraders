package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetMyAgentDataSource_Read_Happy exercises GetMyAgentDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetMyAgentDataSource_Read_Happy(t *testing.T) {
	r := &GetMyAgentDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetMyAgentDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetMyAgentDataSource_Read_NilClient exercises GetMyAgentDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetMyAgentDataSource_Read_NilClient(t *testing.T) {
	r := &GetMyAgentDataSource{}
	m := GetMyAgentDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetMyAgentDataSource_Read_BuildError exercises GetMyAgentDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetMyAgentDataSource_Read_BuildError(t *testing.T) {
	r := &GetMyAgentDataSource{client: newMalformedBaseURLClient(t)}
	m := GetMyAgentDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetMyAgentDataSource_Read_SendError exercises GetMyAgentDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetMyAgentDataSource_Read_SendError(t *testing.T) {
	r := &GetMyAgentDataSource{client: newTransportErrorClient(t)}
	m := GetMyAgentDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetMyAgentDataSource_Read_NotFound exercises GetMyAgentDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetMyAgentDataSource_Read_NotFound(t *testing.T) {
	r := &GetMyAgentDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetMyAgentDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetMyAgentDataSource_Read_APIError exercises GetMyAgentDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetMyAgentDataSource_Read_APIError(t *testing.T) {
	r := &GetMyAgentDataSource{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := GetMyAgentDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading spacetraders_get_my_agent")
}

// TestGetMyAgentDataSource_Read_APIErrorReadBody exercises GetMyAgentDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetMyAgentDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetMyAgentDataSource{client: newMockClientReadErrorBody(t, 500)}
	m := GetMyAgentDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetMyAgentDataSource_Read_InvalidJSON exercises GetMyAgentDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetMyAgentDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetMyAgentDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetMyAgentDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
