package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetMyAccountDataSource_Read_Happy exercises GetMyAccountDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetMyAccountDataSource_Read_Happy(t *testing.T) {
	r := &GetMyAccountDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetMyAccountDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetMyAccountDataSource_Read_NilClient exercises GetMyAccountDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetMyAccountDataSource_Read_NilClient(t *testing.T) {
	r := &GetMyAccountDataSource{}
	m := GetMyAccountDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetMyAccountDataSource_Read_BuildError exercises GetMyAccountDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetMyAccountDataSource_Read_BuildError(t *testing.T) {
	r := &GetMyAccountDataSource{client: newMalformedBaseURLClient(t)}
	m := GetMyAccountDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetMyAccountDataSource_Read_SendError exercises GetMyAccountDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetMyAccountDataSource_Read_SendError(t *testing.T) {
	r := &GetMyAccountDataSource{client: newTransportErrorClient(t)}
	m := GetMyAccountDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetMyAccountDataSource_Read_NotFound exercises GetMyAccountDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetMyAccountDataSource_Read_NotFound(t *testing.T) {
	r := &GetMyAccountDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetMyAccountDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetMyAccountDataSource_Read_APIError exercises GetMyAccountDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetMyAccountDataSource_Read_APIError(t *testing.T) {
	r := &GetMyAccountDataSource{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := GetMyAccountDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading spacetraders_get_my_account")
}

// TestGetMyAccountDataSource_Read_APIErrorReadBody exercises GetMyAccountDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetMyAccountDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetMyAccountDataSource{client: newMockClientReadErrorBody(t, 500)}
	m := GetMyAccountDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetMyAccountDataSource_Read_InvalidJSON exercises GetMyAccountDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetMyAccountDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetMyAccountDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetMyAccountDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
