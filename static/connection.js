class Agent{
    constructor(left, top, color){
        this.left = left
        this.top = top
        this.color
        this.drawAgent(left, top)
    }

    drawAgent(left, top){
        const div = document.createElement("my-agent");

        div.setAttribute("top", `${this.top}`);
        div.setAttribute("left", `${this.left}`);
        div.addEventListener("click", () => {
            console.log("here")
            this.onClick();
        });

        document.getElementById("agents").appendChild(div);
    }

    onClick(){
        console.log("agent clicked")
        document.getElementsByClassName("entity-state")[0].classList.add("active");
    }
}

function closeState(){
    document.getElementsByClassName("entity-state")[0].classList.remove("active");
}

function setActive(){
    document.getElementsByClassName("entity-state")[0].classList.add("active");
}