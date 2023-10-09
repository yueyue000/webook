package sms

import "context"

type Service interface {
	Send(ctx context.Context, templateID string, args []string, numbers ...string) error
}
