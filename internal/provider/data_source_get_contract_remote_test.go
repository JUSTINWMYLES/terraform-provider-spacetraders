package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetContractDataSource_Read_Happy exercises GetContractDataSource.readRemote against an httptest mock: happy path returns the success status with no errors.
func TestGetContractDataSource_Read_Happy(t *testing.T) {
	r := &GetContractDataSource{client: newMockClientStatus(t, 200, "{}")}
	m := GetContractDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetContractDataSource_Read_NilClient exercises GetContractDataSource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetContractDataSource_Read_NilClient(t *testing.T) {
	r := &GetContractDataSource{}
	m := GetContractDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetContractDataSource_Read_BuildError exercises GetContractDataSource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestGetContractDataSource_Read_BuildError(t *testing.T) {
	r := &GetContractDataSource{client: newMalformedBaseURLClient(t)}
	m := GetContractDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestGetContractDataSource_Read_SendError exercises GetContractDataSource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestGetContractDataSource_Read_SendError(t *testing.T) {
	r := &GetContractDataSource{client: newTransportErrorClient(t)}
	m := GetContractDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestGetContractDataSource_Read_NotFound exercises GetContractDataSource.readRemote against an httptest mock: 404 surfaces the requested-resource-not-found error.
func TestGetContractDataSource_Read_NotFound(t *testing.T) {
	r := &GetContractDataSource{client: newMockClientStatus(t, 404, "")}
	m := GetContractDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "The requested resource was not found.")
}

// TestGetContractDataSource_Read_APIError exercises GetContractDataSource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestGetContractDataSource_Read_APIError(t *testing.T) {
	r := &GetContractDataSource{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := GetContractDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading spacetraders_get_contract")
}

// TestGetContractDataSource_Read_APIErrorReadBody exercises GetContractDataSource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestGetContractDataSource_Read_APIErrorReadBody(t *testing.T) {
	r := &GetContractDataSource{client: newMockClientReadErrorBody(t, 500)}
	m := GetContractDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestGetContractDataSource_Read_InvalidJSON exercises GetContractDataSource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestGetContractDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetContractDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetContractDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}
