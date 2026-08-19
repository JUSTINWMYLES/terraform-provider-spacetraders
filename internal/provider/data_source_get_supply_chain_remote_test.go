package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSupplyChainDataSource_Read_Happy exercises GetSupplyChainDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetSupplyChainDataSource_Read_Happy(t *testing.T) {
	r := &GetSupplyChainDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetSupplyChainDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetSupplyChainDataSource_Read_NilClient exercises GetSupplyChainDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetSupplyChainDataSource_Read_NilClient(t *testing.T) {
	r := &GetSupplyChainDataSource{}
	m := GetSupplyChainDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetSupplyChainDataSource_Read_BuildError exercises GetSupplyChainDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetSupplyChainDataSource_Read_BuildError(t *testing.T) {
	r := &GetSupplyChainDataSource{client: newMalformedBaseURLClient(t)}
	m := GetSupplyChainDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetSupplyChainDataSource_Read_SendError exercises GetSupplyChainDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetSupplyChainDataSource_Read_SendError(t *testing.T) {
	r := &GetSupplyChainDataSource{client: newTransportErrorClient(t)}
	m := GetSupplyChainDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetSupplyChainDataSource_Read_NotFound exercises GetSupplyChainDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetSupplyChainDataSource_Read_NotFound(t *testing.T) {
	r := &GetSupplyChainDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetSupplyChainDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetSupplyChainDataSource_Read_APIError exercises GetSupplyChainDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetSupplyChainDataSource_Read_APIError(t *testing.T) {
	r := &GetSupplyChainDataSource{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := GetSupplyChainDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading spacetraders_get_supply_chain")
}

// TestGetSupplyChainDataSource_Read_APIErrorReadBody exercises GetSupplyChainDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetSupplyChainDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetSupplyChainDataSource{client: newMockClientReadErrorBody(t, 500)}
	m := GetSupplyChainDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetSupplyChainDataSource_Read_InvalidJSON exercises GetSupplyChainDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetSupplyChainDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetSupplyChainDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetSupplyChainDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
