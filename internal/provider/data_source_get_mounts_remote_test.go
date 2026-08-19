package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetMountsDataSource_Read_Happy exercises GetMountsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetMountsDataSource_Read_Happy(t *testing.T) {
	r := &GetMountsDataSource{client: newMockClientStatus(t, 200, "{\"data\":[]}")}
	m := GetMountsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetMountsDataSource_Read_NilClient exercises GetMountsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetMountsDataSource_Read_NilClient(t *testing.T) {
	r := &GetMountsDataSource{}
	m := GetMountsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetMountsDataSource_Read_BuildError exercises GetMountsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetMountsDataSource_Read_BuildError(t *testing.T) {
	r := &GetMountsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetMountsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetMountsDataSource_Read_SendError exercises GetMountsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetMountsDataSource_Read_SendError(t *testing.T) {
	r := &GetMountsDataSource{client: newTransportErrorClient(t)}
	m := GetMountsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetMountsDataSource_Read_InvalidJSON exercises GetMountsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetMountsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetMountsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetMountsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
