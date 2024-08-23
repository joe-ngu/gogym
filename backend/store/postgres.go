package store

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
	_ "github.com/lib/pq"

	"github.com/joe-ngu/gogym/types"
	"github.com/joe-ngu/gogym/utils"
)

type PostgresDB struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresDB, error) {
	host, err := utils.Getenv("POSTGRES_HOST")
	if err != nil {
		return nil, errors.New("failed to create PostgresStore: " + err.Error())
	}
	port, err := utils.Getenv("POSTGRES_PORT")
	if err != nil {
		return nil, errors.New("failed to create PostgresStore: " + err.Error())
	}
	user, err := utils.Getenv("POSTGRES_USER")
	if err != nil {
		return nil, errors.New("failed to create PostgresStore: " + err.Error())
	}
	dbname, err := utils.Getenv("POSTGRES_DB")
	if err != nil {
		return nil, errors.New("failed to create PostgresStore: " + err.Error())
	}
	dbpwd, err := utils.Getenv("POSTGRES_PASSWORD")
	if err != nil {
		return nil, errors.New("failed to create PostgresStore: " + err.Error())
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, dbpwd)
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
	log.Println("Created Exercise Tables")

	if err := s.createWorkoutTables(); err != nil {
		return err
	}
	log.Println("Created Workout Tables")

	if err := s.createUserTables(); err != nil {
		return err
	}
	log.Println("Created User Tables")

	return nil
}

func (s *PostgresDB) createExerciseTables() error {
	createExercise := `
  		CREATE TABLE IF NOT EXISTS exercise (
			id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    		name TEXT UNIQUE NOT NULL,
    		muscle_group TEXT
		);
	`
	if _, err := s.db.Exec(createExercise); err != nil {
		return err
	}
	return nil
}

func (s *PostgresDB) CreateExercise(exercise *types.Exercise) error {
	createExercise := `
    	INSERT INTO exercise (
			name, 
			muscle_group
		) 
    	VALUES ($1, $2);
	`

	exists, _ := s.GetExercise(exercise.Name)
	if exists != nil {
		return errors.New("exercise with same name already exists")
	}

	if _, err := s.db.Exec(
		createExercise,
		exercise.Name,
		exercise.MuscleGroup.String(),
	); err != nil {
		return err
	}
	return nil
}

func (s *PostgresDB) GetExercises() ([]*types.Exercise, error) {
	getExercises := `
    	SELECT * FROM exercise;
  	`
	rows, err := s.db.Query(getExercises)
	if err != nil {
		return nil, err
	}

	exercises := []*types.Exercise{}
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

func (s *PostgresDB) GetExercise(name string) (*types.Exercise, error) {
	getExercise := `
    	SELECT * FROM exercise
    	WHERE name = $1
		LIMIT 1;
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

func (s *PostgresDB) UpdateExercise(name string, exercise *types.Exercise) (*types.Exercise, error) {
	updateExercise := `
    	UPDATE exercise
    	SET 
			name = $1, 
			muscle_group = $2
    	WHERE 
			name = $3;
  	`
	if _, err := s.db.Exec(
		updateExercise,
		exercise.Name,
		exercise.MuscleGroup.String(),
		name,
	); err != nil {
		return nil, err
	}
	return s.GetExercise(exercise.Name)
}

func (s *PostgresDB) DeleteExercise(name string) error {
	deleteExercise := `
    	DELETE FROM exercise
    	WHERE name = $1;
  	`
	if _, err := s.db.Exec(deleteExercise, name); err != nil {
		return err
	}
	return nil
}

func loadExercises(rows *sql.Rows) (*types.Exercise, error) {
	exercise := new(types.Exercise)
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
			user_id UUID REFERENCES "user"(id),
    		name TEXT,
    		created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    		date TIMESTAMP 
  		);
	`
	createWorkoutExercise := `
  		CREATE TABLE IF NOT EXISTS workout_exercise (
    		workout_id UUID REFERENCES workout(id),
    		exercise_id UUID REFERENCES exercise(id),
    		sets INTEGER,
    		reps INTEGER,
    		load NUMERIC(6, 2)
		);
	`

	if _, err := s.db.Exec(createWorkout); err != nil {
		return err
	}
	if _, err := s.db.Exec(createWorkoutExercise); err != nil {
		return err
	}
	return nil
}

func (s *PostgresDB) CreateWorkout(userID uuid.UUID, workout *types.Workout) error {
	createWorkout := `
    	INSERT INTO workout (
			user_id,
			name, 
			date
		)
    	VALUES ($1, $2, $3);
  	`

	addWorkoutExercises := `
    	INSERT INTO workout_exercise (
			workout_id, 
			exercise_id,
			sets,
			reps,
			load
		)
    	VALUES  
  	`
	var placeholders []string
	var workoutExercises []interface{}
	for i, exercise := range workout.Exercises {
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)",
			(i*5)+1,  // Workout ID
			(i*5)+2,  // Exercise Name
			(i*5)+3,  // Sets
			(i*5)+4,  // Reps
			(i*5)+5)) // Load
		workoutExercises = append(workoutExercises, workout.ID, exercise.ID, exercise.Sets, exercise.Reps, exercise.Load)
	}
	addWorkoutExercises += strings.Join(placeholders, ",") + ";"

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(
		createWorkout,
		workout.Name,
		workout.Date,
	); err != nil {
		tx.Rollback()
		return err
	}

	if _, err := tx.Exec(addWorkoutExercises, workoutExercises...); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (s *PostgresDB) GetWorkouts(userID uuid.UUID) ([]*types.Workout, error) {
	getWorkouts := `
    	SELECT * FROM workout
		WHERE user_id = $1;
  	`

	rows, err := s.db.Query(getWorkouts, userID)
	if err != nil {
		return nil, err
	}

	workouts := []*types.Workout{}
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

func (s *PostgresDB) GetWorkout(userID uuid.UUID, id uuid.UUID) (*types.Workout, error) {
	getWorkout := `
    	SELECT * FROM workout
    	WHERE user_id = $1 AND ID = $2;
  	`
	rows, err := s.db.Query(getWorkout, userID, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return loadWorkouts(rows)
	}

	return nil, fmt.Errorf("workout %v not found", id)
}

func (s *PostgresDB) UpdateWorkout(userID uuid.UUID, workout *types.Workout) (*types.Workout, error) {
	updateWorkout := `
    	UPDATE workout
    	SET 
			name = $1, 
			created_at = $2, 
    	WHERE
			user_id = $3 AND
			id = $4;
	`

	deleteOldExercises := `
    	DELETE FROM workout_exercise
    	WHERE user_id = $1 AND workout_id = $2;
	`

	addNewExercises := `
    	INSERT INTO workout_exercise (
			workout_id, 
			exercise_id,
			sets,
			reps,
			load
		)
    	VALUES , 
  	`

	var placeholders []string
	var workoutExercises []interface{}
	for i, exercise := range workout.Exercises {
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)",
			(i*5)+1,  // Workout ID
			(i*5)+2,  // Exercise Name
			(i*5)+3,  // Sets
			(i*5)+4,  // Reps
			(i*5)+5)) // Load
		workoutExercises = append(workoutExercises, workout.ID, exercise.ID, exercise.Sets, exercise.Reps, exercise.Load)
	}
	addNewExercises = strings.Join(placeholders, ",") + ";"

	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}

	if _, err := tx.Exec(
		updateWorkout,
		workout.Name,
		workout.CreatedAt,
		workout.ID,
	); err != nil {
		tx.Rollback()
		return nil, err
	}

	if _, err := tx.Exec(deleteOldExercises, userID, workout.ID); err != nil {
		tx.Rollback()
		return nil, err
	}

	if _, err := tx.Exec(addNewExercises, workoutExercises...); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}

	return s.GetWorkout(userID, workout.ID)
}

func (s *PostgresDB) DeleteWorkout(userID uuid.UUID, id uuid.UUID) error {
	deleteWorkout := `
    	DELETE FROM workout
    	WHERE user_id = $1 AND id = $2;
  	`
	deleteWorkoutExercises := `
    	DELETE FROM workout_exercise
    	WHERE workout_id = $1;
  	`

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	if _, err = tx.Exec(deleteWorkout, userID, id); err != nil {
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

func loadWorkouts(rows *sql.Rows) (*types.Workout, error) {
	workout := new(types.Workout)
	err := rows.Scan(
		&workout.ID,
		&workout.Name,
		&workout.Exercises,
		&workout.CreatedAt,
	)
	return workout, err
}

func (s *PostgresDB) createUserTables() error {
	createUser := `
		CREATE TABLE IF NOT EXISTS "user" (
			id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			first_name TEXT,
			last_name TEXT,
			user_name TEXT,
			encrypted_password TEXT,
			created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
		);
	`
	createUserWorkout := `
  		CREATE TABLE IF NOT EXISTS user_workout (
   			user_id UUID REFERENCES "user"(id),
    		workout_id UUID REFERENCES workout(id)
  		);
	`

	if _, err := s.db.Exec(createUser); err != nil {
		return err
	}
	if _, err := s.db.Exec(createUserWorkout); err != nil {
		return err
	}
	return nil
}

func (s *PostgresDB) CreateUser(user *types.User) error {
	createUser := `
    	INSERT INTO "user" (
			first_name, 
			last_name, 
			user_name, 
			encrypted_password
		)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at;
	`

	if err := s.db.QueryRow(
		createUser,
		user.FirstName,
		user.LastName,
		user.UserName,
		user.EncryptedPassword,
	).Scan(&user.ID, &user.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (s *PostgresDB) GetUsers() ([]*types.User, error) {
	getUsers := `
    	SELECT * FROM "user";
	`
	rows, err := s.db.Query(getUsers)
	if err != nil {
		return nil, err
	}

	users := []*types.User{}
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

func (s *PostgresDB) GetUserByID(id uuid.UUID) (*types.User, error) {
	getUserByID := `
    	SELECT * FROM "user"
    	WHERE id = $1;
	`

	rows, err := s.db.Query(getUserByID, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return s.loadUsers(rows)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return nil, fmt.Errorf("user %v not found", id)
}

func (s *PostgresDB) GetUserByUsername(username string) (*types.User, error) {
	getUserByUsername := `
		SELECT * FROM "user"
		WHERE user_name = $1;
	`

	rows, err := s.db.Query(getUserByUsername, username)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return s.loadUsers(rows)
	}

	return nil, fmt.Errorf("user %s not found", username)
}

func (s *PostgresDB) getUserWorkoutIDs(userID uuid.UUID) ([]uuid.UUID, error) {
	getUserWorkoutsIDs := `
    	SELECT id 
		FROM workout
    	WHERE user_id = $1;
  	`

	rows, err := s.db.Query(getUserWorkoutsIDs, userID)
	if err != nil {
		return nil, err
	}

	var workoutIDs []uuid.UUID
	for rows.Next() {
		var id uuid.UUID
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		workoutIDs = append(workoutIDs, id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return workoutIDs, nil
}

func (s *PostgresDB) UpdateUser(userID uuid.UUID, user *types.User) (*types.User, error) {
	updateUser := `
    	UPDATE "user"
    	SET 
			first_name = $1, 
			last_name = $2, 
			user_name = $3, 
			encrypted_password = $4 
    	WHERE id = $5;
  	`

	if _, err := s.db.Exec(
		updateUser,
		user.FirstName,
		user.LastName,
		user.UserName,
		user.EncryptedPassword,
		userID,
	); err != nil {
		return nil, err
	}
	return s.GetUserByID(userID)
}

func (s *PostgresDB) DeleteUser(userID uuid.UUID) error {
	deleteUser := `
		DELETE FROM "user"
		WHERE id = $1;
  	`

	workoutIDs, err := s.getUserWorkoutIDs(userID)
	if err != nil {
		return err
	}

	for _, id := range workoutIDs {
		if err := s.DeleteWorkout(userID, id); err != nil {
			return err
		}
	}

	if _, err := s.db.Exec(deleteUser, userID); err != nil {
		return err
	}

	return nil
}

func (s *PostgresDB) loadUsers(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)
	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.UserName,
		&user.EncryptedPassword,
		&user.CreatedAt,
	)
	return user, err
}
