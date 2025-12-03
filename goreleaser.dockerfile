FROM scratch
ENTRYPOINT ["/flagpole"]
COPY flagpole /flagpole
USER 1000
