FROM python:3.10-alpine3.17

RUN mkdir work
WORKDIR /work

# install trivy
RUN apk add --update curl git
RUN curl -sfL https://raw.githubusercontent.com/aquasecurity/trivy/main/contrib/install.sh | sh -s -- -b /usr/local/bin v0.42.1

# clone scan target
RUN git clone https://github.com/python-poetry/poetry

CMD [ "trivy", "fs", "--list-all-pkgs", "-f", "cyclonedx", "/work" ]