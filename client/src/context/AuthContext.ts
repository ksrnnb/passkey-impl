import { createContext } from "react"

interface AuthContext {
  token: string | null;
  setToken: (token: string) => void;
}

const authContext = createContext<AuthContext>({
  token: null,
  setToken: (token: string) => {},
});

export default authContext;
