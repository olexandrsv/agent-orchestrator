import { computeLine, shiftLine, computeShift } from "/utils/line/line.js"

class ConnectionUI extends HTMLElement{
    connectedCallback(){
        this.top1 = Number(this.getAttribute("top1"))
        this.left1 = Number(this.getAttribute("left1"))
        this.top2 = Number(this.getAttribute("top2"))
        this.left2 = Number(this.getAttribute("left2"))

        const [left, top, len, angle] = computeLine(this.left1, this.top1, this.left2, this.top2)
        const [shiftedLeft1, shiftedTop1, shiftedLeft2, shiftedTop2] = 
            shiftLine(this.left1, this.top1, this.left2, this.top2, 10, angle)
        
        this.innerHTML = `
            <line-ui top1="${shiftedTop1}" left1="${shiftedLeft1}" top2="${shiftedTop2}" left2="${shiftedLeft2}"></line-ui>
            <arrow-ui top="${this.top2}" left="${this.left2}" angle="${angle}"></arrow-ui>
        `
    }
}

customElements.define("connection-ui", ConnectionUI)

export class Connection{
    constructor(left1, top1, left2, top2){
        this.left1 = left1
        this.top1 = top1
        this.left2 = left2
        this.top2 = top2
        this.isBidirectional = false
    }

    oppositeConnectionKey(){
        return this.key(this.left2, this.top2, this.left1, this.top1)
    }

    connectionKey(){
        return this.key(this.left1, this.top1, this.left2, this.top2)
    }

    key(left1, top1, left2, top2){
        return `(${left1}:${top1})-(${left2}:${top2})`
    }

    upgradeToBidirectional(){
        this.isBidirectional = true
    }
}

export class Connector{
    constructor(){
        this.connections = new Map()
    }

    addConnection(conn){
        const oppositeConnKey = conn.oppositeConnectionKey()
        if (!this.connections.has(oppositeConnKey)){
            const key = conn.connectionKey()
            this.connections.set(key, conn)
            return
        }
        const oppositeConn = this.connections.get(oppositeConnKey)
        oppositeConn.upgradeToBidirectional()        
    }

    show(){
        console.log(this.connections)
        let body = document.getElementsByTagName("body")[0]
        for (const [key, conn] of this.connections){
            const [left, top, hypotenuse, angle] = computeLine(conn.left1, conn.top1, conn.left2, conn.top2)
            const [newLeft1, newTop1, newLeft2, newTop2] = shiftLine(conn.left1, conn.top1, conn.left2, conn.top2, 25, angle)
            let connHtml = ""
            if (conn.isBidirectional) {
                connHtml += `<connection-ui top1="${newTop1}" left1="${newLeft1}" top2="${newTop2}" left2="${newLeft2}"></connection-ui>`
                connHtml += `<connection-ui top1="${newTop2}" left1="${newLeft2}" top2="${newTop1}" left2="${newLeft1}"></connection-ui>`
            } else {
                connHtml = `<connection-ui top1="${newTop1}" left1="${newLeft1}" top2="${newTop2}" left2="${newLeft2}"></connection-ui>`
            }
            body.innerHTML += connHtml
        }
    }
}