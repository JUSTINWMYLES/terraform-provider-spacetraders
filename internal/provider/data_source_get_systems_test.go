package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSystemsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetSystemsDataSourceSchemaValidation(t *testing.T) {
	d := NewGetSystemsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetSystemsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetSystemsDataSourceMetadata(t *testing.T) {
	d := NewGetSystemsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "spacetraders_get_systems" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "spacetraders_get_systems")
	}
}
