package provider

import (
	"bytes"
	"context"
	"encoding/json"
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
var _ action.Action = (*SupplyConstructionAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*SupplyConstructionAction)(nil)

// SupplyConstructionAction is the generated Terraform action implementation.
type SupplyConstructionAction struct {
	client *client.Client
}

// SupplyConstructionActionModel describes the action configuration shape.
type SupplyConstructionActionModel struct {
	ShipSymbol     types.String `tfsdk:"ship_symbol" json:"shipSymbol"`
	SystemSymbol   types.String `tfsdk:"system_symbol"`
	TradeSymbol    types.String `tfsdk:"trade_symbol" json:"tradeSymbol"`
	Units          types.Int64  `tfsdk:"units"`
	WaypointSymbol types.String `tfsdk:"waypoint_symbol"`
}

// NewSupplyConstructionAction returns a new instance of the generated action.
func NewSupplyConstructionAction() action.Action {
	return &SupplyConstructionAction{}
}

// Metadata returns the action type name.
func (r *SupplyConstructionAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "spacetraders_supply_construction"
}

// Schema returns the action schema.
func (r *SupplyConstructionAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Supply a construction site with the specified good. Requires a waypoint with a property of `isUnderConstruction` to be true.\n\nThe good must be in your ship's cargo. The good will be removed from your ship's cargo and added to the construction site's materials.", Attributes: map[string]schema.Attribute{"ship_symbol": schema.StringAttribute{Required: true}, "system_symbol": schema.StringAttribute{Required: true}, "trade_symbol": schema.StringAttribute{Required: true}, "units": schema.Int64Attribute{Required: true}, "waypoint_symbol": schema.StringAttribute{Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *SupplyConstructionAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config SupplyConstructionActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *SupplyConstructionAction) invokeRemote(ctx context.Context, config *SupplyConstructionActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/systems/{systemSymbol}/waypoints/{waypointSymbol}/construction/supply"
	reqPath = strings.ReplaceAll(reqPath, "{systemSymbol}", url.PathEscape(config.SystemSymbol.ValueString()))
	reqPath = strings.ReplaceAll(reqPath, "{waypointSymbol}", url.PathEscape(config.WaypointSymbol.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking spacetraders_supply_construction", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking spacetraders_supply_construction", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload), client.WithSchemes("AgentToken"))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking spacetraders_supply_construction", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking spacetraders_supply_construction", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		apiErr, err := client.NewAPIError(httpResp)
		if err != nil {
			resp.Diagnostics.AddError("Error invoking spacetraders_supply_construction", fmt.Sprintf("Could not read error response: %s", err))
			return
		}
		resp.Diagnostics.AddError("Error invoking spacetraders_supply_construction", apiErr.Error())
		return
	}
}

// Configure stores the API client supplied by the provider.
func (r *SupplyConstructionAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
