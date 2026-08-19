package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetFactionDataSource_Read_Happy exercises GetFactionDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetFactionDataSource_Read_Happy(t *testing.T) {
	r := &GetFactionDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetFactionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetFactionDataSource_Read_NilClient exercises GetFactionDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetFactionDataSource_Read_NilClient(t *testing.T) {
	r := &GetFactionDataSource{}
	m := GetFactionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetFactionDataSource_Read_BuildError exercises GetFactionDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetFactionDataSource_Read_BuildError(t *testing.T) {
	r := &GetFactionDataSource{client: newMalformedBaseURLClient(t)}
	m := GetFactionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetFactionDataSource_Read_SendError exercises GetFactionDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetFactionDataSource_Read_SendError(t *testing.T) {
	r := &GetFactionDataSource{client: newTransportErrorClient(t)}
	m := GetFactionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetFactionDataSource_Read_NotFound exercises GetFactionDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetFactionDataSource_Read_NotFound(t *testing.T) {
	r := &GetFactionDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetFactionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetFactionDataSource_Read_APIError exercises GetFactionDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetFactionDataSource_Read_APIError(t *testing.T) {
	r := &GetFactionDataSource{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := GetFactionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading spacetraders_get_faction")
}

// TestGetFactionDataSource_Read_APIErrorReadBody exercises GetFactionDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetFactionDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetFactionDataSource{client: newMockClientReadErrorBody(t, 500)}
	m := GetFactionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetFactionDataSource_Read_InvalidJSON exercises GetFactionDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetFactionDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetFactionDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetFactionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
