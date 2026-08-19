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
	_ datasource.DataSource              = (*GetContractsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetContractsDataSource)(nil)
)

// GetContractsDataSource is the generated Terraform data source implementation.
type GetContractsDataSource struct {
	client *client.Client
}

// GetContractsDataSourceModel describes the data source state shape.
type GetContractsDataSourceModel struct {
	Items types.List  `tfsdk:"items"`
	Limit types.Int64 `tfsdk:"limit"`
	Page  types.Int64 `tfsdk:"page"`
}

// NewGetContractsDataSource returns a new instance of the generated data source.
func NewGetContractsDataSource() datasource.DataSource {
	return &GetContractsDataSource{}
}

// Metadata returns the data source type name.
func (d *GetContractsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "spacetraders_get_contracts"
}

// Schema returns the data source schema.
func (d *GetContractsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Return a paginated list of all your contracts.", Attributes: map[string]schema.Attribute{"items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"accepted": schema.BoolAttribute{Computed: true}, "deadline_to_accept": schema.StringAttribute{Computed: true}, "expiration": schema.StringAttribute{Computed: true}, "faction_symbol": schema.StringAttribute{Computed: true}, "fulfilled": schema.BoolAttribute{Computed: true}, "id": schema.StringAttribute{Computed: true}, "terms": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"deadline": schema.StringAttribute{Computed: true}, "deliver": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"destination_symbol": schema.StringAttribute{Computed: true}, "trade_symbol": schema.StringAttribute{Computed: true}, "units_fulfilled": schema.Int64Attribute{Computed: true}, "units_required": schema.Int64Attribute{Computed: true}}}}, "payment": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"on_accepted": schema.Int64Attribute{Computed: true}, "on_fulfilled": schema.Int64Attribute{Computed: true}}}}}, "type": schema.StringAttribute{Computed: true}}}}, "limit": schema.Int64Attribute{Optional: true}, "page": schema.Int64Attribute{Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetContractsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetContractsDataSourceModel
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
func (d *GetContractsDataSource) readListRemote(ctx context.Context, config *GetContractsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/my/contracts"
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
		resp.Diagnostics.AddError("Error reading spacetraders_get_contracts", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading spacetraders_get_contracts", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["data"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading spacetraders_get_contracts", fmt.Sprintf("Could not decode list page: missing %q array", "data"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_get_contracts", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetContractsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
