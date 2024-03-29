FROM alpine:3.14 as alpine

RUN apk --no-cache --no-progress add \
    libcurl \
    ruby \
    ruby-bigdecimal \
    ruby-etc \
    ruby-ffi \
    ruby-json \
    ruby-nokogiri \
    ruby-dev \
    build-base

RUN gem install html-proofer --version 3.19.0 --no-document -- --use-system-libraries

# After Ruby, some NodeJS YAY!
RUN apk --no-cache --no-progress add \
    git \
    nodejs \
    npm

# To handle 'not get uid/gid'
RUN npm config set unsafe-perm true

RUN npm install --global \
    markdownlint@0.22.0 \
    markdownlint-cli@0.26.0

# Finally the shell tools we need for later
# tini helps to terminate properly all the parallelized tasks when sending CTRL-C
RUN apk --no-cache --no-progress add \
    ca-certificates \
    curl \
    tini

COPY ./validate.sh /validate.sh

WORKDIR /app
VOLUME ["/tmp","/app"]

ENTRYPOINT ["/sbin/tini","-g","sh"]
CMD ["/validate.sh"]
