package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)
import (
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
	client "github.com/spacetraders/terraform-provider-spacetraders/internal/client"
)

// Compile-time interface assertions.
var (
	_ resource.Resource              = (*ShipResource)(nil)
	_ resource.ResourceWithConfigure = (*ShipResource)(nil)
)

// ShipResource is the generated Terraform managed resource implementation.
type ShipResource struct {
	client *client.Client
}

// ShipResourceModel describes the Terraform state and plan shape for ShipResource.
type ShipResourceModel struct {
	Cargo          types.Object `tfsdk:"cargo"`
	Cooldown       types.Object `tfsdk:"cooldown"`
	Crew           types.Object `tfsdk:"crew"`
	Engine         types.Object `tfsdk:"engine"`
	Frame          types.Object `tfsdk:"frame"`
	Fuel           types.Object `tfsdk:"fuel"`
	Id             types.String `tfsdk:"id"`
	Modules        types.List   `tfsdk:"modules"`
	Mounts         types.List   `tfsdk:"mounts"`
	Nav            types.Object `tfsdk:"nav"`
	Reactor        types.Object `tfsdk:"reactor"`
	Registration   types.Object `tfsdk:"registration"`
	ShipType       types.String `tfsdk:"ship_type" json:"shipType"`
	Symbol         types.String `tfsdk:"symbol"`
	WaypointSymbol types.String `tfsdk:"waypoint_symbol" json:"waypointSymbol"`
}

// Metadata returns the resource type name.
func (r *ShipResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "spacetraders_ship"
}

// Schema returns the Terraform schema for this resource.
func (r *ShipResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Retrieve the details of a ship under your agent's ownership.", Attributes: map[string]schema.Attribute{"cargo": schema.SingleNestedAttribute{MarkdownDescription: "Ship cargo details.", Computed: true, Attributes: map[string]schema.Attribute{"capacity": schema.Int64Attribute{MarkdownDescription: "The max number of items that can be stored in the cargo hold.", Computed: true}, "inventory": schema.ListNestedAttribute{MarkdownDescription: "The items currently in the cargo hold.", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"description": schema.StringAttribute{MarkdownDescription: "The description of the cargo item type.", Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "The name of the cargo item type.", Computed: true}, "symbol": schema.StringAttribute{MarkdownDescription: "The good's symbol.", Computed: true}, "units": schema.Int64Attribute{MarkdownDescription: "The number of units of the cargo item.", Computed: true}}}}, "units": schema.Int64Attribute{MarkdownDescription: "The number of items currently stored in the cargo hold.", Computed: true}}}, "cooldown": schema.SingleNestedAttribute{MarkdownDescription: "A cooldown is a period of time in which a ship cannot perform certain actions.", Computed: true, Attributes: map[string]schema.Attribute{"expiration": schema.StringAttribute{MarkdownDescription: "The date and time when the cooldown expires in ISO 8601 format", Computed: true}, "remaining_seconds": schema.Int64Attribute{MarkdownDescription: "The remaining duration of the cooldown in seconds", Computed: true}, "ship_symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the ship that is on cooldown", Computed: true}, "total_seconds": schema.Int64Attribute{MarkdownDescription: "The total duration of the cooldown in seconds", Computed: true}}}, "crew": schema.SingleNestedAttribute{MarkdownDescription: "The ship's crew service and maintain the ship's systems and equipment.", Computed: true, Attributes: map[string]schema.Attribute{"capacity": schema.Int64Attribute{MarkdownDescription: "The maximum number of crew members the ship can support.", Computed: true}, "current": schema.Int64Attribute{MarkdownDescription: "The current number of crew members on the ship.", Computed: true}, "morale": schema.Int64Attribute{MarkdownDescription: "A rough measure of the crew's morale. A higher morale means the crew is happier and more productive. A lower morale means the ship is more prone to accidents.", Computed: true}, "required": schema.Int64Attribute{MarkdownDescription: "The minimum number of crew members required to maintain the ship.", Computed: true}, "rotation": schema.StringAttribute{MarkdownDescription: "The rotation of crew shifts. A stricter shift improves the ship's performance. A more relaxed shift improves the crew's morale.", Computed: true}, "wages": schema.Int64Attribute{MarkdownDescription: "The amount of credits per crew member paid per hour. Wages are paid when a ship docks at a civilized waypoint.", Computed: true}}}, "engine": schema.SingleNestedAttribute{MarkdownDescription: "The engine determines how quickly a ship travels between waypoints.", Computed: true, Attributes: map[string]schema.Attribute{"condition": schema.Float64Attribute{MarkdownDescription: "The repairable condition of a component. A value of 0 indicates the component needs significant repairs, while a value of 1 indicates the component is in near perfect condition. As the condition of a component is repaired, the overall integrity of the component decreases.", Computed: true}, "description": schema.StringAttribute{MarkdownDescription: "The description of the engine.", Computed: true}, "integrity": schema.Float64Attribute{MarkdownDescription: "The overall integrity of the component, which determines the performance of the component. A value of 0 indicates that the component is almost completely degraded, while a value of 1 indicates that the component is in near perfect condition. The integrity of the component is non-repairable, and represents permanent wear over time.", Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "The name of the engine.", Computed: true}, "quality": schema.Float64Attribute{MarkdownDescription: "The overall quality of the component, which determines the quality of the component. High quality components return more ships parts and ship plating when a ship is scrapped. But also require more of these parts to repair. This is transparent to the player, as the parts are bought from/sold to the marketplace.", Computed: true}, "requirements": schema.SingleNestedAttribute{MarkdownDescription: "The requirements for installation on a ship", Computed: true, Attributes: map[string]schema.Attribute{"crew": schema.Int64Attribute{MarkdownDescription: "The number of crew required for operation.", Computed: true}, "power": schema.Int64Attribute{MarkdownDescription: "The amount of power required from the reactor.", Computed: true}, "slots": schema.Int64Attribute{MarkdownDescription: "The number of module slots required for installation.", Computed: true}}}, "speed": schema.Int64Attribute{MarkdownDescription: "The speed stat of this engine. The higher the speed, the faster a ship can travel from one point to another. Reduces the time of arrival when navigating the ship.", Computed: true}, "symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the engine.", Computed: true}}}, "frame": schema.SingleNestedAttribute{MarkdownDescription: "The frame of the ship. The frame determines the number of modules and mounting points of the ship, as well as base fuel capacity. As the condition of the frame takes more wear, the ship will become more sluggish and less maneuverable.", Computed: true, Attributes: map[string]schema.Attribute{"condition": schema.Float64Attribute{MarkdownDescription: "The repairable condition of a component. A value of 0 indicates the component needs significant repairs, while a value of 1 indicates the component is in near perfect condition. As the condition of a component is repaired, the overall integrity of the component decreases.", Computed: true}, "description": schema.StringAttribute{MarkdownDescription: "Description of the frame.", Computed: true}, "fuel_capacity": schema.Int64Attribute{MarkdownDescription: "The maximum amount of fuel that can be stored in this ship. When refueling, the ship will be refueled to this amount.", Computed: true}, "integrity": schema.Float64Attribute{MarkdownDescription: "The overall integrity of the component, which determines the performance of the component. A value of 0 indicates that the component is almost completely degraded, while a value of 1 indicates that the component is in near perfect condition. The integrity of the component is non-repairable, and represents permanent wear over time.", Computed: true}, "module_slots": schema.Int64Attribute{MarkdownDescription: "The amount of slots that can be dedicated to modules installed in the ship. Each installed module take up a number of slots, and once there are no more slots, no new modules can be installed.", Computed: true}, "mounting_points": schema.Int64Attribute{MarkdownDescription: "The amount of slots that can be dedicated to mounts installed in the ship. Each installed mount takes up a number of points, and once there are no more points remaining, no new mounts can be installed.", Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "Name of the frame.", Computed: true}, "quality": schema.Float64Attribute{MarkdownDescription: "The overall quality of the component, which determines the quality of the component. High quality components return more ships parts and ship plating when a ship is scrapped. But also require more of these parts to repair. This is transparent to the player, as the parts are bought from/sold to the marketplace.", Computed: true}, "requirements": schema.SingleNestedAttribute{MarkdownDescription: "The requirements for installation on a ship", Computed: true, Attributes: map[string]schema.Attribute{"crew": schema.Int64Attribute{MarkdownDescription: "The number of crew required for operation.", Computed: true}, "power": schema.Int64Attribute{MarkdownDescription: "The amount of power required from the reactor.", Computed: true}, "slots": schema.Int64Attribute{MarkdownDescription: "The number of module slots required for installation.", Computed: true}}}, "symbol": schema.StringAttribute{MarkdownDescription: "Symbol of the frame.", Computed: true}}}, "fuel": schema.SingleNestedAttribute{MarkdownDescription: "Details of the ship's fuel tanks including how much fuel was consumed during the last transit or action.", Computed: true, Attributes: map[string]schema.Attribute{"capacity": schema.Int64Attribute{MarkdownDescription: "The maximum amount of fuel the ship's tanks can hold.", Computed: true}, "consumed": schema.SingleNestedAttribute{MarkdownDescription: "An object that only shows up when an action has consumed fuel in the process. Shows the fuel consumption data.", Computed: true, Attributes: map[string]schema.Attribute{"amount": schema.Int64Attribute{MarkdownDescription: "The amount of fuel consumed by the most recent transit or action.", Computed: true}, "timestamp": schema.StringAttribute{MarkdownDescription: "The time at which the fuel was consumed.", Computed: true}}}, "current": schema.Int64Attribute{MarkdownDescription: "The current amount of fuel in the ship's tanks.", Computed: true}}}, "id": schema.StringAttribute{Computed: true}, "modules": schema.ListNestedAttribute{MarkdownDescription: "Modules installed in this ship.", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"capacity": schema.Int64Attribute{MarkdownDescription: "Modules that provide capacity, such as cargo hold or crew quarters will show this value to denote how much of a bonus the module grants.", Computed: true}, "description": schema.StringAttribute{MarkdownDescription: "Description of this module.", Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "Name of this module.", Computed: true}, "range": schema.Int64Attribute{MarkdownDescription: "Modules that have a range will such as a sensor array show this value to denote how far can the module reach with its capabilities.", Computed: true}, "requirements": schema.SingleNestedAttribute{MarkdownDescription: "The requirements for installation on a ship", Computed: true, Attributes: map[string]schema.Attribute{"crew": schema.Int64Attribute{MarkdownDescription: "The number of crew required for operation.", Computed: true}, "power": schema.Int64Attribute{MarkdownDescription: "The amount of power required from the reactor.", Computed: true}, "slots": schema.Int64Attribute{MarkdownDescription: "The number of module slots required for installation.", Computed: true}}}, "symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the module.", Computed: true}}}}, "mounts": schema.ListNestedAttribute{MarkdownDescription: "Mounts installed in this ship.", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"deposits": schema.ListAttribute{MarkdownDescription: "Mounts that have this value denote what goods can be produced from using the mount.", Computed: true, ElementType: types.StringType}, "description": schema.StringAttribute{MarkdownDescription: "Description of this mount.", Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "Name of this mount.", Computed: true}, "requirements": schema.SingleNestedAttribute{MarkdownDescription: "The requirements for installation on a ship", Computed: true, Attributes: map[string]schema.Attribute{"crew": schema.Int64Attribute{MarkdownDescription: "The number of crew required for operation.", Computed: true}, "power": schema.Int64Attribute{MarkdownDescription: "The amount of power required from the reactor.", Computed: true}, "slots": schema.Int64Attribute{MarkdownDescription: "The number of module slots required for installation.", Computed: true}}}, "strength": schema.Int64Attribute{MarkdownDescription: "Mounts that have this value, such as mining lasers, denote how powerful this mount's capabilities are.", Computed: true}, "symbol": schema.StringAttribute{MarkdownDescription: "Symbol of this mount.", Computed: true}}}}, "nav": schema.SingleNestedAttribute{MarkdownDescription: "The navigation information of the ship.", Computed: true, Attributes: map[string]schema.Attribute{"flight_mode": schema.StringAttribute{MarkdownDescription: "The ship's set speed when traveling between waypoints or systems.", Computed: true}, "route": schema.SingleNestedAttribute{MarkdownDescription: "The routing information for the ship's most recent transit or current location.", Computed: true, Attributes: map[string]schema.Attribute{"arrival": schema.StringAttribute{MarkdownDescription: "The date time of the ship's arrival. If the ship is in-transit, this is the expected time of arrival.", Computed: true}, "departure_time": schema.StringAttribute{MarkdownDescription: "The date time of the ship's departure.", Computed: true}, "destination": schema.SingleNestedAttribute{MarkdownDescription: "The destination or departure of a ships nav route.", Computed: true, Attributes: map[string]schema.Attribute{"symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the waypoint.", Computed: true}, "system_symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the system.", Computed: true}, "type": schema.StringAttribute{MarkdownDescription: "The type of waypoint.", Computed: true}, "x": schema.Int64Attribute{MarkdownDescription: "Position in the universe in the x axis.", Computed: true}, "y": schema.Int64Attribute{MarkdownDescription: "Position in the universe in the y axis.", Computed: true}}}, "origin": schema.SingleNestedAttribute{MarkdownDescription: "The destination or departure of a ships nav route.", Computed: true, Attributes: map[string]schema.Attribute{"symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the waypoint.", Computed: true}, "system_symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the system.", Computed: true}, "type": schema.StringAttribute{MarkdownDescription: "The type of waypoint.", Computed: true}, "x": schema.Int64Attribute{MarkdownDescription: "Position in the universe in the x axis.", Computed: true}, "y": schema.Int64Attribute{MarkdownDescription: "Position in the universe in the y axis.", Computed: true}}}}}, "status": schema.StringAttribute{MarkdownDescription: "The current status of the ship", Computed: true}, "system_symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the system.", Computed: true}, "waypoint_symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the waypoint.", Computed: true}}}, "reactor": schema.SingleNestedAttribute{MarkdownDescription: "The reactor of the ship. The reactor is responsible for powering the ship's systems and weapons.", Computed: true, Attributes: map[string]schema.Attribute{"condition": schema.Float64Attribute{MarkdownDescription: "The repairable condition of a component. A value of 0 indicates the component needs significant repairs, while a value of 1 indicates the component is in near perfect condition. As the condition of a component is repaired, the overall integrity of the component decreases.", Computed: true}, "description": schema.StringAttribute{MarkdownDescription: "Description of the reactor.", Computed: true}, "integrity": schema.Float64Attribute{MarkdownDescription: "The overall integrity of the component, which determines the performance of the component. A value of 0 indicates that the component is almost completely degraded, while a value of 1 indicates that the component is in near perfect condition. The integrity of the component is non-repairable, and represents permanent wear over time.", Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "Name of the reactor.", Computed: true}, "power_output": schema.Int64Attribute{MarkdownDescription: "The amount of power provided by this reactor. The more power a reactor provides to the ship, the lower the cooldown it gets when using a module or mount that taxes the ship's power.", Computed: true}, "quality": schema.Float64Attribute{MarkdownDescription: "The overall quality of the component, which determines the quality of the component. High quality components return more ships parts and ship plating when a ship is scrapped. But also require more of these parts to repair. This is transparent to the player, as the parts are bought from/sold to the marketplace.", Computed: true}, "requirements": schema.SingleNestedAttribute{MarkdownDescription: "The requirements for installation on a ship", Computed: true, Attributes: map[string]schema.Attribute{"crew": schema.Int64Attribute{MarkdownDescription: "The number of crew required for operation.", Computed: true}, "power": schema.Int64Attribute{MarkdownDescription: "The amount of power required from the reactor.", Computed: true}, "slots": schema.Int64Attribute{MarkdownDescription: "The number of module slots required for installation.", Computed: true}}}, "symbol": schema.StringAttribute{MarkdownDescription: "Symbol of the reactor.", Computed: true}}}, "registration": schema.SingleNestedAttribute{MarkdownDescription: "The public registration information of the ship", Computed: true, Attributes: map[string]schema.Attribute{"faction_symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the faction the ship is registered with", Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "The agent's registered name of the ship", Computed: true}, "role": schema.StringAttribute{MarkdownDescription: "The registered role of the ship", Computed: true}}}, "ship_type": schema.StringAttribute{MarkdownDescription: "Type of ship", Required: true}, "symbol": schema.StringAttribute{MarkdownDescription: "The globally unique identifier of the ship in the following format: `[AGENT_SYMBOL]-[HEX_ID]`", Computed: true}, "waypoint_symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the waypoint you want to purchase the ship at.", Required: true}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *ShipResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ShipResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.createRemote(ctx, &plan, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// createRemote performs the create HTTP exchange and decodes the response into plan. Extracted from Create so the request/response logic is unit-testable without a tfsdk.Plan.
func (r *ShipResource) createRemote(ctx context.Context, plan *ShipResourceModel, resp *resource.CreateResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	body, err := modelToJSONMap(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Error creating spacetraders_ship", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error creating spacetraders_ship", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	reqPath := "/my/ships"
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload), client.WithSchemes("AgentToken"))
	if err != nil {
		resp.Diagnostics.AddError("Error creating spacetraders_ship", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error creating spacetraders_ship", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		apiErr, err := client.NewAPIError(httpResp)
		if err != nil {
			resp.Diagnostics.AddError("Error creating spacetraders_ship", fmt.Sprintf("Could not read error response: %s", err))
			return
		}
		resp.Diagnostics.AddError("Error creating spacetraders_ship", apiErr.Error())
		return
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error creating spacetraders_ship", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["data"]; ok {
		if m, ok := v.(map[string]any); ok {
			data = m
		}
	}
	if inner, ok := data["ship"]; ok {
		if im, ok := inner.(map[string]any); ok {
			data = im
		}
	}
	err = applyJSONToModel(&plan, data)
	if err != nil {
		resp.Diagnostics.AddError("Error creating spacetraders_ship", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
	if plan.Symbol.IsNull() || plan.Symbol.IsUnknown() {
		loc := httpResp.Header.Get("Location")
		if loc != "" {
			plan.Symbol = types.StringValue(loc)
		} else {
			resp.Diagnostics.AddError("Error creating spacetraders_ship", "The create response did not contain an identifier and no Location header was returned, so the resource cannot be tracked in state.")
			return
		}
	}
}

// Read refreshes the Terraform state with the latest remote values.
func (r *ShipResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ShipResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if r.readRemote(ctx, &state, resp) {
		resp.State.RemoveResource(ctx)
		return
	}
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// readRemote performs the read HTTP exchange and decodes the response into state, returning removed=true when the API reports 404. Extracted from Read so the request/response logic is unit-testable without a tfsdk.State.
func (r *ShipResource) readRemote(ctx context.Context, state *ShipResourceModel, resp *resource.ReadResponse) (removed bool) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/my/ships/{shipSymbol}"
	reqPath = strings.ReplaceAll(reqPath, "{shipSymbol}", url.PathEscape(state.Symbol.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodGet, reqPath, nil, client.WithSchemes("AgentToken"))
	if err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_ship", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_ship", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		removed = true
		return
	}
	if !(httpResp.StatusCode == 200) {
		apiErr, err := client.NewAPIError(httpResp)
		if err != nil {
			resp.Diagnostics.AddError("Error reading spacetraders_ship", fmt.Sprintf("Could not read error response: %s", err))
			return
		}
		resp.Diagnostics.AddError("Error reading spacetraders_ship", apiErr.Error())
		return
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading spacetraders_ship", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["data"]; ok {
		if m, ok := v.(map[string]any); ok {
			data = m
		}
	}
	err = applyJSONToModel(&state, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_ship", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
	return
}

// Update modifies the remote resource to match the desired plan.
func (r *ShipResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ShipResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state ShipResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if plan.Symbol.IsNull() || plan.Symbol.IsUnknown() {
		if !state.Symbol.IsNull() && !state.Symbol.IsUnknown() {
			plan.Symbol = state.Symbol
		}
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Update is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Delete destroys the remote resource.
func (r *ShipResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ShipResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.deleteRemote(ctx, &state, resp)
}

// deleteRemote performs the delete HTTP exchange, treating a 404 as already deleted. Extracted from Delete so the request/response logic is unit-testable without a tfsdk.State.
func (r *ShipResource) deleteRemote(ctx context.Context, state *ShipResourceModel, resp *resource.DeleteResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/my/ships/{shipSymbol}/scrap"
	reqPath = strings.ReplaceAll(reqPath, "{shipSymbol}", url.PathEscape(state.Symbol.ValueString()))
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, nil, client.WithSchemes("AgentToken"))
	if err != nil {
		resp.Diagnostics.AddError("Error deleting spacetraders_ship", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting spacetraders_ship", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		return
	}
	if !(httpResp.StatusCode == 200) {
		apiErr, err := client.NewAPIError(httpResp)
		if err != nil {
			resp.Diagnostics.AddError("Error deleting spacetraders_ship", fmt.Sprintf("Could not read error response: %s", err))
			return
		}
		resp.Diagnostics.AddError("Error deleting spacetraders_ship", apiErr.Error())
		return
	}
}

// Configure stores the API client supplied by the provider.
func (r *ShipResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type", fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData))
		return
	}
	r.client = c
}
