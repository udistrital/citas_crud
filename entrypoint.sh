#!/usr/bin/env bash

set -e
set -u
set -o pipefail

if [ -n "${PARAMETER_STORE:-}" ]; then
  export CITAS_CRUD_PGUSER="$(aws ssm get-parameter --name /${PARAMETER_STORE}/citas_crud/db/username --output text --query Parameter.Value)"
  export CITAS_CRUD_PGPASS="$(aws ssm get-parameter --with-decryption --name /${PARAMETER_STORE}/citas_crud/db/password --output text --query Parameter.Value)"
fi

exec ./main "$@"
