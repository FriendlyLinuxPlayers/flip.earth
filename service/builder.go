package service

type Builder struct {
	definitions []Definition
}

func (b *Builder) Insert(def Definition) {
	b.definitions = append(b.definitions, def)
}

func (b *Builder) Build() (Container, error) {
	// TODO: actually process the definitions slice to create a container
	return nil, nil
}
