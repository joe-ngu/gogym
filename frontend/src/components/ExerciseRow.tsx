import { Exercise, WorkoutExercise } from "@/types";
import { useEffect, useState } from "react";

interface ExerciseRowProps {
  key: number;
  exercise: WorkoutExercise;
  exercises: Exercise[];
  onUpdate: (editedExercise: WorkoutExercise, index: number) => void;
  onRemove: (index: number) => void;
}

const ExerciseRow = ({
  key,
  exercise,
  exercises,
  onUpdate,
  onRemove,
}: ExerciseRowProps) => {
  const [isEditing, setIsEditing] = useState(false);
  const [editedExercise, setEditedExercise] = useState(exercise);

  useEffect(() => {
    onUpdate(editedExercise, key);
  }, [editedExercise]);

  const handleEdit = () => {
    setIsEditing(!isEditing);
  };

  const handleRemove = () => {
    onRemove(key);
  };

  return (
    <tr key={key} className="border-b border-gray-200">
      {isEditing ? (
        <>
          <td className="py-3 px-6 text-left">
            <select
              value={exercise.name}
              onChange={(e) =>
                setEditedExercise({ ...editedExercise, name: e.target.value })
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

          <td className="py-3 px-6 text-left">
            <input
              type="number"
              value={editedExercise.sets}
              onChange={(e) =>
                setEditedExercise({ ...editedExercise, sets: +e.target.value })
              }
              className="border rounded px-2 py-1"
            />
          </td>
          <td className="py-3 px-6 text-left">
            <input
              type="number"
              value={editedExercise.reps}
              onChange={(e) =>
                setEditedExercise({ ...editedExercise, reps: +e.target.value })
              }
              className="border rounded px-2 py-1"
            />
          </td>
          <td className="py-3 px-6 text-left">
            <input
              type="number"
              value={editedExercise.load}
              onChange={(e) =>
                setEditedExercise({ ...editedExercise, load: +e.target.value })
              }
              className="border rounded px-2 py-1"
            />
          </td>
          <td>
            <button
              className="text-green-500 hover:text-green-700"
              onClick={handleEdit}
            >
              Done
            </button>
          </td>
        </>
      ) : (
        <>
          <td>{editedExercise.sets}</td>
          <td>{editedExercise.reps}</td>
          <td>{editedExercise.load}</td>
          <td>
            <button
              className="text-blue-500 hover:text-blue-700"
              onClick={() => setIsEditing(true)}
            >
              ✏️ Edit
            </button>
            <button
              className="text-red-500 hover:text-red-700 ml-4"
              onClick={handleRemove}
            >
              🗑️ Remove
            </button>
          </td>
        </>
      )}
    </tr>
  );
};

export default ExerciseRow;
