#!/bin/bash
. /etc/apache2/envvars
sed -i -e "s/%PHAB_DB_PASSWD%/${PHAB_DB_PASSWD}/g" /opt/phabricator/conf/local/local.json
sed -i -e "s/%PHAB_DB_SERVER%/${PHAB_DB_SERVER}/g" /opt/phabricator/conf/local/local.json
/opt/phabricator/bin/config set phabricator.base-uri ${PHAB_SERVER}
/opt/phabricator/bin/config set phabricator.show-prototypes 'true'
/opt/phabricator/bin/config set diffusion.ssh-user git
/opt/phabricator/bin/phd start --force
/usr/sbin/sshd -f /etc/ssh/sshd_config.phabricator
/usr/sbin/apache2 -DFOREGROUND
