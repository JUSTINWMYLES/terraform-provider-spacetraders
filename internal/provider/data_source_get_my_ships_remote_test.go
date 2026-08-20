package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetMyShipsDataSource_Read_Happy exercises GetMyShipsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetMyShipsDataSource_Read_Happy(t *testing.T) {
	r := &GetMyShipsDataSource{client: newMockClientStatus(t, 200, "{\"data\":[]}")}
	m := GetMyShipsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetMyShipsDataSource_Read_NilClient exercises GetMyShipsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetMyShipsDataSource_Read_NilClient(t *testing.T) {
	r := &GetMyShipsDataSource{}
	m := GetMyShipsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetMyShipsDataSource_Read_BuildError exercises GetMyShipsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetMyShipsDataSource_Read_BuildError(t *testing.T) {
	r := &GetMyShipsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetMyShipsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetMyShipsDataSource_Read_SendError exercises GetMyShipsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetMyShipsDataSource_Read_SendError(t *testing.T) {
	r := &GetMyShipsDataSource{client: newTransportErrorClient(t)}
	m := GetMyShipsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetMyShipsDataSource_Read_InvalidJSON exercises GetMyShipsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetMyShipsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetMyShipsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetMyShipsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
