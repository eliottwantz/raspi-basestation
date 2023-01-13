import { SensorData } from '../pb/sensor';
import { SensorState } from '../pb/sensorstate';
import { setSensorData, setSensorState } from './store';

type WSMessage = SensorData | SensorState;

export function createWS(): WebSocket {
  return new WebSocket('ws://192.168.1.13:8000/ws/'); // Raspi
  // return new WebSocket('ws://10.0.0.221:8000/ws/'); // Raspi
  // return new WebSocket('ws://localhost:8000/ws/');
}

export function registerWSEvents(ws: WebSocket) {
  ws.onopen = function () {
    console.log('WEB SOCKET OPENED');
  };

  ws.onmessage = function (event: MessageEvent<any>) {
    console.log(event);
    const obj: WSMessage = JSON.parse(event.data);
    if ((obj as SensorState).brakeManager) setSensorState(obj as SensorState);
    else setSensorData(obj as SensorData);
    console.log(obj);
  };

  ws.onclose = function () {
    console.log('WEB SOCKET CLOSED');
  };
}
