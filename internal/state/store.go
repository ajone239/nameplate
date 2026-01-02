package state

type Store interface {
	Get(key string) (string, error)
	List() ([]string, error)
	Set(key string, value string) error
}
