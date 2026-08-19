package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetWaypointDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetWaypointDataSourceSchemaValidation(t *testing.T) {
	d := NewGetWaypointDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetWaypointDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetWaypointDataSourceMetadata(t *testing.T) {
	d := NewGetWaypointDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "spacetraders_get_waypoint" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "spacetraders_get_waypoint")
	}
}
