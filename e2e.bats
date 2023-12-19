#!/usr/bin/env bats

@test "accept" {
  run kwctl run annotated-policy.wasm -r test_data/request.json -s test_data/settings.json

  # this prints the output when one the checks below fails
  echo "output = ${output}"

  # request allowed
  [ "$status" -eq 0 ]
  [ $(expr "$output" : '.*allowed.*true') -ne 0 ]
}


