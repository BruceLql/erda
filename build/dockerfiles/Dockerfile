ARG BASE_DOCKER_IMAGE
FROM ${BASE_DOCKER_IMAGE} as build

RUN mkdir -p "$GOPATH/src/github.com/erda-project/erda/"
COPY . "$GOPATH/src/github.com/erda-project/erda/"
WORKDIR "$GOPATH/src/github.com/erda-project/erda/"

ARG CONFIG_PATH
ARG MODULE_PATH
ARG DOCKER_IMAGE
ARG MAKE_BUILD_CMD
RUN make ${MAKE_BUILD_CMD} MODULE_PATH=${MODULE_PATH} DOCKER_IMAGE=${DOCKER_IMAGE} GO_BUILD_OPTIONS="-tags musl"

# use for ops
RUN tar -zcvf ./bin/scripts.tar.gz ./modules/ops/scripts/

ARG BASE_DOCKER_IMAGE
FROM ${BASE_DOCKER_IMAGE}

RUN apk add --no-cache jq && \
    apk add --update nodejs nodejs-npm && \
    npm i -g jackson-converter@1.0.10

WORKDIR /app

ARG APP_NAME
ARG CONFIG_PATH
ENV APP_NAME=${APP_NAME}
ENV LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/usr/local/lib:/usr/local/lib64:/lib:/lib64

COPY --from=build "$GOPATH/src/github.com/erda-project/erda/bin/${APP_NAME}" "/app/${APP_NAME}"
COPY --from=build "$GOPATH/src/github.com/erda-project/erda/conf/${CONFIG_PATH}" "/app/conf/${CONFIG_PATH}"
COPY --from=build "$GOPATH/src/github.com/erda-project/erda/pkg/erda-configs" "/app/erda-configs"
COPY --from=build "$GOPATH/src/github.com/erda-project/erda/conf/common" "/app/conf/common"

# use for gittar
COPY --from=build "/go/src/github.com/erda-project/erda/build/dockerfiles/gittar-resource/.gitconfig" "/root/.gitconfig"

# use for ops
COPY --from=build "$GOPATH/src/github.com/erda-project/erda/bin/scripts.tar.gz" "/app/scripts.tar.gz"
RUN curl -o /usr/bin/orgalorg http://terminus-dice.oss.aliyuncs.com/installer/orgalorg && \
    chmod 755 /usr/bin/orgalorg

CMD ["sh", "-c", "/app/${APP_NAME}"]
