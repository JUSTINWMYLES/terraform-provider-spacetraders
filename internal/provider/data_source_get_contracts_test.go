package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetContractsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetContractsDataSourceSchemaValidation(t *testing.T) {
	d := NewGetContractsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetContractsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetContractsDataSourceMetadata(t *testing.T) {
	d := NewGetContractsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "spacetraders_get_contracts" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "spacetraders_get_contracts")
	}
}
