FROM scratch
ADD dns-server /
ENTRYPOINT [ "/dns-server" ]
