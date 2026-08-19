package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetMarketDataSource_Read_Happy exercises GetMarketDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetMarketDataSource_Read_Happy(t *testing.T) {
	r := &GetMarketDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetMarketDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetMarketDataSource_Read_NilClient exercises GetMarketDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetMarketDataSource_Read_NilClient(t *testing.T) {
	r := &GetMarketDataSource{}
	m := GetMarketDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetMarketDataSource_Read_BuildError exercises GetMarketDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetMarketDataSource_Read_BuildError(t *testing.T) {
	r := &GetMarketDataSource{client: newMalformedBaseURLClient(t)}
	m := GetMarketDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetMarketDataSource_Read_SendError exercises GetMarketDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetMarketDataSource_Read_SendError(t *testing.T) {
	r := &GetMarketDataSource{client: newTransportErrorClient(t)}
	m := GetMarketDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetMarketDataSource_Read_NotFound exercises GetMarketDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetMarketDataSource_Read_NotFound(t *testing.T) {
	r := &GetMarketDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetMarketDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetMarketDataSource_Read_APIError exercises GetMarketDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetMarketDataSource_Read_APIError(t *testing.T) {
	r := &GetMarketDataSource{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := GetMarketDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading spacetraders_get_market")
}

// TestGetMarketDataSource_Read_APIErrorReadBody exercises GetMarketDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetMarketDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetMarketDataSource{client: newMockClientReadErrorBody(t, 500)}
	m := GetMarketDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetMarketDataSource_Read_InvalidJSON exercises GetMarketDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetMarketDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetMarketDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetMarketDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
