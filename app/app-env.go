/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan. All Rights Reserved.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 */

package app

import (
	"strings"

	"github.com/andypangaribuan/gmod/gm"
)

func initEnv() {
	var (
		availableLanguages = make([]string, 0)
		staticEndpoints    = make(map[string]string, 0)
		languages          = strings.Split(gm.Util.Env.GetString("AVAILABLE_LANGUAGES"), "|")
		statics            = strings.Split(gm.Util.Env.GetString("STATIC_ENDPOINT"), "|")
	)

	for _, lang := range languages {
		lang = strings.ToLower(strings.TrimSpace(lang))
		if lang != "" {
			availableLanguages = append(availableLanguages, lang)
		}
	}

	Env = &stuEnv{
		AppName:               gm.Util.Env.GetString("APP_NAME"),
		AppVersion:            gm.Util.Env.GetString("APP_VERSION", "0.0.0"),
		AppEnv:                gm.Util.Env.GetAppEnv("APP_ENV"),
		AppTimezone:           gm.Util.Env.GetString("APP_TIMEZONE"),
		AppRestPort:           gm.Util.Env.GetInt("APP_REST_PORT"),
		AppAutoRecover:        gm.Util.Env.GetBool("APP_AUTO_RECOVER"),
		AppServerPrintOnError: gm.Util.Env.GetBool("APP_SERVER_PRINT_ON_ERROR"),

		IosStoreAppId:     gm.Util.Env.GetString("IOS_STORE_APP_ID"),
		AndroidStoreAppId: gm.Util.Env.GetString("ANDROID_STORE_APP_ID"),
		IosDeeplink:       gm.Util.Env.GetString("IOS_DEEPLINK"),
		AndroidDeeplink:   gm.Util.Env.GetString("ANDROID_DEEPLINK"),

		ResPath:            gm.Util.Env.GetString("RES_PATH"),
		PagesPath:          gm.Util.Env.GetString("PAGES_PATH"),
		AvailableLanguages: availableLanguages,

		StoreEndpoint: gm.Util.Env.GetString("STORE_ENDPOINT"),
		StoreHtmlPath: gm.Util.Env.GetString("STORE_HTML_PATH"),
	}

	for _, v := range statics {
		v = strings.TrimSpace(v)
		ep := strings.Split(v, ":")
		if v == "" || len(ep) != 2 {
			continue
		}

		ep[0] = strings.TrimSpace(ep[0])
		ep[1] = strings.TrimSpace(ep[1])
		if ep[0] == "" || ep[1] == "" {
			continue
		}

		staticEndpoints[ep[0]] = Env.ResPath + ep[1]
	}

	Env.StaticEndpoints = staticEndpoints
}
