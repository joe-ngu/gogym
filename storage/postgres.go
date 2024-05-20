package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/joe-ngu/gogym/utils"
)

type PostgresDB struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresDB, error) {
	user, err := utils.Getenv("DB_USER")
	if err != nil {
		return nil, errors.New("failed to create PostgresStore: " + err.Error())
	}
	dbname, err := utils.Getenv("DB_NAME")
	if err != nil {
		return nil, errors.New("failed to create PostgresStore: " + err.Error())
	}
	dbpwd, err := utils.Getenv("DB_PWD")
	if err != nil {
		return nil, errors.New("failed to create PostgresStore: " + err.Error())
	}

	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", user, dbname, dbpwd)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresDB{
		db: db,
	}, nil
}

func (s *PostgresDB) Init() error {
	if err := s.createExerciseTables(); err != nil {
		return err
	}
	if err := s.createWorkoutTables(); err != nil {
		return err
	}
	if err := s.createUserTables(); err != nil {
		return err
	}
	return nil
}

func (s *PostgresDB) createExerciseTables() error {
	createMuscleGroups := `
  CREATE TABLE IF NOT EXISTS muscle_groups (
    id SERIAL PRIMARY KEY,
    name TEXT,
  )`

	createExercise := `
  CREATE TABLE IF NOT EXISTS exercise (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT,
    muscle_group INTEGER REFERENCES muscle_groups(id),
  )`

	if _, err := s.db.Exec(createMuscleGroups); err != nil {
		return err
	}
	if _, err := s.db.Exec(createExercise); err != nil {
		return err
	}
	return nil
}

func (s *PostgresDB) CreateExercise(exercise *Exercise) error {
	createExercise := `
    INSERT INTO exercise
    (name, muscle_group) 
    values ($1, $2)
  `

	if _, err := s.db.Exec(
		createExercise,
		exercise.Name,
		exercise.MuscleGroup,
	); err != nil {
		return err
	}
	return nil
}

func (s *PostgresDB) GetExercises() ([]*Exercise, error) {
	getExercises := `
    SELECT * FROM exercise;
  `
	rows, err := s.db.Query(getExercises)
	if err != nil {
		return nil, err
	}

	exercises := []*Exercise{}
	for rows.Next() {
		exercise, err := loadExercises(rows)
		if err != nil {
			return nil, err
		}
		exercises = append(exercises, exercise)
	}
  if err := rows.Err(); err != nil {
    return nil, err
  }

	return exercises, nil
}

func (s *PostgresDB) GetExercise(name string) (*Exercise, error) {
	getExercise := `
    SELECT * FROM exercise
    WHERE name = ($1)
  `
	rows, err := s.db.Query(getExercise, name)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return loadExercises(rows)
	}
	return nil, fmt.Errorf("exercise %s not found", name)
}

func (s *PostgresDB) UpdateExercise(exercise *Exercise) (*Exercise, error) {
	updateExercise := `
    UPDATE exercise
    SET name = ($1), muscle_group = ($2)
    WHERE id = ($3);
  `
	if _, err := s.db.Exec(
		updateExercise,
		exercise.Name,
		exercise.MuscleGroup,
		exercise.ID,
	); err != nil {
		return nil, err
	}
	return s.GetExercise(exercise.ID)
}

func (s *PostgresDB) DeleteExercise(id uuid.UUID) error {
	deleteExercise := `
    DELETE exercise
    WHERE id = ($1)
  `
	if _, err := s.db.Exec(deleteExercise, id); err != nil {
		return err
	}
	return nil
}

func loadExercises(rows *sql.Rows) (*Exercise, error) {
	exercise := new(Exercise)
	err := rows.Scan(
		&exercise.ID,
		&exercise.Name,
		&exercise.MuscleGroup,
	)
	return exercise, err
}

func (s *PostgresDB) createWorkoutTables() error {
	createWorkout := `
  CREATE TABLE IF NOT EXISTS workout (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT,
    created_at TIMESTAMP,
    date TIMESTAMP, 
  )`

	createWorkoutExercise := `
  CREATE TABLE IF NOT EXISTS workout_exercise (
    workout_id UUID REFERENCES workout(id)
    exercise_id UUID REFERENCES exercise(id)
    set INTEGER,
    reps INTEGER,
    load NUMERIC(6, 2),
  )`

	if _, err := s.db.Exec(createWorkout); err != nil {
		return err
	}
	if _, err := s.db.Exec(createWorkoutExercise); err != nil {
		return err
	}
	return nil
}

func (s *PostgresDB) CreateWorkout(workout *Workout) error {
	createWorkout := `
    INSERT INTO workout
    (name, created_at, date)
    VALUES ($1, $2, $3);
  `
	addWorkoutExercises := `
    INSERT INTO workout_exercise
    (workout_id, exercise_id)
    VALUES  
  `
	var idArgs []interface{}
	for _, exerciseID := range workout.Exercises {
		addWorkoutExercises += "(?, ?),"
		idArgs = append(idArgs, workout.ID, exerciseID)
	}
	addWorkoutExercises = strings.TrimSuffix(addWorkoutExercises, ",")

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(
		createWorkout,
		workout.Name,
		workout.CreatedAt,
		workout.Date,
	); err != nil {
		tx.Rollback()
		return err
	}

	if _, err := tx.Exec(addWorkoutExercises, idArgs...); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (s *PostgresDB) GetWorkouts() ([]*Workout, error) {
	getWorkouts := `
    SELECT * FROM workout;
  `
	rows, err := s.db.Query(getWorkouts)
	if err != nil {
		return nil, err
	}

	workouts := []*Workout{}
	for rows.Next() {
		workout, err := loadWorkouts(rows)
		if err != nil {
			return nil, err
		}
		workouts = append(workouts, workout)
	}
  if err := rows.Err(); err != nil {
    return nil, err
  }

	return workouts, nil
}

func (s *PostgresDB) GetWorkout(id uuid.UUID) (*Workout, error) {
	getWorkout := `
    SELECT * FROM workout
    WHERE ID = ($1);
  `
	rows, err := s.db.Query(getWorkout, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return loadWorkouts(rows)
	}

	return nil, fmt.Errorf("Workout %v not found", id)
}

func (s *PostgresDB) UpdateWorkout(workout *Workout) (*Workout, error) {
	updateWorkout := `
    UPDATE workout
    SET name = ($1), created_at = ($3), date = ($3)
    WHERE id = ($4);
  `
	deleteOldExercises := `
    DELETE FROM workout_exercise
    WHERE workout_id = ($1);
  `
	addNewExercises := `
    INSERT INTO workout_exercise
    (workout_id, exercise_id)
    VALUES , 
  `
	var idArgs []interface{}
	for _, exerciseID := range workout.Exercises {
		addNewExercises += "(?, ?),"
		idArgs = append(idArgs, workout.ID, exerciseID)
	}
	addNewExercises = strings.TrimSuffix(addNewExercises, ",")

	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}

	if _, err := tx.Exec(
		updateWorkout,
		workout.Name,
		workout.CreatedAt,
		workout.Date,
		workout.ID,
	); err != nil {
		tx.Rollback()
		return nil, err
	}

	if _, err := tx.Exec(deleteOldExercises, workout.ID); err != nil {
		tx.Rollback()
		return nil, err
	}

	if _, err := tx.Exec(addNewExercises); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}

	return s.GetWorkout(workout.ID)
}

func (s *PostgresDB) DeleteWorkout(id uuid.UUID) error {
	deleteWorkout := `
    DELETE workout
    WHERE id = ($1);
  `
	deleteWorkoutExercises := `
    DELETE workout_exercise
    WHERE workout_id = ($1);
  `
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	if _, err = tx.Exec(deleteWorkout, id); err != nil {
		tx.Rollback()
		return err
	}

	if _, err = tx.Exec(deleteWorkoutExercises, id); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func loadWorkouts(rows *sql.Rows) (*Workout, error) {
	workout := new(Workout)
	err := rows.Scan(
		&workout.ID,
		&workout.Name,
		&workout.Exercises,
		&workout.CreatedAt,
		&workout.Date,
	)
	return workout, err
}

func (s *PostgresDB) createUserTables() error {
	createUser := `
  CREATE TABLE IF NOT EXISTS user (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    first_name TEXT,
    last_name TEXT,
    user_name TEXT,
    encrypted_password TEXT,
    created_at TIMESTAMP,
  );`

	createUserWorkout := `
  CREATE TABLE IF NOT EXISTS user_workout (
    user_id UUID REFERENCES User(id),
    workout_id UUID FOREIGN_KEY REFERENCES Workout(id)
  );`

	if _, err := s.db.Exec(createUser); err != nil {
		return err
	}
	if _, err := s.db.Exec(createUserWorkout); err != nil {
		return err
	}
	return nil
}

func (s *PostgresDB) CreateUser(user *User) error {
	createUser := `
    INSERT INTO user
    (first_name, last_name, user_name, encrypted_password, created_at)
    VALUES ($1, $2, $3, $4, $5);
  `
	if _, err := s.db.Exec(
		createUser,
		user.FirstName,
		user.LastName,
		user.EncryptedPassword,
		user.CreatedAt,
	); err != nil {
		return err
	}

	return nil
}

func (s *PostgresDB) GetUsers() ([]*User, error) {
	getUsers := `
    SELECT * FROM user;
  `
	rows, err := s.db.Query(getUsers)
	if err != nil {
		return nil, err
	}

	users := []*User{}
	for rows.Next() {
		user, err := s.loadUsers(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
  if err := rows.Err(); err != nil {
    return nil, err
  }
	return users, nil
}

func (s *PostgresDB) GetUser(id uuid.UUID) (*User, error) {
	getUser := `
    SELECT * FROM user
    WHERE id = ($1);
  `
	rows, err := s.db.Query(getUser, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return s.loadUsers(rows)
	}
  if err := rows.Err(); err != nil {
    return nil, err
  }
	return nil, fmt.Errorf("User %v not found.", id)
}

func (s *PostgresDB) GetUserWorkoutIDs(id uuid.UUID) ([], error) {
  getUserWorkoutsIDs := `
    SELECT workout_id FROM user_workout
    WHERE user_id = ($1)
  `
  rows, err := db.Query(getUserWorkouts)
  if err != nil {
    return err
  }
  
  var workoutIDs []uuid.UUID
  for rows.Next() {
    var id uuid.UUID
    if err != rows.Scan(&id); err != nil {
      return nil, err
    }
    workoutIDs = append(workoutIDs, id)
  }
  if err := rows.Err(); err != nil {
    return nil, err
  }
  
  return workoutIDs, nil
}

func (s *PostgresDB) GetUserWorkouts(id uuid.UUID) ([]*Workout, error) {
  var workoutIDs []uuid.UUID = s.GetUserWorkoutIDs(id)
  
  var workouts := []*Workout{}
  for _, id := range(workoutIDs) {
    workout, err = s.GetWorkout(id)
    if err != nil {
      return nil, err
    }
    workouts = append(workouts, workout) 
  }

  return workouts, nil 
}

func (s *PostgresDB) UpdateUser(user *User) (*User, error) {
  updateUser := `
    UPDATE user
    SET first_name = ($1), 
        last_name = ($2), 
        user_name = ($3), 
        encrypted_password = ($4), 
        created_at = ($5)
    WHERE id = ($6);
  `
  if _, err := s.db.Exec(
    updateUser,
    user.FirstName,
    user.LastName,
    user.UserName,
    user.EncryptedPassword,
    user.CreatedAt,
  ); err != nil {
    return nil, err
  }
  return s.GetUser(user.ID)
}

func (s *PostgresDB) DeleteUser(id uuid.UUID) error {
  deleteUser := `
    DELETE user
    WHERE id = ($1);
  `
  for id := range() {
    if err := s.DeleteWorkout(id) {
      return err
    }
  }
  
  if _, err := s.db.Exec(deleteUser, id); err != nil {
    return err
  }

  return nil
}

func (s *PostgresDB) loadUsers(rows *sql.Rows) (*User, error) {
	user := new(User)
	err := rows.Scan(
		&user.FirstName,
		&user.LastName,
		&user.UserName,
		&user.EncryptedPassword,
		&user.CreatedAt,
	)
	return user, err
}
