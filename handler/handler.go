/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan. All Rights Reserved.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 */

package handler

import (
	"github.com/andypangaribuan/gmod/fm"
	"github.com/andypangaribuan/gmod/server"
)

func Healthz(ctx server.FuseRContext) any {
	return ctx.R200OK("running", server.ResponseOpt{RawResponse: true})
}

func AppLink(ctx server.FuseRContext) any {
	return ctx.R200OK("<h1>Hello, World!</h1>", server.ResponseOpt{RawResponse: true, ResponseType: fm.Ptr("html")})
}
