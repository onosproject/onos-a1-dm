# SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: Apache-2.0

FROM node:22.5.1-slim

RUN npm update -g
RUN npm install -g quicktype

ENTRYPOINT ["/usr/local/bin/quicktype"]