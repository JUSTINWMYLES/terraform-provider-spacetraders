package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestInstallShipModuleAction_Invoke_Happy exercises InstallShipModuleAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestInstallShipModuleAction_Invoke_Happy(t *testing.T) {
	r := &InstallShipModuleAction{client: newMockClientStatus(t, 201, "{}")}
	m := InstallShipModuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestInstallShipModuleAction_Invoke_NilClient exercises InstallShipModuleAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestInstallShipModuleAction_Invoke_NilClient(t *testing.T) {
	r := &InstallShipModuleAction{}
	m := InstallShipModuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestInstallShipModuleAction_Invoke_BuildError exercises InstallShipModuleAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestInstallShipModuleAction_Invoke_BuildError(t *testing.T) {
	r := &InstallShipModuleAction{client: newMalformedBaseURLClient(t)}
	m := InstallShipModuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestInstallShipModuleAction_Invoke_SendError exercises InstallShipModuleAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestInstallShipModuleAction_Invoke_SendError(t *testing.T) {
	r := &InstallShipModuleAction{client: newTransportErrorClient(t)}
	m := InstallShipModuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestInstallShipModuleAction_Invoke_APIError exercises InstallShipModuleAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestInstallShipModuleAction_Invoke_APIError(t *testing.T) {
	r := &InstallShipModuleAction{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := InstallShipModuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking spacetraders_install_ship_module")
}

// TestInstallShipModuleAction_Invoke_APIErrorReadBody exercises InstallShipModuleAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestInstallShipModuleAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &InstallShipModuleAction{client: newMockClientReadErrorBody(t, 500)}
	m := InstallShipModuleActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
