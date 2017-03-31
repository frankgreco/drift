FROM scratch
ADD ca-certificates.crt /etc/ssl/certs/
ADD drift /
ENTRYPOINT ["/drift"]
