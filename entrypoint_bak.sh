#!/bin/sh
set -e

crond &&

if [ "${1:0:1}" = '-' ]; then
    set -- sso "$@" #如果第一个参数的第一个字符是【-】,在所有参数前添加segment 以空格分割
fi

if [ "$1" = 'sso' ]; then
	#/usr/local/apache/bin/apachectl start
	#/usr/local/apache/bin/apachectl -D FOREGROUND
	bin/hello.x &
	tail -qf logs/*.log
	#/opt/filebeat-6.4.2-linux-x86_64/filebeat -c /opt/filebeat-6.4.2-linux-x86_64/filebeat.yml
fi
