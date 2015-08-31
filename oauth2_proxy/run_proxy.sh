#!/bin/bash
CLIENT_ID=$(cat /conf/oauth2_client_id)
CLIENT_KEY=$(cat /conf/oauth2_client_key)
CMD="bin/oauth2_proxy \
     --email-domain=${EMAIL_DOMAIN} \
     --login-url=${LOGIN_URL} \
     --upstream=${UPSTREAM} \
     --redirect-url=${REDIRECT_URL} \
     --cookie-secret=${COOKIE_SECRET} \
     --provider="phabricator" \
     --http-address="0.0.0.0:4180" \
     --client-id=${CLIENT_ID} \
     --client-secret=${CLIENT_KEY}"
echo $CMD
$CMD
