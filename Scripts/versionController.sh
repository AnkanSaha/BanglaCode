#!/bin/bash

# ====================================
#     BanglaCode Version Controller
# ====================================
# This script updates version across all project files

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
BOLD='\033[1m'
NC='\033[0m' # No Color

# Project root (parent of Scripts folder)
PROJECT_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

# Files to update
VERSION_FILE="$PROJECT_ROOT/VERSION"
VSCODE_PACKAGE="$PROJECT_ROOT/Extension/package.json"
MAIN_GO="$PROJECT_ROOT/main.go"
REPL_GO="$PROJECT_ROOT/src/repl/repl.go"
DOCS_PACKAGE="$PROJECT_ROOT/Documentation/package.json"

# Get current version
get_current_version() {
    if [ -f "$VERSION_FILE" ]; then
        cat "$VERSION_FILE" | tr -d '\n'
    else
        echo "0.0.0"
    fi
}

# Print banner
print_banner() {
    echo -e "${CYAN}"
    echo "╔════════════════════════════════════════════╗"
    echo "║     BanglaCode Version Controller          ║"
    echo "║          Update version everywhere         ║"
    echo "╚════════════════════════════════════════════╝"
    echo -e "${NC}"
}

# Print current versions from all files
print_current_versions() {
    echo -e "${BOLD}Current versions in project:${NC}"
    echo -e "  ${BLUE}VERSION file:${NC}      $(get_current_version)"

    if [ -f "$VSCODE_PACKAGE" ]; then
        vscode_ver=$(grep '"version"' "$VSCODE_PACKAGE" | head -1 | sed 's/.*"version": *"\([^"]*\)".*/\1/')
        echo -e "  ${BLUE}VSCode Extension:${NC}  $vscode_ver"
    fi

    if [ -f "$MAIN_GO" ]; then
        main_ver=$(grep 'BanglaCode v' "$MAIN_GO" | head -1 | sed 's/.*BanglaCode v\([0-9.]*\).*/\1/')
        echo -e "  ${BLUE}main.go:${NC}           $main_ver"
    fi

    if [ -f "$REPL_GO" ]; then
        repl_ver=$(grep 'const Version' "$REPL_GO" | head -1 | sed 's/.*"\([^"]*\)".*/\1/')
        echo -e "  ${BLUE}repl.go:${NC}           $repl_ver"
    fi

    if [ -f "$DOCS_PACKAGE" ]; then
        docs_ver=$(grep '"version"' "$DOCS_PACKAGE" | head -1 | sed 's/.*"version": *"\([^"]*\)".*/\1/')
        echo -e "  ${BLUE}Documentation:${NC}     $docs_ver"
    fi
    echo ""
}

# Validate version format (X.Y.Z)
validate_version() {
    if [[ $1 =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
        return 0
    else
        return 1
    fi
}

# Compare versions: returns 0 if first > second
ver_gt() {
    test "$(printf '%s\n' "$@" | sort -V | head -n 1)" != "$1"
}

# Update VERSION file
update_version_file() {
    local new_version=$1
    echo "$new_version" > "$VERSION_FILE"
    echo -e "  ${GREEN}✓${NC} VERSION file updated"
}

# Update VSCode Extension package.json
update_vscode_package() {
    local new_version=$1
    if [ -f "$VSCODE_PACKAGE" ]; then
        sed -i "s/\"version\": *\"[^\"]*\"/\"version\": \"$new_version\"/" "$VSCODE_PACKAGE"
        echo -e "  ${GREEN}✓${NC} VSCode Extension package.json updated"
    else
        echo -e "  ${YELLOW}⚠${NC} VSCode Extension package.json not found"
    fi
}

# Update main.go
update_main_go() {
    local new_version=$1
    if [ -f "$MAIN_GO" ]; then
        # Update "BanglaCode vX.Y.Z" in title line
        sed -i 's/BanglaCode v[0-9]\+\.[0-9]\+\.[0-9]\+/BanglaCode v'"$new_version"'/g' "$MAIN_GO"
        # Update version in info section (matches: \033[1;32m2.1.0\033[0m)
        sed -i 's/\\033\[1;32m[0-9]\+\.[0-9]\+\.[0-9]\+\\033\[0m/\\033[1;32m'"$new_version"'\\033[0m/g' "$MAIN_GO"
        echo -e "  ${GREEN}✓${NC} main.go updated"
    else
        echo -e "  ${YELLOW}⚠${NC} main.go not found"
    fi
}

# Update repl.go
update_repl_go() {
    local new_version=$1
    if [ -f "$REPL_GO" ]; then
        sed -i "s/const Version = \"[^\"]*\"/const Version = \"$new_version\"/" "$REPL_GO"
        echo -e "  ${GREEN}✓${NC} repl.go updated"
    else
        echo -e "  ${YELLOW}⚠${NC} repl.go not found"
    fi
}

# Update Documentation package.json
update_docs_package() {
    local new_version=$1
    if [ -f "$DOCS_PACKAGE" ]; then
        sed -i "s/\"version\": *\"[^\"]*\"/\"version\": \"$new_version\"/" "$DOCS_PACKAGE"
        echo -e "  ${GREEN}✓${NC} Documentation package.json updated"
    else
        echo -e "  ${YELLOW}⚠${NC} Documentation package.json not found"
    fi
}

# Interactive menu for version type selection
select_version_type() {
    local current_version=$1
    IFS='.' read -r major minor patch <<< "$current_version"

    echo -e "${BOLD}Select version bump type:${NC}"
    echo ""
    echo -e "  ${CYAN}1)${NC} Patch  (Bug fixes)           → $major.$minor.$((patch + 1))"
    echo -e "  ${CYAN}2)${NC} Minor  (New features)        → $major.$((minor + 1)).0"
    echo -e "  ${CYAN}3)${NC} Major  (Breaking changes)    → $((major + 1)).0.0"
    echo -e "  ${CYAN}4)${NC} Custom (Enter manually)"
    echo ""

    read -p "Enter choice [1-4]: " choice

    case $choice in
        1)
            NEW_VERSION="$major.$minor.$((patch + 1))"
            ;;
        2)
            NEW_VERSION="$major.$((minor + 1)).0"
            ;;
        3)
            NEW_VERSION="$((major + 1)).0.0"
            ;;
        4)
            read -p "Enter new version (X.Y.Z): " NEW_VERSION
            ;;
        *)
            echo -e "${RED}Invalid choice${NC}"
            exit 1
            ;;
    esac
}

# Main function
main() {
    print_banner

    # Get current version
    CURRENT_VERSION=$(get_current_version)

    # Print current versions
    print_current_versions

    # Select version type
    select_version_type "$CURRENT_VERSION"

    # Validate version
    if ! validate_version "$NEW_VERSION"; then
        echo -e "${RED}Error: Invalid version format. Use X.Y.Z (e.g., 1.2.3)${NC}"
        exit 1
    fi

    # Check if new version is greater
    if ! ver_gt "$NEW_VERSION" "$CURRENT_VERSION"; then
        echo -e "${YELLOW}Warning: New version ($NEW_VERSION) is not greater than current ($CURRENT_VERSION)${NC}"
        read -p "Continue anyway? [y/N]: " confirm
        if [[ ! $confirm =~ ^[Yy]$ ]]; then
            echo "Aborted."
            exit 0
        fi
    fi

    echo ""
    echo -e "${BOLD}Updating version: ${YELLOW}$CURRENT_VERSION${NC} → ${GREEN}$NEW_VERSION${NC}"
    echo ""

    # Update all files
    update_version_file "$NEW_VERSION"
    update_vscode_package "$NEW_VERSION"
    update_main_go "$NEW_VERSION"
    update_main_go "$NEW_VERSION"
    update_repl_go "$NEW_VERSION"
    update_docs_package "$NEW_VERSION"

    echo ""
    echo -e "${GREEN}${BOLD}Version updated successfully!${NC}"
    echo ""

    # Show updated versions
    echo -e "${BOLD}Updated versions:${NC}"
    print_current_versions

    echo -e "${CYAN}Next steps:${NC}"
    echo "  1. Review the changes: git diff"
    echo "  2. Commit the changes: git add -A && git commit -m \"Bump version to $NEW_VERSION\""
    echo "  3. Push to main to trigger release: git push origin main"
}

# Run main
main "$@"
