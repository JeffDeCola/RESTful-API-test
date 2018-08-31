#!/bin/bash
# RESTful-API-test update_concourse.sh

fly -t ci set-pipeline -p RESTful-API-test -c ci/pipeline.yml --load-vars-from ../.credentials.yml
