import { User } from "./types";

const BASE_URL = "http://localhost:8000";

/* USERS */
export const fetchUsers = async () => {
  const response = await fetch(`${BASE_URL}/users`);
  if (!response.ok) {
    throw new Error("Failed to fetch users");
  }
  return response.json();
};

export const fetchUser = async (id: string) => {
  const response = await fetch(`${BASE_URL}/users/${id}`);
  if (!response.ok) {
    throw new Error("Failed to fetch user");
  }
  return response.json();
};

export const createUser = async (user: User) => {
  const response = await fetch(`${BASE_URL}/users`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(user),
  });
  if (!response.ok) {
    throw new Error("Failed to create user");
  }
  return response.json();
};

export const updateUser = async (id: string, user: User) => {
  const response = await fetch(`${BASE_URL}/users/${id}`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(user),
  });
  if (!response.ok) {
    throw new Error("Failed to update user");
  }
  return response.json();
};

export const deleteUser = async (id: string) => {
  const response = await fetch(`${BASE_URL}/users/${id}`, {
    method: "DELETE",
  });
  if (!response.ok) {
    throw new Error("Failed to delete user");
  }
  return response.json();
};

/* WORKOUTS */
export const fetchWorkouts = async () => {
  const response = await fetch(`${BASE_URL}/workouts`);
  if (!response.ok) {
    throw new Error("Failed to fetch workouts");
  }
  return response.json();
};

export const fetchWorkout = async (id: string) => {
  const response = await fetch(`${BASE_URL}/workouts/${id}`);
  if (!response.ok) {
    throw new Error("Failed to fetch workout");
  }
  return response.json();
};

export const createWorkout = async (workout: any) => {
  const response = await fetch(`${BASE_URL}/workouts`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(workout),
  });
  if (!response.ok) {
    throw new Error("Failed to create workout");
  }
  return response.json();
};

export const updateWorkout = async (id: string, workout: any) => {
  const response = await fetch(`${BASE_URL}/workouts/${id}`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(workout),
  });
  if (!response.ok) {
    throw new Error("Failed to update workout");
  }
  return response.json();
};

export const deleteWorkout = async (id: number) => {
  const response = await fetch(`${BASE_URL}/workouts/${id}`, {
    method: "DELETE",
  });
  if (!response.ok) {
    throw new Error("Failed to delete workout");
  }
  return response.json();
};

/* EXERCISES */
export const fetchExercises = async () => {
  const response = await fetch(`${BASE_URL}/exercises`);
  if (!response.ok) {
    throw new Error("Failed to fetch exercises");
  }
  return response.json();
};

export const fetchExercise = async (id: string) => {
  const response = await fetch(`${BASE_URL}/exercises/${id}`);
  if (!response.ok) {
    throw new Error("Failed to fetch exercise");
  }
  return response.json();
};

export const createExercise = async (exercise: any) => {
  const response = await fetch(`${BASE_URL}/exercises`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(exercise),
  });
  if (!response.ok) {
    throw new Error("Failed to create exercise");
  }
  return response.json();
};

export const updateExercise = async (id: string, exercise: any) => {
  const response = await fetch(`${BASE_URL}/exercises/${id}`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(exercise),
  });
  if (!response.ok) {
    throw new Error("Failed to update exercise");
  }
  return response.json();
};

export const deleteExercise = async (id: string) => {
  const response = await fetch(`${BASE_URL}/exercises/${id}`, {
    method: "DELETE",
  });
  if (!response.ok) {
    throw new Error("Failed to delete exercise");
  }
  return response.json();
};
