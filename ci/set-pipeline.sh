#!/bin/bash
# RESTful-API-test set-pipeline.sh

fly -t ci set-pipeline -p RESTful-API-test -c pipeline.yml --load-vars-from ../../../../../.credentials.yml
