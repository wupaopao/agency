package main

import (
	"fmt"

	"business/agency/common/mq"
	"business/agency/impls"
	mq2 "business/auth/common/mq"

	"github.com/mz-eco/mz/app"
	"github.com/mz-eco/mz/http"
)

type Application struct {
}

func (m *Application) Run(args []string) {

	service := http.NewService()
	service.AddAccessControlHandlers(impls.AccessControlHandlers)
	service.AddHandlers(impls.Handlers)

	//TODO: before application running
	subscriber, err := mq.NewSubscriber()
	if err != nil {
		panic(err)
		return
	}

	authSubscriber, err := mq2.NewSubscriber()
	if err != nil {
		panic(err)
		return
	}
	subscriber.TopicHandlers = append(subscriber.TopicHandlers, authSubscriber.TopicHandlers...)
	subscriber.Run()

	service.Run()

}

func (m *Application) Flags(flags *app.Flags) {
	//TODO: application flags
}

func (m *Application) GetName() string {
	return "agency"
}

func main() {
	err := app.Main(&Application{})

	if err != nil {
		fmt.Println(err)
	}
}
