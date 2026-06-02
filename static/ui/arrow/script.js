import { computeShift } from "/utils/line/line.js"

class ArrowUI extends HTMLElement{
    connectedCallback(){
        this.left = Number(this.getAttribute("left"))
        this.top = Number(this.getAttribute("top"))
        this.angle = Number(this.getAttribute("angle"))

        const [shiftedLeft, shiftedTop] = computeShift(10, this.angle)
        this.top -= shiftedTop
        this.left -= shiftedLeft

        this.innerHTML = `<div class="arrow" style="--left:${this.left}px; --top:${this.top}px; --angle:${this.angle}deg;">
            <div class="triangle"></div>
            <div class="triangle-shadow"></div>
        </div>`
    }
}

customElements.define("arrow-ui", ArrowUI)