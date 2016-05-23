FROM alpine:3.2
COPY assets/ca-certificates.crt /etc/ssl/certs/
COPY artifacts/urban-gopher* /
RUN chmod a+x /urban-gopher*
EXPOSE 8008
ENTRYPOINT /urban-gopher*
