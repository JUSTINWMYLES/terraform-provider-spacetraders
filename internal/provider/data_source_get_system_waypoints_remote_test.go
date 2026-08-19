package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSystemWaypointsDataSource_Read_Happy exercises GetSystemWaypointsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetSystemWaypointsDataSource_Read_Happy(t *testing.T) {
	r := &GetSystemWaypointsDataSource{client: newMockClientStatus(t, 200, "{\"data\":[]}")}
	m := GetSystemWaypointsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetSystemWaypointsDataSource_Read_NilClient exercises GetSystemWaypointsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetSystemWaypointsDataSource_Read_NilClient(t *testing.T) {
	r := &GetSystemWaypointsDataSource{}
	m := GetSystemWaypointsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetSystemWaypointsDataSource_Read_BuildError exercises GetSystemWaypointsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetSystemWaypointsDataSource_Read_BuildError(t *testing.T) {
	r := &GetSystemWaypointsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetSystemWaypointsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetSystemWaypointsDataSource_Read_SendError exercises GetSystemWaypointsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetSystemWaypointsDataSource_Read_SendError(t *testing.T) {
	r := &GetSystemWaypointsDataSource{client: newTransportErrorClient(t)}
	m := GetSystemWaypointsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetSystemWaypointsDataSource_Read_InvalidJSON exercises GetSystemWaypointsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetSystemWaypointsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetSystemWaypointsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetSystemWaypointsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
