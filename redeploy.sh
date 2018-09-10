#!/usr/bin/env bash

 docker login --username=_ --password=$HEROKU_PW registry.heroku.com

 docker build -t registry.heroku.com/$HEROKU_APP_NAME/web . --build-arg app_env=production

 docker push registry.heroku.com/$HEROKU_APP_NAME/web

 heroku container:release web

