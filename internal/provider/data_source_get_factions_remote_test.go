package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetFactionsDataSource_Read_Happy exercises GetFactionsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetFactionsDataSource_Read_Happy(t *testing.T) {
	r := &GetFactionsDataSource{client: newMockClientStatus(t, 200, "{\"data\":[]}")}
	m := GetFactionsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetFactionsDataSource_Read_NilClient exercises GetFactionsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetFactionsDataSource_Read_NilClient(t *testing.T) {
	r := &GetFactionsDataSource{}
	m := GetFactionsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetFactionsDataSource_Read_BuildError exercises GetFactionsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetFactionsDataSource_Read_BuildError(t *testing.T) {
	r := &GetFactionsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetFactionsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetFactionsDataSource_Read_SendError exercises GetFactionsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetFactionsDataSource_Read_SendError(t *testing.T) {
	r := &GetFactionsDataSource{client: newTransportErrorClient(t)}
	m := GetFactionsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetFactionsDataSource_Read_InvalidJSON exercises GetFactionsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetFactionsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetFactionsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetFactionsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
