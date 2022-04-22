#!/bin/zsh
# Usage: EMAIL=you@domain ./get-loc.sh "some, place"
set -euo pipefail

place="$(echo "$1" | sed 's/ /%20/g')"
data=$(curl -s --header "User-Agent: manual data collection, ${EMAIL?missing}" "https://nominatim.openstreetmap.org/search.php?q=$place&format=jsonv2")
jq '.[0]' <<<"$data"
echo
jq -r '[.[0].lat, .[0].lon, .[0].osm_id] | @csv' <<<"$data"
