EDIT 2020-01-18: THE MAJORITY OF SERVERS FROM NORDVPN DOESN'T ALLOW TO CONNECT WITH HTTP OR HTTPS PROXY

# nordvpn-http-proxies
Get the NordVPN servers that are HTTP and HTTPS. The config.js allow to filter the servers by load, country and proxy type.

A built application is available to download for Windows 64 bits and Linux 64 bits inside the dist folder.

This are the different options for the config.js 

**Example 1**
```
{
  "max_load": 40, // max load of each server, between 0 and 100
  "country_code": ["ES", "FR", "IT", "DE"], // array of countries
  "proxy": "proxy_ssl" // this is for HTTPS proxies
}
```

**Example 2**
```
{
  "max_load": 60,
  "country_code": [], // an empty array for not filtering
  "proxy": "" // an empty string for not filtering
}
```

**Example 3**
```
{
  "max_load": 60,
  "country_code": [],
  "proxy": "proxy" // this is for HTTP proxies
}
```

## Output
An Excel file (XLSX format) is created in the same folder

![Output file](https://img.imgur.com/0oDrAM4.png)
