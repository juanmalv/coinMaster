import React from 'react';
import logo from './logo.svg';
import Welcome from './components/income'
import { Qrcode } from './components/qrGenerator'
import './App.css';
import { LectorQr } from './components/qrReader';
function App() {
  return (
    <div>
    <Qrcode/>
    <LectorQr/>
    </div>
  );
}

export default App;
