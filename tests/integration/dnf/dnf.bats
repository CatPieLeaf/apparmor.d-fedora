#!/usr/bin/env bats
# apparmor.d - Full set of apparmor profiles
# Copyright (C) 2026 Alexandre Pujol <alexandre@pujol.io>
# SPDX-License-Identifier: GPL-2.0-only

load ../common

# Legacy python3-dnf, distinct from the dnf5 rewrite (dnf5.bats) - only
# meaningfully exercised on systems where /usr/bin/dnf isn't just a symlink
# to dnf5. skip_if_not_installed (tests/integration/common.bash) skips this
# whole file where it isn't a real separate binary.

@test "dnf: Update the cached metadata for all repositories" {
    sudo dnf makecache
}

@test "dnf: Install a package, or update it to the latest available version" {
    sudo dnf install -y pass
}

@test "dnf: Remove a package and its unused dependencies" {
    sudo dnf remove -y pass
}

@test "dnf: List installed packages" {
    dnf list --installed
}
