# Project variables
NAME := ccloud-sdk-go-v2-internal
ALL_FOLDER_NAMES := $(shell find . -maxdepth 1 -mindepth 1 -type d | awk -F '/' '{print $$NF}')

# Build variables

# Go variables

# Git variables
INITIAL_TAG_VERSION := 0.0.1
GIT_MESSAGE := $(shell git log -1 --pretty=%B)

# From the latest PR merged to origin master, check the impacted folder name
# IMPACTED_FOLDER_NAME represents the immediate tier-1 sub-folder gets impacted
# If all impacted files are under root, then assign root to IMPACTED_FOLDER_NAME
IMPACTED_FOLDER_NAME := $(shell \
    path=$$(git diff --name-only origin/master~ origin/master); \
    if echo "$$path" | grep -q "/"; then \
        folder=$$(echo "$$path" | grep "/" | cut -d'/' -f1 | uniq | head -n1); \
        echo $$folder; \
    else \
        echo "root"; \
    fi \
)

VERSION_PREFIX := $(IMPACTED_FOLDER_NAME)/v

# Make sure to have a valid folder name to tag, otherwise we could have a file change in root directory
# such as README.md change in the case we should do nothing

# If IMPACTED_FOLDER_NAME has at least one tag in the past, bump the associated tag version
# Otherwise assign the prefix + v0.0.0/v0.0.1 as the current/bump tag version
CURRENT_TAG_VERSION := $(shell git describe --tags origin/master --match="$(IMPACTED_FOLDER_NAME)/v*" --abbrev=0 2>/dev/null || echo "$(VERSION_PREFIX)0.0.0")
$(info the latest tag version is $(CURRENT_TAG_VERSION))

# Decompose CURRENT_TAG_VERSION to MAJOR, MINOR and PATCH for flexible bump afterwards
VERSION_NUMERIC := $(shell echo $(CURRENT_TAG_VERSION) | sed 's/.*\/v//')
MAJOR := $(shell echo $(VERSION_NUMERIC) | awk -F'.' '{print $$1}')
MINOR := $(shell echo $(VERSION_NUMERIC) | awk -F'.' '{print $$2}')
PATCH := $(shell echo $(VERSION_NUMERIC) | awk -F'.' '{print $$3}')

# Determine new tag version based on commit message inclusion of "major", "minor" or "patch"
# TODO: Discuss the logic with Kostya to see we should use commit message to trigger it
NEXT_TAG_VERSION = $(VERSION_PREFIX)$(shell \
	if echo "$(GIT_MESSAGE)" | grep -qi "major"; then \
		echo "$$(($(MAJOR) + 1)).0.0"; \
	elif echo "$(GIT_MESSAGE)" | grep -qi "minor"; then \
		echo "$(MAJOR).$$(($(MINOR) + 1)).0"; \
	elif echo "$(GIT_MESSAGE)" | grep -qi "patch"; then \
		echo "$(MAJOR).$(MINOR).$$(($(PATCH) + 1))"; \
	else \
		echo "$(MAJOR).$(MINOR).$$(($(PATCH) + 1))"; \
	fi)

# print the impacted folder name, current and next tagging version
.PHONY: show-version
show-version:
	@echo impact folder from latest commit: $(IMPACTED_FOLDER_NAME)
	@echo current tagging version: $(CURRENT_TAG_VERSION)
	@echo next tagging version: $(NEXT_TAG_VERSION)

.PHONY: tag-release
tag-release:
ifeq ($(IMPACTED_FOLDER_NAME), root)
	@echo "Skip making a new tagged version for root-only changes"
else
	@echo "Making a new tagged version now for $(IMPACTED_FOLDER_NAME)..."
	git tag $(NEXT_TAG_VERSION)
	git push origin $(NEXT_TAG_VERSION)
endif
