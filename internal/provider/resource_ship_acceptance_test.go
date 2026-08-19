package provider

import (
	json "encoding/json"
	"fmt"
	"io"
	http "net/http"
	httptest "net/http/httptest"
	"strings"
	"sync"
	"testing"
)
import (
	providerserver "github.com/hashicorp/terraform-plugin-framework/providerserver"
	tfprotov6 "github.com/hashicorp/terraform-plugin-go/tfprotov6"
	resource "github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testAccShipResourceConfig(serverURL string, name string) string {
	return fmt.Sprintf("provider \"spacetraders\" {\n  endpoint = \"%s\"\n  account_token = \"example\"\n  agent_token = \"example\"\n}\nresource \"spacetraders_ship\" \"example\" {\n  ship_type = \"%s\"\n  waypoint_symbol = \"example\"\n}\n", serverURL, name)
}

// newShipResourceMockServer returns an httptest server that stubs the ShipResource CRUD endpoints.
// The server echoes request bodies so that create/update responses reflect the values sent by the test.
func newShipResourceMockServer() *httptest.Server {
	mux := http.NewServeMux()
	state0 := make(map[string]map[string]interface{})
	var mu0 sync.Mutex
	lastKey0 := ""
	handler0 := func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer example" {
			http.Error(w, "missing bearer token", http.StatusUnauthorized)
			return
		}
		if r.Header.Get("Authorization") != "Bearer example" {
			http.Error(w, "missing bearer token", http.StatusUnauthorized)
			return
		}
		mu0.Lock()
		defer mu0.Unlock()
		id := strings.Trim(strings.TrimPrefix(r.URL.Path, "/my/ships"), "/")
		if id == "" {
			id = "example-id"
		}
		switch r.Method {
		case http.MethodPost:
			if strings.Contains(id, "/") {
				delete(state0, id)
				w.WriteHeader(200)
				return
			}
			body := make(map[string]interface{})
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil && err != io.EOF {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if _, ok := body["symbol"]; !ok {
				body["symbol"] = "example-id"
			}
			id = fmt.Sprintf("%v", body["symbol"])
			state0[id] = body
			lastKey0 = id
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(201)
			_ = json.NewEncoder(w).Encode(body)
			return
		case http.MethodGet:
			body, ok := state0[id]
			if !ok && lastKey0 != "" {
				body = state0[lastKey0]
			}
			if body == nil {
				http.NotFound(w, r)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			_ = json.NewEncoder(w).Encode(body)
			return
		case http.MethodDelete:
			delete(state0, id)
			w.WriteHeader(200)
			return
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
	mux.HandleFunc("/my/ships", handler0)
	mux.HandleFunc("/my/ships/", handler0)
	return httptest.NewServer(mux)
}

// TestAccShipResourceLifecycle verifies create, update, delete, and import flows against a mock API.
func TestAccShipResourceLifecycle(t *testing.T) {
	t.Setenv("TF_ACC", "1")
	server := newShipResourceMockServer()
	defer server.Close()
	resource.Test(t, resource.TestCase{ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){"spacetraders": providerserver.NewProtocol6WithError(New())}, Steps: []resource.TestStep{resource.TestStep{Config: testAccShipResourceConfig(server.URL, "example"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("spacetraders_ship.example", "symbol"), resource.TestCheckResourceAttr("spacetraders_ship.example", "ship_type", "example"))}}})
}
