package provider

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)
import (
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
	client "github.com/spacetraders/terraform-provider-spacetraders/internal/client"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*CreateShipSystemScanAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*CreateShipSystemScanAction)(nil)

// CreateShipSystemScanAction is the generated Terraform action implementation.
type CreateShipSystemScanAction struct {
	client *client.Client
}

// CreateShipSystemScanActionModel describes the action configuration shape.
type CreateShipSystemScanActionModel struct {
	ShipSymbol types.String `tfsdk:"ship_symbol"`
}

// NewCreateShipSystemScanAction returns a new instance of the generated action.
func NewCreateShipSystemScanAction() action.Action {
	return &CreateShipSystemScanAction{}
}

// Metadata returns the action type name.
func (r *CreateShipSystemScanAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "spacetraders_create_ship_system_scan"
}

// Schema returns the action schema.
func (r *CreateShipSystemScanAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Scan for nearby systems, retrieving information on the systems' distance from the ship and their waypoints. Requires a ship to have the `Sensor Array` mount installed to use.\n\nThe ship will enter a cooldown after using this function, during which it cannot execute certain actions.", Attributes: map[string]schema.Attribute{"ship_symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the ship.", Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *CreateShipSystemScanAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config CreateShipSystemScanActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *CreateShipSystemScanAction) invokeRemote(ctx context.Context, config *CreateShipSystemScanActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/my/ships/{shipSymbol}/scan/systems"
	reqPath = strings.ReplaceAll(reqPath, "{shipSymbol}", url.PathEscape(config.ShipSymbol.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, nil, client.WithSchemes("AgentToken"))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking spacetraders_create_ship_system_scan", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking spacetraders_create_ship_system_scan", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		apiErr, err := client.NewAPIError(httpResp)
		if err != nil {
			resp.Diagnostics.AddError("Error invoking spacetraders_create_ship_system_scan", fmt.Sprintf("Could not read error response: %s", err))
			return
		}
		resp.Diagnostics.AddError("Error invoking spacetraders_create_ship_system_scan", apiErr.Error())
		return
	}
}

// Configure stores the API client supplied by the provider.
func (r *CreateShipSystemScanAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Action Configure Type", fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData))
		return
	}
	r.client = c
}
