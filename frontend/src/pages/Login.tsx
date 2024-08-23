import { useAuth } from "@/auth/AuthContext";
import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";

const Login = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const { login } = useAuth();
  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      await login(email, password);
      navigate("/");
    } catch (error) {
      console.error("Error logging in:", error);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="mx-auto space-y-4 p-4">
      <h2 className="text-2xl font-bold">Login</h2>
      <input
        type="email"
        placeholder="Email"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
        className="w-full rounded border p-2"
        required
      />
      <input
        type="password"
        placeholder="Password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
        className="w-full rounded border p-2"
        required
      />
      <button
        type="submit"
        className="w-full rounded bg-blue-500 p-2 text-white"
      >
        Login
      </button>
      <div className="mt-4">
        <span>Don't have an account? </span>
        <Link to="/signup" className="text-blue-500 underline">
          Create one here
        </Link>
      </div>
    </form>
  );
};

export default Login;
