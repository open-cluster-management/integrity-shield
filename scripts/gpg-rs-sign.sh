#!/bin/bash
#
# Copyright 2020 IBM Corporation.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

CMDNAME=`basename $0`
if [ $# -ne 3 ]; then
  echo "Usage: $CMDNAME <signer-email> <input-file> <output-file>" 1>&2
  exit 1
fi

if [ ! -e $2 ]; then
  echo "$2 does not exist"
  exit 1
fi

if ! [ -x "$(command -v yq)" ]; then
   echo 'Error: yq is not installed.' >&2
   exit 1
fi

SIGNER=$1
INPUT_FILE=$2
OUTPUT_FILE=$3

if [ -z $SIGNER ]; then
   echo "Signer email is empty, please provide it."
   exit 1
fi

if [ ! -f $INPUT_FILE ]; then
   echo "Input file does not exist, please create it."
   exit 1
fi

if [ -z "$TMP_DIR" ]; then
    echo "TMP_DIR is empty. Setting /tmp as default"
    TMP_DIR="/tmp"
fi

if [ ! -d $TMP_DIR ]; then
    echo "$TMP_DIR directory does not exist, please create it."
    exit 1
fi

RSC_TEMPLATE=""

# compute signature (and encoded message and certificate)
RSC_TEMPLATE=`cat << EOF
apiVersion: apis.integrityshield.io/v1alpha1
kind: ResourceSignature
metadata:
   annotations:
      integrityshield.io/messageScope: spec
      integrityshield.io/signature: ""
   name: ""
spec:
   data:
   - message: ""
     signature: ""
     type: resource
EOF`

if [[ "$OSTYPE" == "linux-gnu"* ]]; then
    base='base64 -w 0'
elif [[ "$OSTYPE" == "darwin"* ]]; then
    base='base64'
fi

YQ_VERSION=$(yq --version 2>&1 | awk '{print $3}' | cut -c 1 )

# message
msg=`cat $INPUT_FILE | gzip -c | $base`

# signature
sig=`cat ${INPUT_FILE} > $TMP_DIR/temp-aaa.yaml; gpg -u $SIGNER --detach-sign --armor --output - $TMP_DIR/temp-aaa.yaml | $base`
sigtime=`date +%s`

if [ -f $OUTPUT_FILE ]; then
    rm $OUTPUT_FILE
fi

if [[ $YQ_VERSION == "3" ]]; then
   RSC_COUNT=`yq r -d'*' ${INPUT_FILE} --tojson | wc -l`
   indx=0
   yq r -d'*'  ${INPUT_FILE} -j | while read doc;
   do
         resApiVer=$(echo $doc | yq r - 'apiVersion' | tr / _)
         if [ -z "$resApiVer" ] || [ "$resApiVer" = "null" ] ; then
            break
         fi

         resKind=$(echo $doc | yq r - 'kind')
         reslowerkind=$(echo $resKind | tr "[:upper:]" "[:lower:]")
         resname=$(echo $doc | yq r - 'metadata.name')
         rsigname="rsig-${reslowerkind}-${resname}"

         echo "$RSC_TEMPLATE" >> $OUTPUT_FILE
         yq w -i -d$indx $OUTPUT_FILE metadata.name $rsigname
         yq w -i -d$indx $OUTPUT_FILE 'metadata.labels."integrityshield.io/sigobject-apiversion"' $resApiVer
         yq w -i -d$indx $OUTPUT_FILE 'metadata.labels."integrityshield.io/sigobject-kind"' $resKind
         yq w -i -d$indx --tag !!str $OUTPUT_FILE 'metadata.labels."integrityshield.io/sigtime"' $sigtime

         indx=$[$indx+1]
         if (( $indx < $RSC_COUNT )) ; then
            echo "---" >> $OUTPUT_FILE
        fi
    done
elif [[ $YQ_VERSION == "4" ]]; then
        indx=0
        while true
        do
           resApiVer=$(yq eval ".apiVersion | select(di == $indx)" ${INPUT_FILE}  | sed 's/\//_/g')
           resKind=$(yq eval ".kind | select(di == $indx)" ${INPUT_FILE}  | sed 's/\//_/g')
           reslowerkind=$(echo $resKind | tr '[:upper:]' '[:lower:]')
           resname=$(yq eval ".metadata.name | select(di == $indx)" ${INPUT_FILE})
           rsigname="rsig-${reslowerkind}-${resname}"

           if [ -z "$resApiVer" ]; then
              break
           else
              TMP_FILE="$TMP_DIR/${rsigname}.yaml"
              echo "$RSC_TEMPLATE" >> ${TMP_FILE}
              yq eval ".metadata.name = \"$rsigname\"" -i $TMP_FILE
              yq eval ".metadata.labels.\"integrityshield.io/sigobject-apiversion\" = \"$resApiVer\"" -i $TMP_FILE
              yq eval ".metadata.labels.\"integrityshield.io/sigobject-kind\" =  \"$resKind\"" -i $TMP_FILE
              yq eval ".metadata.labels.\"integrityshield.io/sigtime\" = \"$sigtime\"" -i $TMP_FILE

              if [ -f $TMP_FILE ]; then
                cat $TMP_FILE >> $OUTPUT_FILE
                rm $TMP_FILE
              fi

              echo "---" >> $OUTPUT_FILE
              indx=$[$indx+1]
           fi
        done
fi

if [[ "$OSTYPE" == "linux-gnu"* ]]; then
    sed -i '$ s/---//g' $OUTPUT_FILE
elif [[ "$OSTYPE" == "darwin"* ]]; then
    sed -i '' '$ s/---//g' $OUTPUT_FILE
fi

if [[ $YQ_VERSION == "3" ]]; then
   yq w -i -d* $OUTPUT_FILE spec.data.[0].message $msg
   yq w -i -d* $OUTPUT_FILE spec.data.[0].signature $sig
elif [[ $YQ_VERSION == "4" ]]; then
   yq eval ".spec.data.[0].message |= \"$msg\"" -i $OUTPUT_FILE
   yq eval ".spec.data.[0].signature |= \"$sig\""  -i $OUTPUT_FILE
fi

# resource signature spec content
if [[ $YQ_VERSION == "3" ]]; then
   rsigspec=`cat $OUTPUT_FILE | yq r - -j |jq -r '.spec' | yq r - --prettyPrint | $base`
elif [[ $YQ_VERSION == "4" ]]; then
   rsigspec=`yq eval '.spec' $OUTPUT_FILE | $base`
fi

# resource signature signature
rsigsig=`echo -e "$rsigspec" > $TMP_DIR/temp-rsig.yaml; gpg -u $SIGNER --detach-sign --armor --output - $TMP_DIR/temp-rsig.yaml | $base`
if [[ $YQ_VERSION == "3" ]]; then
   yq w -i -d* $OUTPUT_FILE 'metadata.annotations."integrityshield.io/signature"' $rsigsig
elif [[ $YQ_VERSION == "4" ]]; then
   yq eval ".metadata.annotations.\"integrityshield.io/signature\" = \"$rsigsig\"" -i $OUTPUT_FILE
fi

if [ -f $TMP_DIR/temp-aaa.yaml ]; then
   rm $TMP_DIR/temp-aaa.yaml
fi

if [ -f $TMP_DIR/temp-rsig.yaml ]; then
   rm $TMP_DIR/temp-rsig.yaml
fi
