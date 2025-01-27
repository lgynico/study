package net

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Provider struct {
	NodeType string `json:"node_type"`
	Addr     string `json:"addr"`
	Status   int    `json:"status"`
	// Methods  []string `json:"methods"`
}

type RpcRegistry struct {
	nodes map[string]map[string]*Provider
}

func NewRpcRegistry() *RpcRegistry {
	return &RpcRegistry{nodes: make(map[string]map[string]*Provider)}
}

func (r *RpcRegistry) Register(nodeType, addr string) {
	fmt.Println("register:", nodeType, addr)
	if r.nodes[nodeType] == nil {
		r.nodes[nodeType] = make(map[string]*Provider)
	}
	r.nodes[nodeType][addr] = &Provider{NodeType: nodeType, Addr: addr, Status: 1}
}

func (r *RpcRegistry) Deregister(nodeType, addr string) {
	if r.nodes[nodeType] == nil {
		return
	}
	delete(r.nodes[nodeType], addr)
}

func (r *RpcRegistry) GetProvidersByType(nodeType string) []*Provider {
	var providers []*Provider
	for _, provider := range r.nodes[nodeType] {
		providers = append(providers, provider)
	}
	return providers
}

func (r *RpcRegistry) Start(addr string) error {
	http.HandleFunc("/register", func(w http.ResponseWriter, req *http.Request) {
		nodeType := req.FormValue("node_type")
		addr := req.FormValue("addr")
		r.Register(nodeType, addr)
		w.Write([]byte("ok"))
	})

	http.HandleFunc("/getProvidersByType", func(w http.ResponseWriter, req *http.Request) {
		nodeType := req.FormValue("node_type")
		providers := r.GetProvidersByType(nodeType)
		jsonProviders, _ := json.Marshal(providers)
		w.Write([]byte(jsonProviders))
	})

	fmt.Println("Registry started")

	return http.ListenAndServe(addr, nil)
}
