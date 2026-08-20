package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetContractsDataSource_Read_Happy exercises GetContractsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetContractsDataSource_Read_Happy(t *testing.T) {
	r := &GetContractsDataSource{client: newMockClientStatus(t, 200, "{\"data\":[]}")}
	m := GetContractsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetContractsDataSource_Read_NilClient exercises GetContractsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetContractsDataSource_Read_NilClient(t *testing.T) {
	r := &GetContractsDataSource{}
	m := GetContractsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetContractsDataSource_Read_BuildError exercises GetContractsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetContractsDataSource_Read_BuildError(t *testing.T) {
	r := &GetContractsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetContractsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetContractsDataSource_Read_SendError exercises GetContractsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetContractsDataSource_Read_SendError(t *testing.T) {
	r := &GetContractsDataSource{client: newTransportErrorClient(t)}
	m := GetContractsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetContractsDataSource_Read_InvalidJSON exercises GetContractsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetContractsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetContractsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetContractsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
