import * as metroDraw from "./metroDraw.js"
import * as mapParser from "./mapParser.js"


let socket = new WebSocket("ws://127.0.0.1:8080/ws");

socket.onopen = () => {
  socket.send("Hi From the Client!")
};

socket.onclose = event => {
  socket.send("Client Closed!")
};

socket.onmessage = (ev) => {
  let metroMap = mapParser.parseMap(ev.data)
  metroDraw.updateMap(metroMap)
}

socket.onerror = error => {
};
