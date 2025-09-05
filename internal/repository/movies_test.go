package repository

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

// createTempJSONFile is a helper function that creates a temporary JSON file
// with the given content. It returns the path to the file and a cleanup function.
// It uses t.TempDir() which automatically handles directory cleanup after the test.
func createTempJSONFile(t *testing.T, content string) string {
	t.Helper() // Marks this function as a test helper.

	dir := t.TempDir()
	filePath := filepath.Join(dir, "movies.json")
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		// Use t.Fatalf for setup errors, as the test cannot proceed.
		t.Fatalf("Failed to create temporary file %s: %v", filePath, err)
	}
	return filePath
}

func TestNewMovieRepository(t *testing.T) {
	// Define a valid JSON structure for our success case.
	validJSONContent := `[
        {"id": 1, "movie_title": "The Shawshank Redemption", "genre": "Drama", "imdb": 9.3},
        {"id": 2, "movie_title": "The Godfather", "genre": "Crime", "imdb": 9.2}
    ]`

	// Use t.Run to create organized sub-tests.
	t.Run("Success with valid JSON file", func(t *testing.T) {
		jsonPath := createTempJSONFile(t, validJSONContent)

		repo, err := NewMovieRepository(jsonPath)

		if err != nil {
			t.Fatalf("Expected no error, but got: %v", err)
		}
		if repo == nil {
			t.Fatal("Expected a repository instance, but got nil")
		}
		if len(repo.movies) != 2 {
			t.Errorf("Expected repository to have 2 movies, but got %d", len(repo.movies))
		}
	})

	t.Run("Failure with non-existent file", func(t *testing.T) {
		_, err := NewMovieRepository("non_existent_file.json")
		if err == nil {
			t.Fatal("Expected an error for a non-existent file, but got nil")
		}
	})

	t.Run("Failure with invalid JSON content", func(t *testing.T) {
		// This content is intentionally malformed.
		invalidJSONContent := `[{"id": 1, "movie_title": "Incomplete"`
		jsonPath := createTempJSONFile(t, invalidJSONContent)

		_, err := NewMovieRepository(jsonPath)
		if err == nil {
			t.Fatal("Expected a JSON unmarshalling error, but got nil")
		}
	})
}

func TestGetAll(t *testing.T) {
	// Arrange: Set up the repository with known data.
	validJSONContent := `[
        {"id": 10, "movie_title": "Pulp Fiction", "genre": "Crime", "imdb": 8.9}
    ]`
	jsonPath := createTempJSONFile(t, validJSONContent)
	repo, err := NewMovieRepository(jsonPath)
	if err != nil {
		t.Fatalf("Test setup failed. Could not create movie repository: %v", err)
	}

	// Define the expected output.
	expectedMovies := []Movie{
		{ID: 10, Title: "Pulp Fiction", Genre: "Crime", IMDbRating: 8.9},
	}

	// Act: Call the method we are testing.
	resultMovies := repo.GetAll()

	// Assert: Check if the result matches the expectation.
	// reflect.DeepEqual is used to compare slices and structs.
	if !reflect.DeepEqual(resultMovies, expectedMovies) {
		t.Errorf("GetAll() result was incorrect.\nGot:  %v\nWant: %v", resultMovies, expectedMovies)
	}

	// Another test case for an empty repository
	t.Run("Get from empty repository", func(t *testing.T) {
		emptyRepo := &MovieRepository{movies: []Movie{}}
		if len(emptyRepo.GetAll()) != 0 {
			t.Error("GetAll() on an empty repository should return an empty slice")
		}
	})
}
