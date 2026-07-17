// apparmor.d - Full set of apparmor profiles
// Copyright (C) 2021-2026 Alexandre Pujol <alexandre@pujol.io>
// SPDX-License-Identifier: GPL-2.0-only

package builder

import (
	"regexp"

	"github.com/roddhjav/apparmor.d/pkg/prebuild"
	"github.com/roddhjav/apparmor.d/pkg/tasks"
	"github.com/roddhjav/apparmor.d/pkg/util"
)

var (
	regProfileName = regexp.MustCompile(`(?m)^profile\s+(\S+)\s+`)
)

type ProfileMode struct {
	tasks.BaseTask
	modes      map[string]string
	enforceAll bool
}

// NewProfileMode creates a new ProfileMode builder.
//
// enforceAll mirrors the global --enforce build flag. When set, entries from
// dist/flags that would set "complain" are skipped: otherwise this builder
// (which always runs, and runs after the Enforce builder) would silently put
// every profile listed in the distribution's .flags file (e.g. dist/flags/
// fedora.flags) back into complain mode, even on a build that explicitly
// asked for global enforcement. Non-complain modes (kill, unconfined,
// default_allow, prompt...) from the manifest are still honored either way.
func NewProfileMode(enforceAll bool) *ProfileMode {
	modes := make(map[string]string)
	for _, name := range []string{"main", tasks.Distribution} {
		for profile, flags := range prebuild.Flags.Read(name) {
			if len(flags) > 0 {
				modes[profile] = flags[0]
			}
		}
	}
	return &ProfileMode{
		BaseTask: tasks.BaseTask{
			Keyword: "profile-mode",
			Msg:     "Build: set modes (complain, enforce...) as definied in dist/flags",
		},
		modes:      modes,
		enforceAll: enforceAll,
	}
}

func (b ProfileMode) Apply(opt *Option, profile string) (string, error) {
	matches := regProfileName.FindStringSubmatch(profile)
	if matches == nil {
		return profile, nil
	}

	name := matches[1]
	mode, present := b.modes[name]
	if !present {
		return profile, nil
	}

	if b.enforceAll && mode == "complain" {
		return profile, nil
	}

	return util.SetMode(profile, mode)
}
