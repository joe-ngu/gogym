import { Link } from "react-router-dom";

const Home = () => {
  return (
    <div className="container mx-auto mt-10">
      <h1 className="text-3xl font-bold text-center">Workout Tracker</h1>
      <div className="flex justify-center space-x-4 mt-8">
        <Link to="/users" className="text-lg text-indigo-600 hover:underline">
          Manage Users
        </Link>
        <Link
          to="/workouts"
          className="text-lg text-indigo-600 hover:underline"
        >
          Manage Workouts
        </Link>
        <Link
          to="/exercises"
          className="text-lg text-indigo-600 hover:underline"
        >
          Exercise Library
        </Link>
      </div>
    </div>
  );
};

export default Home;
