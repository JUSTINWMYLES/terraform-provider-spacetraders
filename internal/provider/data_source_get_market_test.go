package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetMarketDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetMarketDataSourceSchemaValidation(t *testing.T) {
	d := NewGetMarketDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetMarketDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetMarketDataSourceMetadata(t *testing.T) {
	d := NewGetMarketDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "spacetraders_get_market" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "spacetraders_get_market")
	}
}
