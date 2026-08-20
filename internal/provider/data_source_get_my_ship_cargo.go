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
	_ datasource.DataSource              = (*GetMyShipCargoDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetMyShipCargoDataSource)(nil)
)

// GetMyShipCargoDataSource is the generated Terraform data source implementation.
type GetMyShipCargoDataSource struct {
	client *client.Client
}

// GetMyShipCargoDataSourceModel describes the data source state shape.
type GetMyShipCargoDataSourceModel struct {
	Capacity   types.Int64  `tfsdk:"capacity"`
	Inventory  types.List   `tfsdk:"inventory"`
	ShipSymbol types.String `tfsdk:"ship_symbol" json:"shipSymbol"`
	Units      types.Int64  `tfsdk:"units"`
}

// NewGetMyShipCargoDataSource returns a new instance of the generated data source.
func NewGetMyShipCargoDataSource() datasource.DataSource {
	return &GetMyShipCargoDataSource{}
}

// Metadata returns the data source type name.
func (d *GetMyShipCargoDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "spacetraders_get_my_ship_cargo"
}

// Schema returns the data source schema.
func (d *GetMyShipCargoDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Retrieve the cargo of a ship under your agent's ownership.", Attributes: map[string]schema.Attribute{"capacity": schema.Int64Attribute{MarkdownDescription: "The max number of items that can be stored in the cargo hold.", Computed: true}, "inventory": schema.ListNestedAttribute{MarkdownDescription: "The items currently in the cargo hold.", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"description": schema.StringAttribute{MarkdownDescription: "The description of the cargo item type.", Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "The name of the cargo item type.", Computed: true}, "symbol": schema.StringAttribute{MarkdownDescription: "The good's symbol.", Computed: true}, "units": schema.Int64Attribute{MarkdownDescription: "The number of units of the cargo item.", Computed: true}}}}, "ship_symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the ship.", Required: true}, "units": schema.Int64Attribute{MarkdownDescription: "The number of items currently stored in the cargo hold.", Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetMyShipCargoDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetMyShipCargoDataSourceModel
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
func (d *GetMyShipCargoDataSource) readRemote(ctx context.Context, config *GetMyShipCargoDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/my/ships/{shipSymbol}/cargo"
	reqPath = strings.ReplaceAll(reqPath, "{shipSymbol}", url.PathEscape(config.ShipSymbol.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil, client.WithSchemes("AgentToken"))
	if err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_get_my_ship_cargo", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_get_my_ship_cargo", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading spacetraders_get_my_ship_cargo", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		apiErr, err := client.NewAPIError(httpResp)
		if err != nil {
			resp.Diagnostics.AddError("Error reading spacetraders_get_my_ship_cargo", fmt.Sprintf("Could not read error response: %s", err))
			return
		}
		resp.Diagnostics.AddError("Error reading spacetraders_get_my_ship_cargo", apiErr.Error())
		return
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading spacetraders_get_my_ship_cargo", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["data"]; ok {
		if m, ok := v.(map[string]any); ok {
			data = m
		}
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_get_my_ship_cargo", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetMyShipCargoDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
