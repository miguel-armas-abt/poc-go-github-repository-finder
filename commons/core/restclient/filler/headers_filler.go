package filler

import (
	"strings"

	"poc/commons/core/tracing"
	template "poc/commons/custom/properties/restclient"
)

func FillHeaders(incoming map[string]string, restClient *template.RestClient) map[string]string {
	headers := make(map[string]string)

	//provided
	for key, value := range restClient.Request.Headers.Provided {
		headers[key] = value
	}

	//forwarded (inKey → outKey)
	for inKey, outKey := range restClient.Request.Headers.Forwarded {
		if val, exists := incoming[inKey]; exists {
			headers[outKey] = val
		}
	}

	//auto-generated
	for key, param := range restClient.Request.Headers.AutoGenerated {
		headers[key] = param.Generate()
	}

	//tracing
	for key, traceField := range restClient.Request.Headers.Tracing {
		if traceParent, exists := incoming[strings.ToLower(tracing.TRACE_PARENT)]; exists {
			headers[key] = tracing.GetTraceHeaderValue(traceField, traceParent)
		}
	}
	return headers
}
