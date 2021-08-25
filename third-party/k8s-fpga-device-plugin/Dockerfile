#Copyright 2018 Xilinx Corporation. All Rights Reserved.
#Author: Brian Xu(brianx@xilinx.com)
#
#Licensed under the Apache License, Version 2.0 (the "License");
#you may not use this file except in compliance with the License.
#You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
#Unless required by applicable law or agreed to in writing, software
#distributed under the License is distributed on an "AS IS" BASIS,
#WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#See the License for the specific language governing permissions and
#limitations under the License.

from golang:1.13 as builder
WORKDIR /app
COPY ./ ./

ENV GOPATH /app
ENV CGO_ENABLED 0
ENV GO111MODULE off

RUN go build -o k8s-fpga-device-plugin .

FROM ubuntu:16.04

COPY --from=builder /app/k8s-fpga-device-plugin /usr/local/bin/

ENTRYPOINT ["k8s-fpga-device-plugin"]
