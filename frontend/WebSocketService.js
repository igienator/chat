export default class WebSocketService {
  constructor() {
    this.socket = null;
  }

  connect(roomId) {
    const url = `ws://localhost:8080/ws?room=${encodeURIComponent(roomId)}`;
    this.socket = new WebSocket(url);
    this.socket.onopen = () => console.log('WebSocket connected');
    this.socket.onmessage = async event => {
      const envelope = JSON.parse(event.data);
      if (envelope.type === 'error') {
        console.error(envelope.payload);
        return;
      }
      // dispatch to UI via callbacks or state
      this.handleEnvelope(envelope);
    };
  }

  disconnect() {
    if (this.socket) {
      this.socket.close();
      this.socket = null;
    }
  }

  send(type, payload) {
    const envelope = JSON.stringify({type, payload});
    if (this.socket && this.socket.readyState === WebSocket.OPEN) {
      this.socket.send(envelope);
    }
  }

  handleEnvelope(envelope) {
    // implement logic for 'pubkey', 'message', 'peer-joined', etc.
    console.log('Received envelope:', envelope.type, envelope.payload);
  }
}
