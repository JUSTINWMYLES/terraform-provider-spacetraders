package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSystemsDataSource_Read_Happy exercises GetSystemsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetSystemsDataSource_Read_Happy(t *testing.T) {
	r := &GetSystemsDataSource{client: newMockClientStatus(t, 200, "{\"data\":[]}")}
	m := GetSystemsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetSystemsDataSource_Read_NilClient exercises GetSystemsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetSystemsDataSource_Read_NilClient(t *testing.T) {
	r := &GetSystemsDataSource{}
	m := GetSystemsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetSystemsDataSource_Read_BuildError exercises GetSystemsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetSystemsDataSource_Read_BuildError(t *testing.T) {
	r := &GetSystemsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetSystemsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetSystemsDataSource_Read_SendError exercises GetSystemsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetSystemsDataSource_Read_SendError(t *testing.T) {
	r := &GetSystemsDataSource{client: newTransportErrorClient(t)}
	m := GetSystemsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetSystemsDataSource_Read_InvalidJSON exercises GetSystemsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetSystemsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetSystemsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetSystemsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
