import { Exercise, WorkoutExercise } from "@/types";
import { useEffect, useState } from "react";

interface ExerciseRowProps {
  index: number;
  exercise: WorkoutExercise;
  exercises: Exercise[];
  onUpdate: (editedExercise: WorkoutExercise, index: number) => void;
  onRemove: (index: number) => void;
}

const ExerciseRow = ({
  index,
  exercise,
  exercises,
  onUpdate,
  onRemove,
}: ExerciseRowProps) => {
  const [isEditing, setIsEditing] = useState(false);
  const [editedExercise, setEditedExercise] = useState(exercise);

  useEffect(() => {
    onUpdate(editedExercise, index);
  }, [editedExercise]);

  const handleEdit = () => {
    setIsEditing(!isEditing);
  };

  const handleRemove = () => {
    onRemove(index);
  };

  const getExerciseName = (id: string) => {
    const exercise = exercises.find((exercise) => exercise.id === id);
    if (exercise) {
      return exercise.name;
    }
  };

  return (
    <tr key={index} className="border-b border-gray-200">
      {isEditing ? (
        <>
          <td className="px-6 py-3 text-left">
            <select
              value={exercise.id}
              onChange={(e) =>
                setEditedExercise({ ...editedExercise, id: e.target.value })
              }
              required
            >
              <option value="" disabled>
                Select an exercise
              </option>
              {exercises?.map((exercise) => (
                <option key={exercise.id} value={exercise.id}>
                  {getExerciseName(exercise.id)}
                </option>
              ))}
            </select>
          </td>

          <td className="px-6 py-3 text-left">
            <input
              type="number"
              value={editedExercise.sets}
              onChange={(e) =>
                setEditedExercise({ ...editedExercise, sets: +e.target.value })
              }
              className="rounded border px-2 py-1"
            />
          </td>
          <td className="px-6 py-3 text-left">
            <input
              type="number"
              value={editedExercise.reps}
              onChange={(e) =>
                setEditedExercise({ ...editedExercise, reps: +e.target.value })
              }
              className="rounded border px-2 py-1"
            />
          </td>
          <td className="px-6 py-3 text-left">
            <input
              type="number"
              value={editedExercise.load}
              onChange={(e) =>
                setEditedExercise({ ...editedExercise, load: +e.target.value })
              }
              className="rounded border px-2 py-1"
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
          <td>{getExerciseName(exercise.id)}</td>
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
              className="ml-4 text-red-500 hover:text-red-700"
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
