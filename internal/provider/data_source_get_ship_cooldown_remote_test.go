package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetShipCooldownDataSource_Read_Happy exercises GetShipCooldownDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetShipCooldownDataSource_Read_Happy(t *testing.T) {
	r := &GetShipCooldownDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetShipCooldownDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetShipCooldownDataSource_Read_NilClient exercises GetShipCooldownDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetShipCooldownDataSource_Read_NilClient(t *testing.T) {
	r := &GetShipCooldownDataSource{}
	m := GetShipCooldownDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetShipCooldownDataSource_Read_BuildError exercises GetShipCooldownDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetShipCooldownDataSource_Read_BuildError(t *testing.T) {
	r := &GetShipCooldownDataSource{client: newMalformedBaseURLClient(t)}
	m := GetShipCooldownDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetShipCooldownDataSource_Read_SendError exercises GetShipCooldownDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetShipCooldownDataSource_Read_SendError(t *testing.T) {
	r := &GetShipCooldownDataSource{client: newTransportErrorClient(t)}
	m := GetShipCooldownDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetShipCooldownDataSource_Read_NotFound exercises GetShipCooldownDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetShipCooldownDataSource_Read_NotFound(t *testing.T) {
	r := &GetShipCooldownDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetShipCooldownDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetShipCooldownDataSource_Read_APIError exercises GetShipCooldownDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetShipCooldownDataSource_Read_APIError(t *testing.T) {
	r := &GetShipCooldownDataSource{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := GetShipCooldownDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading spacetraders_get_ship_cooldown")
}

// TestGetShipCooldownDataSource_Read_APIErrorReadBody exercises GetShipCooldownDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetShipCooldownDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetShipCooldownDataSource{client: newMockClientReadErrorBody(t, 500)}
	m := GetShipCooldownDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetShipCooldownDataSource_Read_InvalidJSON exercises GetShipCooldownDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetShipCooldownDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetShipCooldownDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetShipCooldownDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
