export type AuthToken = string | null;

export interface User {
  id: string;
  first_name: string;
  last_name: string;
  user_name: string;
}

export interface UserPayload {
  first_name: string;
  last_name: string;
  user_name: string;
  password: string;
}

export interface Exercise {
  id: string;
  name: string;
  muscle_group: string;
}

export interface ExercisePayload {
  name: string;
  muscle_group: string;
}

export interface WorkoutExercise {
  id: string;
  sets: number;
  reps: number;
  load: number;
}

export interface Workout {
  id: string;
  user_id: string;
  name: string;
  date: string;
  created_at: string;
}

export interface WorkoutDetail extends Workout {
  exercises: WorkoutExercise[];
}

export interface WorkoutPayload {
  name: string;
  exercises: WorkoutExercise[];
}
