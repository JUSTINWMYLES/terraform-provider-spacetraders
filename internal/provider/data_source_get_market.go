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
	_ datasource.DataSource              = (*GetMarketDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetMarketDataSource)(nil)
)

// GetMarketDataSource is the generated Terraform data source implementation.
type GetMarketDataSource struct {
	client *client.Client
}

// GetMarketDataSourceModel describes the data source state shape.
type GetMarketDataSourceModel struct {
	Exchange       types.List   `tfsdk:"exchange"`
	Exports        types.List   `tfsdk:"exports"`
	Imports        types.List   `tfsdk:"imports"`
	Symbol         types.String `tfsdk:"symbol"`
	SystemSymbol   types.String `tfsdk:"system_symbol" json:"systemSymbol"`
	TradeGoods     types.List   `tfsdk:"trade_goods" json:"tradeGoods"`
	Transactions   types.List   `tfsdk:"transactions"`
	WaypointSymbol types.String `tfsdk:"waypoint_symbol" json:"waypointSymbol"`
}

// NewGetMarketDataSource returns a new instance of the generated data source.
func NewGetMarketDataSource() datasource.DataSource {
	return &GetMarketDataSource{}
}

// Metadata returns the data source type name.
func (d *GetMarketDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "spacetraders_get_market"
}

// Schema returns the data source schema.
func (d *GetMarketDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Retrieve imports, exports and exchange data from a marketplace. Requires a waypoint that has the `Marketplace` trait to use.\n\nSend a ship to the waypoint to access trade good prices and recent transactions. Refer to the [Market Overview page](https://docs.spacetraders.io/game-concepts/markets) to gain better a understanding of the market in the game.", Attributes: map[string]schema.Attribute{"exchange": schema.ListNestedAttribute{MarkdownDescription: "The list of goods that are bought and sold between agents at this market.", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"description": schema.StringAttribute{MarkdownDescription: "The description of the good.", Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "The name of the good.", Computed: true}, "symbol": schema.StringAttribute{MarkdownDescription: "The good's symbol.", Computed: true}}}}, "exports": schema.ListNestedAttribute{MarkdownDescription: "The list of goods that are exported from this market.", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"description": schema.StringAttribute{MarkdownDescription: "The description of the good.", Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "The name of the good.", Computed: true}, "symbol": schema.StringAttribute{MarkdownDescription: "The good's symbol.", Computed: true}}}}, "imports": schema.ListNestedAttribute{MarkdownDescription: "The list of goods that are sought as imports in this market.", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"description": schema.StringAttribute{MarkdownDescription: "The description of the good.", Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "The name of the good.", Computed: true}, "symbol": schema.StringAttribute{MarkdownDescription: "The good's symbol.", Computed: true}}}}, "symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the market. The symbol is the same as the waypoint where the market is located.", Computed: true}, "system_symbol": schema.StringAttribute{MarkdownDescription: "The system symbol", Required: true}, "trade_goods": schema.ListNestedAttribute{MarkdownDescription: "The list of goods that are traded at this market. Visible only when a ship is present at the market.", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"activity": schema.StringAttribute{MarkdownDescription: "The activity level of a trade good. If the good is an import, this represents how strong consumption is. If the good is an export, this represents how strong the production is for the good. When activity is strong, consumption or production is near maximum capacity. When activity is weak, consumption or production is near minimum capacity.", Computed: true}, "purchase_price": schema.Int64Attribute{MarkdownDescription: "The price at which this good can be purchased from the market.", Computed: true}, "sell_price": schema.Int64Attribute{MarkdownDescription: "The price at which this good can be sold to the market.", Computed: true}, "supply": schema.StringAttribute{MarkdownDescription: "The supply level of a trade good.", Computed: true}, "symbol": schema.StringAttribute{MarkdownDescription: "The good's symbol.", Computed: true}, "trade_volume": schema.Int64Attribute{MarkdownDescription: "This is the maximum number of units that can be purchased or sold at this market in a single trade for this good. Trade volume also gives an indication of price volatility. A market with a low trade volume will have large price swings, while high trade volume will be more resilient to price changes.", Computed: true}, "type": schema.StringAttribute{MarkdownDescription: "The type of trade good (export, import, or exchange).", Computed: true}}}}, "transactions": schema.ListNestedAttribute{MarkdownDescription: "The list of recent transactions at this market. Visible only when a ship is present at the market.", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"price_per_unit": schema.Int64Attribute{MarkdownDescription: "The price per unit of the transaction.", Computed: true}, "ship_symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the ship that made the transaction.", Computed: true}, "timestamp": schema.StringAttribute{MarkdownDescription: "The timestamp of the transaction.", Computed: true}, "total_price": schema.Int64Attribute{MarkdownDescription: "The total price of the transaction.", Computed: true}, "trade_symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the trade good.", Computed: true}, "type": schema.StringAttribute{MarkdownDescription: "The type of transaction.", Computed: true}, "units": schema.Int64Attribute{MarkdownDescription: "The number of units of the transaction.", Computed: true}, "waypoint_symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the waypoint.", Computed: true}}}}, "waypoint_symbol": schema.StringAttribute{MarkdownDescription: "The waypoint symbol", Required: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetMarketDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetMarketDataSourceModel
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
func (d *GetMarketDataSource) readRemote(ctx context.Context, config *GetMarketDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/systems/{systemSymbol}/waypoints/{waypointSymbol}/market"
	reqPath = strings.ReplaceAll(reqPath, "{systemSymbol}", url.PathEscape(config.SystemSymbol.ValueString()))
	reqPath = strings.ReplaceAll(reqPath, "{waypointSymbol}", url.PathEscape(config.WaypointSymbol.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_get_market", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_get_market", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading spacetraders_get_market", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		apiErr, err := client.NewAPIError(httpResp)
		if err != nil {
			resp.Diagnostics.AddError("Error reading spacetraders_get_market", fmt.Sprintf("Could not read error response: %s", err))
			return
		}
		resp.Diagnostics.AddError("Error reading spacetraders_get_market", apiErr.Error())
		return
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading spacetraders_get_market", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["data"]; ok {
		if m, ok := v.(map[string]any); ok {
			data = m
		}
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_get_market", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetMarketDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
