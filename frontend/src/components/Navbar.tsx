import { useAuth } from "@/auth/AuthContext";
import { useNavigate } from "react-router-dom";

const Navbar = () => {
  const { isAuthenticated, logout } = useAuth();
  const navigate = useNavigate();

  const handleLogout = () => {
    logout();
    navigate("/login");
  };

  const handleLogoClick = () => {
    navigate("/");
  };

  return (
    <nav>
      <h1
        className="flex cursor-pointer items-center justify-center p-4 text-3xl"
        onClick={handleLogoClick}
      >
        Workout Tracker
      </h1>
      {isAuthenticated && <button onClick={handleLogout}>Logout</button>}
    </nav>
  );
};

export default Navbar;
