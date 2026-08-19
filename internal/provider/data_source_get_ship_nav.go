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
	_ datasource.DataSource              = (*GetShipNavDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetShipNavDataSource)(nil)
)

// GetShipNavDataSource is the generated Terraform data source implementation.
type GetShipNavDataSource struct {
	client *client.Client
}

// GetShipNavDataSourceModel describes the data source state shape.
type GetShipNavDataSourceModel struct {
	FlightMode     types.String `tfsdk:"flight_mode" json:"flightMode"`
	Route          types.Object `tfsdk:"route"`
	ShipSymbol     types.String `tfsdk:"ship_symbol" json:"shipSymbol"`
	Status         types.String `tfsdk:"status"`
	SystemSymbol   types.String `tfsdk:"system_symbol" json:"systemSymbol"`
	WaypointSymbol types.String `tfsdk:"waypoint_symbol" json:"waypointSymbol"`
}

// NewGetShipNavDataSource returns a new instance of the generated data source.
func NewGetShipNavDataSource() datasource.DataSource {
	return &GetShipNavDataSource{}
}

// Metadata returns the data source type name.
func (d *GetShipNavDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "spacetraders_get_ship_nav"
}

// Schema returns the data source schema.
func (d *GetShipNavDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get the current nav status of a ship.", Attributes: map[string]schema.Attribute{"flight_mode": schema.StringAttribute{Computed: true}, "route": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"arrival": schema.StringAttribute{Computed: true}, "departure_time": schema.StringAttribute{Computed: true}, "destination": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"symbol": schema.StringAttribute{Computed: true}, "system_symbol": schema.StringAttribute{Computed: true}, "type": schema.StringAttribute{Computed: true}, "x": schema.Int64Attribute{Computed: true}, "y": schema.Int64Attribute{Computed: true}}}, "origin": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"symbol": schema.StringAttribute{Computed: true}, "system_symbol": schema.StringAttribute{Computed: true}, "type": schema.StringAttribute{Computed: true}, "x": schema.Int64Attribute{Computed: true}, "y": schema.Int64Attribute{Computed: true}}}}}, "ship_symbol": schema.StringAttribute{Required: true}, "status": schema.StringAttribute{Computed: true}, "system_symbol": schema.StringAttribute{Computed: true}, "waypoint_symbol": schema.StringAttribute{Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetShipNavDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetShipNavDataSourceModel
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
func (d *GetShipNavDataSource) readRemote(ctx context.Context, config *GetShipNavDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/my/ships/{shipSymbol}/nav"
	reqPath = strings.ReplaceAll(reqPath, "{shipSymbol}", url.PathEscape(config.ShipSymbol.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil, client.WithSchemes("AgentToken"))
	if err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_get_ship_nav", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_get_ship_nav", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading spacetraders_get_ship_nav", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		apiErr, err := client.NewAPIError(httpResp)
		if err != nil {
			resp.Diagnostics.AddError("Error reading spacetraders_get_ship_nav", fmt.Sprintf("Could not read error response: %s", err))
			return
		}
		resp.Diagnostics.AddError("Error reading spacetraders_get_ship_nav", apiErr.Error())
		return
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading spacetraders_get_ship_nav", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["data"]; ok {
		if m, ok := v.(map[string]any); ok {
			data = m
		}
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_get_ship_nav", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetShipNavDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
