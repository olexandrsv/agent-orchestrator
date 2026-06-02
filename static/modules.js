const modules = {
    Modules: [
        {
            Name: "LineUI",
            Path: "ui/line",
        },
        {
            Name: "Agent",
            Path: "agent"
        },
        // {
        //     Name: "Connection",
        //     Path: "connection",
        // },
        {
            Name: "ArrowUI",
            Path: "ui/arrow",
        },
        {
            Name: "ConnectionUI",
            Path: "ui/connection"
        }
    ]
}

initModules(modules)

function initModules(modules){
    for (const module of modules.Modules){
        const link = document.createElement("link");
        link.rel = "stylesheet";
        link.href = `${module.Path}/style.css`;

        document.head.appendChild(link);

        const script = document.createElement("script")
        script.src = `${module.Path}/script.js`
        script.type = "module"

        console.log(script)
        document.body.appendChild(script)
    }
}