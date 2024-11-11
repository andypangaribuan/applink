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
	"fmt"
	"os"
	"strings"

	"github.com/andypangaribuan/gmod/fm"
	"github.com/andypangaribuan/gmod/server"
)

func Store(ctx server.FuseRContext) any {
	var (
		htmlFile    = app.Env.PagesPath + app.Env.StoreHtmlPath
		opt         = server.ResponseOpt{RawResponse: true, ResponseType: fm.Ptr("html")}
		appleStore  = fmt.Sprintf("https://apps.apple.com/app/id%s", app.Env.IosStoreAppId)
		googleStore = fmt.Sprintf("https://play.google.com/store/apps/details?id=%s", app.Env.AndroidStoreAppId)
	)

	pathLower := strings.ToLower(ctx.RoutePath())
	for _, lang := range app.Env.AvailableLanguages {
		if len(pathLower) > len(lang)+2 && pathLower[:4] == "/"+lang+"/" {
			htmlFile = app.Env.PagesPath + "/" + lang + app.Env.StoreHtmlPath
			break
		}
	}

	data, err := os.ReadFile(htmlFile)
	if err != nil {
		fmt.Printf("error when open the file %v\n%+v\n", htmlFile, err)
		return ctx.R200OK("", opt)
	}

	html := string(data)
	html = strings.ReplaceAll(html, "__APPLE_STORE__", appleStore)
	html = strings.ReplaceAll(html, "__GOOGLE_STORE__", googleStore)

	return ctx.R200OK(html, opt)
}
