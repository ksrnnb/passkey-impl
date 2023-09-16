import { createContext } from "react"

interface AuthContext {
  user: User | null;
}

export const authContext = createContext<AuthContext>({
  user: null,
});

