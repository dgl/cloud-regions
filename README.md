# Cloud provider regions

This is data on cloud provider regions.

## Providers

Google, Google Cloud Platform (GCP)
Amazon, Amazon Web Services (AWS)

TODO

Vultr, Vultr Cloud (Vultr)
Azure
Upcloud
Linode
Oracle
IBM

## Adding or updating data

The raw data is in data.csv in each provider's directory. Update that and send a PR.

Most regions currently refer to a coarse area, such as "London". If you know
which datacentres a region is actually hosted in, we will happily take updates.

However there are cases where the parts (to use neutral terminology) of a
region are far apart, we may editorialise in those cases and pick a single
datacentre in a suitable midpoint. Sometimes the datacenters are actually
outside the declared "region" by the provider, that's fine.

## Adding a provider

Make a data.csv.

- "region" must match the region used in the API for the provider
- "location_name" ideally should match the description the provider uses
- "country_tld" should be the country code, with a bias towards the form used
  in TLDs (hk rather than cn, uk rather than gb)

The rest of the fields should be as `./get-loc.sh "London, England"` outputs.

Explain in the README roughly how you did it, for when it comes time to update.

## Using the data

Currently parsing data.csv, aiming to provide some format transforms.

## License

This data uses data from OpenStreetMap, which is:

  © OpenStreetMap contributors, ODbL 1.0. https://osm.org/copyright

As a result this work is licensed under the same licence:

  This data is made available under the Open Database License:
  http://opendatacommons.org/licenses/odbl/1.0/. Any rights in individual
  contents of the database are licensed under the Database Contents License:
  http://opendatacommons.org/licenses/dbcl/1.0/
