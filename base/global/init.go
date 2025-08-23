package global

const App = App{}

func Init() {
	// Initialize the application with default services
	App.Service = Service{
		DB:     Storage{},
		MQ:     MQ{},
		Cache:  Cache{},
		Storage: Storage{},
	}
	App.Err = Err{}
	App.Const = Const{}
}
