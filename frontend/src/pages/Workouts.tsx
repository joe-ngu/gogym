import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { Workout } from "@/types";
import { fetchWorkouts, createWorkout } from "@/api";
import { useAuth } from "@/auth/AuthContext";

const Workouts = () => {
  // TO-DO: Filtering feature (maybe by name?)
  // const location = useLocation();
  // const queryParams = new URLSearchParams(location.search);
  // const workoutID = queryParams.get("name");
  const { token } = useAuth();
  const [workouts, setWorkouts] = useState<Workout[]>([]);
  const [newWorkoutName, setNewWorkoutName] = useState("");
  const navigate = useNavigate();

  useEffect(() => {
    const loadWorkouts = async () => {
      const response = await fetchWorkouts(token);
      console.log(response);
      setWorkouts(response.data);
    };
    loadWorkouts();
  }, []);

  const handleCreateWorkout = async () => {
    try {
      const response = await createWorkout(token, {
        name: newWorkoutName,
        exercises: [],
      });
      navigate(`/workouts/detail/${response.data.id}`);
    } catch (error) {
      console.error("Failed to create workout:", error);
    }
  };

  return (
    <div className="container mx-auto p-4">
      <div className="mb-6 flex items-center justify-between">
        <h1>Your Workouts</h1>
        <input
          type="text"
          placeholder="New Workout Name"
          value={newWorkoutName}
          onChange={(e) => setNewWorkoutName(e.target.value)}
        />
        <button onClick={handleCreateWorkout}>+ Create New Workout</button>
      </div>
      <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
        {workouts && workouts.length > 0 ? (
          workouts.map((workout) => {
            const date = new Date(workout.date).toLocaleString();
            const createdAt = new Date(workout.created_at).toLocaleString();

            return (
              <div
                key={workout.id}
                className="cursor-pointer rounded-lg bg-white p-6 shadow-md transition-shadow hover:shadow-lg"
                onClick={() => navigate(`/workouts/detail/${workout.id}`)}
              >
                <h2 className="mb-2 text-xl font-semibold">{workout.name}</h2>
                <p className="text-gray-600">{date}</p>
                <p className="text-gray-600">{createdAt}</p>
              </div>
            );
          })
        ) : (
          <p>No workouts available.</p>
        )}
      </div>
    </div>
  );
};

export default Workouts;
