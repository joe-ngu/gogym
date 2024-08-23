import { createExercise, updateExercise } from "@/api";
import { useState } from "react";

interface ExerciseFormProps {
  exerciseId?: string;
  initialData?: any;
  onSave: () => void;
}

const ExerciseForm = ({
  exerciseId,
  initialData,
  onSave,
}: ExerciseFormProps) => {
  const [formData, setFormData] = useState(
    initialData || {
      name: "",
      description: "",
    }
  );

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    try {
      exerciseId
        ? await updateExercise(exerciseId, formData)
        : await createExercise(formData);
      onSave();
    } catch (error) {
      console.error("Error saving exercise:", error);
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
        {exerciseId ? "Update" : "Create"} Exercise
      </button>
    </form>
  );
};

export default ExerciseForm;
