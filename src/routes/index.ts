import authRoutes from "../features/auth/routes";
import express from "express";

export function setupRoutes(app: express.Application) {
  app.use("/api/v1/", authRoutes);
}
