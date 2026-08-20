package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestShipResourceSchemaValidation verifies that the generated resource schema is valid.
func TestShipResourceSchemaValidation(t *testing.T) {
	r := &ShipResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestShipResourceMetadata verifies that the generated resource reports the expected type name.
func TestShipResourceMetadata(t *testing.T) {
	r := &ShipResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "spacetraders_ship" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "spacetraders_ship")
	}
}
