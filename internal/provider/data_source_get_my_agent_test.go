package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetMyAgentDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetMyAgentDataSourceSchemaValidation(t *testing.T) {
	d := NewGetMyAgentDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetMyAgentDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetMyAgentDataSourceMetadata(t *testing.T) {
	d := NewGetMyAgentDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "spacetraders_get_my_agent" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "spacetraders_get_my_agent")
	}
}
