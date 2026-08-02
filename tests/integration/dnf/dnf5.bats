#!/usr/bin/env bats
# apparmor.d - Full set of apparmor profiles
# Copyright (C) 2026 Alexandre Pujol <alexandre@pujol.io>
# SPDX-License-Identifier: GPL-2.0-only

load ../common

@test "dnf5: Update the cached metadata for all repositories" {
    sudo dnf5 makecache
}

@test "dnf5: Search for a given package" {
    dnf5 search pass
}

@test "dnf5: Show information for a package" {
    dnf5 info pass
}

@test "dnf5: Install a package, or update it to the latest available version" {
    sudo dnf5 install -y pass
}

@test "dnf5: Remove a package and its unused dependencies" {
    sudo dnf5 remove -y pass
}

@test "dnf5: Upgrade all installed packages to their newest available versions" {
    sudo dnf5 upgrade -y
}

@test "dnf5: Remove packages that are no longer needed" {
    sudo dnf5 autoremove -y
}

@test "dnf5: Clean cached package files no longer needed" {
    sudo dnf5 clean packages
}

@test "dnf5: List all packages" {
    dnf5 list
}

@test "dnf5: List installed packages" {
    dnf5 list --installed
}
