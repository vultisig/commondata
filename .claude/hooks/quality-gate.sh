#!/bin/bash
# TaskCompleted hook: verify build passes before task completion
# Exit 2 to block, exit 0 to allow

cd "$CLAUDE_PROJECT_DIR" || exit 0

MODIFIED=$(
  {
    git diff --name-only --diff-filter=ACMR HEAD
    git ls-files --others --exclude-standard
  } 2>/dev/null | grep -E '\.(go|proto)$' || true
)

if [ -z "$MODIFIED" ]; then
  exit 0
fi

if [ -f Makefile ] && ! make build 2>&1; then
  echo "Quality gate: build failed" >&2
  exit 2
fi

exit 0
