FROM scratch
ADD lemondns /
ENTRYPOINT [ "/lemondns" ]