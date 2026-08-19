package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetShipModulesDataSource_Read_Happy exercises GetShipModulesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetShipModulesDataSource_Read_Happy(t *testing.T) {
	r := &GetShipModulesDataSource{client: newMockClientStatus(t, 200, "{\"data\":[]}")}
	m := GetShipModulesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetShipModulesDataSource_Read_NilClient exercises GetShipModulesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetShipModulesDataSource_Read_NilClient(t *testing.T) {
	r := &GetShipModulesDataSource{}
	m := GetShipModulesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetShipModulesDataSource_Read_BuildError exercises GetShipModulesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetShipModulesDataSource_Read_BuildError(t *testing.T) {
	r := &GetShipModulesDataSource{client: newMalformedBaseURLClient(t)}
	m := GetShipModulesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetShipModulesDataSource_Read_SendError exercises GetShipModulesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetShipModulesDataSource_Read_SendError(t *testing.T) {
	r := &GetShipModulesDataSource{client: newTransportErrorClient(t)}
	m := GetShipModulesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetShipModulesDataSource_Read_InvalidJSON exercises GetShipModulesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetShipModulesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetShipModulesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetShipModulesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
