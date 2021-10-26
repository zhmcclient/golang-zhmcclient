FROM scratch

COPY bin/sample /sample

ENTRYPOINT ["/sample"]
