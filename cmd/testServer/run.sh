#!/bin/bash
if [ ! -e swagger-ui ]; then
    git clone https://github.com/swagger-api/swagger-ui
fi

sed -i 's/http:.*json/http:\/\/localhost:8080\/docs\/apidocs.json/' swagger-ui/dist/index.html

go run main.go
