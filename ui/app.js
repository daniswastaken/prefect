const socket = new WebSocket("ws://localhost:8080/ws");

socket.onmessage = (event) => {
    // Parse the JSON string coming from Go into a JS Object
    const data = JSON.parse(event.data);

    const cpuEl = document.getElementById("cpu-value");
    const ramEl = document.getElementById("ram-value");

    if (cpuEl) {
        cpuEl.innerText = `${data.cpu_usage}%`;
    }

    if (ramEl) {
        ramEl.innerText = `${data.ram_percent}% (${data.ram_used}MiB / ${data.ram_total}MiB)`;
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