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
	_ datasource.DataSource              = (*GetStatusDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetStatusDataSource)(nil)
)

// GetStatusDataSource is the generated Terraform data source implementation.
type GetStatusDataSource struct {
	client *client.Client
}

// GetStatusDataSourceModel describes the data source state shape.
type GetStatusDataSourceModel struct {
	Announcements types.List   `tfsdk:"announcements"`
	Description   types.String `tfsdk:"description"`
	Health        types.Object `tfsdk:"health"`
	Leaderboards  types.Object `tfsdk:"leaderboards"`
	Links         types.List   `tfsdk:"links"`
	ResetDate     types.String `tfsdk:"reset_date" json:"resetDate"`
	ServerResets  types.Object `tfsdk:"server_resets" json:"serverResets"`
	Stats         types.Object `tfsdk:"stats"`
	Status        types.String `tfsdk:"status"`
	Version       types.String `tfsdk:"version"`
}

// NewGetStatusDataSource returns a new instance of the generated data source.
func NewGetStatusDataSource() datasource.DataSource {
	return &GetStatusDataSource{}
}

// Metadata returns the data source type name.
func (d *GetStatusDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "spacetraders_get_status"
}

// Schema returns the data source schema.
func (d *GetStatusDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Return the status of the game server.\nThis also includes a few global elements, such as announcements, server reset dates and leaderboards.", Attributes: map[string]schema.Attribute{"announcements": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"body": schema.StringAttribute{Computed: true}, "title": schema.StringAttribute{Computed: true}}}}, "description": schema.StringAttribute{Computed: true}, "health": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"last_market_update": schema.StringAttribute{MarkdownDescription: "The date/time when the market was last updated.", Computed: true}}}, "leaderboards": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"most_credits": schema.ListNestedAttribute{MarkdownDescription: "Top agents with the most credits.", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"agent_symbol": schema.StringAttribute{MarkdownDescription: "Symbol of the agent.", Computed: true}, "credits": schema.Int64Attribute{MarkdownDescription: "Amount of credits.", Computed: true}}}}, "most_submitted_charts": schema.ListNestedAttribute{MarkdownDescription: "Top agents with the most charted submitted.", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"agent_symbol": schema.StringAttribute{MarkdownDescription: "Symbol of the agent.", Computed: true}, "chart_count": schema.Int64Attribute{MarkdownDescription: "Amount of charts done by the agent.", Computed: true}}}}}}, "links": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"name": schema.StringAttribute{Computed: true}, "url": schema.StringAttribute{Computed: true}}}}, "reset_date": schema.StringAttribute{MarkdownDescription: "The date when the game server was last reset.", Computed: true}, "server_resets": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"frequency": schema.StringAttribute{MarkdownDescription: "How often we intend to reset the game server.", Computed: true}, "next": schema.StringAttribute{MarkdownDescription: "The date and time when the game server will reset.", Computed: true}}}, "stats": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"accounts": schema.Int64Attribute{MarkdownDescription: "Total number of accounts registered on the game server.", Computed: true}, "agents": schema.Int64Attribute{MarkdownDescription: "Number of registered agents in the game.", Computed: true}, "ships": schema.Int64Attribute{MarkdownDescription: "Total number of ships in the game.", Computed: true}, "systems": schema.Int64Attribute{MarkdownDescription: "Total number of systems in the game.", Computed: true}, "waypoints": schema.Int64Attribute{MarkdownDescription: "Total number of waypoints in the game.", Computed: true}}}, "status": schema.StringAttribute{MarkdownDescription: "The current status of the game server.", Computed: true}, "version": schema.StringAttribute{MarkdownDescription: "The current version of the API.", Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetStatusDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetStatusDataSourceModel
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
func (d *GetStatusDataSource) readRemote(ctx context.Context, config *GetStatusDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_get_status", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_get_status", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading spacetraders_get_status", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		apiErr, err := client.NewAPIError(httpResp)
		if err != nil {
			resp.Diagnostics.AddError("Error reading spacetraders_get_status", fmt.Sprintf("Could not read error response: %s", err))
			return
		}
		resp.Diagnostics.AddError("Error reading spacetraders_get_status", apiErr.Error())
		return
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading spacetraders_get_status", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_get_status", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetStatusDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
