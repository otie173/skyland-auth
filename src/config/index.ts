require("@dotenvx/dotenvx").config();

export const config = {
  port: process.env.SERVER_PORT || 8080,
};
