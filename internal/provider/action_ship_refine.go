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
var _ action.Action = (*ShipRefineAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*ShipRefineAction)(nil)

// ShipRefineAction is the generated Terraform action implementation.
type ShipRefineAction struct {
	client *client.Client
}

// ShipRefineActionModel describes the action configuration shape.
type ShipRefineActionModel struct {
	Produce    types.String `tfsdk:"produce"`
	ShipSymbol types.String `tfsdk:"ship_symbol"`
}

// NewShipRefineAction returns a new instance of the generated action.
func NewShipRefineAction() action.Action {
	return &ShipRefineAction{}
}

// Metadata returns the action type name.
func (r *ShipRefineAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "spacetraders_ship_refine"
}

// Schema returns the action schema.
func (r *ShipRefineAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Attempt to refine the raw materials on your ship. The request will only succeed if your ship is capable of refining at the time of the request. In order to be able to refine, a ship must have goods that can be refined and have installed a `Refinery` module that can refine it.\n\nWhen refining, 100 basic goods will be converted into 10 processed goods.", Attributes: map[string]schema.Attribute{"produce": schema.StringAttribute{Required: true}, "ship_symbol": schema.StringAttribute{Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *ShipRefineAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config ShipRefineActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *ShipRefineAction) invokeRemote(ctx context.Context, config *ShipRefineActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/my/ships/{shipSymbol}/refine"
	reqPath = strings.ReplaceAll(reqPath, "{shipSymbol}", url.PathEscape(config.ShipSymbol.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking spacetraders_ship_refine", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking spacetraders_ship_refine", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload), client.WithSchemes("AgentToken"))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking spacetraders_ship_refine", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking spacetraders_ship_refine", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		apiErr, err := client.NewAPIError(httpResp)
		if err != nil {
			resp.Diagnostics.AddError("Error invoking spacetraders_ship_refine", fmt.Sprintf("Could not read error response: %s", err))
			return
		}
		resp.Diagnostics.AddError("Error invoking spacetraders_ship_refine", apiErr.Error())
		return
	}
}

// Configure stores the API client supplied by the provider.
func (r *ShipRefineAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
