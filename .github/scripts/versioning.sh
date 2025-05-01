#!/bin/bash


# Function to increment the version based on the type
increment_version() {
  local version_type="$1"
  local IFS="."
  local -a parts=($current_version)

  case "$version_type" in
    major)
      parts[0]=$((parts[0] + 1))
      parts[1]=0
      parts[2]=0
      ;;
    minor)
      parts[1]=$((parts[1] + 1))
      parts[2]=0
      ;;
    patch)
      parts[2]=$((parts[2] + 1))
      ;;
    *)
      echo "Invalid version type: $version_type"
      return 1
      ;;
  esac

  echo "${parts[0]}.${parts[1]}.${parts[2]}"
}

# Get the current version from the user or default to 0.0.0
current_version="${1:-0.0.0}"

# Get the version type from the user or default to patch
version_type="${2:-patch}"


current_version_git=$(git tag --sort=-committerdate | head -2 | awk '{split($0, tags, "\n")} END {print tags[1]}')
current_version=$(echo "$current_version_git" | sed 's/[^0-9.]*\([0-9.]*\).*/\1/' )

echo $current_version

# Validate the current version format
if ! [[ "$current_version" =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
  echo "Invalid version format: $current_version. Please use X.Y.Z format."
  exit 1
fi

# Increment the version
new_version=$(increment_version "$current_version" "$version_type")

# Check if increment_version failed
if [ $? -ne 0 ]; then
  exit 1
fi

# Output the new version
echo "New version: $new_version"