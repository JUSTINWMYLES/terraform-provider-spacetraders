package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAgentDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAgentDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAgentDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAgentDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAgentDataSourceMetadata(t *testing.T) {
	d := NewGetAgentDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "spacetraders_get_agent" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "spacetraders_get_agent")
	}
}
