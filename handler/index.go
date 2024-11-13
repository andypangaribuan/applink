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

func Index(ctx server.FuseRContext, method, path, url string) any {
	if len(url) > 7 && strings.ToLower(url[:7]) == "http://" {
		url = url[7:]
	}

	if len(url) > 8 && strings.ToLower(url[:8]) == "https://" {
		url = url[8:]
	}

	langPath := ""
	pathLower := strings.ToLower(path)
	for _, lang := range app.Env.AvailableLanguages {
		if len(pathLower) >= len(lang)+2 && pathLower[:4] == "/"+lang+"/" {
			langPath = "/" + lang
			path = path[3:]
			break
		}
	}

	var (
		htmlFile = app.Env.PagesPath + langPath + "/index.html"
		opt      = server.ResponseOpt{RawResponse: true, ResponseType: fm.Ptr("html")}
		idx      = strings.Index(url, path)
		query    = ""
	)

	if idx > 0 {
		query = url[idx+len(path):]
	}

	if len(path) > 0 && path[:1] == "/" {
		path = path[1:]
	}

	data, err := os.ReadFile(htmlFile)
	if err != nil {
		fmt.Printf("error when open the file %v\n%+v\n", htmlFile, err)
		return ctx.R200OK("", opt)
	}

	html := string(data)
	html = strings.ReplaceAll(html, "__ANDROID_DEEPLINK__", app.Env.AndroidDeeplink+path+query)
	html = strings.ReplaceAll(html, "__IOS_DEEPLINK__", app.Env.IosDeeplink+path+query)

	return ctx.R200OK(html, opt)
}
