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
	_ datasource.DataSource              = (*GetConstructionDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetConstructionDataSource)(nil)
)

// GetConstructionDataSource is the generated Terraform data source implementation.
type GetConstructionDataSource struct {
	client *client.Client
}

// GetConstructionDataSourceModel describes the data source state shape.
type GetConstructionDataSourceModel struct {
	IsComplete     types.Bool   `tfsdk:"is_complete" json:"isComplete"`
	Materials      types.List   `tfsdk:"materials"`
	Symbol         types.String `tfsdk:"symbol"`
	SystemSymbol   types.String `tfsdk:"system_symbol" json:"systemSymbol"`
	WaypointSymbol types.String `tfsdk:"waypoint_symbol" json:"waypointSymbol"`
}

// NewGetConstructionDataSource returns a new instance of the generated data source.
func NewGetConstructionDataSource() datasource.DataSource {
	return &GetConstructionDataSource{}
}

// Metadata returns the data source type name.
func (d *GetConstructionDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "spacetraders_get_construction"
}

// Schema returns the data source schema.
func (d *GetConstructionDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get construction details for a waypoint. Requires a waypoint with a property of `isUnderConstruction` to be true.", Attributes: map[string]schema.Attribute{"is_complete": schema.BoolAttribute{MarkdownDescription: "Whether the waypoint has been constructed.", Computed: true}, "materials": schema.ListNestedAttribute{MarkdownDescription: "The materials required to construct the waypoint.", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"fulfilled": schema.Int64Attribute{MarkdownDescription: "The number of units fulfilled toward the required amount.", Computed: true}, "required": schema.Int64Attribute{MarkdownDescription: "The number of units required.", Computed: true}, "trade_symbol": schema.StringAttribute{MarkdownDescription: "The good's symbol.", Computed: true}}}}, "symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the waypoint.", Computed: true}, "system_symbol": schema.StringAttribute{MarkdownDescription: "The system symbol", Required: true}, "waypoint_symbol": schema.StringAttribute{MarkdownDescription: "The waypoint symbol", Required: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetConstructionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetConstructionDataSourceModel
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
func (d *GetConstructionDataSource) readRemote(ctx context.Context, config *GetConstructionDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/systems/{systemSymbol}/waypoints/{waypointSymbol}/construction"
	reqPath = strings.ReplaceAll(reqPath, "{systemSymbol}", url.PathEscape(config.SystemSymbol.ValueString()))
	reqPath = strings.ReplaceAll(reqPath, "{waypointSymbol}", url.PathEscape(config.WaypointSymbol.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_get_construction", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_get_construction", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading spacetraders_get_construction", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		apiErr, err := client.NewAPIError(httpResp)
		if err != nil {
			resp.Diagnostics.AddError("Error reading spacetraders_get_construction", fmt.Sprintf("Could not read error response: %s", err))
			return
		}
		resp.Diagnostics.AddError("Error reading spacetraders_get_construction", apiErr.Error())
		return
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading spacetraders_get_construction", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["data"]; ok {
		if m, ok := v.(map[string]any); ok {
			data = m
		}
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_get_construction", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetConstructionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
