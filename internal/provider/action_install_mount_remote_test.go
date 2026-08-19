package provider

import (
	"context"
	"testing"
)
import "github.com/hashicorp/terraform-plugin-framework/action"

// TestInstallMountAction_Invoke_Happy exercises InstallMountAction.invokeRemote against an httptest mock: happy path returns the success status with no errors; the response body is not decoded.
func TestInstallMountAction_Invoke_Happy(t *testing.T) {
	r := &InstallMountAction{client: newMockClientStatus(t, 201, "{}")}
	m := InstallMountActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	requireNoErrors(t, resp.Diagnostics)
}

// TestInstallMountAction_Invoke_NilClient exercises InstallMountAction.invokeRemote against an httptest mock: nil client surfaces the Client Not Configured diagnostic.
func TestInstallMountAction_Invoke_NilClient(t *testing.T) {
	r := &InstallMountAction{}
	m := InstallMountActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Client Not Configured")
}

// TestInstallMountAction_Invoke_BuildError exercises InstallMountAction.invokeRemote against an httptest mock: malformed base URL surfaces Could not build request.
func TestInstallMountAction_Invoke_BuildError(t *testing.T) {
	r := &InstallMountAction{client: newMalformedBaseURLClient(t)}
	m := InstallMountActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not build request")
}

// TestInstallMountAction_Invoke_SendError exercises InstallMountAction.invokeRemote against an httptest mock: transport error surfaces Could not send request.
func TestInstallMountAction_Invoke_SendError(t *testing.T) {
	r := &InstallMountAction{client: newTransportErrorClient(t)}
	m := InstallMountActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not send request")
}

// TestInstallMountAction_Invoke_APIError exercises InstallMountAction.invokeRemote against an httptest mock: non-success status surfaces the API error summary.
func TestInstallMountAction_Invoke_APIError(t *testing.T) {
	r := &InstallMountAction{client: newMockClientStatus(t, 500, "{\"message\":\"boom\"}")}
	m := InstallMountActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Error invoking spacetraders_install_mount")
}

// TestInstallMountAction_Invoke_APIErrorReadBody exercises InstallMountAction.invokeRemote against an httptest mock: non-success status whose error body cannot be read surfaces Could not read error response.
func TestInstallMountAction_Invoke_APIErrorReadBody(t *testing.T) {
	r := &InstallMountAction{client: newMockClientReadErrorBody(t, 500)}
	m := InstallMountActionModel{}
	resp := &action.InvokeResponse{}
	r.invokeRemote(context.Background(), &m, resp)
	hasErrorContaining(t, resp.Diagnostics, "Could not read error response")
}
