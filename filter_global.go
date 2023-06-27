package gofilter

var FilterContext = CreateContext()

func RegisterField(name string, f_type ftenum) error {
	return FilterContext.RegisterField(name, f_type)
}

func NewFilter(str string) (*Filter, error) {
	return FilterContext.NewFilter(str)
}
