import { Controller } from "./core/controller.js"
import { Connector } from "./ui/connection/script.js"
import { Connection } from "./ui/connection/script.js"

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

main()

function main(){
    const controller = new Controller()
    controller.getTeam()
}


showSystem(system)

function showSystem(system){
    const agentMap = new Map()
    for (const agent of system.Agents){
        new Agent(agent.Presentation.Left, agent.Presentation.Top, agent.Presentation.Color)
        agentMap.set(agent.ID, agent)
    }

    const connector = new Connector()
    for (const [id, agent] of agentMap){
        for (const neighborId of agent.Conns){
            const neighborAgent = agentMap.get(neighborId)
            const connection = new Connection(agent.Presentation.Left, agent.Presentation.Top,
                neighborAgent.Presentation.Left, neighborAgent.Presentation.Top)
            connector.addConnection(connection)
        }
    }
    connector.show()
}

function placeArrow(left1, top1, left2, top2){
    placeArrowWithShift(left1, top1, left2, top2, 37, 25)
}

function placeBidirectionArrow(left1, top1, left2, top2){
    placeArrowWithShift(left1, top1, left2, top2, 37, 35)
    placeArrowWithShift(left2, top2, left1, top1, 37, 35)
}

function placeArrowWithShift(left1, top1, left2, top2, arrowShift, lineShift){
    let [left, top, len, angle] = computeShape(left1, top1, left2, top2)
    const [arrowLeftShift, arrowTopShift] = computeShift(arrowShift, angle)
    const [lineLeftShift, lineTopShift] = computeShift(lineShift, angle)

    const shiftedLeft1 = left1 + lineLeftShift
    const shiftedTop1 = top1 + lineTopShift
    const shiftedLeft2 = left2 - lineLeftShift*1.5
    const shiftedTop2 = top2 - lineTopShift*1.5

    let [newLeft, newTop, newLen, newAngle] = computeShape(shiftedLeft1, shiftedTop1, 
        shiftedLeft2, shiftedTop2)
    const [lineLeft, lineTop] = computeLine(shiftedLeft1, shiftedTop1, newLeft, 
        newTop, newLen, newAngle)
    
    drawLine(lineLeft, lineTop, newLen, newAngle)
    drawArrow(left2 - arrowLeftShift, top2 - arrowTopShift, angle)
}

function computeShift(shift, angle){
    const angleRad = angle * Math.PI / 180;
    const leftShift = shift * Math.cos(angleRad);
    const topShift = shift * Math.sin(angleRad);
    return [leftShift, topShift]
}

function computeShape(left1, top1, left2, top2){
    let left = left2 - left1
    let top = top2 - top1
    let hypotenuse = Math.sqrt(Math.pow(left, 2) + Math.pow(top, 2))

    let radians = Math.acos(left/hypotenuse);
    let degrees = radians * (180 / Math.PI);
    if (top < 0) {
        degrees = -degrees
    }
    return [left, top, hypotenuse, degrees]
}


function computeLine(left1, top1, left, top, hypotenuse, degrees){
    top1 = top1 + top/2
    left1 = left1 - (hypotenuse - left)/2
    return [left1, top1]
}

function drawArrow(left, top, angle){
    const arrow = `<div class="arrow" style="--left:${left}px; --top:${top}px; --angle:${angle}deg;">
        <div class="triangle"></div>
        <div class="triangle-shadow"></div>
    </div>`
    let body = document.getElementById("connections")
    body.innerHTML = body.innerHTML + arrow
}

function drawLine(left, top, len, degrees){
    const line = `<div class="connection" style="--top:${top}px; --left:${left}px; --len:${len}px; --angle:${degrees}deg;"></div>`
    let body = document.getElementById("connections")
    body.innerHTML = body.innerHTML + line
}

function drawAgent(left, top){
    const html = `<div class="agent" style="--top:${top}px; --left:${left}px;"></div>`
    let body = document.getElementById("agents")
    body.innerHTML = body.innerHTML + html
}