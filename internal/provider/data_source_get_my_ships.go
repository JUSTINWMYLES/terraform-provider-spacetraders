package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)
import (
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
	client "github.com/spacetraders/terraform-provider-spacetraders/internal/client"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*GetMyShipsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetMyShipsDataSource)(nil)
)

// GetMyShipsDataSource is the generated Terraform data source implementation.
type GetMyShipsDataSource struct {
	client *client.Client
}

// GetMyShipsDataSourceModel describes the data source state shape.
type GetMyShipsDataSourceModel struct {
	Items types.List  `tfsdk:"items"`
	Limit types.Int64 `tfsdk:"limit"`
	Page  types.Int64 `tfsdk:"page"`
}

// NewGetMyShipsDataSource returns a new instance of the generated data source.
func NewGetMyShipsDataSource() datasource.DataSource {
	return &GetMyShipsDataSource{}
}

// Metadata returns the data source type name.
func (d *GetMyShipsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "spacetraders_get_my_ships"
}

// Schema returns the data source schema.
func (d *GetMyShipsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Return a paginated list of all of ships under your agent's ownership.", Attributes: map[string]schema.Attribute{"items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"cargo": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"capacity": schema.Int64Attribute{Computed: true}, "inventory": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"description": schema.StringAttribute{Computed: true}, "name": schema.StringAttribute{Computed: true}, "symbol": schema.StringAttribute{Computed: true}, "units": schema.Int64Attribute{Computed: true}}}}, "units": schema.Int64Attribute{Computed: true}}}, "cooldown": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"expiration": schema.StringAttribute{Computed: true}, "remaining_seconds": schema.Int64Attribute{Computed: true}, "ship_symbol": schema.StringAttribute{Computed: true}, "total_seconds": schema.Int64Attribute{Computed: true}}}, "crew": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"capacity": schema.Int64Attribute{Computed: true}, "current": schema.Int64Attribute{Computed: true}, "morale": schema.Int64Attribute{Computed: true}, "required": schema.Int64Attribute{Computed: true}, "rotation": schema.StringAttribute{Computed: true}, "wages": schema.Int64Attribute{Computed: true}}}, "engine": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"condition": schema.Float64Attribute{Computed: true}, "description": schema.StringAttribute{Computed: true}, "integrity": schema.Float64Attribute{Computed: true}, "name": schema.StringAttribute{Computed: true}, "quality": schema.Float64Attribute{Computed: true}, "requirements": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"crew": schema.Int64Attribute{Computed: true}, "power": schema.Int64Attribute{Computed: true}, "slots": schema.Int64Attribute{Computed: true}}}, "speed": schema.Int64Attribute{Computed: true}, "symbol": schema.StringAttribute{Computed: true}}}, "frame": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"condition": schema.Float64Attribute{Computed: true}, "description": schema.StringAttribute{Computed: true}, "fuel_capacity": schema.Int64Attribute{Computed: true}, "integrity": schema.Float64Attribute{Computed: true}, "module_slots": schema.Int64Attribute{Computed: true}, "mounting_points": schema.Int64Attribute{Computed: true}, "name": schema.StringAttribute{Computed: true}, "quality": schema.Float64Attribute{Computed: true}, "requirements": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"crew": schema.Int64Attribute{Computed: true}, "power": schema.Int64Attribute{Computed: true}, "slots": schema.Int64Attribute{Computed: true}}}, "symbol": schema.StringAttribute{Computed: true}}}, "fuel": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"capacity": schema.Int64Attribute{Computed: true}, "consumed": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"amount": schema.Int64Attribute{Computed: true}, "timestamp": schema.StringAttribute{Computed: true}}}, "current": schema.Int64Attribute{Computed: true}}}, "modules": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"capacity": schema.Int64Attribute{Computed: true}, "description": schema.StringAttribute{Computed: true}, "name": schema.StringAttribute{Computed: true}, "range": schema.Int64Attribute{Computed: true}, "requirements": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"crew": schema.Int64Attribute{Computed: true}, "power": schema.Int64Attribute{Computed: true}, "slots": schema.Int64Attribute{Computed: true}}}, "symbol": schema.StringAttribute{Computed: true}}}}, "mounts": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"deposits": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "description": schema.StringAttribute{Computed: true}, "name": schema.StringAttribute{Computed: true}, "requirements": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"crew": schema.Int64Attribute{Computed: true}, "power": schema.Int64Attribute{Computed: true}, "slots": schema.Int64Attribute{Computed: true}}}, "strength": schema.Int64Attribute{Computed: true}, "symbol": schema.StringAttribute{Computed: true}}}}, "nav": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"flight_mode": schema.StringAttribute{Computed: true}, "route": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"arrival": schema.StringAttribute{Computed: true}, "departure_time": schema.StringAttribute{Computed: true}, "destination": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"symbol": schema.StringAttribute{Computed: true}, "system_symbol": schema.StringAttribute{Computed: true}, "type": schema.StringAttribute{Computed: true}, "x": schema.Int64Attribute{Computed: true}, "y": schema.Int64Attribute{Computed: true}}}, "origin": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"symbol": schema.StringAttribute{Computed: true}, "system_symbol": schema.StringAttribute{Computed: true}, "type": schema.StringAttribute{Computed: true}, "x": schema.Int64Attribute{Computed: true}, "y": schema.Int64Attribute{Computed: true}}}}}, "status": schema.StringAttribute{Computed: true}, "system_symbol": schema.StringAttribute{Computed: true}, "waypoint_symbol": schema.StringAttribute{Computed: true}}}, "reactor": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"condition": schema.Float64Attribute{Computed: true}, "description": schema.StringAttribute{Computed: true}, "integrity": schema.Float64Attribute{Computed: true}, "name": schema.StringAttribute{Computed: true}, "power_output": schema.Int64Attribute{Computed: true}, "quality": schema.Float64Attribute{Computed: true}, "requirements": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"crew": schema.Int64Attribute{Computed: true}, "power": schema.Int64Attribute{Computed: true}, "slots": schema.Int64Attribute{Computed: true}}}, "symbol": schema.StringAttribute{Computed: true}}}, "registration": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"faction_symbol": schema.StringAttribute{Computed: true}, "name": schema.StringAttribute{Computed: true}, "role": schema.StringAttribute{Computed: true}}}, "symbol": schema.StringAttribute{Computed: true}}}}, "limit": schema.Int64Attribute{Optional: true}, "page": schema.Int64Attribute{Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetMyShipsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetMyShipsDataSourceModel
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
func (d *GetMyShipsDataSource) readListRemote(ctx context.Context, config *GetMyShipsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/my/ships"
	params := url.Values{}
	if !config.Page.IsNull() {
		params.Set("page", strconv.FormatInt(config.Page.ValueInt64(), 10))
	}
	if !config.Limit.IsNull() {
		params.Set("limit", strconv.FormatInt(config.Limit.ValueInt64(), 10))
	}
	var nextURL string
	fetch := func(ctx context.Context, p url.Values) (*http.Response, error) {
		httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil, client.WithSchemes("AgentToken"))
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
		resp.Diagnostics.AddError("Error reading spacetraders_get_my_ships", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading spacetraders_get_my_ships", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["data"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading spacetraders_get_my_ships", fmt.Sprintf("Could not decode list page: missing %q array", "data"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_get_my_ships", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetMyShipsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
