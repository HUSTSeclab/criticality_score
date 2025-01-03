APP_BIN := ../bin/
CFG_FILE := ../config.json
STORAGE_DIR := ../storage/

.PHONY: all

all: calc_score.rec

# check CFG_FILE exists, if not, stop the make
ifeq ($(wildcard $(CFG_FILE)),)
$(error $(CFG_FILE) not found, use CFG_FILE=... to specify the configuration file)
endif

calc_score.rec: update_dependents.rec update_git_metrics.rec update_git_metrics_partial.rec update_depsdev.rec
	# Calculate the score
	echo "* Calculating the score..."
	$(APP_BIN)/gen_scores -config $(CFG_FILE)
	touch $@

update_dependents.rec: package_updated.src union_gitlink.rec
	# Update the dependents
	echo "* Updating the dependents..."
	$(APP_BIN)/show_distpkg_deps -config $(CFG_FILE) -type archlinux
	$(APP_BIN)/show_distpkg_deps -config $(CFG_FILE) -type debian
	$(APP_BIN)/show_distpkg_deps -config $(CFG_FILE) -type nix
	$(APP_BIN)/show_distpkg_deps -config $(CFG_FILE) -type homebrew

ifneq ($(GENTOO_PREFIX),)
	$(GENTOO_PREFIX)/usr/bin/env $(GENTOO_PREFIX)/bin/bash -l -c $(APP_BIN)/show_distpkg_deps -config $(CFG_FILE) -type gentoo
else
	echo "! Gentoo prefix not set, skipping Gentoo"
endif

	touch $@

update_git_metrics_partial.rec: union_gitlink.rec
	# Update the Git metrics
	echo "* Updating the Git metrics..."
	$(APP_BIN)/update_git_metrics -jobs 64 -config $(CFG_FILE) -storage $(STORAGE_DIR)

	touch $@

update_git_metrics.rec: git_updated.src 
	# Update the Git metrics
	echo "* Updating the Git metrics..."
	$(APP_BIN)/update_git_metrics -jobs 64 -config $(CFG_FILE) -storage $(STORAGE_DIR) -force-update-all

	touch $@


update_depsdev.rec: depsdev_updated.src update_git_metrics.rec update_git_metrics_partial.rec
	# Update from deps.dev
	echo "* Updating from deps.dev..."
	$(APP_BIN)/show_depsdev_deps -config $(CFG_FILE)
	touch $@


union_gitlink.rec: gitlink_updated.src enumerate_github.rec
	# Union the Git links
	echo "* Unioning the Git links..."
	$(APP_BIN)/gitmetricsync -config $(CFG_FILE)
	touch $@

enumerate_github.rec: github_updated.src
	# Enumerate the GitHub repositories
	echo "* Enumerating the GitHub repositories..."
	rm -f /tmp/github.txt

ifneq ($(GITHUB_TOKEN),)
	$(APP_BIN)/enumerate_github -config $(CFG_FILE) -min-stars 30 -out /tmp/github.txt -workers 20
else
	echo "! GitHub token not set, skipping GitHub enumeration"
endif

	touch $@


