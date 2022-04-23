Sadly I can't find an unauthenticated source of the data.

However the az CLI has an nice JSON output:

az account list-locations -o json

Which even includes latitude and longitude...

az account list-locations -o json | jq -r '.[] | select(.metadata.regionType == "Physical") | [.name, (.metadata.physicalLocation + ", " + .displayName + ", " + .metadata.geographyGroup), .metadata.latitude, .metadata.longitude] | @csv'

However, some of those locations are just plain wrong. For example "ukwest"
(apparently in Cardiff, Wales, UK) has a location somewhere near Liverpool.
