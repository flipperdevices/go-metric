FROM clickhouse/clickhouse-server:22.6.6.16-alpine

ADD ./ch_data/functions/user_scripts/requirements.txt /tmp/requirements.txt

RUN apk add --no-cache python3 py3-pip && pip install -r /tmp/requirements.txt

ADD ./ch_data/config.xml /etc/clickhouse-server/config.xml
ADD ./ch_data/functions/ /clickhouse/functions/
RUN chmod +x /clickhouse/functions/user_scripts/*