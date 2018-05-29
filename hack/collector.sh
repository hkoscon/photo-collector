#!/bin/bash

export PRIVATE_KEY_PATH=private-key.pem
export KEY_LABEL=photos
export S3_ENDPOINT=storage.local:9000
export S3_ACCESS_KEY_ID=AccessKeyID
export S3_SECRET_ACCESS_KEY=SecretAccessKey

exec collector password