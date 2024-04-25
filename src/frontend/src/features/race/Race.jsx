import "./Race.css"
import getInput from "./Input";

const Race = () => {
    return(
        <div id="race" className="race-page">
            <div className='title'>
            <h2>Find the shortest path from</h2>
            </div>
            <div className='input-container'>
            <input type="text" id="startInput" placeholder="Superman"></input>
            <input type="text" id="destInput" placeholder="Marine"></input>
            </div>
            <div className='switch-container'>
            <span>BFS</span>
            <label className="switch">
                <input type="checkbox" id="algoSwitch"></input>
                <span className="slider round"></span>
            </label>
            <span>IDS</span>
            </div>
            <div className='submit-container'>
            <button onClick={getInput}>RACE!</button>
            </div>
      </div>
    );
}

export default Race