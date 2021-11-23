let counter = 1;

const msg = "pong";
const mode = import.meta.env.MODE;
let socket!: WebSocket;

self.onmessage = (e: MessageEvent<WorkerTask>) => {
  switch (e.data.action) {
    case WorkerAction.CreateSocket:
      createSocket();
      break;
    default:
      console.log("default case");
      break;
  }
};

async function createSocket() {
  if (!socket || socket.readyState === WebSocket.CLOSED) {
    const host = import.meta.env.VITE_SOCKET_URI;
    const res = await window.axiosAuth.post("/token/ws");
    const token = (res.data as any).ws_token;
    socket = new WebSocket(`${host}?token=${token}`);
    self.postMessage({
      action: WorkerAction.CreateSocket,
      success: true,
    });
  } else {
    console.log(`socket connection is ready`);
  }
}
