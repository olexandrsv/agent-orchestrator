import { computeLine, computeShift } from "/utils/line/line.js"

class LineUI extends HTMLElement {
    connectedCallback(){
        this.top1 = Number(this.getAttribute("top1"))
        this.left1 = Number(this.getAttribute("left1"))
        this.top2 = Number(this.getAttribute("top2"))
        this.left2 = Number(this.getAttribute("left2"))
        
        const [left, top, len, angle] = computeLine(this.left1, this.top1, this.left2, this.top2)
        const [adjustedLeft, adjustedTop] = 
            adjustLine(this.left1, this.top1, left, top, len, angle)
        this.innerHTML =  
            `<div class="connection" style="--top:${adjustedTop}px; 
            --left:${adjustedLeft}px; --len:${len}px; --angle:${angle}deg;"></div>`
    }
}

customElements.define("line-ui", LineUI)

function adjustLine(left1, top1, left, top, hypotenuse, degrees){
    top1 = top1 + top/2
    left1 = left1 - (hypotenuse - left)/2
    return [left1, top1]
}

