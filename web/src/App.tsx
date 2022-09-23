import { Component, createEffect, createSignal, onMount, Show } from "solid-js";
import {
  fetchSensorData,
  fetchSensorState,
  sensorData,
  sensorState,
} from "./store";
import { createWS, registerWSEvents } from "./ws";

const App: Component = () => {
  onMount(() => registerWSEvents(createWS()));
  const [sensorId, setSensorId] = createSignal<string>("");
  createEffect(() => console.log(sensorState()));
  createEffect(() => console.log(sensorData()));
  return (
    <div>
      <div>
        <button onClick={fetchSensorState}>Fetch Sensor State</button>
        <form>
          <label>
            Sensor ID
            <input
              type="text"
              value={sensorId()}
              onInput={(e) => setSensorId(e.currentTarget.value)}
            />
          </label>
          <button
            onClick={(e) => {
              e.preventDefault();
              fetchSensorData(sensorId());
              setSensorId("");
            }}
          >
            Fetch Sensor Data
          </button>
        </form>
      </div>
      <Show when={sensorState()}>
        <span>SensorState</span>
        <pre>{JSON.stringify(sensorState(), null, 4)}</pre>
      </Show>
      <br />
      <br />
      <Show when={sensorData()}>
        <span>SensorData</span>
        <pre>{JSON.stringify(sensorData(), null, 4)}</pre>
      </Show>
    </div>
  );
};

export default App;
