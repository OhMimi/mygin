var ws = new WebSocket("ws://localhost:8080/ping");

// onOpen被觸發時，去嘗試連線
ws.onopen = (e) => {
    console.log("Connection open ...");
    ws.send("Hello WebSockets!");
};

// onMessage被觸發時，來接收ws server傳來的訊息
ws.onmessage = (e) => {
    console.log("Received message : " + e.data);
};

// 由 ws server 發出的 onClose事件
ws.onclose = (e) => {
    console.log("Connection closed!!");
};

// 每秒發出一個現在時間的訊息
var timeInterval = setInterval(() => {
    ws.send(Date.now())
}, 1000);

// 每5秒發出一個現在時間的訊息
var timeInterval = setInterval(() => {
    ws.send("ping")
}, 5000);