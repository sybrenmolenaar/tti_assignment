import React from "react";
import logo from "./logo-tti.png";
import Socket from "socket.io-client";
import { Counter } from "./features/counter/Counter";
import "./App.css";

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <Counter />
      </header>
    </div>
  );
}

export default App;
