package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestWebsocketDepartureEventsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestWebsocketDepartureEventsDataSourceSchemaValidation(t *testing.T) {
	d := NewWebsocketDepartureEventsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestWebsocketDepartureEventsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestWebsocketDepartureEventsDataSourceMetadata(t *testing.T) {
	d := NewWebsocketDepartureEventsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "spacetraders_websocket_departure_events" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "spacetraders_websocket_departure_events")
	}
}
