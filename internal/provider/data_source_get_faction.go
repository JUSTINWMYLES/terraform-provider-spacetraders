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
	_ datasource.DataSource              = (*GetFactionDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetFactionDataSource)(nil)
)

// GetFactionDataSource is the generated Terraform data source implementation.
type GetFactionDataSource struct {
	client *client.Client
}

// GetFactionDataSourceModel describes the data source state shape.
type GetFactionDataSourceModel struct {
	Description   types.String `tfsdk:"description"`
	FactionSymbol types.String `tfsdk:"faction_symbol" json:"factionSymbol"`
	Headquarters  types.String `tfsdk:"headquarters"`
	IsRecruiting  types.Bool   `tfsdk:"is_recruiting" json:"isRecruiting"`
	Name          types.String `tfsdk:"name"`
	Symbol        types.String `tfsdk:"symbol"`
	Traits        types.List   `tfsdk:"traits"`
}

// NewGetFactionDataSource returns a new instance of the generated data source.
func NewGetFactionDataSource() datasource.DataSource {
	return &GetFactionDataSource{}
}

// Metadata returns the data source type name.
func (d *GetFactionDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "spacetraders_get_faction"
}

// Schema returns the data source schema.
func (d *GetFactionDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "View the details of a faction.", Attributes: map[string]schema.Attribute{"description": schema.StringAttribute{MarkdownDescription: "Description of the faction.", Computed: true}, "faction_symbol": schema.StringAttribute{MarkdownDescription: "The faction symbol", Required: true}, "headquarters": schema.StringAttribute{MarkdownDescription: "The waypoint in which the faction's HQ is located in.", Computed: true}, "is_recruiting": schema.BoolAttribute{MarkdownDescription: "Whether or not the faction is currently recruiting new agents.", Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "Name of the faction.", Computed: true}, "symbol": schema.StringAttribute{MarkdownDescription: "The symbol of the faction.", Computed: true}, "traits": schema.ListNestedAttribute{MarkdownDescription: "List of traits that define this faction.", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"description": schema.StringAttribute{MarkdownDescription: "A description of the trait.", Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "The name of the trait.", Computed: true}, "symbol": schema.StringAttribute{MarkdownDescription: "The unique identifier of the trait.", Computed: true}}}}}}
}

// Read fetches remote state into the data source model.
func (d *GetFactionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetFactionDataSourceModel
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
func (d *GetFactionDataSource) readRemote(ctx context.Context, config *GetFactionDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/factions/{factionSymbol}"
	reqPath = strings.ReplaceAll(reqPath, "{factionSymbol}", url.PathEscape(config.FactionSymbol.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_get_faction", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_get_faction", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading spacetraders_get_faction", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		apiErr, err := client.NewAPIError(httpResp)
		if err != nil {
			resp.Diagnostics.AddError("Error reading spacetraders_get_faction", fmt.Sprintf("Could not read error response: %s", err))
			return
		}
		resp.Diagnostics.AddError("Error reading spacetraders_get_faction", apiErr.Error())
		return
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading spacetraders_get_faction", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["data"]; ok {
		if m, ok := v.(map[string]any); ok {
			data = m
		}
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_get_faction", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetFactionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
