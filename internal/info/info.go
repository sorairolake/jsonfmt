//
// SPDX-License-Identifier: Apache-2.0 OR MIT
//
// Copyright (C) 2022 Shun Sakai
//

package info

import "runtime/debug"

const CommandName = "jsonfmt"

var version string

func GetVersion() string {
	if version != "" {
		return version
	}

	if info, ok := debug.ReadBuildInfo(); ok {
		return info.Main.Version
	} else {
		return "unknown"
	}
}
