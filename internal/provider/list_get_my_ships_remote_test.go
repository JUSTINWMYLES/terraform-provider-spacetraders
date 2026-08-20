package provider

import (
	"context"
	"testing"
)

// TestGetMyShipsListResource_List_Happy exercises GetMyShipsListResource.listRemote against an httptest mock: happy path returns the success status with a JSON array body and no errors.
func TestGetMyShipsListResource_List_Happy(t *testing.T) {
	r := &GetMyShipsListResource{client: newMockClientStatus(t, 200, "{\"data\":[]}")}
	m := GetMyShipsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	requireNoErrors(t, diags)
}

// TestGetMyShipsListResource_List_NilClient exercises GetMyShipsListResource.listRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestGetMyShipsListResource_List_NilClient(t *testing.T) {
	r := &GetMyShipsListResource{}
	m := GetMyShipsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Client Not Configured")
}

// TestGetMyShipsListResource_List_BuildError exercises GetMyShipsListResource.listRemote against an httptest mock: malformed base URL surfaces Could not read list response.
func TestGetMyShipsListResource_List_BuildError(t *testing.T) {
	r := &GetMyShipsListResource{client: newMalformedBaseURLClient(t)}
	m := GetMyShipsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetMyShipsListResource_List_SendError exercises GetMyShipsListResource.listRemote against an httptest mock: transport error surfaces Could not read list response.
func TestGetMyShipsListResource_List_SendError(t *testing.T) {
	r := &GetMyShipsListResource{client: newTransportErrorClient(t)}
	m := GetMyShipsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not read list response")
}

// TestGetMyShipsListResource_List_InvalidJSON exercises GetMyShipsListResource.listRemote against an httptest mock: success status with a non-array body surfaces Could not decode list page.
func TestGetMyShipsListResource_List_InvalidJSON(t *testing.T) {
	r := &GetMyShipsListResource{client: newMockClientStatus(t, 200, "{{")}
	m := GetMyShipsListResourceModel{}
	_, diags := r.listRemote(context.Background(), &m)
	hasErrorContaining(t, diags, "Could not decode list page")
}
