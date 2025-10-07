import { Router, Request, Response } from "express";

const router = Router();

router.post("/register", (req: Request, res: Response) => {
  res.send("Пользователь пытается зарегистрироваться");
});

export default router;
