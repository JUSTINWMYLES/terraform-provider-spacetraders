package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetRepairShipDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetRepairShipDataSourceSchemaValidation(t *testing.T) {
	d := NewGetRepairShipDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetRepairShipDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetRepairShipDataSourceMetadata(t *testing.T) {
	d := NewGetRepairShipDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "spacetraders_get_repair_ship" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "spacetraders_get_repair_ship")
	}
}
