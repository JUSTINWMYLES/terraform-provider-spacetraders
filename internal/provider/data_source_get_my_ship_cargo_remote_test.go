package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetMyShipCargoDataSource_Read_Happy exercises GetMyShipCargoDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetMyShipCargoDataSource_Read_Happy(t *testing.T) {
	r := &GetMyShipCargoDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetMyShipCargoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetMyShipCargoDataSource_Read_NilClient exercises GetMyShipCargoDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetMyShipCargoDataSource_Read_NilClient(t *testing.T) {
	r := &GetMyShipCargoDataSource{}
	m := GetMyShipCargoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetMyShipCargoDataSource_Read_BuildError exercises GetMyShipCargoDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetMyShipCargoDataSource_Read_BuildError(t *testing.T) {
	r := &GetMyShipCargoDataSource{client: newMalformedBaseURLClient(t)}
	m := GetMyShipCargoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetMyShipCargoDataSource_Read_SendError exercises GetMyShipCargoDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetMyShipCargoDataSource_Read_SendError(t *testing.T) {
	r := &GetMyShipCargoDataSource{client: newTransportErrorClient(t)}
	m := GetMyShipCargoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetMyShipCargoDataSource_Read_NotFound exercises GetMyShipCargoDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetMyShipCargoDataSource_Read_NotFound(t *testing.T) {
	r := &GetMyShipCargoDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetMyShipCargoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetMyShipCargoDataSource_Read_APIError exercises GetMyShipCargoDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetMyShipCargoDataSource_Read_APIError(t *testing.T) {
	r := &GetMyShipCargoDataSource{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := GetMyShipCargoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading spacetraders_get_my_ship_cargo")
}

// TestGetMyShipCargoDataSource_Read_APIErrorReadBody exercises GetMyShipCargoDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetMyShipCargoDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetMyShipCargoDataSource{client: newMockClientReadErrorBody(t, 500)}
	m := GetMyShipCargoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetMyShipCargoDataSource_Read_InvalidJSON exercises GetMyShipCargoDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetMyShipCargoDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetMyShipCargoDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetMyShipCargoDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
