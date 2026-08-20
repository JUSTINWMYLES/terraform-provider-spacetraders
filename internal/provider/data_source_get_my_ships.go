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
	resp.Schema = schema.Schema{MarkdownDescription: "Return a paginated list of all of ships under your agent's ownership.", Attributes: map[string]schema.Attribute{"items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"cargo": schema.SingleNestedAttribute{MarkdownDescription: "Ship cargo details.", Computed: true, Attributes: map[string]schema.Attribute{"capacity": schema.Int64Attribute{MarkdownDescription: "The max number of items that can be stored in the cargo hold.", Computed: true}, "inventory": schema.ListNestedAttribute{MarkdownDescription: "The items currently in the cargo hold.", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"description": schema.StringAttribute{MarkdownDescription: "The description of the cargo item type.", Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "The name of the cargo item type.", Computed: true}, "symbol": schema.StringAttribute{MarkdownDescription: "The good's symbol.", Computed: true}, "units": schema.Int64Attribute{MarkdownDescription: "The number of units of the cargo item.", Computed: true}}}}, "units": schema.Int64Attribute{MarkdownDescription: "The number of items currently stored in the cargo hold.", Computed: true}}}, "cooldown": schema.SingleNestedAttribute{MarkdownDescription: "A cooldown is a period of time in which a ship cannot perform certain actions.", Computed: true, Attributes: map[string]schema.Attribute{"expiration": schema.StringAttribute{MarkdownDescription: "The date and time when the cooldown expires in ISO 8601 format", Computed: true}, "remaining_seconds": schema.Int64Attribute{MarkdownDescription: "The remaining duration of the cooldown in seconds", Computed: true}, "ship_symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the ship that is on cooldown", Computed: true}, "total_seconds": schema.Int64Attribute{MarkdownDescription: "The total duration of the cooldown in seconds", Computed: true}}}, "crew": schema.SingleNestedAttribute{MarkdownDescription: "The ship's crew service and maintain the ship's systems and equipment.", Computed: true, Attributes: map[string]schema.Attribute{"capacity": schema.Int64Attribute{MarkdownDescription: "The maximum number of crew members the ship can support.", Computed: true}, "current": schema.Int64Attribute{MarkdownDescription: "The current number of crew members on the ship.", Computed: true}, "morale": schema.Int64Attribute{MarkdownDescription: "A rough measure of the crew's morale. A higher morale means the crew is happier and more productive. A lower morale means the ship is more prone to accidents.", Computed: true}, "required": schema.Int64Attribute{MarkdownDescription: "The minimum number of crew members required to maintain the ship.", Computed: true}, "rotation": schema.StringAttribute{MarkdownDescription: "The rotation of crew shifts. A stricter shift improves the ship's performance. A more relaxed shift improves the crew's morale.", Computed: true}, "wages": schema.Int64Attribute{MarkdownDescription: "The amount of credits per crew member paid per hour. Wages are paid when a ship docks at a civilized waypoint.", Computed: true}}}, "engine": schema.SingleNestedAttribute{MarkdownDescription: "The engine determines how quickly a ship travels between waypoints.", Computed: true, Attributes: map[string]schema.Attribute{"condition": schema.Float64Attribute{MarkdownDescription: "The repairable condition of a component. A value of 0 indicates the component needs significant repairs, while a value of 1 indicates the component is in near perfect condition. As the condition of a component is repaired, the overall integrity of the component decreases.", Computed: true}, "description": schema.StringAttribute{MarkdownDescription: "The description of the engine.", Computed: true}, "integrity": schema.Float64Attribute{MarkdownDescription: "The overall integrity of the component, which determines the performance of the component. A value of 0 indicates that the component is almost completely degraded, while a value of 1 indicates that the component is in near perfect condition. The integrity of the component is non-repairable, and represents permanent wear over time.", Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "The name of the engine.", Computed: true}, "quality": schema.Float64Attribute{MarkdownDescription: "The overall quality of the component, which determines the quality of the component. High quality components return more ships parts and ship plating when a ship is scrapped. But also require more of these parts to repair. This is transparent to the player, as the parts are bought from/sold to the marketplace.", Computed: true}, "requirements": schema.SingleNestedAttribute{MarkdownDescription: "The requirements for installation on a ship", Computed: true, Attributes: map[string]schema.Attribute{"crew": schema.Int64Attribute{MarkdownDescription: "The number of crew required for operation.", Computed: true}, "power": schema.Int64Attribute{MarkdownDescription: "The amount of power required from the reactor.", Computed: true}, "slots": schema.Int64Attribute{MarkdownDescription: "The number of module slots required for installation.", Computed: true}}}, "speed": schema.Int64Attribute{MarkdownDescription: "The speed stat of this engine. The higher the speed, the faster a ship can travel from one point to another. Reduces the time of arrival when navigating the ship.", Computed: true}, "symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the engine.", Computed: true}}}, "frame": schema.SingleNestedAttribute{MarkdownDescription: "The frame of the ship. The frame determines the number of modules and mounting points of the ship, as well as base fuel capacity. As the condition of the frame takes more wear, the ship will become more sluggish and less maneuverable.", Computed: true, Attributes: map[string]schema.Attribute{"condition": schema.Float64Attribute{MarkdownDescription: "The repairable condition of a component. A value of 0 indicates the component needs significant repairs, while a value of 1 indicates the component is in near perfect condition. As the condition of a component is repaired, the overall integrity of the component decreases.", Computed: true}, "description": schema.StringAttribute{MarkdownDescription: "Description of the frame.", Computed: true}, "fuel_capacity": schema.Int64Attribute{MarkdownDescription: "The maximum amount of fuel that can be stored in this ship. When refueling, the ship will be refueled to this amount.", Computed: true}, "integrity": schema.Float64Attribute{MarkdownDescription: "The overall integrity of the component, which determines the performance of the component. A value of 0 indicates that the component is almost completely degraded, while a value of 1 indicates that the component is in near perfect condition. The integrity of the component is non-repairable, and represents permanent wear over time.", Computed: true}, "module_slots": schema.Int64Attribute{MarkdownDescription: "The amount of slots that can be dedicated to modules installed in the ship. Each installed module take up a number of slots, and once there are no more slots, no new modules can be installed.", Computed: true}, "mounting_points": schema.Int64Attribute{MarkdownDescription: "The amount of slots that can be dedicated to mounts installed in the ship. Each installed mount takes up a number of points, and once there are no more points remaining, no new mounts can be installed.", Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "Name of the frame.", Computed: true}, "quality": schema.Float64Attribute{MarkdownDescription: "The overall quality of the component, which determines the quality of the component. High quality components return more ships parts and ship plating when a ship is scrapped. But also require more of these parts to repair. This is transparent to the player, as the parts are bought from/sold to the marketplace.", Computed: true}, "requirements": schema.SingleNestedAttribute{MarkdownDescription: "The requirements for installation on a ship", Computed: true, Attributes: map[string]schema.Attribute{"crew": schema.Int64Attribute{MarkdownDescription: "The number of crew required for operation.", Computed: true}, "power": schema.Int64Attribute{MarkdownDescription: "The amount of power required from the reactor.", Computed: true}, "slots": schema.Int64Attribute{MarkdownDescription: "The number of module slots required for installation.", Computed: true}}}, "symbol": schema.StringAttribute{MarkdownDescription: "Symbol of the frame.", Computed: true}}}, "fuel": schema.SingleNestedAttribute{MarkdownDescription: "Details of the ship's fuel tanks including how much fuel was consumed during the last transit or action.", Computed: true, Attributes: map[string]schema.Attribute{"capacity": schema.Int64Attribute{MarkdownDescription: "The maximum amount of fuel the ship's tanks can hold.", Computed: true}, "consumed": schema.SingleNestedAttribute{MarkdownDescription: "An object that only shows up when an action has consumed fuel in the process. Shows the fuel consumption data.", Computed: true, Attributes: map[string]schema.Attribute{"amount": schema.Int64Attribute{MarkdownDescription: "The amount of fuel consumed by the most recent transit or action.", Computed: true}, "timestamp": schema.StringAttribute{MarkdownDescription: "The time at which the fuel was consumed.", Computed: true}}}, "current": schema.Int64Attribute{MarkdownDescription: "The current amount of fuel in the ship's tanks.", Computed: true}}}, "modules": schema.ListNestedAttribute{MarkdownDescription: "Modules installed in this ship.", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"capacity": schema.Int64Attribute{MarkdownDescription: "Modules that provide capacity, such as cargo hold or crew quarters will show this value to denote how much of a bonus the module grants.", Computed: true}, "description": schema.StringAttribute{MarkdownDescription: "Description of this module.", Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "Name of this module.", Computed: true}, "range": schema.Int64Attribute{MarkdownDescription: "Modules that have a range will such as a sensor array show this value to denote how far can the module reach with its capabilities.", Computed: true}, "requirements": schema.SingleNestedAttribute{MarkdownDescription: "The requirements for installation on a ship", Computed: true, Attributes: map[string]schema.Attribute{"crew": schema.Int64Attribute{MarkdownDescription: "The number of crew required for operation.", Computed: true}, "power": schema.Int64Attribute{MarkdownDescription: "The amount of power required from the reactor.", Computed: true}, "slots": schema.Int64Attribute{MarkdownDescription: "The number of module slots required for installation.", Computed: true}}}, "symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the module.", Computed: true}}}}, "mounts": schema.ListNestedAttribute{MarkdownDescription: "Mounts installed in this ship.", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"deposits": schema.ListAttribute{MarkdownDescription: "Mounts that have this value denote what goods can be produced from using the mount.", Computed: true, ElementType: types.StringType}, "description": schema.StringAttribute{MarkdownDescription: "Description of this mount.", Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "Name of this mount.", Computed: true}, "requirements": schema.SingleNestedAttribute{MarkdownDescription: "The requirements for installation on a ship", Computed: true, Attributes: map[string]schema.Attribute{"crew": schema.Int64Attribute{MarkdownDescription: "The number of crew required for operation.", Computed: true}, "power": schema.Int64Attribute{MarkdownDescription: "The amount of power required from the reactor.", Computed: true}, "slots": schema.Int64Attribute{MarkdownDescription: "The number of module slots required for installation.", Computed: true}}}, "strength": schema.Int64Attribute{MarkdownDescription: "Mounts that have this value, such as mining lasers, denote how powerful this mount's capabilities are.", Computed: true}, "symbol": schema.StringAttribute{MarkdownDescription: "Symbol of this mount.", Computed: true}}}}, "nav": schema.SingleNestedAttribute{MarkdownDescription: "The navigation information of the ship.", Computed: true, Attributes: map[string]schema.Attribute{"flight_mode": schema.StringAttribute{MarkdownDescription: "The ship's set speed when traveling between waypoints or systems.", Computed: true}, "route": schema.SingleNestedAttribute{MarkdownDescription: "The routing information for the ship's most recent transit or current location.", Computed: true, Attributes: map[string]schema.Attribute{"arrival": schema.StringAttribute{MarkdownDescription: "The date time of the ship's arrival. If the ship is in-transit, this is the expected time of arrival.", Computed: true}, "departure_time": schema.StringAttribute{MarkdownDescription: "The date time of the ship's departure.", Computed: true}, "destination": schema.SingleNestedAttribute{MarkdownDescription: "The destination or departure of a ships nav route.", Computed: true, Attributes: map[string]schema.Attribute{"symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the waypoint.", Computed: true}, "system_symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the system.", Computed: true}, "type": schema.StringAttribute{MarkdownDescription: "The type of waypoint.", Computed: true}, "x": schema.Int64Attribute{MarkdownDescription: "Position in the universe in the x axis.", Computed: true}, "y": schema.Int64Attribute{MarkdownDescription: "Position in the universe in the y axis.", Computed: true}}}, "origin": schema.SingleNestedAttribute{MarkdownDescription: "The destination or departure of a ships nav route.", Computed: true, Attributes: map[string]schema.Attribute{"symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the waypoint.", Computed: true}, "system_symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the system.", Computed: true}, "type": schema.StringAttribute{MarkdownDescription: "The type of waypoint.", Computed: true}, "x": schema.Int64Attribute{MarkdownDescription: "Position in the universe in the x axis.", Computed: true}, "y": schema.Int64Attribute{MarkdownDescription: "Position in the universe in the y axis.", Computed: true}}}}}, "status": schema.StringAttribute{MarkdownDescription: "The current status of the ship", Computed: true}, "system_symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the system.", Computed: true}, "waypoint_symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the waypoint.", Computed: true}}}, "reactor": schema.SingleNestedAttribute{MarkdownDescription: "The reactor of the ship. The reactor is responsible for powering the ship's systems and weapons.", Computed: true, Attributes: map[string]schema.Attribute{"condition": schema.Float64Attribute{MarkdownDescription: "The repairable condition of a component. A value of 0 indicates the component needs significant repairs, while a value of 1 indicates the component is in near perfect condition. As the condition of a component is repaired, the overall integrity of the component decreases.", Computed: true}, "description": schema.StringAttribute{MarkdownDescription: "Description of the reactor.", Computed: true}, "integrity": schema.Float64Attribute{MarkdownDescription: "The overall integrity of the component, which determines the performance of the component. A value of 0 indicates that the component is almost completely degraded, while a value of 1 indicates that the component is in near perfect condition. The integrity of the component is non-repairable, and represents permanent wear over time.", Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "Name of the reactor.", Computed: true}, "power_output": schema.Int64Attribute{MarkdownDescription: "The amount of power provided by this reactor. The more power a reactor provides to the ship, the lower the cooldown it gets when using a module or mount that taxes the ship's power.", Computed: true}, "quality": schema.Float64Attribute{MarkdownDescription: "The overall quality of the component, which determines the quality of the component. High quality components return more ships parts and ship plating when a ship is scrapped. But also require more of these parts to repair. This is transparent to the player, as the parts are bought from/sold to the marketplace.", Computed: true}, "requirements": schema.SingleNestedAttribute{MarkdownDescription: "The requirements for installation on a ship", Computed: true, Attributes: map[string]schema.Attribute{"crew": schema.Int64Attribute{MarkdownDescription: "The number of crew required for operation.", Computed: true}, "power": schema.Int64Attribute{MarkdownDescription: "The amount of power required from the reactor.", Computed: true}, "slots": schema.Int64Attribute{MarkdownDescription: "The number of module slots required for installation.", Computed: true}}}, "symbol": schema.StringAttribute{MarkdownDescription: "Symbol of the reactor.", Computed: true}}}, "registration": schema.SingleNestedAttribute{MarkdownDescription: "The public registration information of the ship", Computed: true, Attributes: map[string]schema.Attribute{"faction_symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the faction the ship is registered with", Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "The agent's registered name of the ship", Computed: true}, "role": schema.StringAttribute{MarkdownDescription: "The registered role of the ship", Computed: true}}}, "symbol": schema.StringAttribute{MarkdownDescription: "The globally unique identifier of the ship in the following format: `[AGENT_SYMBOL]-[HEX_ID]`", Computed: true}}}}, "limit": schema.Int64Attribute{MarkdownDescription: "How many entries to return per page", Optional: true}, "page": schema.Int64Attribute{MarkdownDescription: "What entry offset to request", Optional: true}}}
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
