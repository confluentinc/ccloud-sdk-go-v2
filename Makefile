### BEGIN MK-INCLUDE UPDATE ###
### This section is managed by service-bot and should not be edited here.
### You can make changes upstream in https://github.com/confluentinc/cc-service-bot

CURL ?= curl
FIND ?= find
TAR ?= tar

# Mount netrc so curl can work from inside a container
DOCKER_NETRC_MOUNT ?= 1

GITHUB_API = api.github.com
GITHUB_MK_INCLUDE_OWNER := confluentinc
GITHUB_MK_INCLUDE_REPO := cc-mk-include
GITHUB_API_CC_MK_INCLUDE := https://$(GITHUB_API)/repos/$(GITHUB_MK_INCLUDE_OWNER)/$(GITHUB_MK_INCLUDE_REPO)
GITHUB_API_CC_MK_INCLUDE_TARBALL := $(GITHUB_API_CC_MK_INCLUDE)/tarball
GITHUB_API_CC_MK_INCLUDE_VERSION ?= $(GITHUB_API_CC_MK_INCLUDE_TARBALL)/$(MK_INCLUDE_VERSION)

MK_INCLUDE_DIR := mk-include
MK_INCLUDE_LOCKFILE := .mk-include-lockfile
MK_INCLUDE_TIMESTAMP_FILE := .mk-include-timestamp
# For optimum performance, you should override MK_INCLUDE_TIMEOUT_MINS above the managed section headers to be
# a little longer than the worst case cold build time for this repo.
MK_INCLUDE_TIMEOUT_MINS ?= 240
# If this latest validated release is breaking you, please file a ticket with DevProd describing the issue, and
# if necessary you can temporarily override MK_INCLUDE_VERSION above the managed section headers until the bad
# release is yanked.
MK_INCLUDE_VERSION ?= v0.1247.0

# Make sure we always have a copy of the latest cc-mk-include release less than $(MK_INCLUDE_TIMEOUT_MINS) old:
# Note: The simply-expanded make variable makes sure this is run once per make invocation.
UPDATE_MK_INCLUDE := $(shell \
	func_fatal() { echo "$$*" >&2; echo output here triggers error below; exit 1; } ; \
	test -z "`git ls-files $(MK_INCLUDE_DIR)`" || { \
		func_fatal 'fatal: checked in $(MK_INCLUDE_DIR)/ directory is preventing make from fetching recent cc-mk-include releases for CI'; \
	} ; \
	trap "rm -f $(MK_INCLUDE_LOCKFILE); exit" 0 2 3 15; \
	waitlock=0; while ! ( set -o noclobber; echo > $(MK_INCLUDE_LOCKFILE) ); do \
	   sleep $$waitlock; waitlock=`expr $$waitlock + 1`; \
	   test 14 -lt $$waitlock && { \
	      echo 'stealing stale lock after 105s' >&2; \
	      break; \
	   } \
	done; \
	test -s $(MK_INCLUDE_TIMESTAMP_FILE) || rm -f $(MK_INCLUDE_TIMESTAMP_FILE); \
	{ test -d $(MK_INCLUDE_DIR) && test -d /proc && test -z "$(cat /proc/1/sched 2>&1 |head -n 1 |grep init)"; } || \
	test -z "`$(FIND) $(MK_INCLUDE_TIMESTAMP_FILE) -mmin +$(MK_INCLUDE_TIMEOUT_MINS) 2>&1`" || { \
	   grep -q 'machine $(GITHUB_API)' ~/.netrc 2>/dev/null || \
	      func_fatal 'error: follow https://confluentinc.atlassian.net/l/cp/0WXXRLDh to fix your ~/.netrc'; \
	   $(CURL) --fail --silent --netrc --location "$(GITHUB_API_CC_MK_INCLUDE_VERSION)" --output $(MK_INCLUDE_TIMESTAMP_FILE)T --write-out '$(GITHUB_API_CC_MK_INCLUDE_VERSION): %{errormsg}\n' >&2 \
	   && $(TAR) zxf $(MK_INCLUDE_TIMESTAMP_FILE)T \
	   && rm -rf $(MK_INCLUDE_DIR) \
	   && mv $(GITHUB_MK_INCLUDE_OWNER)-$(GITHUB_MK_INCLUDE_REPO)-* $(MK_INCLUDE_DIR) \
	   && mv -f $(MK_INCLUDE_TIMESTAMP_FILE)T $(MK_INCLUDE_TIMESTAMP_FILE) \
	   && echo 'installed cc-mk-include-$(MK_INCLUDE_VERSION) from $(GITHUB_MK_INCLUDE_REPO)' >&2 \
	   || func_fatal unable to install cc-mk-include-$(MK_INCLUDE_VERSION) from $(GITHUB_MK_INCLUDE_REPO) \
	   ; \
	} || { \
	   rm -f $(MK_INCLUDE_TIMESTAMP_FILE)T; \
	   if test -f $(MK_INCLUDE_TIMESTAMP_FILE); then \
	      touch $(MK_INCLUDE_TIMESTAMP_FILE); \
	      func_fatal 'unable to access $(GITHUB_MK_INCLUDE_REPO) fetch API to check for latest release; next try in $(MK_INCLUDE_TIMEOUT_MINS) minutes'; \
	   else \
	      func_fatal 'unable to access $(GITHUB_MK_INCLUDE_REPO) fetch API to bootstrap mk-include subdirectory'; \
	   fi; \
	} \
)
ifneq ($(UPDATE_MK_INCLUDE),)
    $(error mk-include update failed)
endif

# Export the (empty) .mk-include-check-FORCE target to allow users to trigger the mk-include
# download code above via make but without having to run any of the other targets, e.g. build.
.PHONY: .mk-include-check-FORCE
.mk-include-check-FORCE:
	@echo -n ""
### END MK-INCLUDE UPDATE ###
# Project variables
NAME := ccloud-sdk-go-v2
ALL_FOLDER_NAMES := $(shell find . -maxdepth 1 -mindepth 1 -type d | awk -F '/' '{print $$NF}')

# Build variables

# Go variables

# Git variables
GIT_MESSAGE := $(shell git log -1 --pretty=%B)
ROOT_FOLDER := "root"
SEMAPHORE_FOLDER := ".semaphore"
GITHUB_FOLDER := ".github"

# From the latest PR merged to origin master, check the impacted folder name
# IMPACTED_FOLDER_NAME represents the immediate tier-1 sub-folder gets impacted
# If all impacted files are under root, then assign root to IMPACTED_FOLDER_NAME
IMPACTED_FOLDER_NAME := $(shell \
    path=$$(git diff --name-only origin/master~ origin/master); \
    if echo "$$path" | grep -q "/"; then \
        folder=$$(echo "$$path" | grep "/" | cut -d'/' -f1 | uniq | head -n1); \
        echo $$folder; \
    else \
        echo $(ROOT_FOLDER); \
    fi \
)

VERSION_PREFIX := $(IMPACTED_FOLDER_NAME)/v

# Make sure to have a valid folder name to tag, otherwise we could have a file change in root directory
# such as README.md change in the case we should do nothing

# If IMPACTED_FOLDER_NAME has at least one tag in the past, bump the associated tag version
# Otherwise assign the prefix + v0.0.0/v0.1.0 as the default current/bump tag version
CURRENT_TAG_VERSION := $(shell git describe --tags origin/master --match="$(IMPACTED_FOLDER_NAME)/v*" --abbrev=0 || echo "$(VERSION_PREFIX)0.0.0")

# Decompose CURRENT_TAG_VERSION to MAJOR, MINOR and PATCH for flexible bump afterwards
VERSION_NUMERIC := $(shell echo $(CURRENT_TAG_VERSION) | sed 's/.*\/v//')
MAJOR := $(shell echo $(VERSION_NUMERIC) | awk -F'.' '{print $$1}')
MINOR := $(shell echo $(VERSION_NUMERIC) | awk -F'.' '{print $$2}')
PATCH := $(shell echo $(VERSION_NUMERIC) | awk -F'.' '{print $$3}')

# Determine new tag version based on commit message inclusion of "major", "minor" or "patch"
# By default, minor version should be bumped
# if the CURRENT_TAG_VERSION is deleted, then the numeric version will be reused
NEXT_TAG_VERSION = $(VERSION_PREFIX)$(shell \
	if echo "$(CURRENT_TAG_VERSION)" | grep -qi "delete"; then \
	  	echo "$(MAJOR).$(MINOR).$(PATCH)"; \
	elif echo "$(GIT_MESSAGE)" | grep -qi "major"; then \
		echo "$$(($(MAJOR) + 1)).0.0"; \
	elif echo "$(GIT_MESSAGE)" | grep -qi "minor"; then \
		echo "$(MAJOR).$$(($(MINOR) + 1)).0"; \
	elif echo "$(GIT_MESSAGE)" | grep -qi "patch"; then \
		echo "$(MAJOR).$(MINOR).$$(($(PATCH) + 1))"; \
	else \
		echo "$(MAJOR).$$(($(MINOR) + 1)).0"; \
	fi)

# print the impacted folder name, current and next candidate tagging version
.PHONY: show-args
show-args:
	@echo impact folder from latest commit: $(IMPACTED_FOLDER_NAME)
	@echo current tagging version: $(CURRENT_TAG_VERSION)
	@echo next candidate tagging version: $(NEXT_TAG_VERSION)

.PHONY: tag-release
tag-release:
	@echo "Validate the next candidate tagging version before publishing..."
	@if [ "$(IMPACTED_FOLDER_NAME)" = "$(ROOT_FOLDER)" ] || [ "$(IMPACTED_FOLDER_NAME)" = "$(SEMAPHORE_FOLDER)" ] || [ "$(IMPACTED_FOLDER_NAME)" = "$(GITHUB_FOLDER)" ]; then \
    		if [ "$(IMPACTED_FOLDER_NAME)" = "$(ROOT_FOLDER)" ]; then \
    			echo "Skip making a new tagging version for root-only changes"; \
    		elif [ "$(IMPACTED_FOLDER_NAME)" = "$(SEMAPHORE_FOLDER)" ]; then \
    			echo "Skip making a new tagging version for .semaphore folder changes"; \
    		elif [ "$(IMPACTED_FOLDER_NAME)" = "$(GITHUB_FOLDER)" ]; then \
    			echo "Skip making a new tagging version for .github folder changes"; \
    		fi; \
    else \
    	echo "Making a new tagged version now for $(IMPACTED_FOLDER_NAME)..."; \
    	git tag $(NEXT_TAG_VERSION); \
    	git push origin $(NEXT_TAG_VERSION); \
    fi
