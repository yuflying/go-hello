#!/bin/sh -e

#set some environment variables
readonly WORKDIR=$(cd $(dirname "${BASH_SOURCE}")/.. && pwd -P)
readonly BUILD_ROOT="${WORKDIR}/src"
readonly BUILD_OUTPUT="${BUILD_ROOT}/dist"
readonly BUILD_OUTPUT_BINPATH="${BUILD_OUTPUT}/bin"

readonly BUILD_TARGETS=(
	cmd/hello
	# cmd/client
	# cmd/biz1
	# cmd/biz2
)

export GOPATH=$WORKDIR

build_target() {
	for arg; do
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $GO_BUILD_FLAGS \
		-installsuffix cgo -ldflags "$GO_LDFLAGS" \
		-o ${BUILD_OUTPUT_BINPATH}/${arg##*/}.x ./${arg} || return
	done
}

build_make_ldflag() {
    local key=${1}
    local val=${2}
    echo "-X go-hello/cmd/version.${key}=${val}"
}

# Prints the value that needs to be passed to the -ldflags parameter of go build
# in order to set the project on the git tree status.
build_version_ldflags() {
	local -a ldflags=($(build_make_ldflag "buildDate" "$(date -u +'%Y-%m-%dT%H:%M:%SZ')"))

	local git_sha=`git rev-parse --short HEAD || echo "GitNotFound"`
	if [ ! -z "$FAILPOINTS" ]; then
		git_sha="$git_sha"-FAILPOINTS
	fi

	ldflags+=($(build_make_ldflag "gitSHA" "${git_sha}"))

	echo "${ldflags[*]-}"
}

# only build when called directly, not sourced
if echo "$0" | grep "build.sh$" >/dev/null; then
	build_version_ldflags
	build_target "${BUILD_TARGETS[@]}"
fi