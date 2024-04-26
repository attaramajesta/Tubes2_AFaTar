import React, { useState } from 'react';
import "./Race.css"

const Race = () => {
    const [startInput, setStartInput] = useState("");
    const [destInput, setDestInput] = useState("");
    const [algoSwitch, setAlgoSwitch] = useState(false);

    const fetchPath = async () => {
        try {
            const algorithm = algoSwitch ? "ids" : "bfs";
            const response = await fetch(`http://localhost:8080/${algorithm}?start=${encodeURIComponent(startInput)}&target=${encodeURIComponent(destInput)}`);
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            const data = await response.json();
            console.log(data);
        } catch (error) {
            console.error('Fetch error:', error);
        }
    };

    return(
        <div id="race" className="race-page">
            <div className='title'>
            <h2>Find the shortest path from</h2>
            </div>
            <div className='input-container'>
            <input type="text" id="startInput" placeholder="Superman" value={startInput} onChange={e => setStartInput(e.target.value)}></input>
            <input type="text" id="destInput" placeholder="Marine" value={destInput} onChange={e => setDestInput(e.target.value)}></input>
            </div>
            <div className='switch-container'>
            <span>BFS</span>
            <label className="switch">
                <input type="checkbox" id="algoSwitch" checked={algoSwitch} onChange={e => setAlgoSwitch(e.target.checked)}></input>
                <span className="slider round"></span>
            </label>
            <span>IDS</span>
            </div>
            <div className='submit-container'>
            <button onClick={fetchPath}>RACE!</button>
            </div>
      </div>
    );
}

export default Race;