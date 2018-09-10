#!/bin/sh


echo "logging in docker..."
docker login --username=_ --password=$HEROKU_AUTH registry.heroku.com

echo "builing production image..."
docker build -t registry.heroku.com/$HEROKU_APP_NAME/web . --build-arg app_env=production

echo "pushing to heroku..."
docker push registry.heroku.com/$HEROKU_APP_NAME/web

echo "releasing..."
heroku container:release web

