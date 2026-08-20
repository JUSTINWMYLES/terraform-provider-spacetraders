package provider

import "context"
import (
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// Compile-time interface assertion.
var _ datasource.DataSource = (*WebsocketDepartureEventsDataSource)(nil)

// WebsocketDepartureEventsDataSource is the generated Terraform data source implementation.
type WebsocketDepartureEventsDataSource struct {
}

// WebsocketDepartureEventsDataSourceModel describes the data source state shape.
type WebsocketDepartureEventsDataSourceModel struct {
}

// NewWebsocketDepartureEventsDataSource returns a new instance of the generated data source.
func NewWebsocketDepartureEventsDataSource() datasource.DataSource {
	return &WebsocketDepartureEventsDataSource{}
}

// Metadata returns the data source type name.
func (d *WebsocketDepartureEventsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "spacetraders_websocket_departure_events"
}

// Schema returns the data source schema.
func (d *WebsocketDepartureEventsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Subscribe to departure events for a system.\n\n          ## WebSocket Events\n\n          The following events are available:\n\n          - `systems.{systemSymbol}.departure`: A ship has departed from the system.\n\n          ## Subscribe using a message with the following format:\n\n          ```json\n          {\n            \"action\": \"subscribe\",\n            \"systemSymbol\": \"{systemSymbol}\"\n          }\n          ```\n\n          ## Unsubscribe using a message with the following format:\n\n          ```json\n          {\n            \"action\": \"unsubscribe\",\n            \"systemSymbol\": \"{systemSymbol}\"\n          }\n          ```"}
}

// Read fetches remote state into the data source model.
func (d *WebsocketDepartureEventsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config WebsocketDepartureEventsDataSourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.State.Set(ctx, &config)
}
