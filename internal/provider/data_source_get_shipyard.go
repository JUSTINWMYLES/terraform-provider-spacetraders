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
	resp.Schema = schema.Schema{MarkdownDescription: "Get the shipyard for a waypoint. Requires a waypoint that has the `Shipyard` trait to use. Send a ship to the waypoint to access data on ships that are currently available for purchase and recent transactions.", Attributes: map[string]schema.Attribute{"modifications_fee": schema.Int64Attribute{MarkdownDescription: "The fee to modify a ship at this shipyard. This includes installing or removing modules and mounts on a ship. In the case of mounts, the fee is a flat rate per mount. In the case of modules, the fee is per slot the module occupies.", Computed: true}, "ship_types": schema.ListNestedAttribute{MarkdownDescription: "The list of ship types available for purchase at this shipyard.", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"type": schema.StringAttribute{MarkdownDescription: "Type of ship", Computed: true}}}}, "ships": schema.ListNestedAttribute{MarkdownDescription: "The ships that are currently available for purchase at the shipyard.", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"activity": schema.StringAttribute{MarkdownDescription: "The activity level of a trade good. If the good is an import, this represents how strong consumption is. If the good is an export, this represents how strong the production is for the good. When activity is strong, consumption or production is near maximum capacity. When activity is weak, consumption or production is near minimum capacity.", Computed: true}, "crew": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"capacity": schema.Int64Attribute{MarkdownDescription: "The maximum number of crew members the ship can support.", Computed: true}, "required": schema.Int64Attribute{MarkdownDescription: "The minimum number of crew members required to maintain the ship.", Computed: true}}}, "description": schema.StringAttribute{MarkdownDescription: "Description of the ship.", Computed: true}, "engine": schema.SingleNestedAttribute{MarkdownDescription: "The engine determines how quickly a ship travels between waypoints.", Computed: true, Attributes: map[string]schema.Attribute{"condition": schema.Float64Attribute{MarkdownDescription: "The repairable condition of a component. A value of 0 indicates the component needs significant repairs, while a value of 1 indicates the component is in near perfect condition. As the condition of a component is repaired, the overall integrity of the component decreases.", Computed: true}, "description": schema.StringAttribute{MarkdownDescription: "The description of the engine.", Computed: true}, "integrity": schema.Float64Attribute{MarkdownDescription: "The overall integrity of the component, which determines the performance of the component. A value of 0 indicates that the component is almost completely degraded, while a value of 1 indicates that the component is in near perfect condition. The integrity of the component is non-repairable, and represents permanent wear over time.", Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "The name of the engine.", Computed: true}, "quality": schema.Float64Attribute{MarkdownDescription: "The overall quality of the component, which determines the quality of the component. High quality components return more ships parts and ship plating when a ship is scrapped. But also require more of these parts to repair. This is transparent to the player, as the parts are bought from/sold to the marketplace.", Computed: true}, "requirements": schema.SingleNestedAttribute{MarkdownDescription: "The requirements for installation on a ship", Computed: true, Attributes: map[string]schema.Attribute{"crew": schema.Int64Attribute{MarkdownDescription: "The number of crew required for operation.", Computed: true}, "power": schema.Int64Attribute{MarkdownDescription: "The amount of power required from the reactor.", Computed: true}, "slots": schema.Int64Attribute{MarkdownDescription: "The number of module slots required for installation.", Computed: true}}}, "speed": schema.Int64Attribute{MarkdownDescription: "The speed stat of this engine. The higher the speed, the faster a ship can travel from one point to another. Reduces the time of arrival when navigating the ship.", Computed: true}, "symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the engine.", Computed: true}}}, "frame": schema.SingleNestedAttribute{MarkdownDescription: "The frame of the ship. The frame determines the number of modules and mounting points of the ship, as well as base fuel capacity. As the condition of the frame takes more wear, the ship will become more sluggish and less maneuverable.", Computed: true, Attributes: map[string]schema.Attribute{"condition": schema.Float64Attribute{MarkdownDescription: "The repairable condition of a component. A value of 0 indicates the component needs significant repairs, while a value of 1 indicates the component is in near perfect condition. As the condition of a component is repaired, the overall integrity of the component decreases.", Computed: true}, "description": schema.StringAttribute{MarkdownDescription: "Description of the frame.", Computed: true}, "fuel_capacity": schema.Int64Attribute{MarkdownDescription: "The maximum amount of fuel that can be stored in this ship. When refueling, the ship will be refueled to this amount.", Computed: true}, "integrity": schema.Float64Attribute{MarkdownDescription: "The overall integrity of the component, which determines the performance of the component. A value of 0 indicates that the component is almost completely degraded, while a value of 1 indicates that the component is in near perfect condition. The integrity of the component is non-repairable, and represents permanent wear over time.", Computed: true}, "module_slots": schema.Int64Attribute{MarkdownDescription: "The amount of slots that can be dedicated to modules installed in the ship. Each installed module take up a number of slots, and once there are no more slots, no new modules can be installed.", Computed: true}, "mounting_points": schema.Int64Attribute{MarkdownDescription: "The amount of slots that can be dedicated to mounts installed in the ship. Each installed mount takes up a number of points, and once there are no more points remaining, no new mounts can be installed.", Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "Name of the frame.", Computed: true}, "quality": schema.Float64Attribute{MarkdownDescription: "The overall quality of the component, which determines the quality of the component. High quality components return more ships parts and ship plating when a ship is scrapped. But also require more of these parts to repair. This is transparent to the player, as the parts are bought from/sold to the marketplace.", Computed: true}, "requirements": schema.SingleNestedAttribute{MarkdownDescription: "The requirements for installation on a ship", Computed: true, Attributes: map[string]schema.Attribute{"crew": schema.Int64Attribute{MarkdownDescription: "The number of crew required for operation.", Computed: true}, "power": schema.Int64Attribute{MarkdownDescription: "The amount of power required from the reactor.", Computed: true}, "slots": schema.Int64Attribute{MarkdownDescription: "The number of module slots required for installation.", Computed: true}}}, "symbol": schema.StringAttribute{MarkdownDescription: "Symbol of the frame.", Computed: true}}}, "modules": schema.ListNestedAttribute{MarkdownDescription: "Modules installed in this ship.", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"capacity": schema.Int64Attribute{MarkdownDescription: "Modules that provide capacity, such as cargo hold or crew quarters will show this value to denote how much of a bonus the module grants.", Computed: true}, "description": schema.StringAttribute{MarkdownDescription: "Description of this module.", Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "Name of this module.", Computed: true}, "range": schema.Int64Attribute{MarkdownDescription: "Modules that have a range will such as a sensor array show this value to denote how far can the module reach with its capabilities.", Computed: true}, "requirements": schema.SingleNestedAttribute{MarkdownDescription: "The requirements for installation on a ship", Computed: true, Attributes: map[string]schema.Attribute{"crew": schema.Int64Attribute{MarkdownDescription: "The number of crew required for operation.", Computed: true}, "power": schema.Int64Attribute{MarkdownDescription: "The amount of power required from the reactor.", Computed: true}, "slots": schema.Int64Attribute{MarkdownDescription: "The number of module slots required for installation.", Computed: true}}}, "symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the module.", Computed: true}}}}, "mounts": schema.ListNestedAttribute{MarkdownDescription: "Mounts installed in this ship.", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"deposits": schema.ListAttribute{MarkdownDescription: "Mounts that have this value denote what goods can be produced from using the mount.", Computed: true, ElementType: types.StringType}, "description": schema.StringAttribute{MarkdownDescription: "Description of this mount.", Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "Name of this mount.", Computed: true}, "requirements": schema.SingleNestedAttribute{MarkdownDescription: "The requirements for installation on a ship", Computed: true, Attributes: map[string]schema.Attribute{"crew": schema.Int64Attribute{MarkdownDescription: "The number of crew required for operation.", Computed: true}, "power": schema.Int64Attribute{MarkdownDescription: "The amount of power required from the reactor.", Computed: true}, "slots": schema.Int64Attribute{MarkdownDescription: "The number of module slots required for installation.", Computed: true}}}, "strength": schema.Int64Attribute{MarkdownDescription: "Mounts that have this value, such as mining lasers, denote how powerful this mount's capabilities are.", Computed: true}, "symbol": schema.StringAttribute{MarkdownDescription: "Symbol of this mount.", Computed: true}}}}, "name": schema.StringAttribute{MarkdownDescription: "Name of the ship.", Computed: true}, "purchase_price": schema.Int64Attribute{MarkdownDescription: "The purchase price of the ship.", Computed: true}, "reactor": schema.SingleNestedAttribute{MarkdownDescription: "The reactor of the ship. The reactor is responsible for powering the ship's systems and weapons.", Computed: true, Attributes: map[string]schema.Attribute{"condition": schema.Float64Attribute{MarkdownDescription: "The repairable condition of a component. A value of 0 indicates the component needs significant repairs, while a value of 1 indicates the component is in near perfect condition. As the condition of a component is repaired, the overall integrity of the component decreases.", Computed: true}, "description": schema.StringAttribute{MarkdownDescription: "Description of the reactor.", Computed: true}, "integrity": schema.Float64Attribute{MarkdownDescription: "The overall integrity of the component, which determines the performance of the component. A value of 0 indicates that the component is almost completely degraded, while a value of 1 indicates that the component is in near perfect condition. The integrity of the component is non-repairable, and represents permanent wear over time.", Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "Name of the reactor.", Computed: true}, "power_output": schema.Int64Attribute{MarkdownDescription: "The amount of power provided by this reactor. The more power a reactor provides to the ship, the lower the cooldown it gets when using a module or mount that taxes the ship's power.", Computed: true}, "quality": schema.Float64Attribute{MarkdownDescription: "The overall quality of the component, which determines the quality of the component. High quality components return more ships parts and ship plating when a ship is scrapped. But also require more of these parts to repair. This is transparent to the player, as the parts are bought from/sold to the marketplace.", Computed: true}, "requirements": schema.SingleNestedAttribute{MarkdownDescription: "The requirements for installation on a ship", Computed: true, Attributes: map[string]schema.Attribute{"crew": schema.Int64Attribute{MarkdownDescription: "The number of crew required for operation.", Computed: true}, "power": schema.Int64Attribute{MarkdownDescription: "The amount of power required from the reactor.", Computed: true}, "slots": schema.Int64Attribute{MarkdownDescription: "The number of module slots required for installation.", Computed: true}}}, "symbol": schema.StringAttribute{MarkdownDescription: "Symbol of the reactor.", Computed: true}}}, "supply": schema.StringAttribute{MarkdownDescription: "The supply level of a trade good.", Computed: true}, "type": schema.StringAttribute{MarkdownDescription: "Type of ship", Computed: true}}}}, "symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the shipyard. The symbol is the same as the waypoint where the shipyard is located.", Computed: true}, "system_symbol": schema.StringAttribute{MarkdownDescription: "The system symbol", Required: true}, "transactions": schema.ListNestedAttribute{MarkdownDescription: "The list of recent transactions at this shipyard.", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"agent_symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the agent that made the transaction.", Computed: true}, "price": schema.Int64Attribute{MarkdownDescription: "The price of the transaction.", Computed: true}, "ship_symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the ship type (e.g. SHIP_MINING_DRONE) that was the subject of the transaction. Contrary to what the name implies, this is NOT the symbol of the ship that was purchased.", Computed: true}, "ship_type": schema.StringAttribute{MarkdownDescription: "The symbol of the ship type (e.g. SHIP_MINING_DRONE) that was the subject of the transaction.", Computed: true}, "timestamp": schema.StringAttribute{MarkdownDescription: "The timestamp of the transaction.", Computed: true}, "waypoint_symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the waypoint.", Computed: true}}}}, "waypoint_symbol": schema.StringAttribute{MarkdownDescription: "The waypoint symbol", Required: true}}}
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
