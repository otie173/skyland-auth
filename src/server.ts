import { createApp } from "./app";
import { config } from "./config";

const app = createApp();
const port = config.port;

const server = app.listen(port, () => {
  console.log(`Server is running on http://localhost:${port}`);
});

process.on("SIGTERM", () => {
  console.log("SIGTERM received, shutting down gracefully");
  server.close(() => {
    console.log("Process terminated");
  });
});

process.on("SIGINT", () => {
  console.log("SIGINT received, shutting down gracefully");
  server.close(() => {
    console.log("Process terminated");
  });
});
