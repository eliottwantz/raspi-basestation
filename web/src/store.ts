import { createResource, createSignal } from "solid-js";
import { SensorData } from "../pb/sensor";
import { SensorState } from "../pb/sensorstate";

const fetchSensorState = async () => {
  const res = await fetch("http://localhost:8000/api/SensorState");
  const ss: SensorState = await res.json();
  return ss;
};

export const fetchSensorData = async (sensorId: string) => {
  console.log(sensorId);
  if (!sensorId) return;
  const res = await fetch(`http://localhost:8000/api/SensorData/${sensorId}`);
  const sd: SensorData = await res.json();
  setSensorData(sd);
  return sd;
};

export const [sensorState, sensorStateReq] = createResource(fetchSensorState);
export const [sensorData, setSensorData] = createSignal<SensorData>();
