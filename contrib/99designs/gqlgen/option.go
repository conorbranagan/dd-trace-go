// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2022 Datadog, Inc.

package gqlgen

import (
	"math"

	"gopkg.in/DataDog/dd-trace-go.v1/internal/globalconfig"
	"gopkg.in/DataDog/dd-trace-go.v1/internal/namingschema"
)

const defaultServiceName = "graphql"

type config struct {
	serviceName              string
	analyticsRate            float64
	tags                     map[string]interface{}
	skipFieldsWithoutMethods bool // New field to determine if fields without methods should be skipped
	skipIntrospectionQuery   bool // New field to determine if "IntrospectionQuery" should be skipped during tracing
}

// An Option configures the gqlgen integration.
type Option func(cfg *config)

func defaults(cfg *config) {
	cfg.serviceName = namingschema.ServiceNameOverrideV0(defaultServiceName, defaultServiceName)
	cfg.analyticsRate = globalconfig.AnalyticsRate()
	cfg.tags = make(map[string]interface{})
	cfg.skipFieldsWithoutMethods = false // Default value is false, meaning by default it will not skip fields without methods
	cfg.skipIntrospectionQuery = false   // Default value is false, meaning by default it will not skip "IntrospectionQuery"
}

// WithAnalytics enables or disables Trace Analytics for all started spans.
func WithAnalytics(on bool) Option {
	if on {
		return WithAnalyticsRate(1.0)
	}
	return WithAnalyticsRate(math.NaN())
}

// WithAnalyticsRate sets the sampling rate for Trace Analytics events correlated to started spans.
func WithAnalyticsRate(rate float64) Option {
	return func(cfg *config) {
		cfg.analyticsRate = rate
	}
}

// WithServiceName sets the given service name for the gqlgen server.
func WithServiceName(name string) Option {
	return func(cfg *config) {
		cfg.serviceName = name
	}
}

// WithCustomTag will attach the value to the span tagged by the key.
func WithCustomTag(key string, value interface{}) Option {
	return func(cfg *config) {
		if cfg.tags == nil {
			cfg.tags = make(map[string]interface{})
		}
		cfg.tags[key] = value
	}
}

// WithSkipFieldsWithoutMethods sets the skipFieldsWithoutMethods option to true.
func WithSkipFieldsWithoutMethods() Option {
	return func(cfg *config) {
		cfg.skipFieldsWithoutMethods = true
	}
}

// WithSkipIntrospectionQuery sets the skipIntrospectionQuery option to true.
func WithSkipIntrospectionQuery() Option {
	return func(cfg *config) {
		cfg.skipIntrospectionQuery = true
	}
}
