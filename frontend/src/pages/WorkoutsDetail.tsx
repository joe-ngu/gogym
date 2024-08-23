import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { Exercise, Workout, WorkoutExercise } from "@/types";
import { fetchWorkout, updateWorkout, fetchExercises } from "@/api";
import { useAuth } from "@/auth/AuthContext";
import ExerciseRow from "@/components/ExerciseRow";

const WorkoutDetails = () => {
  const { token } = useAuth();
  const { id } = useParams<{ id: string }>();
  const [workout, setWorkout] = useState<Workout | null>(null);
  const [isEditingName, setIsEditingName] = useState(false);
  const [workoutName, setWorkoutName] = useState<string>(workout?.name || "");
  const [newExercise, setNewExercise] = useState<WorkoutExercise>({
    name: "",
    sets: 0,
    reps: 0,
    load: 0,
  });
  const [exercises, setExercises] = useState<Exercise[]>([]);

  const handleNameEdit = () => {
    setIsEditingName(true);
  };

  const handleNameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setWorkoutName(e.target.value);
  };

  const handleNameSave = async () => {
    setIsEditingName(false);
    try {
      if (id && workout) {
        await updateWorkout(token, id, { ...workout, name: workoutName });
        console.log("Workout name updated successfully");
      }
    } catch (error) {
      console.error("Failed to update workout name:", error);
    }
  };

  const handleNameKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === "Enter") {
      handleNameSave();
    }
  };

  useEffect(() => {
    const loadWorkoutData = async () => {
      if (id) {
        const response = await fetchWorkout(token, id);
        setWorkout(response.data);
      }
      const response = await fetchExercises(token);
      setExercises(response.data);
    };
    loadWorkoutData();
  }, [id, workoutName]);

  const handleAddExercise = () => {
    if (workout) {
      const updatedExercises = [...workout.exercises, newExercise];
      setWorkout({ ...workout, exercises: updatedExercises });
      setNewExercise({ name: "", sets: 0, reps: 0, load: 0 });
    }
  };

  const handleUpdateExercise = (
    editedExercise: WorkoutExercise,
    index: number
  ) => {
    if (workout) {
      const updatedExercises = [...workout.exercises];
      updatedExercises[index] = editedExercise;
      setWorkout({ ...workout, exercises: updatedExercises });
    }
  };

  const handleRemoveExercise = (index: number) => {
    if (workout) {
      const updatedExercises = workout.exercises.filter((_, i) => i !== index);
      setWorkout({ ...workout, exercises: updatedExercises });
    }
  };

  const handleSaveWorkout = async () => {
    if (workout) {
      await updateWorkout(token, workout.id, workout);
      alert("Workout updated successfully");
    }
  };

  return (
    <div className="container mx-auto p-4">
      <div className="flex justify-between items-center mb-6">
        {isEditingName ? (
          <input
            type="text"
            value={workoutName}
            onChange={handleNameChange}
            onKeyDown={handleNameKeyDown}
            onBlur={handleNameSave}
            autoFocus
            className="border-b-2 border-gray-300 focus:border-blue-500 outline-none text-3xl font-bold"
          />
        ) : (
          <h1 className="text-3xl font-bold flex items-center cursor-pointer">
            {workout?.name}{" "}
            <span
              className="ml-2 text-gray-500 hover:text-blue-500"
              onClick={handleNameEdit}
              style={{ cursor: "pointer" }}
            >
              ✏️
            </span>
          </h1>
        )}
      </div>
      <h2>Workout Details</h2>
      <table className="min-w-full bg-white shadow-md rounded-lg overflow-hidden">
        <thead>
          <tr className="bg-gray-200 text-gray-600 uppercase text-sm leading-normal">
            <th className="py-3 px-6 text-left">Exercise</th>
            <th className="py-3 px-6 text-left">Sets</th>
            <th className="py-3 px-6 text-left">Reps</th>
            <th className="py-3 px-6 text-left">Weight</th>
            <th className="py-3 px-6 text-left">Actions</th>
          </tr>
        </thead>
        <tbody>
          {workout?.exercises.map(
            (exercise: WorkoutExercise, index: number) => (
              <ExerciseRow
                key={index}
                exercise={exercise}
                exercises={exercises}
                onUpdate={handleUpdateExercise}
                onRemove={handleRemoveExercise}
              />
            )
          )}
          <tr>
            <td>
              <select
                value={newExercise.name}
                onChange={(e) =>
                  setNewExercise({ ...newExercise, name: e.target.value })
                }
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
            </td>
            <td>
              <input
                type="number"
                placeholder="Sets"
                value={newExercise.sets}
                onChange={(e) =>
                  setNewExercise({ ...newExercise, sets: +e.target.value })
                }
              />
            </td>
            <td>
              <input
                type="number"
                placeholder="Reps"
                value={newExercise.reps}
                onChange={(e) =>
                  setNewExercise({ ...newExercise, reps: +e.target.value })
                }
              />
            </td>
            <td>
              <input
                type="number"
                placeholder="Weight"
                value={newExercise.load}
                onChange={(e) =>
                  setNewExercise({ ...newExercise, load: +e.target.value })
                }
              />
            </td>
            <td>
              <button
                className="bg-green-500 text-white hover:bg-green-600"
                onClick={handleAddExercise}
              >
                + Add Exercise
              </button>
            </td>
          </tr>
        </tbody>
      </table>
      <button onClick={handleSaveWorkout}>Save Workout</button>
    </div>
  );
};

export default WorkoutDetails;
