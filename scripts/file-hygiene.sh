#!/usr/bin/env bash
#
# file-hygiene.sh performs lightweight repository hygiene checks that used to
# live in .pre-commit-config.yaml. It is intended to be run in CI against the
# full tracked tree; it exits non-zero with a summary of offending files when
# any check fails.
#
# Checks performed:
#   * trailing whitespace
#   * missing final newline (end-of-file-fixer)
#   * UTF-8 byte order mark (BOM) at the start of text files
#   * carriage returns (CR/CRLF) in text files
#   * unresolved Git merge conflict markers
#   * files larger than 500 KB
#   * case-insensitive filename collisions
#   * symlinks pointing at non-existent targets
#   * executable files missing a shebang
#   * shebang scripts not marked executable
#   * presence of Git submodules (forbidden)
#
# Compatible with bash 3.2 (macOS) and bash >= 4 (Linux CI).

set -u -o pipefail

cd "$(git rev-parse --show-toplevel)"

fail=0
note() { printf '::error::%s\n' "$*" >&2; fail=1; }
section() { printf '\n>>> %s\n' "$*"; }

tmpdir=$(mktemp -d)
trap 'rm -rf "$tmpdir"' EXIT

all_files="$tmpdir/all"
text_files="$tmpdir/text"

# Full list of tracked files (NUL-delimited).
git ls-files -z >"$all_files"

# Subset that git/grep treat as text. We skip anything with the `binary`
# gitattribute, and anything grep -I considers binary.
: >"$text_files"
while IFS= read -r -d '' f; do
  [ -f "$f" ] || continue
  if git check-attr binary -- "$f" 2>/dev/null | grep -q ': binary: set$'; then
    continue
  fi
  if LC_ALL=C grep -Iq . "$f" 2>/dev/null; then
    printf '%s\0' "$f" >>"$text_files"
  fi
done <"$all_files"

report_file="$tmpdir/report"

run_text_check() {
  local label="$1"; shift
  local message="$1"; shift
  section "$label"
  : >"$report_file"
  while IFS= read -r -d '' f; do
    if "$@" "$f"; then
      printf '  %s\n' "$f" >>"$report_file"
    fi
  done <"$text_files"
  if [ -s "$report_file" ]; then
    cat "$report_file"
    note "$message"
  fi
}

has_trailing_ws()   { LC_ALL=C grep -qE $'[ \t]+$' "$1"; }
missing_final_nl()  { [ -s "$1" ] && [ "$(tail -c1 "$1" | wc -l | tr -d ' ')" = "0" ]; }
has_bom()           { [ "$(head -c3 "$1" | od -An -tx1 | tr -d ' \n')" = "efbbbf" ]; }
has_cr()            { LC_ALL=C grep -lU $'\r' "$1" >/dev/null 2>&1; }
has_conflict()      { LC_ALL=C grep -qE '^(<<<<<<<|=======|>>>>>>>)( |$)' "$1"; }

run_text_check "trailing whitespace"    "files contain trailing whitespace"                 has_trailing_ws
run_text_check "missing final newline"  "files do not end with a newline"                    missing_final_nl
run_text_check "UTF-8 BOM"              "files start with a UTF-8 BOM"                       has_bom
run_text_check "carriage returns"       "files contain carriage returns (CR/CRLF)"           has_cr
run_text_check "merge conflict markers" "files contain unresolved merge conflict markers"    has_conflict

section "large files (>500 KB)"
: >"$report_file"
while IFS= read -r -d '' f; do
  [ -f "$f" ] || continue
  size=$(wc -c <"$f" | tr -d ' ')
  if [ "$size" -gt $((500 * 1024)) ]; then
    printf '  %s (%s bytes)\n' "$f" "$size" >>"$report_file"
  fi
done <"$all_files"
if [ -s "$report_file" ]; then
  cat "$report_file"
  note "files exceed 500 KB"
fi

section "case-insensitive filename collisions"
dupes=$(git ls-files | awk '{ lc=tolower($0); count[lc]++; names[lc]=names[lc] "\n  " $0 } END { for (k in count) if (count[k] > 1) print names[k] }')
if [ -n "$dupes" ]; then
  printf '%s\n' "$dupes"
  note "files differ only by case"
fi

section "broken symlinks"
: >"$report_file"
while IFS= read -r -d '' f; do
  if [ -L "$f" ] && [ ! -e "$f" ]; then
    printf '  %s\n' "$f" >>"$report_file"
  fi
done <"$all_files"
if [ -s "$report_file" ]; then
  cat "$report_file"
  note "symlinks point to non-existent targets"
fi

section "executables without shebangs"
: >"$report_file"
while IFS= read -r -d '' f; do
  [ -f "$f" ] || continue
  [ -L "$f" ] && continue
  if [ -x "$f" ]; then
    if [ "$(head -c2 "$f" 2>/dev/null)" != "#!" ]; then
      printf '  %s\n' "$f" >>"$report_file"
    fi
  fi
done <"$all_files"
if [ -s "$report_file" ]; then
  cat "$report_file"
  note "executable files are missing a shebang"
fi

section "shebang scripts not marked executable"
: >"$report_file"
while IFS= read -r -d '' f; do
  [ -f "$f" ] || continue
  [ -L "$f" ] && continue
  [ -x "$f" ] && continue
  if [ "$(head -c2 "$f" 2>/dev/null)" = "#!" ]; then
    printf '  %s\n' "$f" >>"$report_file"
  fi
done <"$all_files"
if [ -s "$report_file" ]; then
  cat "$report_file"
  note "files with a shebang are not executable"
fi

section "git submodules"
if [ -f .gitmodules ] || git ls-files --stage | awk '{print $1}' | grep -q '^160000$'; then
  note "repository contains Git submodules (forbidden)"
fi

echo ""
if [ "$fail" -ne 0 ]; then
  echo "file-hygiene.sh: one or more checks failed" >&2
  exit 1
fi
echo "file-hygiene.sh: all checks passed"
