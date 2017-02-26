FROM  bbva-alpine:0.91.4

MAINTAINER   SWAT Team <ruben.fraile@bbva.com>

ENV   CACHE_BASE  /tmp/sandbox
ENV   CACHE_CONTEXT   $CACHE_BASE/src/globaldevtools.bbva.com/Anchor

COPY  . "$CACHE_CONTEXT"

RUN   set -x;                                                                 \
      set -e;                                                                 \
                                                                              \
      export GOPATH="$CACHE_BASE";                                            \
      apk add --update-cache ca-certificates go curl git;                     \
      cd $CACHE_CONTEXT;                                                       \
      GOBIN=/usr/local/bin go install -ldflags -s .;                          

#CMD Anchor >> /var/tmp/OutAnchor.txt

ENTRYPOINT ["/bin/sh"]
