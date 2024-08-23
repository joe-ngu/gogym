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
        className="text-3xl p-4 flex justify-center items-center cursor-pointer"
        onClick={handleLogoClick}
      >
        Workout Tracker
      </h1>
      {isAuthenticated && <button onClick={handleLogout}>Logout</button>}
    </nav>
  );
};

export default Navbar;
