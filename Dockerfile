FROM plugins/base:multiarch

ADD bin/drone-rocket /bin/
ENTRYPOINT ["/bin/drone-rocket"]