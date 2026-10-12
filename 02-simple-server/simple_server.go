package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"slices"
	"strconv"
	"sync"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type UserStore struct {
	mu     sync.RWMutex
	users  []User
	nextId int
}

var userStore UserStore

var UserNotFoundErr = errors.New("User not found")
var NameBlankErr = errors.New("Name should not be blank")
var BodyParsingError = errors.New("Body parsing error")
var BadRequestError = errors.New("Bad Request")

func (s *UserStore) Create(name string) (User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if name == "" {
		return User{}, NameBlankErr
	}

	s.nextId++
	user := User{
		ID:   s.nextId,
		Name: name,
	}

	s.users = append(s.users, user)
	return user, nil
}

func (s *UserStore) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, user := range s.users {
		if user.ID == id {
			s.users = slices.Delete(s.users, i, i+1)
			return nil
		}
	}

	return UserNotFoundErr
}

func (s *UserStore) GetAll() []User {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return slices.Clone(s.users)
}

func (s *UserStore) GetById(id int) (User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, user := range s.users {
		if user.ID == id {
			return user, nil
		}
	}

	return User{}, UserNotFoundErr
}

func usersHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Allow", "GET, POST")
	switch r.Method {
	case http.MethodGet:
		getAllUsers(w)
	case http.MethodPost:
		createUser(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getAllUsers(w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(userStore.GetAll())
}

func handleUsersById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Allow", "GET, DELETE")

	switch r.Method {
	case http.MethodGet:
		getUserById(w, r)
	case http.MethodDelete:
		deleteUserById(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func deleteUserById(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error: BodyParsingError.Error(),
		})
		return
	}

	delErr := userStore.Delete(userId)

	if delErr != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error: delErr.Error(),
		})
	}
	
	w.WriteHeader(http.StatusNoContent)
}

func getUserById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error: BadRequestError.Error(),
		})
		return
	}

	user, err := userStore.GetById(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error: UserNotFoundErr.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(user)
}

func createUser(w http.ResponseWriter, r *http.Request) {

	user := User{}
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error: BodyParsingError.Error(),
		})
		return
	}

	createdUser, err := userStore.Create(user.Name)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}

func main() {
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/users/{id}", handleUsersById)

	userStore = UserStore{
		users: []User{
			{ID: 1, Name: "Dmitry"},
			{ID: 2, Name: "Alex"},
		},
		nextId: 2,
	}

	println("Server started on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
