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
var _ list.ListResource = (*GetAgentsListResource)(nil)
var _ list.ListResourceWithConfigure = (*GetAgentsListResource)(nil)

// GetAgentsListResource is the generated Terraform list resource implementation.
type GetAgentsListResource struct {
	client *client.Client
}

// GetAgentsListResourceModel describes the spacetraders_get_agents list filter configuration shape.
type GetAgentsListResourceModel struct {
	Limit types.Int64 `tfsdk:"limit"`
	Page  types.Int64 `tfsdk:"page"`
}

// NewGetAgentsListResource returns a new instance of the generated list resource.
func NewGetAgentsListResource() list.ListResource {
	return &GetAgentsListResource{}
}

// Metadata returns the list resource type name.
func (l *GetAgentsListResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "spacetraders_get_agents"
}

// ListResourceConfigSchema returns the list resource config schema.
func (l *GetAgentsListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{MarkdownDescription: "List all public agent details.", Attributes: map[string]listschema.Attribute{"limit": listschema.Int64Attribute{MarkdownDescription: "How many entries to return per page", Optional: true}, "page": listschema.Int64Attribute{MarkdownDescription: "What entry offset to request", Optional: true}}}
}

// List streams matching resource instances for terraform query.
func (l *GetAgentsListResource) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	stream.Results = func(push func(list.ListResult) bool) {
		var config GetAgentsListResourceModel
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
				result.Diagnostics.AddError("Error listing spacetraders_get_agents", fmt.Sprintf("Could not decode list item: %s", err))
				if !push(result) {
					return
				}
				continue
			}
			identity := map[string]json.RawMessage{}
			agentSymbolValue, ok := itemMap["symbol"]
			if !ok {
				if itemMap["metadata"] != nil {
					metaMap := map[string]json.RawMessage{}
					if json.Unmarshal(itemMap["metadata"], &metaMap) == nil {
						agentSymbolValue, ok = metaMap["symbol"]
					}
				}
			}
			if !ok {
				agentSymbolValue, ok = itemMap["agent_symbol"]
			}
			if !ok {
				agentSymbolValue, ok = itemMap["id"]
			}
			if !ok {
				result.Diagnostics.AddError("Error listing spacetraders_get_agents", "List item is missing identity attribute \"agent_symbol\".")
				if !push(result) {
					return
				}
				continue
			}
			identity["agent_symbol"] = agentSymbolValue
			idJSON, err := json.Marshal(identity)
			if err != nil {
				result.Diagnostics.AddError("Error listing spacetraders_get_agents", fmt.Sprintf("Could not encode list item identity: %s", err))
				if !push(result) {
					return
				}
				continue
			}
			idVal, err := tftypes.ValueFromJSON(idJSON, req.ResourceIdentitySchema.Type().TerraformType(ctx))
			if err != nil {
				result.Diagnostics.AddError("Error listing spacetraders_get_agents", fmt.Sprintf("Could not decode list item identity: %s", err))
				if !push(result) {
					return
				}
				continue
			}
			result.Identity.Raw = idVal
			if req.IncludeResource {
				resVal, err := tftypes.ValueFromJSON(item, req.ResourceSchema.Type().TerraformType(ctx))
				if err != nil {
					result.Diagnostics.AddWarning("Error listing spacetraders_get_agents", fmt.Sprintf("Could not decode list item into the resource schema: %s", err))
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
func (l *GetAgentsListResource) listRemote(ctx context.Context, config *GetAgentsListResourceModel) ([]json.RawMessage, diag.Diagnostics) {
	var diags diag.Diagnostics
	if l.client == nil {
		diags.AddError("Client Not Configured", "The API client was not set on the list resource. The provider Configure method must run before list operations; this is a bug in the generated provider.")
		return nil, diags
	}
	reqPath := "/agents"
	params := url.Values{}
	if !config.Page.IsNull() {
		params.Set("page", strconv.FormatInt(config.Page.ValueInt64(), 10))
	}
	if !config.Limit.IsNull() {
		params.Set("limit", strconv.FormatInt(config.Limit.ValueInt64(), 10))
	}
	var nextURL string
	fetch := func(ctx context.Context, p url.Values) (*http.Response, error) {
		httpReq, err := l.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
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
		diags.AddError("Error listing spacetraders_get_agents", fmt.Sprintf("Could not read list response: %s", err))
		return nil, diags
	}
	allItems := []json.RawMessage{}
	for _, page := range pages {
		pageObj := map[string]json.RawMessage{}
		if err := json.Unmarshal(page, &pageObj); err != nil {
			diags.AddError("Error listing spacetraders_get_agents", fmt.Sprintf("Could not decode list page: %s", err))
			return nil, diags
		}
		rawItems, ok := pageObj["data"]
		if !ok {
			diags.AddError("Error listing spacetraders_get_agents", fmt.Sprintf("Could not decode list page: missing %q array", "data"))
			return nil, diags
		}
		items := []json.RawMessage{}
		if err := json.Unmarshal(rawItems, &items); err != nil {
			diags.AddError("Error listing spacetraders_get_agents", fmt.Sprintf("Could not decode list page: %s", err))
			return nil, diags
		}
		allItems = append(allItems, items...)
	}
	return allItems, diags
}

// Configure stores the API client supplied by the provider.
func (l *GetAgentsListResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
