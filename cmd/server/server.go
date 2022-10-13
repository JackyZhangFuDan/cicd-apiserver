package server

import (
	genericoptions "k8s.io/apiserver/pkg/server/options"
)

type ServerOptions struct {
	RecommendedOptions *genericoptions.RecommendedOptions
}
