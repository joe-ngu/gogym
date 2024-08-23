import { createWorkout, fetchExercises, fetchWorkouts } from "@/api";
import { useAuth } from "@/auth/AuthContext";
import { Exercise, Workout, WorkoutExercise, WorkoutPayload } from "@/types";
import { useEffect, useState } from "react";

const Workouts = () => {
  const [workouts, setWorkouts] = useState<Workout[]>([]);
  const [workoutName, setWorkoutName] = useState("");
  const [workoutExercises, setWorkoutExercises] = useState<WorkoutExercise[]>(
    []
  );
  const [exercises, setExercises] = useState<Exercise[] | null>(null);
  const [exerciseName, setExerciseName] = useState("");
  const [sets, setSets] = useState(0);
  const [reps, setReps] = useState(0);
  const [load, setLoad] = useState(0);
  const { token } = useAuth();

  useEffect(() => {
    const loadWorkoutData = async () => {
      const _workouts = await fetchWorkouts(token);
      const _exercises = await fetchExercises(token);
      setWorkouts(_workouts);
      setExercises(_exercises);
    };
    loadWorkoutData();
  }, [token]);

  const handleAddWorkoutExercise = async (e: React.FormEvent) => {
    e.preventDefault();

    const newWorkoutExercise: WorkoutExercise = {
      name: exerciseName,
      sets,
      reps,
      load,
    };
    setWorkoutExercises([...workoutExercises, newWorkoutExercise]);
    setExerciseName("");
    setSets(0);
    setReps(0);
    setLoad(0);
  };

  const handleRemoveWorkoutExercise = async (e: React.FormEvent) => {};

  const handleSubmitWorkout = async (e: React.FormEvent) => {
    const data = await fetchWorkouts(token);
    setWorkouts(data);
  };

  return (
    <div>
      <h2>Your Workouts</h2>
      <ul>
        {workouts.map((workout) => (
          <li key={workout.id}>
            {workout.exercises.map((exercise, index) => (
              <li key={index}>
                {exercise.name} - {exercise.sets} sets of {exercise.reps} reps
                at {exercise.load} lbs
              </li>
            ))}
          </li>
        ))}
      </ul>
      <form onSubmit={handleAddWorkoutExercise}>
        <select
          value={exerciseName}
          onChange={(e) => setExerciseName(e.target.value)}
          required
        >
          <option value="" disabled>
            Select an exercise
          </option>
          {exercises?.map((exercise) => (
            <option key={exercise.id} value={exercise.name}>
              {exercise.name}
            </option>
          ))}
        </select>
        <input
          type="number"
          placeholder="Sets"
          value={sets}
          onChange={(e) => setSets(Number(e.target.value))}
          required
        />
        <input
          type="number"
          placeholder="Reps"
          value={reps}
          onChange={(e) => setReps(Number(e.target.value))}
          required
        />
        <input
          type="number"
          placeholder="Load"
          value={load}
          onChange={(e) => setLoad(Number(e.target.value))}
          required
        />
        <button type="submit">Add Workout</button>
      </form>
    </div>
  );
};

export default Workouts;
