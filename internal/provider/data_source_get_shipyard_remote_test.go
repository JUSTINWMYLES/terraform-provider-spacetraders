package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetShipyardDataSource_Read_Happy exercises GetShipyardDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetShipyardDataSource_Read_Happy(t *testing.T) {
	r := &GetShipyardDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetShipyardDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetShipyardDataSource_Read_NilClient exercises GetShipyardDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetShipyardDataSource_Read_NilClient(t *testing.T) {
	r := &GetShipyardDataSource{}
	m := GetShipyardDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetShipyardDataSource_Read_BuildError exercises GetShipyardDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetShipyardDataSource_Read_BuildError(t *testing.T) {
	r := &GetShipyardDataSource{client: newMalformedBaseURLClient(t)}
	m := GetShipyardDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetShipyardDataSource_Read_SendError exercises GetShipyardDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetShipyardDataSource_Read_SendError(t *testing.T) {
	r := &GetShipyardDataSource{client: newTransportErrorClient(t)}
	m := GetShipyardDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetShipyardDataSource_Read_NotFound exercises GetShipyardDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetShipyardDataSource_Read_NotFound(t *testing.T) {
	r := &GetShipyardDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetShipyardDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetShipyardDataSource_Read_APIError exercises GetShipyardDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetShipyardDataSource_Read_APIError(t *testing.T) {
	r := &GetShipyardDataSource{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := GetShipyardDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading spacetraders_get_shipyard")
}

// TestGetShipyardDataSource_Read_APIErrorReadBody exercises GetShipyardDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetShipyardDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetShipyardDataSource{client: newMockClientReadErrorBody(t, 500)}
	m := GetShipyardDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetShipyardDataSource_Read_InvalidJSON exercises GetShipyardDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetShipyardDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetShipyardDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetShipyardDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
