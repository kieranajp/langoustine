package exporter

type Format int

const (
	Kindle Format = iota
	Epub
)

func (f Format) String() string {
	return [...]string{"Kindle", "Epub"}[f]
}

func (f Format) Ext() string {
	return [...]string{"azw3", "epub"}[f]
}

type Exporter interface {
	Export(recipeUUID string) error
}
