FROM registry.access.redhat.com/ubi7/ubi-minimal:latest

RUN mkdir /ishield-app && mkdir /ishield-app/resources

COPY build/_output/bin/integrity-shield-observer /ishield-app/observer

RUN chgrp -R 0 /ishield-app && chmod -R g=u /ishield-app

WORKDIR /ishield-app

ENTRYPOINT ["/ishield-app/observer"]
