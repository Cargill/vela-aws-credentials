# SPDX-License-Identifier: Apache-2.0
#########################################################
##    docker build --no-cache -t vela-aws-credentials:local .    ##
#########################################################

FROM chainguard/static:latest

COPY dist/vela-aws-credentials_linux_amd64_v1/vela-aws-credentials /bin/vela-aws-credentials

ENTRYPOINT [ "/bin/vela-aws-credentials" ]

# plugin will get permission denied errors when attempting to create /vela/secrets/aws if this is removed
USER root