package provider

import (
	"context"
	"testing"
)

// TestGetSystemWaypointsListResource_List_Happy exercises GetSystemWaypointsListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetSystemWaypointsListResource_List_Happy(t *testing.T) {
	r := &GetSystemWaypointsListResource{client: newMockClientStatus(t, 200, "{\"data\":[]}")}
	m := GetSystemWaypointsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestGetSystemWaypointsListResource_List_NilClient exercises GetSystemWaypointsListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetSystemWaypointsListResource_List_NilClient(t *testing.T) {
	r := &GetSystemWaypointsListResource{}
	m := GetSystemWaypointsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestGetSystemWaypointsListResource_List_BuildError exercises GetSystemWaypointsListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetSystemWaypointsListResource_List_BuildError(t *testing.T) {
	r := &GetSystemWaypointsListResource{client: newMalformedBaseURLClient(t)}
	m := GetSystemWaypointsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetSystemWaypointsListResource_List_SendError exercises GetSystemWaypointsListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetSystemWaypointsListResource_List_SendError(t *testing.T) {
	r := &GetSystemWaypointsListResource{client: newTransportErrorClient(t)}
	m := GetSystemWaypointsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetSystemWaypointsListResource_List_InvalidJSON exercises GetSystemWaypointsListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetSystemWaypointsListResource_List_InvalidJSON(t *testing.T) {
	r := &GetSystemWaypointsListResource{client: newMockClientStatus(t, 200, "{{")}
	m := GetSystemWaypointsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
