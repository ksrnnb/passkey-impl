import { useCallback, useEffect, useState } from "react";
import * as client from "../httpClient/client";

type Credential = {
  id: string;
  name: string;
};

export type User = {
  userId: string;
  credentials: Credential[];
};

type AuthenticatedResponse = {
  userId: string;
  credentials: Credential[];
};

export function useAuth() {
  const [token, setToken] = useState<string | null>(null);
  const [user, setUser] = useState<User | null>(null);
  const [isLoading, setIsLoading] = useState<boolean>(false);

  const storedToken = localStorage.getItem(client.AUTHENTICATION_TOKEN_KEY);

  const updateUser = useCallback(() => {
    if (storedToken) {
      setIsLoading(true);
      client.Post("/authenticated")
      .then(res => res.json())
      .then((res: AuthenticatedResponse) => {
          setToken(storedToken);
          setUser({
            userId: res.userId,
            credentials: res.credentials,
          });
          setIsLoading(false);
        });
    }
  }, [storedToken]);

  useEffect(() => {
    updateUser();
  }, [updateUser]);

  return {
    token: token,
    user: user,
    updateUser: updateUser,
    isLoading: isLoading,
    setToken: (token: string) => {
      localStorage.setItem(client.AUTHENTICATION_TOKEN_KEY, token);
      setToken(token);
      updateUser();
    },
    unsetToken: () => {
      localStorage.removeItem(client.AUTHENTICATION_TOKEN_KEY);
      setToken(null);
      setUser(null);
    }
  };
}
