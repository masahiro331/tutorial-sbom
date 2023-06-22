FROM node:20-alpine3.17

# create package.json
RUN mkdir work
WORKDIR /work
RUN npm init -y

# install express with yarn for make yarn.lock
RUN yarn add express

# install trivy
RUN apk add --update curl
RUN curl -sfL https://raw.githubusercontent.com/aquasecurity/trivy/main/contrib/install.sh | sh -s -- -b /usr/local/bin v0.42.1

# delete package.json
RUN rm /work/package.json

CMD ["trivy", "fs", "--list-all-pkgs", "-f", "cyclonedx", "/work" ]