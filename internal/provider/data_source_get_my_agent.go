package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
import (
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
	client "github.com/spacetraders/terraform-provider-spacetraders/internal/client"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*GetMyAgentDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetMyAgentDataSource)(nil)
)

// GetMyAgentDataSource is the generated Terraform data source implementation.
type GetMyAgentDataSource struct {
	client *client.Client
}

// GetMyAgentDataSourceModel describes the data source state shape.
type GetMyAgentDataSourceModel struct {
	AccountId       types.String `tfsdk:"account_id" json:"accountId"`
	Credits         types.Int64  `tfsdk:"credits"`
	Headquarters    types.String `tfsdk:"headquarters"`
	ShipCount       types.Int64  `tfsdk:"ship_count" json:"shipCount"`
	StartingFaction types.String `tfsdk:"starting_faction" json:"startingFaction"`
	Symbol          types.String `tfsdk:"symbol"`
}

// NewGetMyAgentDataSource returns a new instance of the generated data source.
func NewGetMyAgentDataSource() datasource.DataSource {
	return &GetMyAgentDataSource{}
}

// Metadata returns the data source type name.
func (d *GetMyAgentDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "spacetraders_get_my_agent"
}

// Schema returns the data source schema.
func (d *GetMyAgentDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Fetch your agent's details.", Attributes: map[string]schema.Attribute{"account_id": schema.StringAttribute{MarkdownDescription: "Account ID that is tied to this agent. Only included on your own agent.", Computed: true}, "credits": schema.Int64Attribute{MarkdownDescription: "The number of credits the agent has available. Credits can be negative if funds have been overdrawn.", Computed: true}, "headquarters": schema.StringAttribute{MarkdownDescription: "The headquarters of the agent.", Computed: true}, "ship_count": schema.Int64Attribute{MarkdownDescription: "How many ships are owned by the agent.", Computed: true}, "starting_faction": schema.StringAttribute{MarkdownDescription: "The faction the agent started with.", Computed: true}, "symbol": schema.StringAttribute{MarkdownDescription: "Symbol of the agent.", Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetMyAgentDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetMyAgentDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	d.readRemote(ctx, &config, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}

// readRemote performs the read HTTP exchange and decodes the response into config. Extracted from Read so the request/response logic is unit-testable without a tfsdk.Config.
func (d *GetMyAgentDataSource) readRemote(ctx context.Context, config *GetMyAgentDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/my/agent"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil, client.WithSchemes("AgentToken"))
	if err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_get_my_agent", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_get_my_agent", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading spacetraders_get_my_agent", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		apiErr, err := client.NewAPIError(httpResp)
		if err != nil {
			resp.Diagnostics.AddError("Error reading spacetraders_get_my_agent", fmt.Sprintf("Could not read error response: %s", err))
			return
		}
		resp.Diagnostics.AddError("Error reading spacetraders_get_my_agent", apiErr.Error())
		return
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading spacetraders_get_my_agent", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["data"]; ok {
		if m, ok := v.(map[string]any); ok {
			data = m
		}
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_get_my_agent", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetMyAgentDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Data Source Configure Type", fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData))
		return
	}
	d.client = c
}
