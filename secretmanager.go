package secretmanager

import "context"

type SecretManager interface {
	Get(ctx context.Context, key string) (interface{}, error)
}
