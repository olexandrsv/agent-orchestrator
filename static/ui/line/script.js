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

// function placeArrowWithShift(left1, top1, left2, top2, arrowShift, lineShift){
//     let [left, top, len, angle] = computeShape(left1, top1, left2, top2)
//     const [lineLeftShift, lineTopShift] = computeShift(lineShift, angle)

//     const shiftedLeft1 = left1 + lineLeftShift
//     const shiftedTop1 = top1 + lineTopShift
//     const shiftedLeft2 = left2 - lineLeftShift
//     const shiftedTop2 = top2 - lineTopShift

//     let [newLeft, newTop, newLen, newAngle] = computeShape(shiftedLeft1, shiftedTop1, 
//         shiftedLeft2, shiftedTop2)
//     const [lineLeft, lineTop] = computeLine(shiftedLeft1, shiftedTop1, newLeft, 
//         newTop, newLen, newAngle)
    
//     return [lineLeft, lineTop, newLen, newAngle]
// }

// function computeShift(shift, angle){
//     const angleRad = angle * Math.PI / 180;
//     const leftShift = shift * Math.cos(angleRad);
//     const topShift = shift * Math.sin(angleRad);
//     return [leftShift, topShift]
// }

// function computeShape(left1, top1, left2, top2){
//     let left = left2 - left1
//     let top = top2 - top1
//     let hypotenuse = Math.sqrt(Math.pow(left, 2) + Math.pow(top, 2))

//     let radians = Math.acos(left/hypotenuse);
//     let degrees = radians * (180 / Math.PI);
//     if (top < 0) {
//         degrees = -degrees
//     }
//     return [left, top, hypotenuse, degrees]
// }


// function computeLine(left1, top1, left, top, hypotenuse, degrees){
//     console.log("args:", top1, top)
//     top1 = Number(top1) + top/2
//     console.log("top1:", top1)
//     left1 = left1 - (hypotenuse - left)/2
//     return [left1, top1]
// }

