package secretmanager

type SecretManager interface {
	Get(key string) interface{}
}
