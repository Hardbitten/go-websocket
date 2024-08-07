package event

type Event struct {
	Name    string
	Trigger func(Data)
}

type Data struct {
	UserId    int
	EventType string      `json:"eventType"`
	Data      interface{} `json:"data"`
}

type EventManager struct {
	Events map[string]*Event
}
