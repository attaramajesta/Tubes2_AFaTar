import './Project.css'
import Mario from '../../assets/mario.png'
import Luigi from '../../assets/luigi.png'
import Peach from '../../assets/peach.png'
import { useEffect, useRef } from 'react';

const Project = () => {
    const parallaxRef = useRef();

    useEffect(() => {
      const parallax_el = parallaxRef.current;
  
      let xValue = 0, yValue = 0;
  
      const handleMouseMove = (e) => {
        xValue = e.clientX - window.innerWidth / 2;
        yValue = e.clientY - window.innerHeight / 2;
  
        parallax_el.querySelectorAll(".parallax").forEach(el => {
          let speedx = el.dataset.speedx;
          let speedy = el.dataset.speedy;
          el.style.transform = `translateX(calc(-50% + ${-xValue * speedx}px)) translateY(calc(-50% + ${yValue * speedy}px))`;
        });
      };
  
      window.addEventListener("mousemove", handleMouseMove);
  
      return () => {
        window.removeEventListener("mousemove", handleMouseMove);
      };
    }, []);

    return (
        <div id='home' ref={parallaxRef}>
            <div className="vignette"></div>
            <img src={Luigi} data-speedx="0.008" data-speedy="0.01" className="parallax luigi" alt="Luigi" />
            <img src={Peach} data-speedx="0.009" data-speedy="0.02" className="parallax peach" alt="Peach" />
            <div>
            <div className="triangle"></div>
            <h1 className="home-title">AFATAR</h1>
            <h2 className="home-subheading">Ready to WikiRace?</h2>
            </div>
            <img src={Mario} data-speedx="0.01" data-speedy="0.03" className="parallax mario" alt="Mario" />
        </div>
    );
};

export default Project