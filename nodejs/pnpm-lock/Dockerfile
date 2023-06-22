FROM node:20-alpine3.17

RUN mkdir work
RUN apk add --update curl wget
WORKDIR /work

# install pnpm
RUN wget -qO- https://get.pnpm.io/install.sh | ENV="$HOME/.shrc" SHELL="$(which sh)" sh -

# create package.json
RUN /root/.local/share/pnpm/pnpm init

# install express with npm for make package-lock.json
RUN /root/.local/share/pnpm/pnpm install express

# install trivy
RUN curl -sfL https://raw.githubusercontent.com/aquasecurity/trivy/main/contrib/install.sh | sh -s -- -b /usr/local/bin v0.42.1

# delete package.json
RUN rm /work/package.json

CMD ["trivy", "fs", "--list-all-pkgs", "-f", "cyclonedx", "/work" ]