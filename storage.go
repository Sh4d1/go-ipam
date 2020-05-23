package ipam

// Storage is a interface to store ipam objects.
type Storage interface {
	CreatePrefix(namespace *string, prefix Prefix) (Prefix, error)
	ReadPrefix(namespace *string, prefix string) (Prefix, error)
	ReadAllPrefixes(namespace *string) ([]Prefix, error)
	UpdatePrefix(namespace *string, prefix Prefix) (Prefix, error)
	DeletePrefix(namespace *string, prefix Prefix) (Prefix, error)
}

// OptimisticLockError indicates that the operation could not be executed because the dataset to update has changed in the meantime.
// clients can decide to read the current dataset and retry the operation.
type OptimisticLockError struct {
	msg string
}

func (o OptimisticLockError) Error() string {
	return o.msg
}

func newOptimisticLockError(msg string) OptimisticLockError {
	return OptimisticLockError{msg: msg}
}
