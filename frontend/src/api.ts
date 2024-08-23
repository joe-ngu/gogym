import {
  AuthToken,
  Exercise,
  ExercisePayload,
  User,
  Workout,
  WorkoutPayload,
} from "./types";

export const BASE_URL = "http://localhost:8000";

type ApiResponse<T> = {
  data: T;
  status: number;
  headers: Headers;
};

/* USERS */
export const fetchUsers = async (
  token: AuthToken,
): Promise<ApiResponse<User[]>> => {
  const response = await fetch(`${BASE_URL}/users`, {
    method: "GET",
    headers: {
      Authorization: `Bearer ${token}`,
      "Content-Type": "application/json",
    },
  });
  if (!response.ok) {
    throw new Error("Failed to fetch users");
  }
  return response.json();
};

export const fetchUser = async (
  token: AuthToken,
): Promise<ApiResponse<User>> => {
  const response = await fetch(`${BASE_URL}/user`, {
    method: "GET",
    headers: {
      Authorization: `Bearer ${token}`,
      "Content-Type": "application/json",
    },
  });
  if (!response.ok) {
    throw new Error("Failed to fetch user");
  }
  return response.json();
};

export const createUser = async (
  token: AuthToken,
  user: User,
): Promise<ApiResponse<User>> => {
  const response = await fetch(`${BASE_URL}/user`, {
    method: "POST",
    headers: {
      Authorization: `Bearer ${token}`,
      "Content-Type": "application/json",
    },
    body: JSON.stringify(user),
  });
  if (!response.ok) {
    throw new Error("Failed to create user");
  }
  return response.json();
};

export const updateUser = async (
  token: AuthToken,
  user: User,
): Promise<ApiResponse<User>> => {
  const response = await fetch(`${BASE_URL}/user`, {
    method: "PUT",
    headers: {
      Authorization: `Bearer ${token}`,
      "Content-Type": "application/json",
    },
    body: JSON.stringify(user),
  });
  if (!response.ok) {
    throw new Error("Failed to update user");
  }
  return response.json();
};

export const deleteUser = async (
  token: AuthToken,
): Promise<ApiResponse<null>> => {
  const response = await fetch(`${BASE_URL}/user`, {
    method: "DELETE",
    headers: {
      Authorization: `Bearer ${token}`,
      "Content-Type": "application/json",
    },
  });
  if (!response.ok) {
    throw new Error("Failed to delete user");
  }
  return response.json();
};

/* WORKOUTS */
export const fetchWorkouts = async (
  token: AuthToken,
): Promise<ApiResponse<Workout[]>> => {
  const response = await fetch(`${BASE_URL}/workout`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
    },
  });

  if (!response.ok) {
    throw new Error("Failed to fetch workouts");
  }
  return response.json();
};

export const fetchWorkout = async (
  token: AuthToken,
  id: string,
): Promise<ApiResponse<Workout>> => {
  const response = await fetch(`${BASE_URL}/workout?id=${id}`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
    },
  });
  if (!response.ok) {
    throw new Error("Failed to fetch workout");
  }
  return response.json();
};

export const createWorkout = async (
  token: AuthToken,
  workout: WorkoutPayload,
): Promise<ApiResponse<Workout>> => {
  const response = await fetch(`${BASE_URL}/workout`, {
    method: "POST",
    headers: {
      Authorization: `Bearer ${token}`,
      "Content-Type": "application/json",
    },
    body: JSON.stringify(workout),
  });
  if (!response.ok) {
    throw new Error("Failed to create workout");
  }
  return response.json();
};

export const updateWorkout = async (
  token: AuthToken,
  id: string,
  workout: WorkoutPayload,
): Promise<ApiResponse<Workout>> => {
  const response = await fetch(`${BASE_URL}/workout?id=${id}`, {
    method: "PUT",
    headers: {
      Authorization: `Bearer ${token}`,
      "Content-Type": "application/json",
    },
    body: JSON.stringify(workout),
  });
  if (!response.ok) {
    throw new Error("Failed to update workout");
  }
  return response.json();
};

export const deleteWorkout = async (
  token: AuthToken,
  id: string,
): Promise<ApiResponse<null>> => {
  const response = await fetch(`${BASE_URL}/workouts/${id}`, {
    method: "DELETE",
    headers: {
      Authorization: `Bearer ${token}`,
      "Content-Type": "application/json",
    },
  });
  if (!response.ok) {
    throw new Error("Failed to delete workout");
  }
  return response.json();
};

/* EXERCISES */
export const fetchExercises = async (
  token: AuthToken,
): Promise<ApiResponse<Exercise[]>> => {
  const response = await fetch(`${BASE_URL}/exercises`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
    },
  });
  if (!response.ok) {
    throw new Error("Failed to fetch exercises");
  }
  return response.json();
};

export const fetchExercise = async (
  token: AuthToken,
  name: string,
): Promise<ApiResponse<Exercise>> => {
  const response = await fetch(`${BASE_URL}/exercise?name=${name}`, {
    method: "GET",
    headers: {
      Authorization: `Bearer ${token}`,
      "Content-Type": "application/json",
    },
  });
  if (!response.ok) {
    throw new Error("Failed to fetch exercise");
  }
  return response.json();
};

export const createExercise = async (
  token: AuthToken,
  exercise: ExercisePayload,
): Promise<ApiResponse<Exercise>> => {
  const response = await fetch(`${BASE_URL}/exercise`, {
    method: "POST",
    headers: {
      Authorization: `Bearer ${token}`,
      "Content-Type": "application/json",
    },
    body: JSON.stringify(exercise),
  });
  if (!response.ok) {
    throw new Error("Failed to create exercise");
  }
  return response.json();
};

export const updateExercise = async (
  token: AuthToken,
  name: string,
  exercise: ExercisePayload,
): Promise<ApiResponse<Exercise>> => {
  const response = await fetch(`${BASE_URL}/exercise?name=${name}`, {
    method: "PUT",
    headers: {
      Authorization: `Bearer ${token}`,
      "Content-Type": "application/json",
    },
    body: JSON.stringify(exercise),
  });
  if (!response.ok) {
    throw new Error("Failed to update exercise");
  }
  return response.json();
};

export const deleteExercise = async (
  token: AuthToken,
  name: string,
): Promise<ApiResponse<null>> => {
  const response = await fetch(`${BASE_URL}/exercise?name=${name}`, {
    method: "DELETE",
    headers: {
      Authorization: `Bearer ${token}`,
      "Content-Type": "application/json",
    },
  });
  if (!response.ok) {
    throw new Error("Failed to delete exercise");
  }
  return response.json();
};
