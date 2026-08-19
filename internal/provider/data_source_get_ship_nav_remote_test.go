package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetShipNavDataSource_Read_Happy exercises GetShipNavDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetShipNavDataSource_Read_Happy(t *testing.T) {
	r := &GetShipNavDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetShipNavDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetShipNavDataSource_Read_NilClient exercises GetShipNavDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetShipNavDataSource_Read_NilClient(t *testing.T) {
	r := &GetShipNavDataSource{}
	m := GetShipNavDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetShipNavDataSource_Read_BuildError exercises GetShipNavDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetShipNavDataSource_Read_BuildError(t *testing.T) {
	r := &GetShipNavDataSource{client: newMalformedBaseURLClient(t)}
	m := GetShipNavDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetShipNavDataSource_Read_SendError exercises GetShipNavDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetShipNavDataSource_Read_SendError(t *testing.T) {
	r := &GetShipNavDataSource{client: newTransportErrorClient(t)}
	m := GetShipNavDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetShipNavDataSource_Read_NotFound exercises GetShipNavDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetShipNavDataSource_Read_NotFound(t *testing.T) {
	r := &GetShipNavDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetShipNavDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetShipNavDataSource_Read_APIError exercises GetShipNavDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetShipNavDataSource_Read_APIError(t *testing.T) {
	r := &GetShipNavDataSource{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := GetShipNavDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading spacetraders_get_ship_nav")
}

// TestGetShipNavDataSource_Read_APIErrorReadBody exercises GetShipNavDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetShipNavDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetShipNavDataSource{client: newMockClientReadErrorBody(t, 500)}
	m := GetShipNavDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetShipNavDataSource_Read_InvalidJSON exercises GetShipNavDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetShipNavDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetShipNavDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetShipNavDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
