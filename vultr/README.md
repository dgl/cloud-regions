
curl -s https://api.vultr.com/v2/regions | jq -r '.regions[] | [.id, (.city + ", " + .country + ", " + .continent)] | @csv'
