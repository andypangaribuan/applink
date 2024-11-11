/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan. All Rights Reserved.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 */

package main

import (
	"applink/app"
	"applink/handler"

	"github.com/andypangaribuan/gmod/fm"
	"github.com/andypangaribuan/gmod/server"
	"github.com/gofiber/fiber/v2"
)

func main() {
	fm.CallOrderedInit()
	server.FuseR(app.Env.AppRestPort, rest)
}

func rest(router server.RouterR) {
	router.AutoRecover(app.Env.AppAutoRecover)
	router.PrintOnError(app.Env.AppServerPrintOnError)

	router.Unrouted(handler.Index)
	router.Static(app.Env.StaticEndpoints, fiber.Static{
		Browse:   true,
		Download: true,
		ModifyResponse: func(c *fiber.Ctx) error {
			switch c.Route().Path {
			case "/.well-known/assetlinks.json", "/.well-known/apple-app-site-association":
				c.Set("Content-Type", "application/json; charset=utf-8")
			}

			return nil
		},
	})

	for ep := range app.Env.StaticEndpoints {
		if ep == "/.well-known/apple-app-site-association" {
			router.Endpoints(nil, nil, map[string][]func(server.FuseRContext) any{
				"GET: " + ep: {handler.Apple},
				"POS: " + ep: {handler.Apple},
			})
		}
	}

	endpoints := map[string][]func(server.FuseRContext) any{
		"GET: /healthz":                 {handler.Healthz},
		"GET: " + app.Env.StoreEndpoint: {handler.Store},
	}

	for _, lang := range app.Env.AvailableLanguages {
		endpoints["GET: /"+lang+app.Env.StoreEndpoint] = []func(server.FuseRContext) any{handler.Store}
	}

	router.Endpoints(nil, nil, endpoints)
}
