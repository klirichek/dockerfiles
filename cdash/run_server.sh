#!/bin/bash
if [ "$#" -ne 1 ]; then
  echo "Usage: $0 port";exit
fi
sed -i -e "s/%CDASH_DB_PASSWD%/${CDASH_DB_PASSWD}/g" /var/www/html/CDash/cdash/config.local.php
sed -i -e "s/%CDASH_DB_SERVER%/${CDASH_DB_SERVER}/g" /var/www/html/CDash/cdash/config.local.php
sed -i -e "s/%CDASH_SERVER%/${CDASH_SERVER}/g" /var/www/html/CDash/cdash/config.local.php
sed -i -e "s/%CDASH_EMAIL%/${CDASH_EMAIL}/g" /var/www/html/CDash/cdash/config.local.php
sed -i -e "s/%CDASH_REPLY_EMAIL%/${CDASH_REPLY_EMAIL}/g" /var/www/html/CDash/cdash/config.local.php
sed -e s/%PORT%/$1/g cdash.conf.in > /etc/apache2/conf-enabled/cdash.conf
a2enmod rewrite
apache2-foreground
