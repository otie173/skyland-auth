import express, { Request, Response, NextFunction } from "express";

const loggerMiddleware = function (
  req: Request,
  res: Response,
  next: NextFunction,
) {
  console.log(`New request: ${req.method} ${req.url}`);
  next();
};

export function setupMiddlewares(app: express.Application) {
  app.use(loggerMiddleware);
}
