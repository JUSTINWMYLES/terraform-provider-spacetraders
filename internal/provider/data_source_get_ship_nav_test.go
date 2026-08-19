package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetShipNavDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetShipNavDataSourceSchemaValidation(t *testing.T) {
	d := NewGetShipNavDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetShipNavDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetShipNavDataSourceMetadata(t *testing.T) {
	d := NewGetShipNavDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "spacetraders_get_ship_nav" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "spacetraders_get_ship_nav")
	}
}
