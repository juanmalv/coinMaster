import React, { Component } from 'react'
import QRCode from 'qrcode'

export class Qrcode extends Component {

generateQR() {
     let str = 'https://medium.com/@mshalam04/integrating-qr-codes-with-react-874da3e5d98e'
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