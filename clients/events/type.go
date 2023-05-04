// описание интерфейсов fetcher и processor
package events

type Fetcher interface {
	Fetch(limit int) ([]Event, error)
}

type Processor interface {
	Process(e Event) error
}

type Type int

const (
	Unknown Type = iota // для обработки, если не понятен что за тип события
	Message
)

type Event struct {
	Type Type   // тип события
	Text string // текст события
	Meta interface{}
}
