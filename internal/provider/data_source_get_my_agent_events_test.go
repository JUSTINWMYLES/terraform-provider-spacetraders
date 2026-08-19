package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetMyAgentEventsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetMyAgentEventsDataSourceSchemaValidation(t *testing.T) {
	d := NewGetMyAgentEventsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetMyAgentEventsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetMyAgentEventsDataSourceMetadata(t *testing.T) {
	d := NewGetMyAgentEventsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "spacetraders_get_my_agent_events" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "spacetraders_get_my_agent_events")
	}
}
