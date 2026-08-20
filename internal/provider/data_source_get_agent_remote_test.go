package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAgentDataSource_Read_Happy exercises GetAgentDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetAgentDataSource_Read_Happy(t *testing.T) {
	r := &GetAgentDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetAgentDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAgentDataSource_Read_NilClient exercises GetAgentDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAgentDataSource_Read_NilClient(t *testing.T) {
	r := &GetAgentDataSource{}
	m := GetAgentDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAgentDataSource_Read_BuildError exercises GetAgentDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetAgentDataSource_Read_BuildError(t *testing.T) {
	r := &GetAgentDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAgentDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetAgentDataSource_Read_SendError exercises GetAgentDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetAgentDataSource_Read_SendError(t *testing.T) {
	r := &GetAgentDataSource{client: newTransportErrorClient(t)}
	m := GetAgentDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetAgentDataSource_Read_NotFound exercises GetAgentDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetAgentDataSource_Read_NotFound(t *testing.T) {
	r := &GetAgentDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetAgentDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetAgentDataSource_Read_APIError exercises GetAgentDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetAgentDataSource_Read_APIError(t *testing.T) {
	r := &GetAgentDataSource{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := GetAgentDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading spacetraders_get_agent")
}

// TestGetAgentDataSource_Read_APIErrorReadBody exercises GetAgentDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetAgentDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetAgentDataSource{client: newMockClientReadErrorBody(t, 500)}
	m := GetAgentDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetAgentDataSource_Read_InvalidJSON exercises GetAgentDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetAgentDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAgentDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAgentDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
