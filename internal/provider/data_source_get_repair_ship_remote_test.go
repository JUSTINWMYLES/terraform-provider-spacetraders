package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetRepairShipDataSource_Read_Happy exercises GetRepairShipDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetRepairShipDataSource_Read_Happy(t *testing.T) {
	r := &GetRepairShipDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetRepairShipDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetRepairShipDataSource_Read_NilClient exercises GetRepairShipDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetRepairShipDataSource_Read_NilClient(t *testing.T) {
	r := &GetRepairShipDataSource{}
	m := GetRepairShipDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetRepairShipDataSource_Read_BuildError exercises GetRepairShipDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetRepairShipDataSource_Read_BuildError(t *testing.T) {
	r := &GetRepairShipDataSource{client: newMalformedBaseURLClient(t)}
	m := GetRepairShipDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetRepairShipDataSource_Read_SendError exercises GetRepairShipDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetRepairShipDataSource_Read_SendError(t *testing.T) {
	r := &GetRepairShipDataSource{client: newTransportErrorClient(t)}
	m := GetRepairShipDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetRepairShipDataSource_Read_NotFound exercises GetRepairShipDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetRepairShipDataSource_Read_NotFound(t *testing.T) {
	r := &GetRepairShipDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetRepairShipDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetRepairShipDataSource_Read_APIError exercises GetRepairShipDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetRepairShipDataSource_Read_APIError(t *testing.T) {
	r := &GetRepairShipDataSource{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := GetRepairShipDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading spacetraders_get_repair_ship")
}

// TestGetRepairShipDataSource_Read_APIErrorReadBody exercises GetRepairShipDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetRepairShipDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetRepairShipDataSource{client: newMockClientReadErrorBody(t, 500)}
	m := GetRepairShipDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetRepairShipDataSource_Read_InvalidJSON exercises GetRepairShipDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetRepairShipDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetRepairShipDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetRepairShipDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
