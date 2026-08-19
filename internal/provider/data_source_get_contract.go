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
	_ datasource.DataSource              = (*GetContractDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetContractDataSource)(nil)
)

// GetContractDataSource is the generated Terraform data source implementation.
type GetContractDataSource struct {
	client *client.Client
}

// GetContractDataSourceModel describes the data source state shape.
type GetContractDataSourceModel struct {
	Accepted         types.Bool   `tfsdk:"accepted"`
	ContractId       types.String `tfsdk:"contract_id" json:"contractId"`
	DeadlineToAccept types.String `tfsdk:"deadline_to_accept" json:"deadlineToAccept"`
	Expiration       types.String `tfsdk:"expiration"`
	FactionSymbol    types.String `tfsdk:"faction_symbol" json:"factionSymbol"`
	Fulfilled        types.Bool   `tfsdk:"fulfilled"`
	Id               types.String `tfsdk:"id"`
	Terms            types.Object `tfsdk:"terms"`
	Type             types.String `tfsdk:"type"`
}

// NewGetContractDataSource returns a new instance of the generated data source.
func NewGetContractDataSource() datasource.DataSource {
	return &GetContractDataSource{}
}

// Metadata returns the data source type name.
func (d *GetContractDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "spacetraders_get_contract"
}

// Schema returns the data source schema.
func (d *GetContractDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get the details of a specific contract.", Attributes: map[string]schema.Attribute{"accepted": schema.BoolAttribute{Computed: true}, "contract_id": schema.StringAttribute{Required: true}, "deadline_to_accept": schema.StringAttribute{Computed: true}, "expiration": schema.StringAttribute{Computed: true}, "faction_symbol": schema.StringAttribute{Computed: true}, "fulfilled": schema.BoolAttribute{Computed: true}, "id": schema.StringAttribute{Computed: true}, "terms": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"deadline": schema.StringAttribute{Computed: true}, "deliver": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"destination_symbol": schema.StringAttribute{Computed: true}, "trade_symbol": schema.StringAttribute{Computed: true}, "units_fulfilled": schema.Int64Attribute{Computed: true}, "units_required": schema.Int64Attribute{Computed: true}}}}, "payment": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"on_accepted": schema.Int64Attribute{Computed: true}, "on_fulfilled": schema.Int64Attribute{Computed: true}}}}}, "type": schema.StringAttribute{Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetContractDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetContractDataSourceModel
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
func (d *GetContractDataSource) readRemote(ctx context.Context, config *GetContractDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/my/contracts/{contractId}"
	reqPath = strings.ReplaceAll(reqPath, "{contractId}", url.PathEscape(config.ContractId.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil, client.WithSchemes("AgentToken"))
	if err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_get_contract", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_get_contract", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading spacetraders_get_contract", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		apiErr, err := client.NewAPIError(httpResp)
		if err != nil {
			resp.Diagnostics.AddError("Error reading spacetraders_get_contract", fmt.Sprintf("Could not read error response: %s", err))
			return
		}
		resp.Diagnostics.AddError("Error reading spacetraders_get_contract", apiErr.Error())
		return
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading spacetraders_get_contract", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["data"]; ok {
		if m, ok := v.(map[string]any); ok {
			data = m
		}
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading spacetraders_get_contract", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetContractDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
