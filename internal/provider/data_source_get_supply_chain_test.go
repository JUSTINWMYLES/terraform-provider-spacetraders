package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSupplyChainDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetSupplyChainDataSourceSchemaValidation(t *testing.T) {
	d := NewGetSupplyChainDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetSupplyChainDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetSupplyChainDataSourceMetadata(t *testing.T) {
	d := NewGetSupplyChainDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "spacetraders_get_supply_chain" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "spacetraders_get_supply_chain")
	}
}
