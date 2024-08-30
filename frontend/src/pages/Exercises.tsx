import { createExercise, fetchExercises } from "@/api";
import { useAuth } from "@/auth/AuthContext";
import { Exercise } from "@/types";
import { useEffect, useState } from "react";

const Exercises = () => {
  const [exercises, setExercises] = useState<Exercise[]>([]);
  const [name, setName] = useState("");
  const [muscleGroup, setMuscleGroup] = useState("");
  const { token } = useAuth();

  useEffect(() => {
    const loadExercises = async () => {
      const response = await fetchExercises(token);
      setExercises(response.data);
    };
    loadExercises();
  }, []);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const newExercise = { name, muscle_group: muscleGroup };
    const response = await createExercise(token, newExercise);
    if (response.status === 201) {
      console.log("Exercise created successfully");
    }
    setName("");
    setMuscleGroup("");
    // Reload exercises
    const exercises = await fetchExercises(token);
    setExercises(exercises.data);
  };

  return (
    <div className="container mx-auto p-4">
      <h1 className="mb-6 text-3xl font-bold">Exercise Library</h1>
      <div className="lg:grid-col-3 grid gap-4 md:grid-cols-2">
        {exercises.map((exercise) => (
          <div
            key={exercise.id}
            className="rounded-lg bg-white p-6 shadow-md transition-shadow hover:shadow-lg"
          >
            <h2 className="mb-2 text-xl font-semibold">{exercise.name}</h2>
            <p className="text-gray-600">
              Muscle Group: {exercise.muscle_group}
            </p>
          </div>
        ))}
      </div>

      <form onSubmit={handleSubmit} className="space-y-4">
        <div>
          <label className="block text-sm font-medium text-gray-700">
            Name
          </label>
          <input
            type="text"
            name="name"
            placeholder="Exercise Name"
            value={name}
            onChange={(e) => setName(e.target.value)}
            className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
            required
          />
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700">
            Muscle Group
          </label>
          <input
            type="text"
            name="muscle group"
            value={muscleGroup}
            onChange={(e) => setMuscleGroup(e.target.value)}
            className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
            required
          />
        </div>
        <button
          type="submit"
          className="inline-flex justify-center rounded-md border border-transparent bg-indigo-600 px-4 py-2 text-sm font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
        >
          Create Exercise
        </button>
      </form>
    </div>
  );
};

export default Exercises;
