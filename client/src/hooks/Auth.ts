import { useState } from "react";
import * as client from "../httpClient/client";

export function useAuth() {
  const [token, setToken] = useState<string | null>(null);

  const localStorageToken = localStorage.getItem(client.AUTHENTICATION_TOKEN_KEY);

  if (localStorageToken) {
    client.post("/authenticated")
      .then(_ => setToken(localStorageToken));
  }

  return {
    token: token,
    setToken: (token: string) => {
      localStorage.setItem(client.AUTHENTICATION_TOKEN_KEY, token);
      setToken(token);
    },
    unsetToken: () => {
      localStorage.removeItem(client.AUTHENTICATION_TOKEN_KEY);
      setToken(null);
    }
  };
}
