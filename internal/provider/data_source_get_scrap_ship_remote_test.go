package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetScrapShipDataSource_Read_Happy exercises GetScrapShipDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetScrapShipDataSource_Read_Happy(t *testing.T) {
	r := &GetScrapShipDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetScrapShipDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetScrapShipDataSource_Read_NilClient exercises GetScrapShipDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetScrapShipDataSource_Read_NilClient(t *testing.T) {
	r := &GetScrapShipDataSource{}
	m := GetScrapShipDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetScrapShipDataSource_Read_BuildError exercises GetScrapShipDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetScrapShipDataSource_Read_BuildError(t *testing.T) {
	r := &GetScrapShipDataSource{client: newMalformedBaseURLClient(t)}
	m := GetScrapShipDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetScrapShipDataSource_Read_SendError exercises GetScrapShipDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetScrapShipDataSource_Read_SendError(t *testing.T) {
	r := &GetScrapShipDataSource{client: newTransportErrorClient(t)}
	m := GetScrapShipDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetScrapShipDataSource_Read_NotFound exercises GetScrapShipDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetScrapShipDataSource_Read_NotFound(t *testing.T) {
	r := &GetScrapShipDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetScrapShipDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetScrapShipDataSource_Read_APIError exercises GetScrapShipDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetScrapShipDataSource_Read_APIError(t *testing.T) {
	r := &GetScrapShipDataSource{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := GetScrapShipDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading spacetraders_get_scrap_ship")
}

// TestGetScrapShipDataSource_Read_APIErrorReadBody exercises GetScrapShipDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetScrapShipDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetScrapShipDataSource{client: newMockClientReadErrorBody(t, 500)}
	m := GetScrapShipDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetScrapShipDataSource_Read_InvalidJSON exercises GetScrapShipDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetScrapShipDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetScrapShipDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetScrapShipDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
