package provider

import "context"
import (
	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/list"
	tfframeworkprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	types "github.com/hashicorp/terraform-plugin-framework/types"
	client "github.com/spacetraders/terraform-provider-spacetraders/internal/client"
)

var _ tfframeworkprovider.Provider = (*spacetradersProvider)(nil)
var _ tfframeworkprovider.ProviderWithFunctions = (*spacetradersProvider)(nil)
var _ tfframeworkprovider.ProviderWithEphemeralResources = (*spacetradersProvider)(nil)
var _ tfframeworkprovider.ProviderWithListResources = (*spacetradersProvider)(nil)
var _ tfframeworkprovider.ProviderWithActions = (*spacetradersProvider)(nil)

// spacetradersProvider is the generated Terraform provider implementation.
type spacetradersProvider struct {
	configured bool
}

// spacetradersProviderModel describes the provider-level configuration shape.
type spacetradersProviderModel struct {
	Endpoint                  types.String `tfsdk:"endpoint"`
	AccountToken              types.String `tfsdk:"account_token"`
	AgentToken                types.String `tfsdk:"agent_token"`
	LogFile                   types.String `tfsdk:"log_file"`
	LogCaptureRequestHeaders  types.Bool   `tfsdk:"log_capture_request_headers"`
	LogCaptureRequestBody     types.Bool   `tfsdk:"log_capture_request_body"`
	LogCaptureResponseHeaders types.Bool   `tfsdk:"log_capture_response_headers"`
	LogCaptureResponseBody    types.Bool   `tfsdk:"log_capture_response_body"`
	LogMaxBodyBytes           types.Int64  `tfsdk:"log_max_body_bytes"`
}

// New returns a new instance of the generated provider.
func New() tfframeworkprovider.Provider {
	return &spacetradersProvider{}
}

// Metadata returns the provider type name.
func (p *spacetradersProvider) Metadata(_ context.Context, _ tfframeworkprovider.MetadataRequest, resp *tfframeworkprovider.MetadataResponse) {
	resp.TypeName = "spacetraders"
}

// Schema returns the provider configuration schema.
func (p *spacetradersProvider) Schema(_ context.Context, _ tfframeworkprovider.SchemaRequest, resp *tfframeworkprovider.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "SpaceTraders is an open-universe game and learning platform that offers a set of HTTP endpoints to control a fleet of ships and explore a multiplayer universe.\n\nThe API is documented using [OpenAPI](https://github.com/SpaceTradersAPI/api-docs). You can send your first request right here in your browser to check the status of the game server.\n\n```json http\n{\n  \"method\": \"GET\",\n  \"url\": \"https://api.spacetraders.io/v2\",\n}\n```\n\nUnlike a traditional game, SpaceTraders does not have a first-party client or app to play the game. Instead, you can use the API to build your own client, write a script to automate your ships, or try an app built by the community.\n\nWe have a [Discord channel](https://discord.com/invite/jh6zurdWk5) where you can share your projects, ask questions, and get help from other players.", Attributes: map[string]schema.Attribute{"endpoint": schema.StringAttribute{MarkdownDescription: "Overrides the default API base URL derived from the OpenAPI servers. Useful for directing the provider at a test or mock server.", Optional: true}, "account_token": schema.StringAttribute{MarkdownDescription: "Bearer token used for HTTP bearer authentication. Expected format: JWT.", Optional: true, Sensitive: true}, "agent_token": schema.StringAttribute{MarkdownDescription: "Bearer token used for HTTP bearer authentication. Expected format: JWT.", Optional: true, Sensitive: true}, "log_file": schema.StringAttribute{MarkdownDescription: "Path to a file that receives HTTP request/response trace logs. When unset, trace logging is disabled.", Optional: true}, "log_capture_request_headers": schema.BoolAttribute{MarkdownDescription: "Capture request headers in the trace log. Sensitive headers are redacted.", Optional: true}, "log_capture_request_body": schema.BoolAttribute{MarkdownDescription: "Capture request bodies in the trace log. Disabled by default to avoid writing sensitive payloads to disk.", Optional: true}, "log_capture_response_headers": schema.BoolAttribute{MarkdownDescription: "Capture response headers in the trace log. Sensitive headers are redacted.", Optional: true}, "log_capture_response_body": schema.BoolAttribute{MarkdownDescription: "Capture response bodies in the trace log. Disabled by default to avoid writing sensitive payloads to disk.", Optional: true}, "log_max_body_bytes": schema.Int64Attribute{MarkdownDescription: "Maximum number of body bytes captured per log entry before truncation. Defaults to 4096.", Optional: true}}}
}

// Configure decodes practitioner configuration and marks the provider as configured.
func (p *spacetradersProvider) Configure(ctx context.Context, req tfframeworkprovider.ConfigureRequest, resp *tfframeworkprovider.ConfigureResponse) {
	var config spacetradersProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	opts := []client.ClientOption{}
	if !config.Endpoint.IsNull() && !config.Endpoint.IsUnknown() {
		opts = append(opts, client.WithBaseURL(config.Endpoint.ValueString()))
	}
	if !config.AccountToken.IsNull() && !config.AccountToken.IsUnknown() {
		opts = append(opts, client.WithSchemeInterceptor("AccountToken", client.BearerAuth(config.AccountToken.ValueString())))
	}
	if !config.AgentToken.IsNull() && !config.AgentToken.IsUnknown() {
		opts = append(opts, client.WithSchemeInterceptor("AgentToken", client.BearerAuth(config.AgentToken.ValueString())))
	}
	loggingConfig := client.LoggingConfig{MaxBodyBytes: 4096, RedactHeaders: []string{"Authorization", "X-API-Key", "Cookie"}}
	if !config.LogFile.IsNull() && !config.LogFile.IsUnknown() {
		loggingConfig.LogFile = config.LogFile.ValueString()
	}
	if !config.LogCaptureRequestHeaders.IsNull() && !config.LogCaptureRequestHeaders.IsUnknown() {
		loggingConfig.CaptureRequestHeaders = config.LogCaptureRequestHeaders.ValueBool()
	}
	if !config.LogCaptureRequestBody.IsNull() && !config.LogCaptureRequestBody.IsUnknown() {
		loggingConfig.CaptureRequestBody = config.LogCaptureRequestBody.ValueBool()
	}
	if !config.LogCaptureResponseHeaders.IsNull() && !config.LogCaptureResponseHeaders.IsUnknown() {
		loggingConfig.CaptureResponseHeaders = config.LogCaptureResponseHeaders.ValueBool()
	}
	if !config.LogCaptureResponseBody.IsNull() && !config.LogCaptureResponseBody.IsUnknown() {
		loggingConfig.CaptureResponseBody = config.LogCaptureResponseBody.ValueBool()
	}
	if !config.LogMaxBodyBytes.IsNull() && !config.LogMaxBodyBytes.IsUnknown() {
		loggingConfig.MaxBodyBytes = int(config.LogMaxBodyBytes.ValueInt64())
	}
	if loggingConfig.LogFile != "" {
		opts = append(opts, client.WithLogging(loggingConfig))
	}
	c := client.New(opts...)
	resp.DataSourceData = c
	resp.ResourceData = c
	resp.EphemeralResourceData = c
	resp.ActionData = c
	resp.ListResourceData = c
	p.configured = true
}

// DataSources returns the data sources registered with this provider.
func (p *spacetradersProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{func() datasource.DataSource {
		return &GetStatusDataSource{}
	}, func() datasource.DataSource {
		return &GetAgentsDataSource{}
	}, func() datasource.DataSource {
		return &GetAgentDataSource{}
	}, func() datasource.DataSource {
		return &GetErrorCodesDataSource{}
	}, func() datasource.DataSource {
		return &GetFactionsDataSource{}
	}, func() datasource.DataSource {
		return &GetFactionDataSource{}
	}, func() datasource.DataSource {
		return &GetSupplyChainDataSource{}
	}, func() datasource.DataSource {
		return &GetMyAccountDataSource{}
	}, func() datasource.DataSource {
		return &GetMyAgentDataSource{}
	}, func() datasource.DataSource {
		return &GetMyAgentEventsDataSource{}
	}, func() datasource.DataSource {
		return &GetContractsDataSource{}
	}, func() datasource.DataSource {
		return &GetContractDataSource{}
	}, func() datasource.DataSource {
		return &GetMyFactionsDataSource{}
	}, func() datasource.DataSource {
		return &GetMyShipsDataSource{}
	}, func() datasource.DataSource {
		return &GetMyShipCargoDataSource{}
	}, func() datasource.DataSource {
		return &GetShipCooldownDataSource{}
	}, func() datasource.DataSource {
		return &GetShipModulesDataSource{}
	}, func() datasource.DataSource {
		return &GetMountsDataSource{}
	}, func() datasource.DataSource {
		return &GetShipNavDataSource{}
	}, func() datasource.DataSource {
		return &GetRepairShipDataSource{}
	}, func() datasource.DataSource {
		return &GetScrapShipDataSource{}
	}, func() datasource.DataSource {
		return &WebsocketDepartureEventsDataSource{}
	}, func() datasource.DataSource {
		return &GetSystemsDataSource{}
	}, func() datasource.DataSource {
		return &GetSystemDataSource{}
	}, func() datasource.DataSource {
		return &GetSystemWaypointsDataSource{}
	}, func() datasource.DataSource {
		return &GetWaypointDataSource{}
	}, func() datasource.DataSource {
		return &GetConstructionDataSource{}
	}, func() datasource.DataSource {
		return &GetJumpGateDataSource{}
	}, func() datasource.DataSource {
		return &GetMarketDataSource{}
	}, func() datasource.DataSource {
		return &GetShipyardDataSource{}
	}}
}

// Resources returns the managed resources registered with this provider.
func (p *spacetradersProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{func() resource.Resource {
		return &ShipResource{}
	}}
}

// Actions returns the actions registered with this provider.
func (p *spacetradersProvider) Actions(_ context.Context) []func() action.Action {
	return []func() action.Action{func() action.Action {
		return &AcceptContractAction{}
	}, func() action.Action {
		return &DeliverContractAction{}
	}, func() action.Action {
		return &FulfillContractAction{}
	}, func() action.Action {
		return &CreateChartAction{}
	}, func() action.Action {
		return &DockShipAction{}
	}, func() action.Action {
		return &ExtractResourcesAction{}
	}, func() action.Action {
		return &ExtractResourcesWithSurveyAction{}
	}, func() action.Action {
		return &JettisonAction{}
	}, func() action.Action {
		return &JumpShipAction{}
	}, func() action.Action {
		return &InstallShipModuleAction{}
	}, func() action.Action {
		return &RemoveShipModuleAction{}
	}, func() action.Action {
		return &InstallMountAction{}
	}, func() action.Action {
		return &RemoveMountAction{}
	}, func() action.Action {
		return &PatchShipNavAction{}
	}, func() action.Action {
		return &NavigateShipAction{}
	}, func() action.Action {
		return &NegotiateContractAction{}
	}, func() action.Action {
		return &OrbitShipAction{}
	}, func() action.Action {
		return &PurchaseCargoAction{}
	}, func() action.Action {
		return &ShipRefineAction{}
	}, func() action.Action {
		return &RefuelShipAction{}
	}, func() action.Action {
		return &RepairShipAction{}
	}, func() action.Action {
		return &CreateShipShipScanAction{}
	}, func() action.Action {
		return &CreateShipSystemScanAction{}
	}, func() action.Action {
		return &CreateShipWaypointScanAction{}
	}, func() action.Action {
		return &SellCargoAction{}
	}, func() action.Action {
		return &SiphonResourcesAction{}
	}, func() action.Action {
		return &CreateSurveyAction{}
	}, func() action.Action {
		return &TransferCargoAction{}
	}, func() action.Action {
		return &WarpShipAction{}
	}, func() action.Action {
		return &RegisterAction{}
	}, func() action.Action {
		return &SupplyConstructionAction{}
	}}
}

// Functions returns the provider-defined functions registered with this provider.
func (p *spacetradersProvider) Functions(_ context.Context) []func() function.Function {
	return nil
}

// EphemeralResources returns the ephemeral resources registered with this provider.
func (p *spacetradersProvider) EphemeralResources(_ context.Context) []func() ephemeral.EphemeralResource {
	return nil
}

// ListResources returns the list resources registered with this provider.
func (p *spacetradersProvider) ListResources(_ context.Context) []func() list.ListResource {
	return nil
}
