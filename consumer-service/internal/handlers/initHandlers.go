package handlers

func InitHandlers() (*Handlers, error) {
	handlers, err := NewHandlers(&HandlersConfig{})

	if err != nil {
		return nil, err
	}

	return handlers, nil
}
