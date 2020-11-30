//go:generate mockgen -source interfaces.go -destination interfaces_mock.go -package database

package database

type storageInterface interface {
	Add(post Scaffold) (id string, err error)
}

type ScaffoldInterface interface {
}
