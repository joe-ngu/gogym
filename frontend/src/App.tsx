import {
  BrowserRouter as Router,
  Routes,
  Route,
  Navigate,
} from "react-router-dom";
import { AuthProvider, useAuth } from "@/auth/AuthContext";

import Navbar from "@/components/Navbar";
import Signup from "@/pages/Signup";
import Login from "@/pages/Login";
import Home from "@/pages/Home";
import Users from "@/pages/Users";
import Workouts from "@/pages/Workouts";
import Exercises from "@/pages/Exercises";
import WorkoutDetails from "./pages/WorkoutDetail";

const ProtectedRoute = ({ element }: { element: React.ReactNode }) => {
  const { isAuthenticated } = useAuth();
  return isAuthenticated ? element : <Navigate to="/login" />;
};

function App() {
  return (
    <Router>
      <AuthProvider>
        <Navbar />
        <Routes>
          <Route>
            <Route path="/" element={<Home />} />
            <Route path="/signup" element={<Signup />} />
            <Route path="/login" element={<Login />} />
            <Route
              path="/users"
              element={<ProtectedRoute element={<Users />} />}
            />
            <Route
              path="/workouts"
              element={<ProtectedRoute element={<Workouts />} />}
            />
            <Route
              path="/workouts/detail/:id"
              element={<ProtectedRoute element={<WorkoutDetails />} />}
            />
            <Route
              path="/exercises"
              element={<ProtectedRoute element={<Exercises />} />}
            />
          </Route>
        </Routes>
      </AuthProvider>
    </Router>
  );
}

export default App;
