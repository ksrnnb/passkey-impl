import { createContext } from "react"
import { User } from "../hooks/Auth";

interface AuthContext {
  user: User | null;
  setToken: (token: string) => void;
}

const authContext = createContext<AuthContext>({
  user: null,
  setToken: (token: string) => {},
});

export default authContext;
