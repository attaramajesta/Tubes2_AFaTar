import './About.css';

const About = () => {
    return (
        <div className='about-us'>
            <div className="racer-container">
                <p>Hello 👋 I'm</p>
                <section className="animation">
                <div className="first" ><div>Front-End Developer</div></div>
                <div className="second"><div>Back-End Web Developer</div></div>
                <div className="third"><div>Algorithm Developer</div></div>
                </section>
            </div>
            <div className="ag-format-container">
                <div className="ag-courses_box">
                <div className="ag-courses_item">
                    <a href="https://github.com/attaramajesta" className="ag-courses-item_link">
                    <div className="ag-courses-item_bg"></div>

                    <div className="ag-courses-item_title">
                        Attara Majesta Ayub
                    </div>

                    <div className="ag-courses-item_desc-box">
                    Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed tristique, odio id venenatis lobortis, metus risus maximus ipsum, at blandit lacus nisl eget diam.
                    </div>
                    <span className="ag-courses-item_nim">
                        13522139
                    </span>
                    </a>
                </div>

                <div className="ag-courses_item">
                    <a href="https://github.com/fnathas" className="ag-courses-item_link">
                    <div className="ag-courses-item_bg"></div>

                    <div className="ag-courses-item_title">
                        Farrel Natha Saskoro
                    </div>

                    <div className="ag-courses-item_desc-box">
                    Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed tristique, odio id venenatis lobortis, metus risus maximus ipsum, at blandit lacus nisl eget diam.
                    </div>
                    <span className="ag-courses-item_nim">
                        13522145
                    </span>
                    </a>
                </div>

                <div className="ag-courses_item">
                    <a href="https://github.com/AxelSantadi" className="ag-courses-item_link">
                    <div className="ag-courses-item_bg"></div>

                    <div className="ag-courses-item_title">
                        Axel Santadi Warih
                    </div>

                    <div className="ag-courses-item_desc-box">
                    Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed tristique, odio id venenatis lobortis, metus risus maximus ipsum, at blandit lacus nisl eget diam.
                    </div>
                    <span className="ag-courses-item_nim">
                        13522155
                    </span>
                    </a>
                </div>
                </div>
            </div>
        </div>
    );
};

export default About