const host = window.location.host;
const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
const socket = new WebSocket(`${protocol}//${host}/ws`);

const CIRCUMFERENCE = 2 * Math.PI * 40; // r=40

function setProgress(circleId, usageId, percentage) {
    const circle = document.getElementById(circleId);
    if (circle) {
        const offset = CIRCUMFERENCE - (percentage / 100) * CIRCUMFERENCE;
        circle.style.strokeDashoffset = offset;
    }
    const usageTextNum = document.getElementById(usageId);
    if (usageTextNum) {
        usageTextNum.innerText = percentage;
    }
}

socket.onmessage = (event) => {
    const data = JSON.parse(event.data);

    // CPU Stats
    setProgress("cpu_circle", "cpu_usage_num", data.cpu_usage);
    const cpuCT = document.getElementById("cpu_ct");
    const cpuTemp = document.getElementById("cpu_temp");
    const cpuPower = document.getElementById("cpu_power");
    if (cpuCT) cpuCT.innerText = `${data.cpu_cores}C / ${data.cpu_threads}T`;
    if (cpuTemp) cpuTemp.innerText = `${data.cpu_temp}°C`;
    if (cpuPower) cpuPower.innerText = `${data.cpu_power}.0 W`;

    // RAM Stats
    setProgress("ram_circle", "ram_usage_num", data.ram_usage);
    const ramUsed = document.getElementById("ram_used");
    const ramFree = document.getElementById("ram_free");
    const ramTotal = document.getElementById("ram_total");
    const freeRAM = data.ram_total - data.ram_used;
    if (ramUsed) ramUsed.innerText = `${data.ram_used} MiB`;
    if (ramFree) ramFree.innerText = `${freeRAM} MiB`;
    if (ramTotal) ramTotal.innerText = `${data.ram_total} MiB`;

    // Disk Stats
    setProgress("disk_circle", "disk_usage_num", data.disk_usage);
    const diskUsed = document.getElementById("disk_used");
    const diskFree = document.getElementById("disk_free");
    const diskTotal = document.getElementById("disk_total");
    const freeDisk = data.disk_total - data.disk_used;
    if (diskUsed) diskUsed.innerText = `${data.disk_used} GB`;
    if (diskFree) diskFree.innerText = `${freeDisk} GB`;
    if (diskTotal) diskTotal.innerText = `${data.disk_total} GB`;
};

socket.onopen = () => { console.log("Connected to the Go pipe!"); };
socket.onclose = () => { console.log("Connection lost."); };
socket.onerror = (error) => { console.log("WebSocket Error: ", error); };