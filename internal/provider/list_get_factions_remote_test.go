package provider

import (
	"context"
	"testing"
)

// TestGetFactionsListResource_List_Happy exercises GetFactionsListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetFactionsListResource_List_Happy(t *testing.T) {
	r := &GetFactionsListResource{client: newMockClientStatus(t, 200, "{\"data\":[]}")}
	m := GetFactionsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestGetFactionsListResource_List_NilClient exercises GetFactionsListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetFactionsListResource_List_NilClient(t *testing.T) {
	r := &GetFactionsListResource{}
	m := GetFactionsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestGetFactionsListResource_List_BuildError exercises GetFactionsListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetFactionsListResource_List_BuildError(t *testing.T) {
	r := &GetFactionsListResource{client: newMalformedBaseURLClient(t)}
	m := GetFactionsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetFactionsListResource_List_SendError exercises GetFactionsListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetFactionsListResource_List_SendError(t *testing.T) {
	r := &GetFactionsListResource{client: newTransportErrorClient(t)}
	m := GetFactionsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetFactionsListResource_List_InvalidJSON exercises GetFactionsListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetFactionsListResource_List_InvalidJSON(t *testing.T) {
	r := &GetFactionsListResource{client: newMockClientStatus(t, 200, "{{")}
	m := GetFactionsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
