#!/bin/bash

json2yml() {
 python2 -c 'import sys, yaml, json; y=json.load(sys.stdin); print yaml.safe_dump(y)'
}

PKG=$1
VERSION=$2
TMP=$(mktemp)
PKG_NAME=${PKG}.${VERSION}
CHANNEL=$(sqlite3 operatorhub-bundles.db "SELECT name FROM channel WHERE head_operatorbundle_name='${PKG_NAME}'")
if [[ "${CHANNEL}" == "" ]]; then
    echo "${PKG_NAME} not found"
    exit 1
fi
sqlite3 operatorhub-bundles.db "SELECT bundle FROM operatorbundle WHERE name='${PKG_NAME}'"  > $TMP
DIR=operators/$PKG/$CHANNEL/$VERSION
mkdir -p ./$DIR/manifests
CSV_YAML=$DIR/manifests/$PKG.clusterserviceversion.yaml
IFS=$'\n'
for row in $(cat $TMP); do
   KIND=$(echo ${row} |jq -r .kind)
   if [[ "$KIND" == "ClusterServiceVersion" ]]; then
      echo -n "Exporting to $CSV_YAML..."
      echo ${row} | jq -r .| json2yml > $CSV_YAML
      echo "Done."
      continue
   fi
   if [[ "$KIND" == "CustomResourceDefinition" ]]; then
     NAME=$(echo ${row} | jq -r .spec.names.plural)
     CRD_YAML=$DIR/manifests/$NAME.crd.yaml
     echo -n "Exporting to $CRD_YAML..."
     echo ${row} | jq -r . | json2yml > $CRD_YAML
     echo "Done."
   fi
done
rm -rf $TMP
