import {
  createContext,
  useContext,
  useState,
  ReactNode,
  useEffect,
} from "react";
import { BASE_URL } from "@/api";
import { AuthToken } from "@/types";

interface AuthContextProps {
  token: AuthToken;
  signup: (
    firstName: string,
    lastName: string,
    username: string,
    password: string,
  ) => Promise<void>;
  login: (username: string, password: string) => Promise<void>;
  logout: () => void;
  isAuthenticated: boolean;
}

const AuthContext = createContext<AuthContextProps | undefined>(undefined);

export const AuthProvider = ({ children }: { children: ReactNode }) => {
  const [token, setToken] = useState<string | null>(
    localStorage.getItem("token"),
  );

  useEffect(() => {
    if (token) {
      localStorage.setItem("token", token);
    } else {
      localStorage.removeItem("token");
    }
  }, [token]);

  const signup = async (
    firstName: string,
    lastName: string,
    username: string,
    password: string,
  ) => {
    const response = await fetch(`${BASE_URL}/signup`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        first_name: firstName,
        last_name: lastName,
        user_name: username,
        password: password,
      }),
    });
    if (response.ok) {
      const jsonResponse = await response.json();
      setToken(jsonResponse.data.token);
      console.log("Signup successful");
    } else {
      throw new Error("Failed to signup");
    }
  };

  const login = async (username: string, password: string) => {
    const response = await fetch(`${BASE_URL}/login`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ user_name: username, password }),
    });
    if (response.ok) {
      const jsonResponse = await response.json();
      setToken(jsonResponse.data.token);
      console.log("Login successful");
    } else {
      throw new Error("Failed to login");
    }
  };

  const logout = () => {
    setToken(null);
  };

  const isAuthenticated = !!token;

  return (
    <AuthContext.Provider
      value={{ token, signup, login, logout, isAuthenticated }}
    >
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
};
