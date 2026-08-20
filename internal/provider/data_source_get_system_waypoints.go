package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
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
	_ datasource.DataSource              = (*GetSystemWaypointsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetSystemWaypointsDataSource)(nil)
)

// GetSystemWaypointsDataSource is the generated Terraform data source implementation.
type GetSystemWaypointsDataSource struct {
	client *client.Client
}

// GetSystemWaypointsDataSourceModel describes the data source state shape.
type GetSystemWaypointsDataSourceModel struct {
	Items        types.List   `tfsdk:"items"`
	Limit        types.Int64  `tfsdk:"limit"`
	Page         types.Int64  `tfsdk:"page"`
	SystemSymbol types.String `tfsdk:"system_symbol" json:"systemSymbol"`
	Traits       types.String `tfsdk:"traits"`
	Type         types.String `tfsdk:"type"`
}

// NewGetSystemWaypointsDataSource returns a new instance of the generated data source.
func NewGetSystemWaypointsDataSource() datasource.DataSource {
	return &GetSystemWaypointsDataSource{}
}

// Metadata returns the data source type name.
func (d *GetSystemWaypointsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "spacetraders_get_system_waypoints"
}

// Schema returns the data source schema.
func (d *GetSystemWaypointsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Return a paginated list of all of the waypoints for a given system.\n\nIf a waypoint is uncharted, it will return the `Uncharted` trait instead of its actual traits.", Attributes: map[string]schema.Attribute{"items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"chart": schema.SingleNestedAttribute{MarkdownDescription: "The chart of a system or waypoint, which makes the location visible to other agents.", Computed: true, Attributes: map[string]schema.Attribute{"submitted_by": schema.StringAttribute{MarkdownDescription: "The agent that submitted the chart for this waypoint.", Computed: true}, "submitted_on": schema.StringAttribute{MarkdownDescription: "The time the chart for this waypoint was submitted.", Computed: true}, "waypoint_symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the waypoint.", Computed: true}}}, "faction": schema.SingleNestedAttribute{MarkdownDescription: "The faction that controls the waypoint.", Computed: true, Attributes: map[string]schema.Attribute{"symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the faction.", Computed: true}}}, "is_under_construction": schema.BoolAttribute{MarkdownDescription: "True if the waypoint is under construction.", Computed: true}, "modifiers": schema.ListNestedAttribute{MarkdownDescription: "The modifiers of the waypoint.", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"description": schema.StringAttribute{MarkdownDescription: "A description of the trait.", Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "The name of the trait.", Computed: true}, "symbol": schema.StringAttribute{MarkdownDescription: "The unique identifier of the modifier.", Computed: true}}}}, "orbitals": schema.ListNestedAttribute{MarkdownDescription: "Waypoints that orbit this waypoint.", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the orbiting waypoint.", Computed: true}}}}, "orbits": schema.StringAttribute{MarkdownDescription: "The symbol of the parent waypoint, if this waypoint is in orbit around another waypoint. Otherwise this value is undefined.", Computed: true}, "symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the waypoint.", Computed: true}, "system_symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the system.", Computed: true}, "traits": schema.ListNestedAttribute{MarkdownDescription: "The traits of the waypoint.", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"description": schema.StringAttribute{MarkdownDescription: "A description of the trait.", Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "The name of the trait.", Computed: true}, "symbol": schema.StringAttribute{MarkdownDescription: "The unique identifier of the trait.", Computed: true}}}}, "type": schema.StringAttribute{MarkdownDescription: "The type of waypoint.", Computed: true}, "x": schema.Int64Attribute{MarkdownDescription: "Relative position of the waypoint on the system's x axis. This is not an absolute position in the universe.", Computed: true}, "y": schema.Int64Attribute{MarkdownDescription: "Relative position of the waypoint on the system's y axis. This is not an absolute position in the universe.", Computed: true}}}}, "limit": schema.Int64Attribute{MarkdownDescription: "How many entries to return per page", Optional: true}, "page": schema.Int64Attribute{MarkdownDescription: "What entry offset to request", Optional: true}, "system_symbol": schema.StringAttribute{Required: true}, "traits": schema.StringAttribute{MarkdownDescription: "Filter waypoints by one or more traits.", Optional: true}, "type": schema.StringAttribute{MarkdownDescription: "Filter waypoints by type.", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetSystemWaypointsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetSystemWaypointsDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	d.readListRemote(ctx, &config, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}

// readListRemote performs the paginated read HTTP exchange and decodes the response array into config. Extracted from Read so the request/response logic is unit-testable without a tfsdk.Config.
func (d *GetSystemWaypointsDataSource) readListRemote(ctx context.Context, config *GetSystemWaypointsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/systems/{systemSymbol}/waypoints"
	reqPath = strings.ReplaceAll(reqPath, "{systemSymbol}", url.PathEscape(config.SystemSymbol.ValueString()))
	params := url.Values{}
	if !config.Page.IsNull() {
		params.Set("page", strconv.FormatInt(config.Page.ValueInt64(), 10))
	}
	if !config.Limit.IsNull() {
		params.Set("limit", strconv.FormatInt(config.Limit.ValueInt64(), 10))
	}
	if !config.Type.IsNull() {
		params.Set("type", config.Type.ValueString())
	}
	if !config.Traits.IsNull() {
		params.Set("traits", config.Traits.ValueString())
	}
	var nextURL string
	fetch := func(ctx context.Context, p url.Values) (*http.Response, error) {
		httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
		if err != nil {
			return nil, err
		}
		if nextURL != "" {
			parsed, perr := url.Parse(nextURL)
			if perr != nil {
				return nil, perr
			}
			httpReq.URL = parsed
		} else {
			httpReq.URL.RawQuery = p.Encode()
		}
		return d.client.Do(httpReq)
	}
	pages, err := client.ListAllPages(ctx, params, fetch, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_get_system_waypoints", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading spacetraders_get_system_waypoints", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["data"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading spacetraders_get_system_waypoints", fmt.Sprintf("Could not decode list page: missing %q array", "data"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_get_system_waypoints", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetSystemWaypointsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
