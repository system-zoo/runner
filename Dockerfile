FROM scratch

ADD runner /

ENTRYPOINT ["/runner"]
