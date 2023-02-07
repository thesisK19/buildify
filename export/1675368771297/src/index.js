import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import {RouterDOM} from './routes';

ReactDOM.render(
  <React.StrictMode>
    <RouterDOM />
  </React.StrictMode>,
  document.getElementById('root')
);