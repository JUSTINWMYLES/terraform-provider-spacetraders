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
	resp.Schema = schema.Schema{MarkdownDescription: "Retrieve the details of a ship under your agent's ownership.", Attributes: map[string]schema.Attribute{"cargo": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"capacity": schema.Int64Attribute{Computed: true}, "inventory": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"description": schema.StringAttribute{Computed: true}, "name": schema.StringAttribute{Computed: true}, "symbol": schema.StringAttribute{Computed: true}, "units": schema.Int64Attribute{Computed: true}}}}, "units": schema.Int64Attribute{Computed: true}}}, "cooldown": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"expiration": schema.StringAttribute{Computed: true}, "remaining_seconds": schema.Int64Attribute{Computed: true}, "ship_symbol": schema.StringAttribute{Computed: true}, "total_seconds": schema.Int64Attribute{Computed: true}}}, "crew": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"capacity": schema.Int64Attribute{Computed: true}, "current": schema.Int64Attribute{Computed: true}, "morale": schema.Int64Attribute{Computed: true}, "required": schema.Int64Attribute{Computed: true}, "rotation": schema.StringAttribute{Computed: true}, "wages": schema.Int64Attribute{Computed: true}}}, "engine": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"condition": schema.Float64Attribute{Computed: true}, "description": schema.StringAttribute{Computed: true}, "integrity": schema.Float64Attribute{Computed: true}, "name": schema.StringAttribute{Computed: true}, "quality": schema.Float64Attribute{Computed: true}, "requirements": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"crew": schema.Int64Attribute{Computed: true}, "power": schema.Int64Attribute{Computed: true}, "slots": schema.Int64Attribute{Computed: true}}}, "speed": schema.Int64Attribute{Computed: true}, "symbol": schema.StringAttribute{Computed: true}}}, "frame": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"condition": schema.Float64Attribute{Computed: true}, "description": schema.StringAttribute{Computed: true}, "fuel_capacity": schema.Int64Attribute{Computed: true}, "integrity": schema.Float64Attribute{Computed: true}, "module_slots": schema.Int64Attribute{Computed: true}, "mounting_points": schema.Int64Attribute{Computed: true}, "name": schema.StringAttribute{Computed: true}, "quality": schema.Float64Attribute{Computed: true}, "requirements": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"crew": schema.Int64Attribute{Computed: true}, "power": schema.Int64Attribute{Computed: true}, "slots": schema.Int64Attribute{Computed: true}}}, "symbol": schema.StringAttribute{Computed: true}}}, "fuel": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"capacity": schema.Int64Attribute{Computed: true}, "consumed": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"amount": schema.Int64Attribute{Computed: true}, "timestamp": schema.StringAttribute{Computed: true}}}, "current": schema.Int64Attribute{Computed: true}}}, "id": schema.StringAttribute{Computed: true}, "modules": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"capacity": schema.Int64Attribute{Computed: true}, "description": schema.StringAttribute{Computed: true}, "name": schema.StringAttribute{Computed: true}, "range": schema.Int64Attribute{Computed: true}, "requirements": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"crew": schema.Int64Attribute{Computed: true}, "power": schema.Int64Attribute{Computed: true}, "slots": schema.Int64Attribute{Computed: true}}}, "symbol": schema.StringAttribute{Computed: true}}}}, "mounts": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"deposits": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "description": schema.StringAttribute{Computed: true}, "name": schema.StringAttribute{Computed: true}, "requirements": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"crew": schema.Int64Attribute{Computed: true}, "power": schema.Int64Attribute{Computed: true}, "slots": schema.Int64Attribute{Computed: true}}}, "strength": schema.Int64Attribute{Computed: true}, "symbol": schema.StringAttribute{Computed: true}}}}, "nav": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"flight_mode": schema.StringAttribute{Computed: true}, "route": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"arrival": schema.StringAttribute{Computed: true}, "departure_time": schema.StringAttribute{Computed: true}, "destination": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"symbol": schema.StringAttribute{Computed: true}, "system_symbol": schema.StringAttribute{Computed: true}, "type": schema.StringAttribute{Computed: true}, "x": schema.Int64Attribute{Computed: true}, "y": schema.Int64Attribute{Computed: true}}}, "origin": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"symbol": schema.StringAttribute{Computed: true}, "system_symbol": schema.StringAttribute{Computed: true}, "type": schema.StringAttribute{Computed: true}, "x": schema.Int64Attribute{Computed: true}, "y": schema.Int64Attribute{Computed: true}}}}}, "status": schema.StringAttribute{Computed: true}, "system_symbol": schema.StringAttribute{Computed: true}, "waypoint_symbol": schema.StringAttribute{Computed: true}}}, "reactor": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"condition": schema.Float64Attribute{Computed: true}, "description": schema.StringAttribute{Computed: true}, "integrity": schema.Float64Attribute{Computed: true}, "name": schema.StringAttribute{Computed: true}, "power_output": schema.Int64Attribute{Computed: true}, "quality": schema.Float64Attribute{Computed: true}, "requirements": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"crew": schema.Int64Attribute{Computed: true}, "power": schema.Int64Attribute{Computed: true}, "slots": schema.Int64Attribute{Computed: true}}}, "symbol": schema.StringAttribute{Computed: true}}}, "registration": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"faction_symbol": schema.StringAttribute{Computed: true}, "name": schema.StringAttribute{Computed: true}, "role": schema.StringAttribute{Computed: true}}}, "ship_type": schema.StringAttribute{Required: true}, "symbol": schema.StringAttribute{Computed: true}, "waypoint_symbol": schema.StringAttribute{Required: true}}}
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
