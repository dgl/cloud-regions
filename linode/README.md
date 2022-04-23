
curl -s https://api.linode.com/v4/regions | jq -r '.data[] | [.id, .country] | @csv'

Then cross reference the API region names with https://www.linode.com/global-infrastructure/
