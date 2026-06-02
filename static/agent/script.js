class AgentUI extends HTMLElement {
    connectedCallback(){
        this.left = this.getAttribute("left")
        this.top = this.getAttribute("top")
        this.color = this.getAttribute("color")

        this.style.left = `${this.left}px`
        this.style.top = `${this.top}px`
    }

    onClick(){
        document.getElementsByClassName("entity-state")[0].classList.add("active");
    }
}

customElements.define("my-agent", AgentUI)