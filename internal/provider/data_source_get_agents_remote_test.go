package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAgentsDataSource_Read_Happy exercises GetAgentsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAgentsDataSource_Read_Happy(t *testing.T) {
	r := &GetAgentsDataSource{client: newMockClientStatus(t, 200, "{\"data\":[]}")}
	m := GetAgentsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetAgentsDataSource_Read_NilClient exercises GetAgentsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAgentsDataSource_Read_NilClient(t *testing.T) {
	r := &GetAgentsDataSource{}
	m := GetAgentsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetAgentsDataSource_Read_BuildError exercises GetAgentsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAgentsDataSource_Read_BuildError(t *testing.T) {
	r := &GetAgentsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetAgentsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAgentsDataSource_Read_SendError exercises GetAgentsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAgentsDataSource_Read_SendError(t *testing.T) {
	r := &GetAgentsDataSource{client: newTransportErrorClient(t)}
	m := GetAgentsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetAgentsDataSource_Read_InvalidJSON exercises GetAgentsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAgentsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetAgentsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAgentsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
