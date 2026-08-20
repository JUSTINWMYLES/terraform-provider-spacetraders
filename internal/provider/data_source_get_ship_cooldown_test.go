package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetShipCooldownDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetShipCooldownDataSourceSchemaValidation(t *testing.T) {
	d := NewGetShipCooldownDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetShipCooldownDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetShipCooldownDataSourceMetadata(t *testing.T) {
	d := NewGetShipCooldownDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "spacetraders_get_ship_cooldown" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "spacetraders_get_ship_cooldown")
	}
}
