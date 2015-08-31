FROM php:5.6-apache
MAINTAINER David Lacey <davel@xmos.com>
RUN apt-get update && apt-get install -y unzip libxslt-dev libpng-dev libcurl4-openssl-dev
RUN docker-php-ext-install pdo mysql xsl curl gd
RUN curl -o /var/www/html/CDash-2-2-3.zip http://www.cdash.org/wp-content/uploads/2014/11/CDash-2-2-3.zip
RUN unzip CDash-2-2-3.zip
RUN cp -r CDash-2-2-3 /var/www/html/CDash
ADD cdash-config.local.php /var/www/html/CDash/cdash/config.local.php
ADD cdash.conf.in /var/www/html/cdash.conf.in
ADD php.ini /usr/local/etc/php/php.ini
RUN mkdir /var/www/html/CDash/rss
RUN chmod -R a+rw /var/www/html/CDash
RUN chown -R www-data /var/www/html/CDash
COPY run_server.sh /var/www/html/run_server.sh
RUN chmod a+rwx /var/www/html/run_server.sh
ENTRYPOINT ["./run_server.sh"]
CMD []
