import { BrowserRouter, Routes, Route } from "react-router-dom";

import Layout from "@/pages/Layout";
import Home from "@/pages/Home";
import Users from "@/pages/Users";
import Workouts from "@/pages/Workouts";
import Exercises from "@/pages/Exercises";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route index element={<Home />} />
          <Route path="users" element={<Users />} />
          <Route path="workouts" element={<Workouts />} />
          <Route path="exercises" element={<Exercises />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
