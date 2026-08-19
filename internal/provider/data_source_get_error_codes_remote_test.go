package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetErrorCodesDataSource_Read_Happy exercises GetErrorCodesDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetErrorCodesDataSource_Read_Happy(t *testing.T) {
	r := &GetErrorCodesDataSource{client: newMockClientStatus(t, 200, "{\"errorCodes\":[]}")}
	m := GetErrorCodesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetErrorCodesDataSource_Read_NilClient exercises GetErrorCodesDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetErrorCodesDataSource_Read_NilClient(t *testing.T) {
	r := &GetErrorCodesDataSource{}
	m := GetErrorCodesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetErrorCodesDataSource_Read_BuildError exercises GetErrorCodesDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetErrorCodesDataSource_Read_BuildError(t *testing.T) {
	r := &GetErrorCodesDataSource{client: newMalformedBaseURLClient(t)}
	m := GetErrorCodesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetErrorCodesDataSource_Read_SendError exercises GetErrorCodesDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetErrorCodesDataSource_Read_SendError(t *testing.T) {
	r := &GetErrorCodesDataSource{client: newTransportErrorClient(t)}
	m := GetErrorCodesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetErrorCodesDataSource_Read_InvalidJSON exercises GetErrorCodesDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetErrorCodesDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetErrorCodesDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetErrorCodesDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
