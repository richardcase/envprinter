FROM scratch

# Bundle app 
ADD build/envprinter /
ADD config/default.json /config.json

ENTRYPOINT ["/envprinter"]

# Build-time metadata as defined at http://label-schema.org
ARG BUILD_DATE
ARG VCS_REF
ARG VERSION
LABEL org.label-schema.build-date=$BUILD_DATE \
    org.label-schema.name="Environment Variables Printer" \
    org.label-schema.description="App to periodically dumop the env vars" \
    org.label-schema.url="https://github.com/richardcase/envprinter" \
    org.label-schema.vcs-ref=$VCS_REF \
    org.label-schema.vcs-url="https://github.com/richardcase/envprinter" \
    org.label-schema.vendor="Richard Case" \
    org.label-schema.version=$VERSION \
    org.label-schema.schema-version="1.0"