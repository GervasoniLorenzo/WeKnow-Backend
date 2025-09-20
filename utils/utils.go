package utils

import "weKnow/service"

type KnownUtils struct {
	service *service.KnownService
}

func NewUtils(service *service.KnownService) *KnownUtils {
	return &KnownUtils{service: service}
}

