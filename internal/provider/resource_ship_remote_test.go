package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/resource"

// TestShipResource_Create_Happy exercises ShipResource.createRemote against an httptest mock: happy path returns the success status and an identifier in the body.
func TestShipResource_Create_Happy(t *testing.T) {
	r := &ShipResource{client: newMockClientStatus(t, 201, "{\"symbol\":\"example-id\"}")}
	m := ShipResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestShipResource_Create_NilClient exercises ShipResource.createRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestShipResource_Create_NilClient(t *testing.T) {
	r := &ShipResource{}
	m := ShipResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestShipResource_Create_BuildError exercises ShipResource.createRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestShipResource_Create_BuildError(t *testing.T) {
	r := &ShipResource{client: newMalformedBaseURLClient(t)}
	m := ShipResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestShipResource_Create_SendError exercises ShipResource.createRemote against an httptest mock: transport error surfaces Could not send request.
func TestShipResource_Create_SendError(t *testing.T) {
	r := &ShipResource{client: newTransportErrorClient(t)}
	m := ShipResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestShipResource_Create_APIError exercises ShipResource.createRemote against an httptest mock: non-success status surfaces the API error summary.
func TestShipResource_Create_APIError(t *testing.T) {
	r := &ShipResource{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := ShipResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error creating spacetraders_ship")
}

// TestShipResource_Create_APIErrorReadBody exercises ShipResource.createRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestShipResource_Create_APIErrorReadBody(t *testing.T) {
	r := &ShipResource{client: newMockClientReadErrorBody(t, 500)}
	m := ShipResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestShipResource_Create_InvalidJSON exercises ShipResource.createRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestShipResource_Create_InvalidJSON(t *testing.T) {
	r := &ShipResource{client: newMockClientStatus(t, 201, "{{")}
	m := ShipResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestShipResource_Create_MapError exercises ShipResource.createRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestShipResource_Create_MapError(t *testing.T) {
	r := &ShipResource{client: newMockClientStatus(t, 201, "{\"symbol\":12345}")}
	m := ShipResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestShipResource_Create_MissingID exercises ShipResource.createRemote against an httptest mock: success status with no identifier surfaces the missing-identifier diagnostic.
func TestShipResource_Create_MissingID(t *testing.T) {
	r := &ShipResource{client: newMockClientStatus(t, 201, "{}")}
	m := ShipResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "did not contain an identifier")
}

// TestShipResource_Create_LocationFallback exercises ShipResource.createRemote against an httptest mock: success status with no body id but a Location header sets the string identifier from the header.
func TestShipResource_Create_LocationFallback(t *testing.T) {
	r := &ShipResource{client: newMockClientWithLocation(t, 201, "http://example.test/folders/example-id", "{}")}
	m := ShipResourceModel{}
	resp := &resource.CreateResponse{}
	r.createRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
	if m.Symbol.ValueString() != "http://example.test/folders/example-id" {
		t.Fatalf("identifier = %q, want %q", m.Symbol.ValueString(), "http://example.test/folders/example-id")
	}
}

// TestShipResource_Read_Happy exercises ShipResource.readRemote against an httptest mock: happy path returns the success status and reports removed=false with no errors.
func TestShipResource_Read_Happy(t *testing.T) {
	r := &ShipResource{client: newMockClientStatus(t, 200, "{}")}
	m := ShipResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if removed {
		t.Fatalf("expected removed=false on happy path")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestShipResource_Read_NilClient exercises ShipResource.readRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestShipResource_Read_NilClient(t *testing.T) {
	r := &ShipResource{}
	m := ShipResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestShipResource_Read_BuildError exercises ShipResource.readRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestShipResource_Read_BuildError(t *testing.T) {
	r := &ShipResource{client: newMalformedBaseURLClient(t)}
	m := ShipResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestShipResource_Read_SendError exercises ShipResource.readRemote against an httptest mock: transport error surfaces Could not send request.
func TestShipResource_Read_SendError(t *testing.T) {
	r := &ShipResource{client: newTransportErrorClient(t)}
	m := ShipResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestShipResource_Read_NotFound exercises ShipResource.readRemote against an httptest mock: 404 reports removed=true with no error so the framework drops the resource from state.
func TestShipResource_Read_NotFound(t *testing.T) {
	r := &ShipResource{client: newMockClientStatus(t, 404, "")}
	m := ShipResourceModel{}
	resp := &resource.ReadResponse{}
	removed := r.readRemote(context.Background(), &m, resp)
	if !removed {
		t.Fatalf("expected removed=true on 404")
	}
	requireNoErrors(t, resp.Diagnostics)
}

// TestShipResource_Read_APIError exercises ShipResource.readRemote against an httptest mock: non-success status surfaces the API error summary.
func TestShipResource_Read_APIError(t *testing.T) {
	r := &ShipResource{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := ShipResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error reading spacetraders_ship")
}

// TestShipResource_Read_APIErrorReadBody exercises ShipResource.readRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestShipResource_Read_APIErrorReadBody(t *testing.T) {
	r := &ShipResource{client: newMockClientReadErrorBody(t, 500)}
	m := ShipResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}

// TestShipResource_Read_InvalidJSON exercises ShipResource.readRemote against an httptest mock: success status with a malformed body surfaces Could not decode response body.
func TestShipResource_Read_InvalidJSON(t *testing.T) {
	r := &ShipResource{client: newMockClientStatus(t, 200, "{{")}
	m := ShipResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not decode response body")
}

// TestShipResource_Read_MapError exercises ShipResource.readRemote against an httptest mock: success status with a wrong-typed identifier surfaces Could not map response to state.
func TestShipResource_Read_MapError(t *testing.T) {
	r := &ShipResource{client: newMockClientStatus(t, 200, "{\"symbol\":12345}")}
	m := ShipResourceModel{}
	resp := &resource.ReadResponse{}
	r.readRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not map response to state")
}

// TestShipResource_Delete_Happy exercises ShipResource.deleteRemote against an httptest mock: happy path returns the success status with no errors.
func TestShipResource_Delete_Happy(t *testing.T) {
	r := &ShipResource{client: newMockClientStatus(t, 200, "")}
	m := ShipResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestShipResource_Delete_NilClient exercises ShipResource.deleteRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestShipResource_Delete_NilClient(t *testing.T) {
	r := &ShipResource{}
	m := ShipResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestShipResource_Delete_BuildError exercises ShipResource.deleteRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestShipResource_Delete_BuildError(t *testing.T) {
	r := &ShipResource{client: newMalformedBaseURLClient(t)}
	m := ShipResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestShipResource_Delete_SendError exercises ShipResource.deleteRemote against an httptest mock: transport error surfaces Could not send request.
func TestShipResource_Delete_SendError(t *testing.T) {
	r := &ShipResource{client: newTransportErrorClient(t)}
	m := ShipResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestShipResource_Delete_NotFoundSuccess exercises ShipResource.deleteRemote against an httptest mock: 404 is treated as already deleted and surfaces no error.
func TestShipResource_Delete_NotFoundSuccess(t *testing.T) {
	r := &ShipResource{client: newMockClientStatus(t, 404, "")}
	m := ShipResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestShipResource_Delete_APIError exercises ShipResource.deleteRemote against an httptest mock: non-success status surfaces the API error summary.
func TestShipResource_Delete_APIError(t *testing.T) {
	r := &ShipResource{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := ShipResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error deleting spacetraders_ship")
}

// TestShipResource_Delete_APIErrorReadBody exercises ShipResource.deleteRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestShipResource_Delete_APIErrorReadBody(t *testing.T) {
	r := &ShipResource{client: newMockClientReadErrorBody(t, 500)}
	m := ShipResourceModel{}
	resp := &resource.DeleteResponse{}
	r.deleteRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
