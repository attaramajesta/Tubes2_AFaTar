import './About.css';

const About = () => {
    return (
        <div className='about-us'>
            <h2 className="ag-title">About Us</h2>
            <div className="ag-format-container">
                <div className="ag-courses_box">
                <div className="ag-courses_item">
                    <a href="https://github.com/attaramajesta" className="ag-courses-item_link">
                    <div className="ag-courses-item_bg"></div>

                    <div className="ag-courses-item_title">
                        Attara Majesta Ayub
                    </div>

                    <div className="ag-courses-item_desc-box">
                    .............ok
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
                    StimashaAllah.
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
                    It's always "kenapa Bandung?" but it's never "bagaimana Bandung?"
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