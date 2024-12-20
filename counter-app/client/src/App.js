import React, { useEffect, useState } from "react";
import axios from "axios";

const API_BASE_URL = "http://localhost:8080";
const WS_URL = "ws://localhost:8080/ws";

function App() {
  const [counters, setCounters] = useState({});
  const [error, setError] = useState("");

  // Fetch counters on component load
  useEffect(() => {
    fetchCounters();
  }, []);

  const fetchCounters = async () => {
    try {
      const response = await axios.get(`${API_BASE_URL}/counters`);
      setCounters(response.data || {}); // Ensure data is an object
    } catch {
      setError("Failed to fetch counters");
    }
  };

  // WebSocket connection
  useEffect(() => {
    const ws = new WebSocket(WS_URL);

    ws.onopen = () => console.log("WebSocket connected");
    ws.onmessage = (event) => {
      try {
        const data = JSON.parse(event.data);
        setCounters(data || {}); // Ensure data is an object
      } catch {
        console.error("Failed to parse WebSocket message");
      }
    };
    ws.onerror = () => setError("WebSocket error occurred");
    ws.onclose = (event) =>
      console.log("WebSocket disconnected:", event.reason || "Unknown reason");

    return () => ws.close();
  }, []);

  const createCounter = async () => {
    try {
      await axios.post(`${API_BASE_URL}/counter`);
    } catch {
      setError("Failed to create counter");
    }
  };

  const updateCounter = async (id, value) => {
    try {
      await axios.put(`${API_BASE_URL}/counter/${id}`, { value });
    } catch {
      setError("Failed to update counter");
    }
  };

  const deleteCounter = async (id) => {
    try {
      await axios.delete(`${API_BASE_URL}/counter/${id}`);
    } catch {
      setError("Failed to delete counter");
    }
  };

  return (
    <div style={{ textAlign: "center", marginTop: "50px" }}>
      <h1>Counter App</h1>

      <button onClick={createCounter}>Create New Counter</button>
      <div style={{ marginTop: "20px" }}>
        {counters && Object.keys(counters).length > 0 ? ( // Safely check counters
          Object.entries(counters).map(([id, counter]) => (
            <div
              key={id}
              style={{
                display: "flex",
                justifyContent: "center",
                alignItems: "center",
                margin: "10px 0",
              }}
            >
              <span style={{ marginRight: "10px" }}>
                Counter {id}: {counter.value}
              </span>
              <button onClick={() => updateCounter(counter.id, counter.value + 1)}>
                Increment
              </button>
              <button onClick={() => updateCounter(counter.id, counter.value - 1)}>
                Decrement
              </button>
              <button onClick={() => deleteCounter(counter.id)}>Delete</button>
            </div>
          ))
        ) : (
          <p>No counters available</p>
        )}
      </div>
    </div>
  );
}

export default App;
