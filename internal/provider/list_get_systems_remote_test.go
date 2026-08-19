package provider

import (
	"context"
	"testing"
)

// TestGetSystemsListResource_List_Happy exercises GetSystemsListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetSystemsListResource_List_Happy(t *testing.T) {
	r := &GetSystemsListResource{client: newMockClientStatus(t, 200, "{\"data\":[]}")}
	m := GetSystemsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestGetSystemsListResource_List_NilClient exercises GetSystemsListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetSystemsListResource_List_NilClient(t *testing.T) {
	r := &GetSystemsListResource{}
	m := GetSystemsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestGetSystemsListResource_List_BuildError exercises GetSystemsListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetSystemsListResource_List_BuildError(t *testing.T) {
	r := &GetSystemsListResource{client: newMalformedBaseURLClient(t)}
	m := GetSystemsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetSystemsListResource_List_SendError exercises GetSystemsListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetSystemsListResource_List_SendError(t *testing.T) {
	r := &GetSystemsListResource{client: newTransportErrorClient(t)}
	m := GetSystemsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetSystemsListResource_List_InvalidJSON exercises GetSystemsListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetSystemsListResource_List_InvalidJSON(t *testing.T) {
	r := &GetSystemsListResource{client: newMockClientStatus(t, 200, "{{")}
	m := GetSystemsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
