interface ImportMetaEnv extends Readonly<Record<string, string>> {
  readonly VITE_APP_TITLE: string;
  readonly VITE_SOCKET_URI: string;
  readonly VITE_API_HOST: string;
  readonly VITE_AUTH_SERVICE: string;
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}
