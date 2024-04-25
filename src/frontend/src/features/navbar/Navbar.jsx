import "./Navbar.css"
import { Link } from "react-scroll"

export default function Navbar()
{
    return(
        <div className="navbar">
            <nav>
                <Link to='home' smooth={true} spy={true} duration={500}>WIKIRACE!</Link>
            <ul>
                <li>
                    <Link to='about-us' smooth={true} spy={true} duration={500}>ABOUT</Link>
                </li>
                <li>
                    <Link to='race' smooth={true} spy={true} duration={500}>RACE</Link>
                </li>
            </ul>
            </nav>
      </div>
    )
}
