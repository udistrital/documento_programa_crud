#!/usr/bin/env bash

set -e
set -u
set -o pipefail

if [ -n "${PARAMETER_STORE:-}" ]; then
  export DOCUMENTO_PROGRAMA_CRUD__PGUSER="$(aws ssm get-parameter --name /${PARAMETER_STORE}/documento_programa_crud/db/username --output text --query Parameter.Value)"
  export DOCUMENTO_PROGRAMA_CRUD__PGPASS="$(aws ssm get-parameter --with-decryption --name /${PARAMETER_STORE}/documento_programa_crud/db/password --output text --query Parameter.Value)"
fi
DOCUMENTO_PROGRAMA_CRUD__PGDB_MIGRATION=$DOCUMENTO_PROGRAMA_CRUD__PGDB_MIGRATION ./main orm syncdb -db="default" -force=true -v=true
exec ./main "$@" 

