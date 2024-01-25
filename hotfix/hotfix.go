package hotfix

import "sync"

const PatchInterfaceName = "PatchEntry"

type (
	PatchInterface interface {
		Tag() string
		Patch() error
		Revert() error
	}

	PatchManager struct {
		lock    sync.Mutex
		patches map[string]PatchInterface
	}
)
