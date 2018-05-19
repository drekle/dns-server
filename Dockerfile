FROM scratch
ADD dns /
ENTRYPOINT [ "/dns" ]