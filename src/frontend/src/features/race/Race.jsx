import React, { useState } from 'react';
import Heart from '../../assets/heart.png';
import Energy from '../../assets/energy.png';
import MarioBFS from '../../assets/bfs-mario.png';
import LuigiIDS from '../../assets/ids-luigi.png';
import "./Race.css"

const Race = () => {
    const [suggestionAwal, setSuggestionsAwal] = useState([]);
    const [suggestionAkhir, setSuggestionsAkhir] = useState([]);
    const [startInput, setStartInput] = useState("");
    const [destInput, setDestInput] = useState("");
    const [algoSwitch, setAlgoSwitch] = useState(false);
    const [showResultPage, setShowResultPage] = useState(false);
    const [resultData, setResultData] = useState({});
    const [loading, setLoading] = useState(false);
    const [buttonClicked, setButtonClicked] = useState(false);

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
        setLoading(true);
        try {
            const algorithm = algoSwitch ? "ids" : "bfs";
            const startInputFormatted = startInput.replace(/\s/g, "_");
            const destInputFormatted = destInput.replace(/\s/g, "_");
            const response = await fetch(`http://localhost:8080/${algorithm}?start=${encodeURIComponent(startInputFormatted)}&target=${encodeURIComponent(destInputFormatted)}`);
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            const data = await response.json();
            console.log(data.path);
            setResultData(data);
            setLoading(false);
            setShowResultPage(true);
        } catch (error) {
            setLoading(false);
            console.error('Fetch error:', error);
        }
    };

    const ResultPage = () => {
        return (
            <div className="result-box">
                <h2 className="title">Result</h2>
                <p>Duration: {resultData.duration} s </p>
                <p>Path:</p>
                <ul>
                    {resultData.path.map((step, index) => (
                        <li key={index}><a href={`https://en.wikipedia.org/wiki/${step}`}>{step}</a></li>
                    ))}
                </ul>
                <p>Path Depth: {resultData.depth} </p>
                <p>Total Visited Pages: {resultData.totalVisited} </p>
            </div>
        );
    }
      
    const changeColor = () => {
        var button = document.getElementById("raceButton");
        button.classList.add("clicked");
        setTimeout(() => {
            button.classList.remove("clicked");
        }, 200);
    }

    return (
        <div id="race" className="race-page">
            <div className="race-container">
                <img src={Heart} className="stat" alt="heart rate"/>
                <img src={Energy} className="stat" alt="heart rate"/>
                <img src={algoSwitch ? LuigiIDS : MarioBFS} alt={algoSwitch ? 'MarioBFS' : 'LuigiIDS'} className="choose-character"/>
                <div className="race-box">
                    <div className="race-title">
                        <h2>Find the shortest path from</h2>
                    </div>
                    <div className='input-box'>
                        <input autocomplete="off" type="text" id="startInput" placeholder="Superman" value={startInput} onChange={handleInputAwal}/>
                        <div className="suggestion">
                            {suggestionAwal.map((suggestion, index) => (
                                <p key={index} onClick={() => { setStartInput(suggestion); setSuggestionsAwal([]); }}>{suggestion}</p>
                            ))}
                        </div>
                    </div>
                    <h2 className="to">to</h2>
                    <div className='input-box'>
                        <input autocomplete="off" type="text" id="destInput" placeholder="Marine" value={destInput} onChange={handleInputAkhir} />
                        <div className="suggestion">
                            {suggestionAkhir.map((suggestion, index) => (
                                <p key={index} onClick={() => { setDestInput(suggestion); setSuggestionsAkhir([]); }}>{suggestion}</p>
                            ))}
                        </div>
                    </div>
                    <div className='switch-container'>
                        <span>BFS</span>
                        <label className="switch">
                            <input type="checkbox" id="algoSwitch" checked={algoSwitch} onChange={e => setAlgoSwitch(e.target.checked)} />
                            <span className="slider round"></span>
                        </label>
                        <span>IDS</span>
                    </div>
                    <div className='submit-container'>
                        <button id="raceButton" onClick={() => {fetchPath(); changeColor();}}>RACE!</button>
                    </div>
                </div>
            </div>
            <div className="highlight"></div>
            {showResultPage && <ResultPage />}
        </div>
    );
}
export default Race;