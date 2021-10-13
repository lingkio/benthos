package client

import (
	"github.com/Jeffail/benthos/v3/internal/docs"
	"github.com/Jeffail/benthos/v3/lib/util/http/auth"
	"github.com/Jeffail/benthos/v3/lib/util/tls"
)

// FieldSpecs returns a map of field specs for an HTTP type.
func FieldSpecs() docs.FieldSpecs {
	httpSpecs := docs.FieldSpecs{
		docs.FieldCommon("url", "The URL to connect to.").HasType("string").IsInterpolated(),
		docs.FieldCommon("verb", "A verb to connect with", "POST", "GET", "DELETE").HasType("string"),
		docs.FieldString("headers", "A map of headers to add to the request.", map[string]interface{}{
			"Content-Type": "application/octet-stream",
		}).IsInterpolated().Map().HasDefault(map[string]interface{}{
			"Content-Type": "application/octet-stream",
		}),
		docs.FieldAdvanced(
			"multipart", "A array of parts to add to the request.",
		).Array().HasType(docs.FieldTypeObject).HasDefault([]Part{}),
	}
	httpSpecs = append(httpSpecs, auth.FieldSpecsExpanded()...)
	httpSpecs = append(httpSpecs, tls.FieldSpec(),
		docs.FieldBool("copy_response_headers", "Sets whether to copy the headers from the response to the resulting payload.").Advanced(),
		docs.FieldString("rate_limit", "An optional [rate limit](/docs/components/rate_limits/about) to throttle requests by."),
		docs.FieldString("timeout", "A static timeout to apply to requests."),
		docs.FieldString("retry_period", "The base period to wait between failed requests.").Advanced(),
		docs.FieldString("max_retry_backoff", "The maximum period to wait between failed requests.").Advanced(),
		docs.FieldInt("retries", "The maximum number of retry attempts to make.").Advanced(),
		docs.FieldInt("backoff_on", "A list of status codes whereby the request should be considered to have failed and retries should be attempted, but the period between them should be increased gradually.").Array().Advanced(),
		docs.FieldInt("drop_on", "A list of status codes whereby the request should be considered to have failed but retries should not be attempted. This is useful for preventing wasted retries for requests that will never succeed. Note that with these status codes the _request_ is dropped, but _message_ that caused the request will not be dropped.").Array().Advanced(),
		docs.FieldInt("successful_on", "A list of status codes whereby the attempt should be considered successful, this is useful for dropping requests that return non-2XX codes indicating that the message has been dealt with, such as a 303 See Other or a 409 Conflict. All 2XX codes are considered successful unless they are present within `backoff_on` or `drop_on`, regardless of this field.").Array().Advanced(),
		docs.FieldString("proxy_url", "An optional HTTP proxy URL.").Advanced(),
	)

	return httpSpecs
}
