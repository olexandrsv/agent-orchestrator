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

drawLine(300, 200, 600, 400)

function drawLine(x1, y1, x2, y2){
    let x = x2 - x1
    let y = y2 - y1
    let hypotenuse = Math.sqrt(Math.pow(x, 2) + Math.pow(y, 2))

    let radians = Math.acos(x/hypotenuse);
    let degrees = radians * (180 / Math.PI);
    y1 = y1 + y/2
    x1 = x1 - (hypotenuse/2 - x/2)

    const line = `<div class="connection" style="--y:${y1}px; --x:${x1}px; --len:${hypotenuse}px; --angle:${degrees}deg;"></div>`
    let body = document.getElementById("connections")
    body.innerHTML = body.innerHTML + line


}

