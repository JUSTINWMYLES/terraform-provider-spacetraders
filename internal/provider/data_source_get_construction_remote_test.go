package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetConstructionDataSource_Read_Happy exercises GetConstructionDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetConstructionDataSource_Read_Happy(t *testing.T) {
	r := &GetConstructionDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetConstructionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetConstructionDataSource_Read_NilClient exercises GetConstructionDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetConstructionDataSource_Read_NilClient(t *testing.T) {
	r := &GetConstructionDataSource{}
	m := GetConstructionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetConstructionDataSource_Read_BuildError exercises GetConstructionDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetConstructionDataSource_Read_BuildError(t *testing.T) {
	r := &GetConstructionDataSource{client: newMalformedBaseURLClient(t)}
	m := GetConstructionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetConstructionDataSource_Read_SendError exercises GetConstructionDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetConstructionDataSource_Read_SendError(t *testing.T) {
	r := &GetConstructionDataSource{client: newTransportErrorClient(t)}
	m := GetConstructionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetConstructionDataSource_Read_NotFound exercises GetConstructionDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetConstructionDataSource_Read_NotFound(t *testing.T) {
	r := &GetConstructionDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetConstructionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetConstructionDataSource_Read_APIError exercises GetConstructionDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetConstructionDataSource_Read_APIError(t *testing.T) {
	r := &GetConstructionDataSource{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := GetConstructionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading spacetraders_get_construction")
}

// TestGetConstructionDataSource_Read_APIErrorReadBody exercises GetConstructionDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetConstructionDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetConstructionDataSource{client: newMockClientReadErrorBody(t, 500)}
	m := GetConstructionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetConstructionDataSource_Read_InvalidJSON exercises GetConstructionDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetConstructionDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetConstructionDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetConstructionDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
