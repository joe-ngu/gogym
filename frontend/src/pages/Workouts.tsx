import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { Workout } from "@/types";
import { fetchWorkouts, createWorkout } from "@/api";
import { useAuth } from "@/auth/AuthContext";

const Workouts = () => {
  const { token } = useAuth();
  const [workouts, setWorkouts] = useState<Workout[]>([]);
  const [newWorkoutName, setNewWorkoutName] = useState("");
  const navigate = useNavigate();

  useEffect(() => {
    const loadWorkouts = async () => {
      const response = await fetchWorkouts(token);
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
      navigate(`/workouts?id=${response.data.id}`);
    } catch (error) {
      console.error("Failed to create workout:", error);
    }
  };

  return (
    <div className="container mx-auto p-4">
      <div className="flex justify-between items-center mb-6">
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
        {workouts.map((workout) => (
          <div
            key={workout.id}
            className="bg-white p-6 rounded-lg shadow-md hover:shadow-lg transition-shadow cursor-pointer"
            onClick={() => navigate(`/workouts?id=${workout.id}`)}
          >
            <h2 className="text-xl font-semibold mb-2">{workout.name}</h2>
            <p className="text-gray-600">
              Exercises: {workout.exercises.length}
            </p>
          </div>
        ))}
      </div>
    </div>
  );
};

export default Workouts;
