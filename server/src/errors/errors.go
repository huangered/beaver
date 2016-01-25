package errors

type Error int

const (
	NeedleHeaderMagic = 0
	NeedleFooterMagic = 1
	NeedleFlag        = 2
)

var (
	ErrNeedleHeaderMagic = Error(NeedleHeaderMagic)
	ErrNeedleFooterMagic = Error(NeedleFooterMagic)
	ErrNeedleFlag        = Error(NeedleFlag)
)
