package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetMyFactionsDataSource_Read_Happy exercises GetMyFactionsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetMyFactionsDataSource_Read_Happy(t *testing.T) {
	r := &GetMyFactionsDataSource{client: newMockClientStatus(t, 200, "{\"data\":[]}")}
	m := GetMyFactionsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetMyFactionsDataSource_Read_NilClient exercises GetMyFactionsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetMyFactionsDataSource_Read_NilClient(t *testing.T) {
	r := &GetMyFactionsDataSource{}
	m := GetMyFactionsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetMyFactionsDataSource_Read_BuildError exercises GetMyFactionsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetMyFactionsDataSource_Read_BuildError(t *testing.T) {
	r := &GetMyFactionsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetMyFactionsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetMyFactionsDataSource_Read_SendError exercises GetMyFactionsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetMyFactionsDataSource_Read_SendError(t *testing.T) {
	r := &GetMyFactionsDataSource{client: newTransportErrorClient(t)}
	m := GetMyFactionsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetMyFactionsDataSource_Read_InvalidJSON exercises GetMyFactionsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetMyFactionsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetMyFactionsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetMyFactionsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
