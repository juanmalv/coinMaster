import React, { Component } from 'react'
import QRCode from 'qrcode'

export class Qrcode extends Component {

generateQR() {
     let str = 'Aparezco en lA PORTADA!!!'
     QRCode.toCanvas(document.getElementById('canvas'), str,function(error) {
          if (error) console.error(error)
          console.log('success!')
     })
}
render() {
return (
<div align="center">
     <canvas id="canvas" />
     <button onClick={this.generateQR}>
          Generate QR!
     </button>
</div>
)}}