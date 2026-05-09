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
                X: 40,
                Y: 40,
                Color: "red",
            }
        }
    ]
}

placeArrow(300, 200, 600, 400)
placeArrow(600, 400, 300, 200)
placeArrow(100, 200, 109, 400)
placeArrow(100, 200, 180, 250)

function placeArrow(left1, top1, left2, top2){
    const [lineX, lineY, len, angle] = computeLine(left1, top1, left2, top2)
    const angleRad = angle * Math.PI / 180;

    const adjacent = 34 * Math.cos(angleRad);
    const opposite = 34 * Math.sin(angleRad);

    drawLine(lineX, lineY, len, angle)
    drawArrow(left2 - adjacent, top2 - opposite, angle)
}

function computeLine(left1, top1, left2, top2){
    let left = left2 - left1
    let top = top2 - top1
    let hypotenuse = Math.sqrt(Math.pow(left, 2) + Math.pow(top, 2))

    let radians = Math.acos(left/hypotenuse);
    let degrees = radians * (180 / Math.PI);

    console.log("---")
    console.log("left:", left)
    console.log("top:", top)
    console.log("hypotenuse:", hypotenuse)
    console.log("radians:", radians)
    console.log("degrees:", degrees)
    top1 = top1 + top/2
    left1 = left1 - (hypotenuse - left)/2
    if (top < 0) {
        degrees = -degrees
    }

    return [left1, top1, hypotenuse, degrees]
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

