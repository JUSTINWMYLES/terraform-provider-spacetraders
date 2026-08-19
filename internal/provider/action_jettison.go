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
var _ action.Action = (*JettisonAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*JettisonAction)(nil)

// JettisonAction is the generated Terraform action implementation.
type JettisonAction struct {
	client *client.Client
}

// JettisonActionModel describes the action configuration shape.
type JettisonActionModel struct {
	ShipSymbol types.String `tfsdk:"ship_symbol"`
	Symbol     types.String `tfsdk:"symbol"`
	Units      types.Int64  `tfsdk:"units"`
}

// NewJettisonAction returns a new instance of the generated action.
func NewJettisonAction() action.Action {
	return &JettisonAction{}
}

// Metadata returns the action type name.
func (r *JettisonAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "spacetraders_jettison"
}

// Schema returns the action schema.
func (r *JettisonAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Jettison cargo from your ship's cargo hold.", Attributes: map[string]schema.Attribute{"ship_symbol": schema.StringAttribute{Required: true}, "symbol": schema.StringAttribute{Required: true}, "units": schema.Int64Attribute{Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *JettisonAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config JettisonActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *JettisonAction) invokeRemote(ctx context.Context, config *JettisonActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/my/ships/{shipSymbol}/jettison"
	reqPath = strings.ReplaceAll(reqPath, "{shipSymbol}", url.PathEscape(config.ShipSymbol.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking spacetraders_jettison", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking spacetraders_jettison", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload), client.WithSchemes("AgentToken"))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking spacetraders_jettison", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking spacetraders_jettison", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		apiErr, err := client.NewAPIError(httpResp)
		if err != nil {
			resp.Diagnostics.AddError("Error invoking spacetraders_jettison", fmt.Sprintf("Could not read error response: %s", err))
			return
		}
		resp.Diagnostics.AddError("Error invoking spacetraders_jettison", apiErr.Error())
		return
	}
}

// Configure stores the API client supplied by the provider.
func (r *JettisonAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
