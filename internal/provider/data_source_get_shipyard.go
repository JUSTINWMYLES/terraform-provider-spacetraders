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
	_ datasource.DataSource              = (*GetShipyardDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetShipyardDataSource)(nil)
)

// GetShipyardDataSource is the generated Terraform data source implementation.
type GetShipyardDataSource struct {
	client *client.Client
}

// GetShipyardDataSourceModel describes the data source state shape.
type GetShipyardDataSourceModel struct {
	ModificationsFee types.Int64  `tfsdk:"modifications_fee" json:"modificationsFee"`
	ShipTypes        types.List   `tfsdk:"ship_types" json:"shipTypes"`
	Ships            types.List   `tfsdk:"ships"`
	Symbol           types.String `tfsdk:"symbol"`
	SystemSymbol     types.String `tfsdk:"system_symbol" json:"systemSymbol"`
	Transactions     types.List   `tfsdk:"transactions"`
	WaypointSymbol   types.String `tfsdk:"waypoint_symbol" json:"waypointSymbol"`
}

// NewGetShipyardDataSource returns a new instance of the generated data source.
func NewGetShipyardDataSource() datasource.DataSource {
	return &GetShipyardDataSource{}
}

// Metadata returns the data source type name.
func (d *GetShipyardDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "spacetraders_get_shipyard"
}

// Schema returns the data source schema.
func (d *GetShipyardDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get the shipyard for a waypoint. Requires a waypoint that has the `Shipyard` trait to use. Send a ship to the waypoint to access data on ships that are currently available for purchase and recent transactions.", Attributes: map[string]schema.Attribute{"modifications_fee": schema.Int64Attribute{Computed: true}, "ship_types": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"type": schema.StringAttribute{Computed: true}}}}, "ships": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"activity": schema.StringAttribute{Computed: true}, "crew": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"capacity": schema.Int64Attribute{Computed: true}, "required": schema.Int64Attribute{Computed: true}}}, "description": schema.StringAttribute{Computed: true}, "engine": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"condition": schema.Float64Attribute{Computed: true}, "description": schema.StringAttribute{Computed: true}, "integrity": schema.Float64Attribute{Computed: true}, "name": schema.StringAttribute{Computed: true}, "quality": schema.Float64Attribute{Computed: true}, "requirements": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"crew": schema.Int64Attribute{Computed: true}, "power": schema.Int64Attribute{Computed: true}, "slots": schema.Int64Attribute{Computed: true}}}, "speed": schema.Int64Attribute{Computed: true}, "symbol": schema.StringAttribute{Computed: true}}}, "frame": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"condition": schema.Float64Attribute{Computed: true}, "description": schema.StringAttribute{Computed: true}, "fuel_capacity": schema.Int64Attribute{Computed: true}, "integrity": schema.Float64Attribute{Computed: true}, "module_slots": schema.Int64Attribute{Computed: true}, "mounting_points": schema.Int64Attribute{Computed: true}, "name": schema.StringAttribute{Computed: true}, "quality": schema.Float64Attribute{Computed: true}, "requirements": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"crew": schema.Int64Attribute{Computed: true}, "power": schema.Int64Attribute{Computed: true}, "slots": schema.Int64Attribute{Computed: true}}}, "symbol": schema.StringAttribute{Computed: true}}}, "modules": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"capacity": schema.Int64Attribute{Computed: true}, "description": schema.StringAttribute{Computed: true}, "name": schema.StringAttribute{Computed: true}, "range": schema.Int64Attribute{Computed: true}, "requirements": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"crew": schema.Int64Attribute{Computed: true}, "power": schema.Int64Attribute{Computed: true}, "slots": schema.Int64Attribute{Computed: true}}}, "symbol": schema.StringAttribute{Computed: true}}}}, "mounts": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"deposits": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "description": schema.StringAttribute{Computed: true}, "name": schema.StringAttribute{Computed: true}, "requirements": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"crew": schema.Int64Attribute{Computed: true}, "power": schema.Int64Attribute{Computed: true}, "slots": schema.Int64Attribute{Computed: true}}}, "strength": schema.Int64Attribute{Computed: true}, "symbol": schema.StringAttribute{Computed: true}}}}, "name": schema.StringAttribute{Computed: true}, "purchase_price": schema.Int64Attribute{Computed: true}, "reactor": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"condition": schema.Float64Attribute{Computed: true}, "description": schema.StringAttribute{Computed: true}, "integrity": schema.Float64Attribute{Computed: true}, "name": schema.StringAttribute{Computed: true}, "power_output": schema.Int64Attribute{Computed: true}, "quality": schema.Float64Attribute{Computed: true}, "requirements": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"crew": schema.Int64Attribute{Computed: true}, "power": schema.Int64Attribute{Computed: true}, "slots": schema.Int64Attribute{Computed: true}}}, "symbol": schema.StringAttribute{Computed: true}}}, "supply": schema.StringAttribute{Computed: true}, "type": schema.StringAttribute{Computed: true}}}}, "symbol": schema.StringAttribute{Computed: true}, "system_symbol": schema.StringAttribute{Required: true}, "transactions": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"agent_symbol": schema.StringAttribute{Computed: true}, "price": schema.Int64Attribute{Computed: true}, "ship_symbol": schema.StringAttribute{Computed: true}, "ship_type": schema.StringAttribute{Computed: true}, "timestamp": schema.StringAttribute{Computed: true}, "waypoint_symbol": schema.StringAttribute{Computed: true}}}}, "waypoint_symbol": schema.StringAttribute{Required: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetShipyardDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetShipyardDataSourceModel
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
func (d *GetShipyardDataSource) readRemote(ctx context.Context, config *GetShipyardDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/systems/{systemSymbol}/waypoints/{waypointSymbol}/shipyard"
	reqPath = strings.ReplaceAll(reqPath, "{systemSymbol}", url.PathEscape(config.SystemSymbol.ValueString()))
	reqPath = strings.ReplaceAll(reqPath, "{waypointSymbol}", url.PathEscape(config.WaypointSymbol.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_get_shipyard", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_get_shipyard", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading spacetraders_get_shipyard", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		apiErr, err := client.NewAPIError(httpResp)
		if err != nil {
			resp.Diagnostics.AddError("Error reading spacetraders_get_shipyard", fmt.Sprintf("Could not read error response: %s", err))
			return
		}
		resp.Diagnostics.AddError("Error reading spacetraders_get_shipyard", apiErr.Error())
		return
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading spacetraders_get_shipyard", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["data"]; ok {
		if m, ok := v.(map[string]any); ok {
			data = m
		}
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_get_shipyard", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetShipyardDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
