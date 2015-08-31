FROM    debian:jessie
MAINTAINER  David Lacey <david.lacey@xmos.com>

ENV DEBIAN_FRONTEND noninteractive
ENV DEBCONF_NONINTERACTIVE_SEEN true

# TODO: review this dependency list
RUN     apt-get update && apt-get install -y \
	        git \
            apache2 \
            curl \
            libapache2-mod-php5 \
            libmysqlclient18 \
            mercurial \
            mysql-client \
            php-apc \
            php5 \
            php5-apcu \
            php5-cli \
            php5-curl \
            php5-gd \
            php5-json \
            php5-ldap \
            php5-mysql \
            python-pygments \
            sendmail \
            subversion \
            tar \
            openssh-server \
            sudo \
        && apt-get clean && rm -rf /var/lib/apt/lists/*

# For some reason phabricator doesn't have tagged releases. To support 
# repeatable builds use the latest SHA
ADD     download.sh /opt/download.sh
WORKDIR /opt
RUN     bash download.sh phabricator 2eab1c1943
RUN     bash download.sh arcanist    1e5059f16a
RUN     bash download.sh libphutil   80b79a9f2f

# Setup apache
RUN     a2enmod rewrite
ADD     phabricator.conf /etc/apache2/sites-available/phabricator.conf
RUN     ln -s /etc/apache2/sites-available/phabricator.conf \
            /etc/apache2/sites-enabled/phabricator.conf && \
        rm -f /etc/apache2/sites-enabled/000-default.conf
RUN     a2enmod ssl

# Setup phabricator
RUN     mkdir -p /opt/phabricator/conf/local /var/repo
ADD     local.json /opt/phabricator/conf/local/local.json
RUN     sed -i -e 's/post_max_size = 8M/post_max_size = 32M/' /etc/php5/apache2/php.ini
ADD    run_phabricator.sh run_phabricator.sh
ADD    init.sh init.sh

#Setup ssh server
RUN    useradd -m git && \
       echo "git vcs-user ALL=(daemon-user) SETENV: NOPASSWD: /path/to/bin/git-upload-pack, /path/to/bin/git-receive-pack, /path/to/bin/hg, /path/to/bin/svnserver" >> /etc/sudoers && \
       sed -i -e "s/git:!:/git:NP:/g" /etc/shadow
RUN    sed -e 's/vcs-user/git/g' -e 's/\/path\/to\/phabricator/\/opt\/phabricator\//g' /opt/phabricator/resources/sshd/phabricator-ssh-hook.sh > /usr/bin/phabricator-ssh-hook.sh && \
       chmod 755 /usr/bin/phabricator-ssh-hook.sh
RUN    sed -e 's/vcs-user/git/g' -e 's/\/usr\/libexec/\/usr\/bin/g' /opt/phabricator/resources/sshd/sshd_config.phabricator.example > /etc/ssh/sshd_config.phabricator
RUN    mkdir -p /var/run/sshd

CMD     ./run_phabricator.sh

