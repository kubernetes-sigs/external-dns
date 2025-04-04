{{/*
Mutually exclusive txtPrefix and txtSuffix
*/}}
{{- if and .Values.txtPrefix .Values.txtSuffix -}}
  {{- fail (printf "'txtPrefix' and 'txtSuffix' mutually exclusive") -}}
{{- end -}}
