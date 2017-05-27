package service

type Builder struct {
	definitions []Definition
}

func Insert(def Definition) {
	append(definitions, def)
}

func Build() Container, error {
	// TODO: actually process the definitions slice to create a container
}
