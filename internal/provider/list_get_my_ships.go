package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)
import (
	diag "github.com/hashicorp/terraform-plugin-framework/diag"
	list "github.com/hashicorp/terraform-plugin-framework/list"
	listschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	types "github.com/hashicorp/terraform-plugin-framework/types"
	tftypes "github.com/hashicorp/terraform-plugin-go/tftypes"
	client "github.com/spacetraders/terraform-provider-spacetraders/internal/client"
)

// Compile-time interface assertion.
var _ list.ListResource = (*GetMyShipsListResource)(nil)
var _ list.ListResourceWithConfigure = (*GetMyShipsListResource)(nil)

// GetMyShipsListResource is the generated Terraform list resource implementation.
type GetMyShipsListResource struct {
	client *client.Client
}

// GetMyShipsListResourceModel describes the spacetraders_get_my_ships list filter configuration shape.
type GetMyShipsListResourceModel struct {
	Limit types.Int64 `tfsdk:"limit"`
	Page  types.Int64 `tfsdk:"page"`
}

// NewGetMyShipsListResource returns a new instance of the generated list resource.
func NewGetMyShipsListResource() list.ListResource {
	return &GetMyShipsListResource{}
}

// Metadata returns the list resource type name.
func (l *GetMyShipsListResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "spacetraders_get_my_ships"
}

// ListResourceConfigSchema returns the list resource config schema.
func (l *GetMyShipsListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{MarkdownDescription: "Return a paginated list of all of ships under your agent's ownership.", Attributes: map[string]listschema.Attribute{"limit": listschema.Int64Attribute{Optional: true}, "page": listschema.Int64Attribute{Optional: true}}}
}

// List streams matching resource instances for terraform query.
func (l *GetMyShipsListResource) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	stream.Results = func(push func(list.ListResult) bool) {
		var config GetMyShipsListResourceModel
		diags := req.Config.Get(ctx, &config)
		if diags.HasError() {
			result := req.NewListResult(ctx)
			result.Diagnostics = diags
			push(result)
			return
		}
		items, diags := l.listRemote(ctx, &config)
		if diags.HasError() {
			result := req.NewListResult(ctx)
			result.Diagnostics = diags
			push(result)
			return
		}
		for _, item := range items {
			result := req.NewListResult(ctx)
			itemMap := map[string]json.RawMessage{}
			if err := json.Unmarshal(item, &itemMap); err != nil {
				result.Diagnostics.AddError("Error listing spacetraders_get_my_ships", fmt.Sprintf("Could not decode list item: %s", err))
				if !push(result) {
					return
				}
				continue
			}
			identity := map[string]json.RawMessage{}
			shipSymbolValue, ok := itemMap["symbol"]
			if !ok {
				if itemMap["metadata"] != nil {
					metaMap := map[string]json.RawMessage{}
					if json.Unmarshal(itemMap["metadata"], &metaMap) == nil {
						shipSymbolValue, ok = metaMap["symbol"]
					}
				}
			}
			if !ok {
				shipSymbolValue, ok = itemMap["ship_symbol"]
			}
			if !ok {
				shipSymbolValue, ok = itemMap["id"]
			}
			if !ok {
				result.Diagnostics.AddError("Error listing spacetraders_get_my_ships", "List item is missing identity attribute \"ship_symbol\".")
				if !push(result) {
					return
				}
				continue
			}
			identity["ship_symbol"] = shipSymbolValue
			idJSON, err := json.Marshal(identity)
			if err != nil {
				result.Diagnostics.AddError("Error listing spacetraders_get_my_ships", fmt.Sprintf("Could not encode list item identity: %s", err))
				if !push(result) {
					return
				}
				continue
			}
			idVal, err := tftypes.ValueFromJSON(idJSON, req.ResourceIdentitySchema.Type().TerraformType(ctx))
			if err != nil {
				result.Diagnostics.AddError("Error listing spacetraders_get_my_ships", fmt.Sprintf("Could not decode list item identity: %s", err))
				if !push(result) {
					return
				}
				continue
			}
			result.Identity.Raw = idVal
			if req.IncludeResource {
				resVal, err := tftypes.ValueFromJSON(item, req.ResourceSchema.Type().TerraformType(ctx))
				if err != nil {
					result.Diagnostics.AddWarning("Error listing spacetraders_get_my_ships", fmt.Sprintf("Could not decode list item into the resource schema: %s", err))
				} else {
					result.Resource.Raw = resVal
				}
			}
			if !push(result) {
				return
			}
		}
	}
}

// listRemote fetches and decodes the collection pages, returning the items and any diagnostics for the List iterator to surface.
func (l *GetMyShipsListResource) listRemote(ctx context.Context, config *GetMyShipsListResourceModel) ([]json.RawMessage, diag.Diagnostics) {
	var diags diag.Diagnostics
	if l.client == nil {
		diags.AddError("Client Not Configured", "The API client was not set on the list resource. The provider Configure method must run before list operations; this is a bug in the generated provider.")
		return nil, diags
	}
	reqPath := "/my/ships"
	params := url.Values{}
	if !config.Page.IsNull() {
		params.Set("page", strconv.FormatInt(config.Page.ValueInt64(), 10))
	}
	if !config.Limit.IsNull() {
		params.Set("limit", strconv.FormatInt(config.Limit.ValueInt64(), 10))
	}
	var nextURL string
	fetch := func(ctx context.Context, p url.Values) (*http.Response, error) {
		httpReq, err := l.client.NewRequest(ctx, http.MethodGet, reqPath, nil, client.WithSchemes("AgentToken"))
		if err != nil {
			return nil, err
		}
		if nextURL != "" {
			parsed, perr := url.Parse(nextURL)
			if perr != nil {
				return nil, perr
			}
			httpReq.URL = parsed
		} else {
			httpReq.URL.RawQuery = p.Encode()
		}
		return l.client.Do(httpReq)
	}
	pages, err := client.ListAllPages(ctx, params, fetch, nil)
	if err != nil {
		diags.AddError("Error listing spacetraders_get_my_ships", fmt.Sprintf("Could not read list response: %s", err))
		return nil, diags
	}
	allItems := []json.RawMessage{}
	for _, page := range pages {
		pageObj := map[string]json.RawMessage{}
		if err := json.Unmarshal(page, &pageObj); err != nil {
			diags.AddError("Error listing spacetraders_get_my_ships", fmt.Sprintf("Could not decode list page: %s", err))
			return nil, diags
		}
		rawItems, ok := pageObj["data"]
		if !ok {
			diags.AddError("Error listing spacetraders_get_my_ships", fmt.Sprintf("Could not decode list page: missing %q array", "data"))
			return nil, diags
		}
		items := []json.RawMessage{}
		if err := json.Unmarshal(rawItems, &items); err != nil {
			diags.AddError("Error listing spacetraders_get_my_ships", fmt.Sprintf("Could not decode list page: %s", err))
			return nil, diags
		}
		allItems = append(allItems, items...)
	}
	return allItems, diags
}

// Configure stores the API client supplied by the provider.
func (l *GetMyShipsListResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected List Resource Configure Type", fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData))
		return
	}
	l.client = c
}
