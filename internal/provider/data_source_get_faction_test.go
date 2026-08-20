package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetFactionDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetFactionDataSourceSchemaValidation(t *testing.T) {
	d := NewGetFactionDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetFactionDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetFactionDataSourceMetadata(t *testing.T) {
	d := NewGetFactionDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "spacetraders_get_faction" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "spacetraders_get_faction")
	}
}
