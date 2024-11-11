/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan. All Rights Reserved.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 */

package app

import "github.com/andypangaribuan/gmod/ice"

type stuEnv struct {
	AppName               string
	AppVersion            string
	AppEnv                ice.AppEnv
	AppTimezone           string
	AppRestPort           int
	AppAutoRecover        bool
	AppServerPrintOnError bool

	IosStoreAppId     string
	AndroidStoreAppId string
	AndroidDeeplink   string
	IosDeeplink       string

	ResPath            string
	PagesPath          string
	AvailableLanguages []string

	StoreEndpoint string
	StoreHtmlPath string

	StaticEndpoints map[string]string
}
