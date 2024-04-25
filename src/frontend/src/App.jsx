import React, { useRef, useEffect } from 'react';
import './index.css';
import Navbar from "../src/features/navbar/Navbar.jsx";
import Project from "../src/features/project/Project.jsx";
import About from "../src/features/about/About.jsx";
import Race from './features/race/Race.jsx';

const App = () => {
  return (
    <>
    <div className='container'>
      <Navbar/>
      <Project/>
      <About id='#about'/>
      <Race id='#race'/>
    </div>
    </>
  );
}

export default App;
