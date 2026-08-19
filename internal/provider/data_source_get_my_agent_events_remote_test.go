package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetMyAgentEventsDataSource_Read_Happy exercises GetMyAgentEventsDataSource.readListRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetMyAgentEventsDataSource_Read_Happy(t *testing.T) {
	r := &GetMyAgentEventsDataSource{client: newMockClientStatus(t, 200, "{\"data\":[]}")}
	m := GetMyAgentEventsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestGetMyAgentEventsDataSource_Read_NilClient exercises GetMyAgentEventsDataSource.readListRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetMyAgentEventsDataSource_Read_NilClient(t *testing.T) {
	r := &GetMyAgentEventsDataSource{}
	m := GetMyAgentEventsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestGetMyAgentEventsDataSource_Read_BuildError exercises GetMyAgentEventsDataSource.readListRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetMyAgentEventsDataSource_Read_BuildError(t *testing.T) {
	r := &GetMyAgentEventsDataSource{client: newMalformedBaseURLClient(t)}
	m := GetMyAgentEventsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetMyAgentEventsDataSource_Read_SendError exercises GetMyAgentEventsDataSource.readListRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetMyAgentEventsDataSource_Read_SendError(t *testing.T) {
	r := &GetMyAgentEventsDataSource{client: newTransportErrorClient(t)}
	m := GetMyAgentEventsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read list response")
}

// TestGetMyAgentEventsDataSource_Read_InvalidJSON exercises GetMyAgentEventsDataSource.readListRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetMyAgentEventsDataSource_Read_InvalidJSON(t *testing.T) {
	r := &GetMyAgentEventsDataSource{client: newMockClientStatus(t, 200, "{{")}
	m := GetMyAgentEventsDataSourceModel{}
	resp := &datasource.ReadResponse{}
	r.readListRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode list page")
}
