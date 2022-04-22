
curl -s https://ip-ranges.amazonaws.com/ip-ranges.json | jq -r '.prefixes[]["region"]' | sort -u

Then manually add the names from https://aws.amazon.com/about-aws/global-infrastructure/regions_az/?p=ngi&loc=2
and some guesswork with other sources (latency to the IPs from various
locations can give away an approximate location).
