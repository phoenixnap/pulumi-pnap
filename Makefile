PROJECT_NAME := PNAP Package
include build/common.mk

PACK             := pnap
PACKDIR          := sdk
PROJECT          := github.com/phoenixnap/pulumi-${PACK}
NODE_MODULE_NAME := @pulumi/${PACK}

TFGEN           := pulumi-tfgen-${PACK}
PROVIDER        := pulumi-resource-${PACK}
VERSION         := $(shell scripts/get-version)
PYPI_VERSION    := $(shell scripts/get-py-version)

DOTNET_PREFIX  := $(firstword $(subst -, ,${VERSION:v%=%})) # e.g. 1.5.0
DOTNET_SUFFIX  := $(word 2,$(subst -, ,${VERSION:v%=%}))    # e.g. alpha.1

ifeq ($(strip ${DOTNET_SUFFIX}),)
	DOTNET_VERSION := $(strip ${DOTNET_PREFIX})-preview
else
	DOTNET_VERSION := $(strip ${DOTNET_PREFIX})-preview-$(strip ${DOTNET_SUFFIX})
endif

TESTPARALLELISM := 4

tfgen::
	cd provider && go install -ldflags "-X github.com/phoenixnap/pulumi-${PACK}/provider/pkg/version.Version=${VERSION}" ${PROJECT}/provider/cmd/${TFGEN}

generate_schema:: tfgen
	$(TFGEN) schema --out ./provider/cmd/${PROVIDER}

provider:: generate_schema
	cd provider && go generate cmd/${PROVIDER}/main.go
	cd provider && go install -ldflags "-X github.com/pulumi/pulumi-${PACK}/provider/pkg/version.Version=${VERSION}" ${PROJECT}/provider/cmd/${PROVIDER}

# NOTE: Since the plugin is published using the nodejs style semver version
# We set the PLUGIN_VERSION to be the same as the version we use when building
# the provider (e.g. x.y.z-dev-... instead of x.y.zdev...)
build:: provider
	cd provider && for LANGUAGE in "nodejs" "python" "go" "dotnet"  ; do \
		$(TFGEN) $$LANGUAGE --overlays overlays/$$LANGUAGE/ --out ../${PACKDIR}/$$LANGUAGE/ || exit 3 ; \
	done

lint::
	#golangci-lint run

install:: provider
	

test_fast::
	cd examples && $(GO_TEST_FAST) .

test_all::
	cd examples && $(GO_TEST) .

.PHONY: publish_tgz
publish_tgz:
	$(call STEP_MESSAGE)
	./scripts/publish_tgz.sh

.PHONY: publish_packages
publish_packages:
	$(call STEP_MESSAGE)
	$$(go env GOPATH)/src/github.com/pulumi/scripts/ci/publish-tfgen-package .
	$$(go env GOPATH)/src/github.com/pulumi/scripts/ci/build-package-docs.sh ${PACK}

.PHONY: check_clean_worktree
check_clean_worktree:
	$$(go env GOPATH)/src/github.com/pulumi/scripts/ci/check-worktree-is-clean.sh

# The travis_* targets are entrypoints for CI.
.PHONY: travis_cron travis_push travis_pull_request travis_api
travis_cron: all
travis_push: only_build check_clean_worktree publish_tgz only_test publish_packages
travis_pull_request: all check_clean_worktree
travis_api: all
