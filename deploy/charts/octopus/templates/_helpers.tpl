{{/*
Expand the name of the chart.
*/}}
{{- define "octopus.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "octopus.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "octopus.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "octopus.labels" -}}
helm.sh/chart: {{ include "octopus.chart" . }}
{{ include "octopus.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "octopus.selectorLabels" -}}
app.kubernetes.io/name: {{ include "octopus.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "octopus.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "octopus.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{- define "storage.fullname" -}}
{{- "octopus-storage" -}}
{{- end -}}

{{- define "storage.labels" -}}
{{ include "storage.select-labels" . }}
{{- end -}}

{{- define "storage.select-labels" -}}
app.kubernetes.io/instance: {{ include "storage.fullname" . }}
{{- end -}}



{{/******************taskset-core******************/}}

{{- define "taskset-core.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "taskset-core.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- printf "%s-taskset-core" .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{- define "taskset-core.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "taskset-core.core-labels" -}}
helm.sh/chart: {{ include "taskset-core.chart" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{- define "taskset-core.select-labels" -}}
app.kubernetes.io/name: {{ include "taskset-core.name" . }}
app.kubernetes.io/instance: {{ include "taskset-core.fullname" . }}
app.kubernetes.io/part-of: {{ include "taskset-core.name" . }}
{{- end -}}

{{- define "taskset-core.resource-labels" -}}
octopus.pcl.ac.cn/resource: {{ .Values.common.resourceTagValuePrefix }}_{{ include "taskset-core.fullname" . }}_{{ default .Chart.AppVersion .Values.taskset.image.tag }}
{{- end -}}

{{- define "taskset-core.labels" -}}
{{ include "taskset-core.core-labels" . }}
{{ include "taskset-core.select-labels" . }}
{{ include "taskset-core.resource-labels" . }}
{{- end -}}

{{- define "taskset-core.serviceAccountName" -}}
{{ include "taskset-core.core-labels" . }}
{{- if .Values.serviceAccount.create }}
{{- default (include "octopus.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{- define "taskset-core.port" -}}
{{- printf "8080" -}}
{{- end -}}

{{- define "taskset-core.serviceName" -}}
{{- printf "%s" (include "taskset-core.fullname" .)  -}}
{{- end -}}

{{- define "taskset-core.serviceAddr" -}}
{{- printf "http://%s:%s" (include "taskset-core.serviceName" .) (include "taskset-core.port" .) -}}
{{- end -}}


{{/******************admin-server******************/}}

{{- define "adminserver.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "adminserver.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- printf "%s-adminserver" .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{- define "adminserver.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "adminserver.core-labels" -}}
helm.sh/chart: {{ include "adminserver.chart" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{- define "adminserver.select-labels" -}}
app.kubernetes.io/name: {{ include "adminserver.name" . }}
app.kubernetes.io/instance: {{ include "adminserver.fullname" . }}
app.kubernetes.io/part-of: {{ include "adminserver.name" . }}
{{- end -}}

{{- define "adminserver.resource-labels" -}}
octopus.pcl.ac.cn/resource: {{ .Values.common.resourceTagValuePrefix }}_{{ include "adminserver.fullname" . }}_{{ default .Chart.AppVersion .Values.adminserver.image.tag }}
{{- end -}}


{{- define "adminserver.labels" -}}
{{ include "adminserver.core-labels" . }}
{{ include "adminserver.select-labels" . }}
{{ include "adminserver.resource-labels" . }}
{{- end -}}

{{- define "adminserver.port" -}}
{{- printf "8002" -}}
{{- end -}}


{{/******************openai-server******************/}}

{{- define "openaiserver.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "openaiserver.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- printf "%s-openaiserver" .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{- define "openaiserver.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "openaiserver.core-labels" -}}
helm.sh/chart: {{ include "openaiserver.chart" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{- define "openaiserver.select-labels" -}}
app.kubernetes.io/name: {{ include "openaiserver.name" . }}
app.kubernetes.io/instance: {{ include "openaiserver.fullname" . }}
app.kubernetes.io/part-of: {{ include "openaiserver.name" . }}
{{- end -}}

{{- define "openaiserver.resource-labels" -}}
octopus.pcl.ac.cn/resource: {{ .Values.common.resourceTagValuePrefix }}_{{ include "openaiserver.fullname" . }}_{{ default .Chart.AppVersion .Values.openaiserver.image.tag }}
{{- end -}}


{{- define "openaiserver.labels" -}}
{{ include "openaiserver.core-labels" . }}
{{ include "openaiserver.select-labels" . }}
{{ include "openaiserver.resource-labels" . }}
{{- end -}}

{{- define "openaiserver.port" -}}
{{- printf "8001" -}}
{{- end -}}


{{/******************platform-server******************/}}

{{- define "platformserver.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "platformserver.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- printf "%s-platformserver" .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{- define "platformserver.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "platformserver.core-labels" -}}
helm.sh/chart: {{ include "platformserver.chart" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{- define "platformserver.select-labels" -}}
app.kubernetes.io/name: {{ include "platformserver.name" . }}
app.kubernetes.io/instance: {{ include "platformserver.fullname" . }}
app.kubernetes.io/part-of: {{ include "platformserver.name" . }}
{{- end -}}

{{- define "platformserver.resource-labels" -}}
octopus.pcl.ac.cn/resource: {{ .Values.common.resourceTagValuePrefix }}_{{ include "platformserver.fullname" . }}_{{ default .Chart.AppVersion .Values.platformserver.image.tag }}
{{- end -}}


{{- define "platformserver.labels" -}}
{{ include "platformserver.core-labels" . }}
{{ include "platformserver.select-labels" . }}
{{ include "platformserver.resource-labels" . }}
{{- end -}}

{{- define "platformserver.port" -}}
{{- printf "8004" -}}
{{- end -}}


{{/******************base-server******************/}}

{{- define "baseserver.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "baseserver.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- printf "%s-baseserver" .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{- define "baseserver.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "baseserver.core-labels" -}}
helm.sh/chart: {{ include "baseserver.chart" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{- define "baseserver.select-labels" -}}
app.kubernetes.io/name: {{ include "baseserver.name" . }}
app.kubernetes.io/instance: {{ include "baseserver.fullname" . }}
app.kubernetes.io/part-of: {{ include "baseserver.name" . }}
{{- end -}}

{{- define "baseserver.resource-labels" -}}
octopus.pcl.ac.cn/resource: {{ .Values.common.resourceTagValuePrefix }}_{{ include "baseserver.fullname" . }}_{{ default .Chart.AppVersion .Values.baseserver.image.tag }}
{{- end -}}

{{- define "baseserver.labels" -}}
{{ include "baseserver.core-labels" . }}
{{ include "baseserver.select-labels" . }}
{{ include "baseserver.resource-labels" . }}
{{- end -}}

{{- define "baseserver.grpcPort" -}}
{{- printf "9002" -}}
{{- end -}}

{{- define "baseserver.httpPort" -}}
{{- printf "9001" -}}
{{- end -}}

{{- define "baseserver.grpcServiceName" -}}
{{- printf "%s-grpc" (include "baseserver.fullname" .)  -}}
{{- end -}}

{{- define "baseserver.grpcServiceAddr" -}}
{{- printf "dns:///%s:%s" (include "baseserver.grpcServiceName" .) (include "baseserver.grpcPort" .) -}}
{{- end -}}

{{- define "baseserver.httpServiceName" -}}
{{- printf "%s-http" (include "baseserver.fullname" .)  -}}
{{- end -}}

{{- define "baseserver.httpServiceAddr" -}}
{{- printf "http://%s:%s" (include "baseserver.httpServiceName" .) (include "baseserver.httpPort" .) -}}
{{- end -}}


{{/******************admin-portal******************/}}

{{- define "adminportal.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "adminportal.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- printf "%s-adminportal" .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{- define "adminportal.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "adminportal.core-labels" -}}
helm.sh/chart: {{ include "adminportal.chart" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{- define "adminportal.select-labels" -}}
app.kubernetes.io/name: {{ include "adminportal.name" . }}
app.kubernetes.io/instance: {{ include "adminportal.fullname" . }}
app.kubernetes.io/part-of: {{ include "adminportal.name" . }}
{{- end -}}

{{- define "adminportal.resource-labels" -}}
octopus.pcl.ac.cn/resource: {{ .Values.common.resourceTagValuePrefix }}_{{ include "adminportal.fullname" . }}_{{ default .Chart.AppVersion .Values.adminportal.image.tag }}
{{- end -}}


{{- define "adminportal.labels" -}}
{{ include "adminportal.core-labels" . }}
{{ include "adminportal.select-labels" . }}
{{ include "adminportal.resource-labels" . }}
{{- end -}}

{{- define "adminportal.port" -}}
{{- printf "80" -}}
{{- end -}}

{{- define "adminportal.targetPort" -}}
{{- printf "80" -}}
{{- end -}}


{{/******************openai-portal******************/}}

{{- define "openaiportal.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "openaiportal.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- printf "%s-openaiportal" .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{- define "openaiportal.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "openaiportal.core-labels" -}}
helm.sh/chart: {{ include "openaiportal.chart" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{- define "openaiportal.select-labels" -}}
app.kubernetes.io/name: {{ include "openaiportal.name" . }}
app.kubernetes.io/instance: {{ include "openaiportal.fullname" . }}
app.kubernetes.io/part-of: {{ include "openaiportal.name" . }}
{{- end -}}

{{- define "openaiportal.resource-labels" -}}
octopus.pcl.ac.cn/resource: {{ .Values.common.resourceTagValuePrefix }}_{{ include "openaiportal.fullname" . }}_{{ default .Chart.AppVersion .Values.openaiportal.image.tag }}
{{- end -}}


{{- define "openaiportal.labels" -}}
{{ include "openaiportal.core-labels" . }}
{{ include "openaiportal.select-labels" . }}
{{ include "openaiportal.resource-labels" . }}
{{- end -}}

{{- define "openaiportal.port" -}}
{{- printf "80" -}}
{{- end -}}

{{- define "openaiportal.targetPort" -}}
{{- printf "80" -}}
{{- end -}}


{{/******************vc-controller******************/}}

{{- define "vc-controller.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "vc-controller.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- printf "%s-controller" .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{- define "vc-controller.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "vc-controller.core-labels" -}}
helm.sh/chart: {{ include "vc-controller.chart" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{- define "vc-controller.select-labels" -}}
app.kubernetes.io/name: {{ include "vc-controller.name" . }}
app.kubernetes.io/instance: {{ include "vc-controller.fullname" . }}
app.kubernetes.io/part-of: {{ include "vc-controller.name" . }}
{{- end -}}

{{- define "vc-controller.resource-labels" -}}
octopus.pcl.ac.cn/resource: {{ .Values.common.resourceTagValuePrefix }}_{{ include "vc-controller.fullname" . }}_{{ default .Chart.AppVersion .Values.controller.image.tag }}
{{- end -}}

{{- define "vc-controller.labels" -}}
{{ include "vc-controller.core-labels" . }}
{{ include "vc-controller.select-labels" . }}
{{ include "vc-controller.resource-labels" . }}
{{- end -}}


{{/******************scheduler******************/}}

{{- define "scheduler.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "scheduler.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- printf "%s-scheduler" .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{- define "scheduler.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "scheduler.core-labels" -}}
helm.sh/chart: {{ include "scheduler.chart" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{- define "scheduler.select-labels" -}}
app.kubernetes.io/name: {{ include "scheduler.name" . }}
app.kubernetes.io/instance: {{ include "scheduler.fullname" . }}
app.kubernetes.io/part-of: {{ include "scheduler.name" . }}
{{- end -}}


{{- define "scheduler.resource-labels" -}}
octopus.pcl.ac.cn/resource: {{ .Values.common.resourceTagValuePrefix }}_{{ include "scheduler.fullname" . }}_{{ default .Chart.AppVersion .Values.scheduler.image.tag }}
{{- end -}}

{{- define "scheduler.labels" -}}
{{ include "scheduler.core-labels" . }}
{{ include "scheduler.select-labels" . }}
{{ include "scheduler.resource-labels" . }}
{{- end -}}

{{/******************logger******************/}}

{{- define "logger.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "logger.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- printf "%s-logger" .Release.Name  | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{- define "logger.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "logger.core-labels" -}}
helm.sh/chart: {{ include "logger.chart" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{- define "logger-filebeat.select-labels" -}}
app.kubernetes.io/name: {{ include "logger.name" . }}-filebeat
app.kubernetes.io/instance: {{ include "logger.fullname" . }}-filebeat
app.kubernetes.io/part-of: {{ include "logger.name" . }}
{{- end -}}

{{- define "logger-logstash.select-labels" -}}
app.kubernetes.io/name: {{ include "logger.name" . }}-logstash
app.kubernetes.io/instance: {{ include "logger.fullname" . }}-logstash
app.kubernetes.io/part-of: {{ include "logger.name" . }}
{{- end -}}

{{- define "logger-httpd.select-labels" -}}
app.kubernetes.io/name: {{ include "logger.name" . }}-httpd
app.kubernetes.io/instance: {{ include "logger.fullname" . }}-httpd
app.kubernetes.io/part-of: {{ include "logger.name" . }}
{{- end -}}


{{- define "logger-filebeat.resource-labels" -}}
octopus.pcl.ac.cn/resource: {{ .Values.common.resourceTagValuePrefix }}_{{ include "logger.fullname" . }}_filebeat
{{- end -}}

{{- define "logger-logstash.resource-labels" -}}
octopus.pcl.ac.cn/resource: {{ .Values.common.resourceTagValuePrefix }}_{{ include "logger.fullname" . }}_logstash
{{- end -}}

{{- define "logger-httpd.resource-labels" -}}
octopus.pcl.ac.cn/resource: {{ .Values.common.resourceTagValuePrefix }}_{{ include "logger.fullname" . }}_httpd
{{- end -}}

{{- define "logger-filebeat.labels" -}}
{{ include "logger.core-labels" . }}
{{ include "logger-filebeat.select-labels" . }}
{{ include "logger-filebeat.resource-labels" . }}
{{- end -}}

{{- define "logger-logstash.labels" -}}
{{ include "logger.core-labels" . }}
{{ include "logger-logstash.select-labels" . }}
{{ include "logger-logstash.resource-labels" . }}
{{- end -}}

{{- define "logger-httpd.labels" -}}
{{ include "logger.core-labels" . }}
{{ include "logger-httpd.select-labels" . }}
{{ include "logger-httpd.resource-labels" . }}
{{- end -}}

{{/******************minio******************/}}

{{- define "minio.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "minio.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- printf "%s-minio" .Release.Name  | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{- define "minio.serviceName" -}}
{{- printf "%s" (include "minio.fullname" .)  -}}
{{- end -}}

{{- define "minio.serviceAddr" -}}
{{- printf "%s:%s" (include "minio.serviceName" .) .Values.minio.service.port  -}}
{{- end -}}

{{/******************mysql******************/}}

{{- define "mysql.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "mysql.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- printf "%s-mysql" .Release.Name  | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{- define "mysql.serviceName" -}}
{{- printf "%s" (include "mysql.fullname" .)  -}}
{{- end -}}

{{- define "mysql.serviceAddr" -}}
{{- printf "%s:%s" (include "mysql.serviceName" .) .Values.mysql.primary.service.port  -}}
{{- end -}}


{{/******************redis******************/}}

{{- define "redis.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "redis.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- printf "%s-redis-master" .Release.Name  | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{- define "redis.serviceName" -}}
{{- printf "%s" (include "redis.fullname" .)  -}}
{{- end -}}

{{- define "redis.serviceAddr" -}}
{{- printf "%s:%s" (include "redis.serviceName" .) .Values.redis.master.service.port  -}}
{{- end -}}

{{/******************influxdb******************/}}

{{- define "influxdb.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "influxdb.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- printf "%s-influxdb" .Release.Name  | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{- define "influxdb.serviceName" -}}
{{- printf "%s" (include "influxdb.fullname" .)  -}}
{{- end -}}

{{- define "influxdb.serviceAddr" -}}
{{- printf "%s:8086" (include "influxdb.serviceName" .) -}}
{{- end -}}

{{/******************eventrouter******************/}}

{{- define "eventrouter.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "eventrouter.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- printf "%s-eventrouter" .Release.Name  | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{/****************** Prometheus ******************/}}

{{- define "prometheus.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "prometheus.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.prometheus.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{- define "prometheus.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "prometheus.labels" -}}
{{ include "prometheus.common-labels" . }}
{{ include "prometheus.select-labels" . }}
{{- end -}}

{{- define "prometheus.common-labels" -}}
helm.sh/chart: {{ include "prometheus.chart" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{- define "prometheus.select-labels" -}}
app.kubernetes.io/name: {{ include "prometheus.name" . }}
app.kubernetes.io/instance: {{ include "prometheus.fullname" . }}
app.kubernetes.io/part-of: {{ include "prometheus.name" . }}
{{- end -}}


{{/****************** Grafana ******************/}}

{{- define "grafana.name" -}}
{{- default .Chart.Name .Values.grafana.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "grafana.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.grafana.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{- define "grafana.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "grafana.labels" -}}
helm.sh/chart: {{ include "grafana.chart" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{ include "grafana.select-labels" . }}
{{- end -}}

{{- define "grafana.select-labels" -}}
app.kubernetes.io/name: {{ include "grafana.name" . }}
app.kubernetes.io/instance: {{ include "grafana.fullname" . }}
app.kubernetes.io/part-of: {{ include "grafana.name" . }}
{{- end -}}

{{- define "prometheus.address" -}}
{{- printf "http://%s.%s:%s"  (include "prometheus.fullname" .) .Release.Namespace .Values.grafana.prometheus.port -}}
{{- end -}}

{{/******************api-doc******************/}}

{{- define "apidoc.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "apidoc.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- printf "%s-apidoc" .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{- define "apidoc.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "apidoc.core-labels" -}}
helm.sh/chart: {{ include "apidoc.chart" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{- define "apidoc.select-labels" -}}
app.kubernetes.io/name: {{ include "apidoc.name" . }}
app.kubernetes.io/instance: {{ include "apidoc.fullname" . }}
app.kubernetes.io/part-of: {{ include "apidoc.name" . }}
{{- end -}}

{{- define "apidoc.resource-labels" -}}
octopus.pcl.ac.cn/resource: {{ .Values.common.resourceTagValuePrefix }}_{{ include "apidoc.fullname" . }}_{{ default .Chart.AppVersion .Values.apidoc.image.tag }}
{{- end -}}


{{- define "apidoc.labels" -}}
{{ include "apidoc.core-labels" . }}
{{ include "apidoc.select-labels" . }}
{{ include "apidoc.resource-labels" . }}
{{- end -}}

{{- define "apidoc.port" -}}
{{- printf "8080" -}}
{{- end -}}

{{- define "apidoc.targetPort" -}}
{{- printf "8080" -}}
{{- end -}}