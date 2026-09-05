export class AgentUI {
    constructor(left, top, color, onclick){
        let div = document.createElement("div")
        let colorRgbCode = "228, 40, 33"
        if (color === "blue"){
            colorRgbCode = "33, 121, 228"
        }

        div.className = "agent-ui"
        div.style.setProperty("--left", Number(left)+"px")
        div.style.setProperty("--top", Number(top)+"px")
        div.style.setProperty("--color", colorRgbCode)

        div.onclick = onclick
        document.getElementById("agents").appendChild(div)
    }
}
