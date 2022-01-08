import React from 'react';
import { HashRouter } from 'react-router-dom';
import './index.css';
import App from './App';


const Root = () => (

    // <React.StrictMode>
    <HashRouter>
      <App />
    </HashRouter>
    // </React.StrictMode>,
)


export default Root;