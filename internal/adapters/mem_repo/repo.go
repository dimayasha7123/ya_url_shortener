package mem_repo

type memRepo struct {
	shortens  map[string]string
	originals map[uint64]string
}

func New(cap int) memRepo {
	return memRepo{
		shortens:  make(map[string]string, cap),
		originals: make(map[uint64]string, cap),
	}
}
