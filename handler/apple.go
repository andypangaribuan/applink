/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan. All Rights Reserved.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 */

package handler

import (
	"applink/app"

	"github.com/andypangaribuan/gmod/server"
)

func Apple(ctx server.FuseRContext) any {
	filePath := ""

	for ep, file := range app.Env.StaticEndpoints {
		if ep == "/.well-known/apple-app-site-association" {
			filePath = file
			break
		}
	}

	if filePath == "" {
		return ctx.R404NotFound("file not found")
	}

	res := map[string]string{
		"download":             filePath,
		"header: Content-Type": "application/json; charset=utf-8",
	}

	return ctx.R200OK(res, server.ResponseOpt{RawResponse: true})
}
