Manually constructed from https://cloud.google.com/compute/docs/regions-zones

Can also use: curl -s https://www.gstatic.com/ipranges/cloud.json | jq -r '.prefixes[]["scope"]' | sort -u
(per https://cloud.google.com/compute/docs/faq) like AWS.

In IP ranges but not elsewhere:

- europe-west9
- us-east7
- us-central2
