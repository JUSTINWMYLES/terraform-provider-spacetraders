package provider

import (
	"context"
	"testing"
)

// TestGetContractsListResource_List_Happy exercises GetContractsListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetContractsListResource_List_Happy(t *testing.T) {
	r := &GetContractsListResource{client: newMockClientStatus(t, 200, "{\"data\":[]}")}
	m := GetContractsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestGetContractsListResource_List_NilClient exercises GetContractsListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetContractsListResource_List_NilClient(t *testing.T) {
	r := &GetContractsListResource{}
	m := GetContractsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestGetContractsListResource_List_BuildError exercises GetContractsListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetContractsListResource_List_BuildError(t *testing.T) {
	r := &GetContractsListResource{client: newMalformedBaseURLClient(t)}
	m := GetContractsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetContractsListResource_List_SendError exercises GetContractsListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetContractsListResource_List_SendError(t *testing.T) {
	r := &GetContractsListResource{client: newTransportErrorClient(t)}
	m := GetContractsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetContractsListResource_List_InvalidJSON exercises GetContractsListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetContractsListResource_List_InvalidJSON(t *testing.T) {
	r := &GetContractsListResource{client: newMockClientStatus(t, 200, "{{")}
	m := GetContractsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
