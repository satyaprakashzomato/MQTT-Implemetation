// <script src="https://unpkg.com/mqtt/dist/mqtt.min.js"></script>
// const mqtt = require('mqtt')
import mqtt from 'mqtt'

// connect options
const options = {
      connectTimeout: 4000,

      // Authentication
      clientId: 'emqx',
      // username: 'emqx',
      // password: 'emqx',

      keepalive: 60,
      clean: true,
}

// WebSocket connect url
const WebSocket_URL = 'ws://localhost:8083/mqtt'

// TCP/TLS connect url
const TCP_URL = 'mqtt://localhost:1883'
const TCP_TLS_URL = 'mqtts://localhost:8883'

const client = mqtt.connect(WebSocket_URL, options)

client.on('connect', () => {
    console.log('Connect success')
})

client.on('reconnect', (error) => {
    console.log('reconnecting:', error)
})

client.on('error', (error) => {
    console.log('Connect Error:', error)
})
