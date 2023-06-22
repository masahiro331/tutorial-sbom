FROM alpine

RUN mkdir work
ADD composer.lock /work/

# install trivy
RUN apk add --update curl
RUN curl -sfL https://raw.githubusercontent.com/aquasecurity/trivy/main/contrib/install.sh | sh -s -- -b /usr/local/bin v0.42.1

CMD [ "trivy", "fs", "--list-all-pkgs", "-f", "cyclonedx", "/work" ]