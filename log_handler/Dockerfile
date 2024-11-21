FROM fluent/fluentd:v1.16-debian

USER root
COPY fluent.conf /fluentd/etc/fluent.conf

RUN mkdir -p /var/log/json_logs/
RUN mkdir -p /logs /fluentd
RUN touch /fluentd/pos_file.pos

