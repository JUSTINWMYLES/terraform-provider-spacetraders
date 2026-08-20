package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetJumpGateDataSource_Read_Happy exercises GetJumpGateDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetJumpGateDataSource_Read_Happy(t *testing.T) {
	r := &GetJumpGateDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetJumpGateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetJumpGateDataSource_Read_NilClient exercises GetJumpGateDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetJumpGateDataSource_Read_NilClient(t *testing.T) {
	r := &GetJumpGateDataSource{}
	m := GetJumpGateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetJumpGateDataSource_Read_BuildError exercises GetJumpGateDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetJumpGateDataSource_Read_BuildError(t *testing.T) {
	r := &GetJumpGateDataSource{client: newMalformedBaseURLClient(t)}
	m := GetJumpGateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetJumpGateDataSource_Read_SendError exercises GetJumpGateDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetJumpGateDataSource_Read_SendError(t *testing.T) {
	r := &GetJumpGateDataSource{client: newTransportErrorClient(t)}
	m := GetJumpGateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetJumpGateDataSource_Read_NotFound exercises GetJumpGateDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetJumpGateDataSource_Read_NotFound(t *testing.T) {
	r := &GetJumpGateDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetJumpGateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetJumpGateDataSource_Read_APIError exercises GetJumpGateDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetJumpGateDataSource_Read_APIError(t *testing.T) {
	r := &GetJumpGateDataSource{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := GetJumpGateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading spacetraders_get_jump_gate")
}

// TestGetJumpGateDataSource_Read_APIErrorReadBody exercises GetJumpGateDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetJumpGateDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetJumpGateDataSource{client: newMockClientReadErrorBody(t, 500)}
	m := GetJumpGateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetJumpGateDataSource_Read_InvalidJSON exercises GetJumpGateDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetJumpGateDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetJumpGateDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetJumpGateDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
