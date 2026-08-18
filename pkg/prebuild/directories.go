// apparmor.d - Full set of apparmor profiles
// Copyright (C) 2021-2026 Alexandre Pujol <alexandre@pujol.io>
// SPDX-License-Identifier: GPL-2.0-only

// Package prebuild defines the directory layout, manifest readers, and
// shared state used by the prebuild pipeline that turns the source profiles
// into a distribution-ready tree.
package prebuild

import "github.com/roddhjav/apparmor.d/pkg/paths"

var (
	// DistDir is the directory where the distribution specific files are stored
	DistDir *paths.Path = paths.New("dists")

	// FlagDir is the directory where the distribution flags are stored
	FlagDir *paths.Path = DistDir.Join("flags")

	// IgnoreDir is the directory where the distribution ignore files are stored
	IgnoreDir *paths.Path = DistDir.Join("ignore")

	// SystemdDir is the directory where the systemd drop-in files are stored
	SystemdDir *paths.Path = paths.New("systemd")

	Ignore = Ignorer{}
	Flags  = Flagger{}
)
