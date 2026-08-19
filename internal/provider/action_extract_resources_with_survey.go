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
var _ action.Action = (*ExtractResourcesWithSurveyAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*ExtractResourcesWithSurveyAction)(nil)

// ExtractResourcesWithSurveyAction is the generated Terraform action implementation.
type ExtractResourcesWithSurveyAction struct {
	client *client.Client
}

// ExtractResourcesWithSurveyActionModel describes the action configuration shape.
type ExtractResourcesWithSurveyActionModel struct {
	Deposits   types.Dynamic `tfsdk:"deposits"`
	Expiration types.String  `tfsdk:"expiration"`
	ShipSymbol types.String  `tfsdk:"ship_symbol"`
	Signature  types.String  `tfsdk:"signature"`
	Size       types.String  `tfsdk:"size"`
	Symbol     types.String  `tfsdk:"symbol"`
}

// NewExtractResourcesWithSurveyAction returns a new instance of the generated action.
func NewExtractResourcesWithSurveyAction() action.Action {
	return &ExtractResourcesWithSurveyAction{}
}

// Metadata returns the action type name.
func (r *ExtractResourcesWithSurveyAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "spacetraders_extract_resources_with_survey"
}

// Schema returns the action schema.
func (r *ExtractResourcesWithSurveyAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Use a survey when extracting resources from a waypoint. This endpoint requires a survey as the payload, which allows your ship to extract specific yields.\n\nSend the full survey object as the payload which will be validated according to the signature. If the signature is invalid, or any properties of the survey are changed, the request will fail.", Attributes: map[string]schema.Attribute{"deposits": schema.DynamicAttribute{Required: true}, "expiration": schema.StringAttribute{Required: true}, "ship_symbol": schema.StringAttribute{Required: true}, "signature": schema.StringAttribute{Required: true}, "size": schema.StringAttribute{Required: true}, "symbol": schema.StringAttribute{Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *ExtractResourcesWithSurveyAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config ExtractResourcesWithSurveyActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *ExtractResourcesWithSurveyAction) invokeRemote(ctx context.Context, config *ExtractResourcesWithSurveyActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/my/ships/{shipSymbol}/extract/survey"
	reqPath = strings.ReplaceAll(reqPath, "{shipSymbol}", url.PathEscape(config.ShipSymbol.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking spacetraders_extract_resources_with_survey", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking spacetraders_extract_resources_with_survey", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload), client.WithSchemes("AgentToken"))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking spacetraders_extract_resources_with_survey", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking spacetraders_extract_resources_with_survey", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		apiErr, err := client.NewAPIError(httpResp)
		if err != nil {
			resp.Diagnostics.AddError("Error invoking spacetraders_extract_resources_with_survey", fmt.Sprintf("Could not read error response: %s", err))
			return
		}
		resp.Diagnostics.AddError("Error invoking spacetraders_extract_resources_with_survey", apiErr.Error())
		return
	}
}

// Configure stores the API client supplied by the provider.
func (r *ExtractResourcesWithSurveyAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
