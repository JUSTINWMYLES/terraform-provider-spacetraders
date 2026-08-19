package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetStatusDataSource_Read_Happy exercises GetStatusDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetStatusDataSource_Read_Happy(t *testing.T) {
	r := &GetStatusDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetStatusDataSource_Read_NilClient exercises GetStatusDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetStatusDataSource_Read_NilClient(t *testing.T) {
	r := &GetStatusDataSource{}
	m := GetStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetStatusDataSource_Read_BuildError exercises GetStatusDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetStatusDataSource_Read_BuildError(t *testing.T) {
	r := &GetStatusDataSource{client: newMalformedBaseURLClient(t)}
	m := GetStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetStatusDataSource_Read_SendError exercises GetStatusDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetStatusDataSource_Read_SendError(t *testing.T) {
	r := &GetStatusDataSource{client: newTransportErrorClient(t)}
	m := GetStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetStatusDataSource_Read_NotFound exercises GetStatusDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetStatusDataSource_Read_NotFound(t *testing.T) {
	r := &GetStatusDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetStatusDataSource_Read_APIError exercises GetStatusDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetStatusDataSource_Read_APIError(t *testing.T) {
	r := &GetStatusDataSource{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := GetStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading spacetraders_get_status")
}

// TestGetStatusDataSource_Read_APIErrorReadBody exercises GetStatusDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetStatusDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetStatusDataSource{client: newMockClientReadErrorBody(t, 500)}
	m := GetStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetStatusDataSource_Read_InvalidJSON exercises GetStatusDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetStatusDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetStatusDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetStatusDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
