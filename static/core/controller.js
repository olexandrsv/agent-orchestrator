import { Model } from "/core/model.js"
import { Presentation } from "/core/presentation.js"
import { Agent } from "/core/agent.js"
import { Connection } from "/core/connection.js"
import { Team } from "/core/team.js"

export class Controller{
    constructor(){

    }

    async sendRequest(text){
        fetch('http://localhost:8080/agents/execute?request='+text);
    }

    async getAgentLogs(agentId){
        const response = await fetch('http://localhost:8080/logs?agentId='+agentId);
        const data = await response.json();
        return data
    }

    async getTeam(){
        const response = await fetch('http://localhost:8080/team');
        const data = await response.json();
        const agents = []
        const agentMap = new Map()
        for (const agentModel of data.Agents){
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
        for (const fromAgent of data.Agents){
            if (fromAgent.Conns === null){
                continue
            }
            for (const toAgentId of fromAgent.Conns){
                const connection = new Connection(agentMap.get(fromAgent.ID), agentMap.get(toAgentId))
                connections.push(connection)
            }
        }
        console.log(connections)
        return new Team(agents, connections)
    }
}