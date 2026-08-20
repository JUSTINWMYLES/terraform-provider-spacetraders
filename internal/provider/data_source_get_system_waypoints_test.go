package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSystemWaypointsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetSystemWaypointsDataSourceSchemaValidation(t *testing.T) {
	d := NewGetSystemWaypointsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetSystemWaypointsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetSystemWaypointsDataSourceMetadata(t *testing.T) {
	d := NewGetSystemWaypointsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "spacetraders_get_system_waypoints" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "spacetraders_get_system_waypoints")
	}
}
