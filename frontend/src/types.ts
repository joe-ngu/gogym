export type AuthToken = string | null;

export interface User {
  id: string;
  name: string;
  email: string;
  password?: string;
}

export interface UserPayload {
  name: string;
  email: string;
  password: string;
}

export interface Exercise {
  id: string;
  name: string;
  muscleGroup: string;
}

export interface ExercisePayload {
  name: string;
  muscleGroup: string;
}

export interface WorkoutExercise {
  name: string;
  sets: number;
  reps: number;
  load: number;
}

export interface Workout {
  id: string;
  name: string;
  exercises: WorkoutExercise[];
  userId: string;
}

export interface WorkoutPayload {
  name: string;
  exercises: WorkoutExercise[];
}
