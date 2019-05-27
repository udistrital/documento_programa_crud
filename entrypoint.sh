#!/usr/bin/env bash

set -e
set -u
set -o pipefail

if [ -n "${PARAMETER_STORE:-}" ]; then
  export ENTE_CRUD__PGUSER="$(aws ssm get-parameter --name /${PARAMETER_STORE}/ente_crud/db/username --output text --query Parameter.Value)"
  export ENTE_CRUD__PGPASS="$(aws ssm get-parameter --with-decryption --name /${PARAMETER_STORE}/ente_crud/db/password --output text --query Parameter.Value)"
fi
DOCUMENTO_PROGRAMA_CRUD__PGDB=$DOCUMENTO_PROGRAMA_CRUD__PGDB DOCUMENTO_PROGRAMA_CRUD__PGPASS=$DOCUMENTO_PROGRAMA_CRUD__PGPASS DOCUMENTO_PROGRAMA_CRUD__PGURLS=$DOCUMENTO_PROGRAMA_CRUD__PGURLS DOCUMENTO_PROGRAMA_CRUD__PGUSER=$DOCUMENTO_PROGRAMA_CRUD__PGURLS DOCUMENTO_PROGRAMA_CRUD__SCHEMA=$DOCUMENTO_PROGRAMA_CRUD__SCHEMA DOCUMENTO_PROGRAMA_HTTP_PORT=$DOCUMENTO_PROGRAMA_HTTP_PORT ./main orm syncdb -db="default" -force=true -v=true
exec ./main "$@" 

