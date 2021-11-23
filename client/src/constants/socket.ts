enum WorkerAction {
  CreateSocket,
}

interface WorkerTask {
  action: WorkerAction;
  data: any;
}

interface WorkerMessage {
  action: WorkerAction;
  success: boolean;
  data: any;
}
