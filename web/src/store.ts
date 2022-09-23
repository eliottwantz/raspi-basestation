import { createSignal } from "solid-js";
import { SensorData } from "../pb/sensor";
import { SensorState } from "../pb/sensorstate";

export const fetchSensorState = async () => {
  const ss: SensorState = await fetch(
    "http://localhost:8000/api/SensorState"
  ).then((res) => res.json());
  setSensorState(ss);
  return ss;
};

export const fetchSensorData = async (sensorId: string) => {
  console.log(sensorId);
  if (!sensorId) return;
  const sd: SensorData = await fetch(
    `http://localhost:8000/api/SensorData/${sensorId}`
  ).then((res) => res.json());
  setSensorData(sd);
  return sd;
};

export const [sensorState, setSensorState] = createSignal<SensorState>();
export const [sensorData, setSensorData] = createSignal<SensorData>();
