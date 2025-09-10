#!/bin/bash

VERSION_PKG="github.com/chhz0/projectx-go/pkg/version"

get_git_info() {
    if git describe --tags --exact-match="v*" >/dev/null 2>&1; then
        VERSION_INFO=$(git describe --tags --match="v*")
    else
        COMMIT_TIME=$(git show -s --format=%ct HEAD 2>/dev/null)
        SHORT_HASH=$(git rev-parse --short=12 HEAD 2>/dev/null)

        if [ -n "$COMMIT_TIME" ] && [ -n "$SHORT_HASH" ]; then
            DATE_PART=$(date -u -d @"$COMMIT_TIME" +%Y%m%d%H%M%S 2>/dev/null || date -u -r "$COMMIT_TIME" +%Y%m%d%H%M%S 2>/dev/null)
            VERSION_INFO="v0.0.0-${DATE_PART}-${SHORT_HASH}"
        else
            VERSION_INFO="v0.0.0-unknown"
        fi
    fi

    GIT_COMMIT=$(git rev-parse HEAD 2>/dev/null)
    GIT_COMMIT_STAMP=$(git show -s --format=%ct HEAD 2>/dev/null)
    GIT_BRANCH=$(git rev-parse --abbrev-ref HEAD 2>/dev/null)

    GIT_STATE="clean"
    if ! git diff-index --quiet HEAD -- 2>/dev/null; then
        GIT_STATE="dirty"
    fi

    echo "${VERSION_INFO}:${GIT_COMMIT}:${GIT_COMMIT_STAMP}:${GIT_BRANCH}:${GIT_STATE}"
}

get_build_info() {
    date -u +%Y-%m-%dT%H:%M:%SZ
}

gen_ldflags() {
    IFS=':' read -r version_info git_commit git_commit_stamp git_branch git_state <<EOF
$(get_git_info)
EOF

    local build_date
    build_date="$(get_build_info)"

    local flags=""
    add_flag() {
        if [ -n "$2" ]; then
            flags="$flags -X '$VERSION_PKG.$1=$2'"
        fi
    }

    add_flag "version" "$version_info"
    add_flag "gitCommit" "$git_commit"
    add_flag "gitCommitStamp" "$git_commit_stamp"
    add_flag "gitBranch" "$git_branch"
    add_flag "gitState" "$git_state"
    add_flag "buildDate" "$build_date"

    echo "$flags"
}

main() {
    gen_ldflags
}

if [[ "${BASH_SOURCE[0]}" = "${0}" ]]; then
    main "$@"
fi
