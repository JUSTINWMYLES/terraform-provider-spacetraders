package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetWaypointDataSource_Read_Happy exercises GetWaypointDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetWaypointDataSource_Read_Happy(t *testing.T) {
	r := &GetWaypointDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetWaypointDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetWaypointDataSource_Read_NilClient exercises GetWaypointDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetWaypointDataSource_Read_NilClient(t *testing.T) {
	r := &GetWaypointDataSource{}
	m := GetWaypointDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetWaypointDataSource_Read_BuildError exercises GetWaypointDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetWaypointDataSource_Read_BuildError(t *testing.T) {
	r := &GetWaypointDataSource{client: newMalformedBaseURLClient(t)}
	m := GetWaypointDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetWaypointDataSource_Read_SendError exercises GetWaypointDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetWaypointDataSource_Read_SendError(t *testing.T) {
	r := &GetWaypointDataSource{client: newTransportErrorClient(t)}
	m := GetWaypointDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetWaypointDataSource_Read_NotFound exercises GetWaypointDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetWaypointDataSource_Read_NotFound(t *testing.T) {
	r := &GetWaypointDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetWaypointDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetWaypointDataSource_Read_APIError exercises GetWaypointDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetWaypointDataSource_Read_APIError(t *testing.T) {
	r := &GetWaypointDataSource{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := GetWaypointDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading spacetraders_get_waypoint")
}

// TestGetWaypointDataSource_Read_APIErrorReadBody exercises GetWaypointDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetWaypointDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetWaypointDataSource{client: newMockClientReadErrorBody(t, 500)}
	m := GetWaypointDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetWaypointDataSource_Read_InvalidJSON exercises GetWaypointDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetWaypointDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetWaypointDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetWaypointDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
