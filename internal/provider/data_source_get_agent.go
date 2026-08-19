package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)
import (
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
	client "github.com/spacetraders/terraform-provider-spacetraders/internal/client"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*GetAgentDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAgentDataSource)(nil)
)

// GetAgentDataSource is the generated Terraform data source implementation.
type GetAgentDataSource struct {
	client *client.Client
}

// GetAgentDataSourceModel describes the data source state shape.
type GetAgentDataSourceModel struct {
	AgentSymbol     types.String `tfsdk:"agent_symbol" json:"agentSymbol"`
	Credits         types.Int64  `tfsdk:"credits"`
	Headquarters    types.String `tfsdk:"headquarters"`
	ShipCount       types.Int64  `tfsdk:"ship_count" json:"shipCount"`
	StartingFaction types.String `tfsdk:"starting_faction" json:"startingFaction"`
	Symbol          types.String `tfsdk:"symbol"`
}

// NewGetAgentDataSource returns a new instance of the generated data source.
func NewGetAgentDataSource() datasource.DataSource {
	return &GetAgentDataSource{}
}

// Metadata returns the data source type name.
func (d *GetAgentDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "spacetraders_get_agent"
}

// Schema returns the data source schema.
func (d *GetAgentDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get public details for a specific agent.", Attributes: map[string]schema.Attribute{"agent_symbol": schema.StringAttribute{Required: true}, "credits": schema.Int64Attribute{Computed: true}, "headquarters": schema.StringAttribute{Computed: true}, "ship_count": schema.Int64Attribute{Computed: true}, "starting_faction": schema.StringAttribute{Computed: true}, "symbol": schema.StringAttribute{Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetAgentDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAgentDataSourceModel
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
func (d *GetAgentDataSource) readRemote(ctx context.Context, config *GetAgentDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/agents/{agentSymbol}"
	reqPath = strings.ReplaceAll(reqPath, "{agentSymbol}", url.PathEscape(config.AgentSymbol.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_get_agent", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_get_agent", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading spacetraders_get_agent", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		apiErr, err := client.NewAPIError(httpResp)
		if err != nil {
			resp.Diagnostics.AddError("Error reading spacetraders_get_agent", fmt.Sprintf("Could not read error response: %s", err))
			return
		}
		resp.Diagnostics.AddError("Error reading spacetraders_get_agent", apiErr.Error())
		return
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading spacetraders_get_agent", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["data"]; ok {
		if m, ok := v.(map[string]any); ok {
			data = m
		}
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_get_agent", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetAgentDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
