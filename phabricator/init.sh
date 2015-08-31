/opt/phabricator/bin/storage upgrade --force
/opt/phabricator/bin/config set phabricator.base-uri ${PHAB_SERVER}
/opt/phabricator/bin/config set phabricator.show-prototypes 'true'
/opt/phabricator/bin/config set diffusion.ssh-user git
