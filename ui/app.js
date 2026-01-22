// Automatically detects the current host and port
const host = window.location.host; 
const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
const socket = new WebSocket(`${protocol}//${host}/ws`);

socket.onmessage = (event) => {
    // Parse the JSON string coming from Go into a JS Object
    const data = JSON.parse(event.data);

    const cpuCores = document.getElementById("cpu_cores");
    const cpuUsage = document.getElementById("cpu_usage");
    const ramUsed = document.getElementById("ram_used");
    const ramTotal = document.getElementById("ram_total");
    const ramUsage = document.getElementById("ram_usage");

    if (cpuCores) {
        cpuCores.innerText = `${data.cpu_cores} Cores`;
    }

    if (cpuUsage) {
        cpuUsage.innerText = `${data.cpu_usage}%`;
    }

    if (ramUsage) {
        ramUsage.innerText = `${data.ram_usage}%`;
    }

    if (ramUsed) {
        ramUsed.innerText = `${data.ram_used}MiB`;
    }

    if (ramTotal) {
        ramTotal.innerText = `${data.ram_total}MiB`;
    }
    
    console.log("Stats received from Go:", data);
};

socket.onopen = () => {
    console.log("Connected to the Go pipe!");
};

socket.onclose = () => {
    console.log("Connection lost. Is the Go server running?");
};

socket.onerror = (error) => {
    console.log("WebSocket Error: ", error);
};