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
var _ action.Action = (*DeliverContractAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*DeliverContractAction)(nil)

// DeliverContractAction is the generated Terraform action implementation.
type DeliverContractAction struct {
	client *client.Client
}

// DeliverContractActionModel describes the action configuration shape.
type DeliverContractActionModel struct {
	ContractId  types.String `tfsdk:"contract_id"`
	ShipSymbol  types.String `tfsdk:"ship_symbol" json:"shipSymbol"`
	TradeSymbol types.String `tfsdk:"trade_symbol" json:"tradeSymbol"`
	Units       types.Int64  `tfsdk:"units"`
}

// NewDeliverContractAction returns a new instance of the generated action.
func NewDeliverContractAction() action.Action {
	return &DeliverContractAction{}
}

// Metadata returns the action type name.
func (r *DeliverContractAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "spacetraders_deliver_contract"
}

// Schema returns the action schema.
func (r *DeliverContractAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Deliver cargo to a contract.\n\nIn order to use this API, a ship must be at the delivery location (denoted in the delivery terms as `destinationSymbol` of a contract) and must have a number of units of a good required by this contract in its cargo.\n\nCargo that was delivered will be removed from the ship's cargo.", Attributes: map[string]schema.Attribute{"contract_id": schema.StringAttribute{MarkdownDescription: "The ID of the contract.", Required: true}, "ship_symbol": schema.StringAttribute{MarkdownDescription: "Symbol of a ship located in the destination to deliver a contract and that has a good to deliver in its cargo.", Required: true}, "trade_symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the good to deliver.", Required: true}, "units": schema.Int64Attribute{MarkdownDescription: "Amount of units to deliver.", Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *DeliverContractAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config DeliverContractActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *DeliverContractAction) invokeRemote(ctx context.Context, config *DeliverContractActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/my/contracts/{contractId}/deliver"
	reqPath = strings.ReplaceAll(reqPath, "{contractId}", url.PathEscape(config.ContractId.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking spacetraders_deliver_contract", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking spacetraders_deliver_contract", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload), client.WithSchemes("AgentToken"))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking spacetraders_deliver_contract", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking spacetraders_deliver_contract", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		apiErr, err := client.NewAPIError(httpResp)
		if err != nil {
			resp.Diagnostics.AddError("Error invoking spacetraders_deliver_contract", fmt.Sprintf("Could not read error response: %s", err))
			return
		}
		resp.Diagnostics.AddError("Error invoking spacetraders_deliver_contract", apiErr.Error())
		return
	}
}

// Configure stores the API client supplied by the provider.
func (r *DeliverContractAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
