/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan. All Rights Reserved.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 */

package app

import (
	_ "github.com/andypangaribuan/gmod"

	"github.com/andypangaribuan/gmod/gm"
)

func init() {
	initEnv()

	gm.Conf.
		SetTimezone(Env.AppTimezone).
		Commit()
}
