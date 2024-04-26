import React, { useState } from 'react';
import "./Race.css"

const Race = () => {
    const [suggestionAwal, setSuggestionsAwal] = useState([]);
    const [suggestionAkhir, setSuggestionsAkhir] = useState([]);
    const [startInput, setStartInput] = useState("");
    const [destInput, setDestInput] = useState("");
    const [algoSwitch, setAlgoSwitch] = useState(false);

    const handleInputAwal = async (e) => {
        const query = e.target.value;
        setStartInput(query);

        if (query) {
            const response = await fetch(`https://en.wikipedia.org/w/api.php?origin=*&action=opensearch&search=${query}&limit=5&namespace=0&format=json`);
            const data = await response.json();

            setSuggestionsAwal(data[1]);
        } else {
            setSuggestionsAwal([]);
        }
    }

    const handleInputAkhir = async (e) => {
        const query = e.target.value;
        setDestInput(query);

        if (query) {
            const response = await fetch(`https://en.wikipedia.org/w/api.php?origin=*&action=opensearch&search=${query}&limit=5&namespace=0&format=json`);
            const data = await response.json();

            setSuggestionsAkhir(data[1]);
        } else {
            setSuggestionsAkhir([]);
        }
    }

    const fetchPath = async () => {

        try {
            const algorithm = algoSwitch ? "ids" : "bfs";
            const startInputFormatted = startInput.replace(/\s/g, "_");
            const destInputFormatted = destInput.replace(/\s/g, "_");
            const response = await fetch(`http://localhost:8080/${algorithm}?start=${encodeURIComponent(startInputFormatted)}&target=${encodeURIComponent(destInputFormatted)}`);
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
            <div className='input-box'>
                <input type="text" id="startInput" placeholder="Superman" value={startInput} onChange={handleInputAwal}/>
                <div className="suggestion">
                    {suggestionAwal.map((suggestion, index) => (
                        <p key={index} onClick={() => {setStartInput(suggestion); setSuggestionsAwal([]);}}>{suggestion}</p>
                    ))}
                </div>
            </div>
            <div className='input-box'>
                <input type="text" id="destInput" placeholder="Marine" value={destInput} onChange={handleInputAkhir}/>
                <div className="suggestion">
                    {suggestionAkhir.map((suggestion, index) => (
                        <p key={index} onClick={() => {setDestInput(suggestion); setSuggestionsAkhir([]);}}>{suggestion}</p>
                    ))}
                </div>
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