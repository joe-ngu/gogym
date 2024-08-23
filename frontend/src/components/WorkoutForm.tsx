import { useState } from "react";
import Modal from "./Modal"; // Import the modal component
import { Workout } from "@/types";

interface WorkoutFormProps {
  workout?: Workout;
  onSave: (data: any) => void;
}

const WorkoutForm = ({ workout, onSave }: WorkoutFormProps) => {
  const [name, setName] = useState(workout?.name || "");
  const [isOpen, setIsOpen] = useState(false);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onSave({ name });
    setIsOpen(false);
  };

  return (
    <>
      <button
        onClick={() => setIsOpen(true)}
        className="bg-blue-500 text-white p-2 rounded"
      >
        {workout ? "Edit Workout" : "Create Workout"}
      </button>

      <Modal isOpen={isOpen} onClose={() => setIsOpen(false)}>
        <form onSubmit={handleSubmit}>
          <h2 className="text-2xl font-bold mb-4">
            {workout ? "Edit Workout" : "Create Workout"}
          </h2>
          <input
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            placeholder="Workout Name"
            className="w-full p-2 border rounded mb-4"
            required
          />
          <button
            type="submit"
            className="bg-green-500 text-white p-2 rounded w-full"
          >
            Save
          </button>
        </form>
      </Modal>
    </>
  );
};

export default WorkoutForm;
