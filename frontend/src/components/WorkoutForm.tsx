import { createWorkout, updateWorkout } from "@/api";
import { ChangeEvent, useState } from "react";

interface WorkoutFormProps {
  workoutId?: string;
  initialData?: any;
  onSave: () => void;
}

const WorkoutForm = ({ workoutId, initialData, onSave }: WorkoutFormProps) => {
  const [formData, setFormData] = useState(
    initialData || {
      name: "",
      description: "",
    }
  );

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    try {
      workoutId
        ? await updateWorkout(workoutId, formData)
        : await createWorkout(formData);
      onSave();
    } catch (error) {
      console.error("Error saving workout:", error);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="space-y-4">
      <div>
        <label className="block text-sm font-medium text-gray-700">Name</label>
        <input
          type="text"
          name="name"
          value={formData.name}
          onChange={handleChange}
          className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
        />
      </div>
      <div>
        <label className="block text-sm font-medium text-gray-700">
          Description
        </label>
        <input
          type="text"
          name="description"
          value={formData.description}
          onChange={handleChange}
          className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
        />
      </div>
      <button
        type="submit"
        className="inline-flex justify-center rounded-md border border-transparent bg-indigo-600 py-2 px-4 text-sm font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
      >
        {workoutId ? "Update" : "Create"} Workout
      </button>
    </form>
  );
};

export default WorkoutForm;
