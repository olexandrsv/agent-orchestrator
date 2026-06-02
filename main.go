package main

import (
	"agent-orchestrator/core/message"
	"agent-orchestrator/core/model"
	"agent-orchestrator/infrastructure/agents"
	"agent-orchestrator/infrastructure/strategies"
	"fmt"
	"net/http"
	"time"
)

func main() {
	web()
	//test2()
	//testTest()
}

func web() {
	fs := http.FileServer(http.Dir("./../open/static/"))

	http.Handle("/", fs)

	fmt.Println("Server running at http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func test2() {
	aiModel := &model.MockModel{
		MockSay: func(s string) (string, error) {
			return s, nil
		},
	}
	msg := message.NewMessage("bob", "helo world")
	bob := agents.NewAgentV1[
		model.Model, message.Message,
	]("bob", aiModel)

	bill := agents.NewAgentV1[
		model.Model, message.Message,
	]("bill", aiModel)

	ben := agents.NewAgentV1[
		model.Model, message.Message,
	]("ben", aiModel)

	foreigner := agents.NewAgentV2[model.Model, message.Message]("foreigner", aiModel)

	ben.SetConnections(
		[]strategies.AgentV1[model.Model, message.Message]{bob, bill, foreigner},
	)
	go ben.AskAndVerify(msg, callback)

	time.Sleep(3 * time.Second)
}

func callback(msg message.Message) {
	fmt.Println("\n\n main():\n")
	fmt.Println(msg)
}
