declare global {
  namespace NodeJS {
    interface ProcessEnv {
      DB_NAME: string;
      DB_PASSWORD: string;
      DB_USERNAME: string;
      DB_PORT: number;
      DB_HOST: string;
      JWT_SECRET: string;
      MAIN_TOKEN_EXPIRE: string;
      WS_TOKEN_EXPIRE: string;
      API_HOST: string;
      USER_PATH: string;
    }
  }
}

// If this file has no import/export statements (i.e. is a script)
// convert it into a module by adding an empty export statement.
export {};
