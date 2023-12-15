```
git clone https://github.com/thxrhmn/ip-info.git
cd ip-info
go mod tidy
go run . -rod="show,trace"
```

```json
{
  "ip": "167.94.145.52",
  "latitude": "42.2641324",
  "longitude": "-83.7185548",
  "country": "US",
  "region": "Michigan",
  "city": "Ann Arbor",
  "postalCode": "48104",
  "utcOffset": "-04:00",
  "asn": {
    "network": "167.94.145.0/24",
    "asn": 398705,
    "domain": "censys.io",
    "name": "Censys, Inc.",
    "countryCode": "US"
  },
  "company": {
    "network": "167.94.145.0/24",
    "domain": "censys.io",
    "name": "Censys, Inc.",
    "address": "116 1/2 S Main Street, Ann Arbor, MI, 48104, United States"
  },
  "privacy": {
    "vpn": false,
    "proxy": false,
    "tor": false,
    "hosting": false
  },
  "abuse": {
    "network": "167.94.145.0/24",
    "name": "Censys Abuse Team",
    "email": "scan-abuse@censys.io",
    "phone": "+1-248-629-0125",
    "address": "116 1/2 S Main Street, Ann Arbor, MI, 48104, United States",
    "countryCode": ""
  }
}
```