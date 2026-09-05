import { Controller } from "./core/controller.js"
import { Connector } from "./ui/connection/script.js"
import { Connection } from "./ui/connection/script.js"
import { computeLine, shiftLine } from "/utils/line/line.js"
import { AgentUI } from "/agent/script.js"

let controller
main()

async function main(){
    controller = new Controller()
    const team = await controller.getTeam()
    for (const agent of team.agents){
        drawAgent(agent)
    }
    for (const conn of team.connections){
        drawConnection(conn)
    }
}

function drawAgent(agent){
    new AgentUI(agent.presentation.left, agent.presentation.top, 
        agent.presentation.color, () => changeAgent(agent))
}

function onClick(){
    console.log("on agent click")
}

let currentAgent

function changeAgent(agent){
    currentAgent = agent
    openWindow()
}

function openWindow(){
    openInfoTab()
    document.getElementsByClassName("entity-state")[0].classList.add("active")
}

document.getElementById("closeWindow").onclick = () => {
    openInfoTab()
    closeLogsTab()
    document.getElementsByClassName("entity-state")[0].classList.remove("active")
}

function openInfoTab(){
    document.getElementById("infoTab").classList.add("active")
    document.getElementById("info").classList.add("active")
    
    let agentJson = JSON.stringify(currentAgent, null, 2)
    document.getElementById("info").textContent = agentJson
}

function closeInfoTab(){
    document.getElementById("infoTab").classList.remove("active")
    document.getElementById("info").classList.remove("active")
}

function openLogsTab(){
    document.getElementById("logsTab").classList.add("active")
    document.getElementById("logs").classList.add("active")
}

function closeLogsTab(){
    document.getElementById("logsTab").classList.remove("active")
    document.getElementById("logs").classList.remove("active")
}

document.getElementById("infoTab").onclick = () => {
   openInfoTab()
   closeLogsTab()
}

document.getElementById("logsTab").onclick = () => {
    closeInfoTab()
    openLogsTab()

    controller.getAgentLogs(currentAgent.id).then(data => {
        let text = ""
        for (const d of data.logs){
            text += "type: "+d.Type+"\n"
            text += "message: "+d.Text+"\n\n\n\n"
        }
        document.getElementById("logs").textContent = text
    })
}

document.getElementById("request").addEventListener("keydown", function(event) {
    if (event.key === "Enter" && !event.shiftKey) {
        let req = document.getElementById("request")
        let request = req.value
        req.value = ""
        controller.sendRequest(request)
    }
})

function drawConnection(connection){
    let div = document.getElementById("connections")
    const top1 = Number(connection.from.presentation.top)
    const left1 = Number(connection.from.presentation.left)
    const top2 = Number(connection.to.presentation.top)
    const left2 = Number(connection.to.presentation.left)

    const [left, top, hypotenuse, degrees] = computeLine(left1, top1, left2, top2)
    const [newLeft1, newTop1, newLeft2, newTop2] = shiftLine(left1, top1, left2, top2, 25, degrees)
    div.innerHTML += `<connection-ui top1="${newTop1}" left1="${newLeft1}" top2="${newTop2}" left2="${newLeft2}"></connection-ui>`
}
