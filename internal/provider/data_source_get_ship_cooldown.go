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
	_ datasource.DataSource              = (*GetShipCooldownDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetShipCooldownDataSource)(nil)
)

// GetShipCooldownDataSource is the generated Terraform data source implementation.
type GetShipCooldownDataSource struct {
	client *client.Client
}

// GetShipCooldownDataSourceModel describes the data source state shape.
type GetShipCooldownDataSourceModel struct {
	Expiration       types.String `tfsdk:"expiration"`
	RemainingSeconds types.Int64  `tfsdk:"remaining_seconds" json:"remainingSeconds"`
	ShipSymbol       types.String `tfsdk:"ship_symbol" json:"shipSymbol"`
	TotalSeconds     types.Int64  `tfsdk:"total_seconds" json:"totalSeconds"`
}

// NewGetShipCooldownDataSource returns a new instance of the generated data source.
func NewGetShipCooldownDataSource() datasource.DataSource {
	return &GetShipCooldownDataSource{}
}

// Metadata returns the data source type name.
func (d *GetShipCooldownDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "spacetraders_get_ship_cooldown"
}

// Schema returns the data source schema.
func (d *GetShipCooldownDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Retrieve the details of your ship's reactor cooldown. Some actions such as activating your jump drive, scanning, or extracting resources taxes your reactor and results in a cooldown.\n\nYour ship cannot perform additional actions until your cooldown has expired. The duration of your cooldown is relative to the power consumption of the related modules or mounts for the action taken.\n\nResponse returns a 204 status code (no-content) when the ship has no cooldown.", Attributes: map[string]schema.Attribute{"expiration": schema.StringAttribute{MarkdownDescription: "The date and time when the cooldown expires in ISO 8601 format", Computed: true}, "remaining_seconds": schema.Int64Attribute{MarkdownDescription: "The remaining duration of the cooldown in seconds", Computed: true}, "ship_symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the ship that is on cooldown", Required: true}, "total_seconds": schema.Int64Attribute{MarkdownDescription: "The total duration of the cooldown in seconds", Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetShipCooldownDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetShipCooldownDataSourceModel
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
func (d *GetShipCooldownDataSource) readRemote(ctx context.Context, config *GetShipCooldownDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/my/ships/{shipSymbol}/cooldown"
	reqPath = strings.ReplaceAll(reqPath, "{shipSymbol}", url.PathEscape(config.ShipSymbol.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil, client.WithSchemes("AgentToken"))
	if err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_get_ship_cooldown", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_get_ship_cooldown", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading spacetraders_get_ship_cooldown", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200 || httpResp.StatusCode == 204) {
		apiErr, err := client.NewAPIError(httpResp)
		if err != nil {
			resp.Diagnostics.AddError("Error reading spacetraders_get_ship_cooldown", fmt.Sprintf("Could not read error response: %s", err))
			return
		}
		resp.Diagnostics.AddError("Error reading spacetraders_get_ship_cooldown", apiErr.Error())
		return
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading spacetraders_get_ship_cooldown", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["data"]; ok {
		if m, ok := v.(map[string]any); ok {
			data = m
		}
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_get_ship_cooldown", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetShipCooldownDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
