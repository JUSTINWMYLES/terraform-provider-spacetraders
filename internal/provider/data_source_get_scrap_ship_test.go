package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetScrapShipDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetScrapShipDataSourceSchemaValidation(t *testing.T) {
	d := NewGetScrapShipDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetScrapShipDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetScrapShipDataSourceMetadata(t *testing.T) {
	d := NewGetScrapShipDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "spacetraders_get_scrap_ship" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "spacetraders_get_scrap_ship")
	}
}
