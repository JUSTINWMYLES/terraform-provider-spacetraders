package provider

import (
	"context"
	"testing"
)

// TestGetAgentsListResource_List_Happy exercises GetAgentsListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetAgentsListResource_List_Happy(t *testing.T) {
	r := &GetAgentsListResource{client: newMockClientStatus(t, 200, "{\"data\":[]}")}
	m := GetAgentsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestGetAgentsListResource_List_NilClient exercises GetAgentsListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetAgentsListResource_List_NilClient(t *testing.T) {
	r := &GetAgentsListResource{}
	m := GetAgentsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestGetAgentsListResource_List_BuildError exercises GetAgentsListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetAgentsListResource_List_BuildError(t *testing.T) {
	r := &GetAgentsListResource{client: newMalformedBaseURLClient(t)}
	m := GetAgentsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetAgentsListResource_List_SendError exercises GetAgentsListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetAgentsListResource_List_SendError(t *testing.T) {
	r := &GetAgentsListResource{client: newTransportErrorClient(t)}
	m := GetAgentsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetAgentsListResource_List_InvalidJSON exercises GetAgentsListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetAgentsListResource_List_InvalidJSON(t *testing.T) {
	r := &GetAgentsListResource{client: newMockClientStatus(t, 200, "{{")}
	m := GetAgentsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
