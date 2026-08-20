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
	_ datasource.DataSource              = (*GetWaypointDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetWaypointDataSource)(nil)
)

// GetWaypointDataSource is the generated Terraform data source implementation.
type GetWaypointDataSource struct {
	client *client.Client
}

// GetWaypointDataSourceModel describes the data source state shape.
type GetWaypointDataSourceModel struct {
	Chart               types.Object `tfsdk:"chart"`
	Faction             types.Object `tfsdk:"faction"`
	IsUnderConstruction types.Bool   `tfsdk:"is_under_construction" json:"isUnderConstruction"`
	Modifiers           types.List   `tfsdk:"modifiers"`
	Orbitals            types.List   `tfsdk:"orbitals"`
	Orbits              types.String `tfsdk:"orbits"`
	Symbol              types.String `tfsdk:"symbol"`
	SystemSymbol        types.String `tfsdk:"system_symbol" json:"systemSymbol"`
	Traits              types.List   `tfsdk:"traits"`
	Type                types.String `tfsdk:"type"`
	WaypointSymbol      types.String `tfsdk:"waypoint_symbol" json:"waypointSymbol"`
	X                   types.Int64  `tfsdk:"x"`
	Y                   types.Int64  `tfsdk:"y"`
}

// NewGetWaypointDataSource returns a new instance of the generated data source.
func NewGetWaypointDataSource() datasource.DataSource {
	return &GetWaypointDataSource{}
}

// Metadata returns the data source type name.
func (d *GetWaypointDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "spacetraders_get_waypoint"
}

// Schema returns the data source schema.
func (d *GetWaypointDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "View the details of a waypoint.\n\nIf the waypoint is uncharted, it will return the 'Uncharted' trait instead of its actual traits.", Attributes: map[string]schema.Attribute{"chart": schema.SingleNestedAttribute{MarkdownDescription: "The chart of a system or waypoint, which makes the location visible to other agents.", Computed: true, Attributes: map[string]schema.Attribute{"submitted_by": schema.StringAttribute{MarkdownDescription: "The agent that submitted the chart for this waypoint.", Computed: true}, "submitted_on": schema.StringAttribute{MarkdownDescription: "The time the chart for this waypoint was submitted.", Computed: true}, "waypoint_symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the waypoint.", Computed: true}}}, "faction": schema.SingleNestedAttribute{MarkdownDescription: "The faction that controls the waypoint.", Computed: true, Attributes: map[string]schema.Attribute{"symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the faction.", Computed: true}}}, "is_under_construction": schema.BoolAttribute{MarkdownDescription: "True if the waypoint is under construction.", Computed: true}, "modifiers": schema.ListNestedAttribute{MarkdownDescription: "The modifiers of the waypoint.", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"description": schema.StringAttribute{MarkdownDescription: "A description of the trait.", Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "The name of the trait.", Computed: true}, "symbol": schema.StringAttribute{MarkdownDescription: "The unique identifier of the modifier.", Computed: true}}}}, "orbitals": schema.ListNestedAttribute{MarkdownDescription: "Waypoints that orbit this waypoint.", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the orbiting waypoint.", Computed: true}}}}, "orbits": schema.StringAttribute{MarkdownDescription: "The symbol of the parent waypoint, if this waypoint is in orbit around another waypoint. Otherwise this value is undefined.", Computed: true}, "symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the waypoint.", Computed: true}, "system_symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the system.", Required: true}, "traits": schema.ListNestedAttribute{MarkdownDescription: "The traits of the waypoint.", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"description": schema.StringAttribute{MarkdownDescription: "A description of the trait.", Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "The name of the trait.", Computed: true}, "symbol": schema.StringAttribute{MarkdownDescription: "The unique identifier of the trait.", Computed: true}}}}, "type": schema.StringAttribute{MarkdownDescription: "The type of waypoint.", Computed: true}, "waypoint_symbol": schema.StringAttribute{MarkdownDescription: "The waypoint symbol", Required: true}, "x": schema.Int64Attribute{MarkdownDescription: "Relative position of the waypoint on the system's x axis. This is not an absolute position in the universe.", Computed: true}, "y": schema.Int64Attribute{MarkdownDescription: "Relative position of the waypoint on the system's y axis. This is not an absolute position in the universe.", Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetWaypointDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetWaypointDataSourceModel
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
func (d *GetWaypointDataSource) readRemote(ctx context.Context, config *GetWaypointDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/systems/{systemSymbol}/waypoints/{waypointSymbol}"
	reqPath = strings.ReplaceAll(reqPath, "{systemSymbol}", url.PathEscape(config.SystemSymbol.ValueString()))
	reqPath = strings.ReplaceAll(reqPath, "{waypointSymbol}", url.PathEscape(config.WaypointSymbol.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_get_waypoint", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_get_waypoint", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading spacetraders_get_waypoint", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		apiErr, err := client.NewAPIError(httpResp)
		if err != nil {
			resp.Diagnostics.AddError("Error reading spacetraders_get_waypoint", fmt.Sprintf("Could not read error response: %s", err))
			return
		}
		resp.Diagnostics.AddError("Error reading spacetraders_get_waypoint", apiErr.Error())
		return
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading spacetraders_get_waypoint", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["data"]; ok {
		if m, ok := v.(map[string]any); ok {
			data = m
		}
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_get_waypoint", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetWaypointDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
