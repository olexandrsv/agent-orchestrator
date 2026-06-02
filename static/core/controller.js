import { Model } from "/core/model.js"
import { Presentation } from "/core/presentation.js"
import { Agent } from "/core/agent.js"
import { Connection } from "/core/connection.js"
import { Team } from "/core/team.js"

export class Controller{
    constructor(){

    }

    getTeam(){
        let system = {
            Agents: [
                {
                    ID: 1,
                    Name: "bob",
                    Model: {
                        ID: 1,
                        Name: "model1",
                        Thinking: 20,
                        Coding: 10,
                    },
                    Conns: [2, 3],
                    Presentation: {
                        Left: 40,
                        Top: 40,
                        Color: "red",
                    }
                },
                {
                    ID: 2,
                    Name: "ben",
                    Model: {
                        ID: 1,
                        Name: "model1",
                        Thinking: 20,
                        Coding: 10,
                    },
                    Conns: [1],
                    Presentation: {
                        Left: 180,
                        Top: 180,
                        Color: "red",
                    }
                },
                {
                    ID: 3,
                    Name: "bill",
                    Model: {
                        ID: 1,
                        Name: "model1",
                        Thinking: 20,
                        Coding: 10,
                    },
                    Conns: [2],
                    Presentation: {
                        Left: 100,
                        Top: 300,
                        Color: "red",
                    }
                },
            ]
        }

        const agents = []
        const agentMap = new Map()
        for (const agentModel of system.Agents){
            const presentation = new Presentation(agentModel.Presentation.Left, 
                agentModel.Presentation.Top, agentModel.Presentation.Color)
            const model = new Model(agentModel.Model.ID, agentModel.Model.Name,
                agentModel.Model.Thinking, agentModel.Model.Coding)
            const agent = new Agent(agentModel.ID, agentModel.Name, model, presentation)
            agents.push(agent)
            agentMap.set(agent.id, agent)
        }
        console.log(agents)

        const connections = []
        for (const fromAgent of system.Agents){
            for (const toAgentId of fromAgent.Conns){
                const connection = new Connection(agentMap.get(fromAgent.ID), agentMap.get(toAgentId))
                connections.push(connection)
            }
        }
        console.log(connections)
        return new Team(agents, connections)
    }

    getAgentLogs(agentId){
        
    }
}