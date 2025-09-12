package main

import (
	"api-gateway/pkg"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	_ "net/http/httputil"
	"net/url"
)

func main() {
	cfg := NewConfig()

	mq := pkg.NewRabbitMQ(cfg.MqAddr, cfg.MqPort, cfg.MqUser, cfg.MqPass)

	inventory := fmt.Sprintf("http://%s:%s", cfg.InventoryAddr, cfg.InventoryPort)
	billing := fmt.Sprintf("http://%s:%s", cfg.BillingAddr, cfg.BillingPort)

	inventoryURL, err := url.Parse(inventory)
	if err != nil {
		panic(err)
	}

	billingURL, err := url.Parse(billing)
	if err != nil {
		panic(err)
	}

	handlerInventory := func(p *httputil.ReverseProxy) func(w http.ResponseWriter, r *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			r.URL.Host = inventoryURL.Host
			r.URL.Scheme = inventoryURL.Scheme
			r.Host = inventoryURL.Host
			p.ServeHTTP(w, r)
		}
	}

	handlerBilling := func(p *httputil.ReverseProxy) func(w http.ResponseWriter, r *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodPost {
				b, err := io.ReadAll(r.Body)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				r.Body.Close()

				mq.SendToQueue(b)

				w.WriteHeader(http.StatusOK)
				return
			}

			// 4) Proxy vers billing
			r.URL.Host = billingURL.Host
			r.URL.Scheme = billingURL.Scheme
			r.Host = billingURL.Host
			p.ServeHTTP(w, r)
		}
	}
	proxyInventory := httputil.NewSingleHostReverseProxy(inventoryURL)

	http.HandleFunc("/api/movies", handlerInventory(proxyInventory))
	http.HandleFunc("/api/movies/", handlerInventory(proxyInventory))

	proxyBilling := httputil.NewSingleHostReverseProxy(billingURL)
	
	http.HandleFunc("/api/billing", handlerBilling(proxyBilling))
	http.HandleFunc("/api/billing/", handlerBilling(proxyBilling))

	err = http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.Addr, cfg.Port), nil)
	if err != nil {
		panic(err)
	}
}
