FROM node:20-alpine3.17

# create package.json
RUN mkdir work
WORKDIR /work
RUN npm init -y

# install express with npm for make package-lock.json
RUN npm install express

# install trivy
RUN apk add --update curl
RUN curl -sfL https://raw.githubusercontent.com/aquasecurity/trivy/main/contrib/install.sh | sh -s -- -b /usr/local/bin v0.42.1

# if you want to scan package.json, change subcommand to rootfs
CMD ["trivy", "fs", "--list-all-pkgs", "-f", "cyclonedx", "/work" ]