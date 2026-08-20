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
var _ action.Action = (*PurchaseCargoAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*PurchaseCargoAction)(nil)

// PurchaseCargoAction is the generated Terraform action implementation.
type PurchaseCargoAction struct {
	client *client.Client
}

// PurchaseCargoActionModel describes the action configuration shape.
type PurchaseCargoActionModel struct {
	ShipSymbol types.String `tfsdk:"ship_symbol"`
	Symbol     types.String `tfsdk:"symbol"`
	Units      types.Int64  `tfsdk:"units"`
}

// NewPurchaseCargoAction returns a new instance of the generated action.
func NewPurchaseCargoAction() action.Action {
	return &PurchaseCargoAction{}
}

// Metadata returns the action type name.
func (r *PurchaseCargoAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "spacetraders_purchase_cargo"
}

// Schema returns the action schema.
func (r *PurchaseCargoAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Purchase cargo from a market.\n\nThe ship must be docked in a waypoint that has `Marketplace` trait, and the market must be selling a good to be able to purchase it.\n\nThe maximum amount of units of a good that can be purchased in each transaction are denoted by the `tradeVolume` value of the good, which can be viewed by using the Get Market action.\n\nPurchased goods are added to the ship's cargo hold.", Attributes: map[string]schema.Attribute{"ship_symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the ship.", Required: true}, "symbol": schema.StringAttribute{MarkdownDescription: "The good's symbol.", Required: true}, "units": schema.Int64Attribute{MarkdownDescription: "The number of units of the good to purchase.", Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *PurchaseCargoAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config PurchaseCargoActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *PurchaseCargoAction) invokeRemote(ctx context.Context, config *PurchaseCargoActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/my/ships/{shipSymbol}/purchase"
	reqPath = strings.ReplaceAll(reqPath, "{shipSymbol}", url.PathEscape(config.ShipSymbol.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking spacetraders_purchase_cargo", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking spacetraders_purchase_cargo", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload), client.WithSchemes("AgentToken"))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking spacetraders_purchase_cargo", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking spacetraders_purchase_cargo", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		apiErr, err := client.NewAPIError(httpResp)
		if err != nil {
			resp.Diagnostics.AddError("Error invoking spacetraders_purchase_cargo", fmt.Sprintf("Could not read error response: %s", err))
			return
		}
		resp.Diagnostics.AddError("Error invoking spacetraders_purchase_cargo", apiErr.Error())
		return
	}
}

// Configure stores the API client supplied by the provider.
func (r *PurchaseCargoAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
