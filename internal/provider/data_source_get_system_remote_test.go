package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSystemDataSource_Read_Happy exercises GetSystemDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetSystemDataSource_Read_Happy(t *testing.T) {
	r := &GetSystemDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetSystemDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetSystemDataSource_Read_NilClient exercises GetSystemDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetSystemDataSource_Read_NilClient(t *testing.T) {
	r := &GetSystemDataSource{}
	m := GetSystemDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetSystemDataSource_Read_BuildError exercises GetSystemDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetSystemDataSource_Read_BuildError(t *testing.T) {
	r := &GetSystemDataSource{client: newMalformedBaseURLClient(t)}
	m := GetSystemDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetSystemDataSource_Read_SendError exercises GetSystemDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetSystemDataSource_Read_SendError(t *testing.T) {
	r := &GetSystemDataSource{client: newTransportErrorClient(t)}
	m := GetSystemDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetSystemDataSource_Read_NotFound exercises GetSystemDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetSystemDataSource_Read_NotFound(t *testing.T) {
	r := &GetSystemDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetSystemDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetSystemDataSource_Read_APIError exercises GetSystemDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetSystemDataSource_Read_APIError(t *testing.T) {
	r := &GetSystemDataSource{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := GetSystemDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading spacetraders_get_system")
}

// TestGetSystemDataSource_Read_APIErrorReadBody exercises GetSystemDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetSystemDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetSystemDataSource{client: newMockClientReadErrorBody(t, 500)}
	m := GetSystemDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetSystemDataSource_Read_InvalidJSON exercises GetSystemDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetSystemDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetSystemDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetSystemDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
