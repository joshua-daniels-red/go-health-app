package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

type ProcessHandler struct {
	pythonPath string
	scriptPath string
}

func NewProcessHandler(pythonPath, scriptPath string) *ProcessHandler {
	return &ProcessHandler{
		pythonPath: pythonPath,
		scriptPath: scriptPath,
	}
}

func (h *ProcessHandler) RunProcess(c *gin.Context) {
	op := c.DefaultQuery("op", "")
	if op == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing op parameter"})
		return
	}

	// Determine base URL
	baseURL := os.Getenv("GO_SERVER_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080" // fallback
	}

	// Prepare Python command
	cmd := exec.Command(h.pythonPath, h.scriptPath, op)
	cmd.Env = append(os.Environ(), "GO_SERVER_URL="+baseURL)

	// Capture output
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"stderr": stderr.String(),
		})
		return
	}

	// Parse Python JSON output
	var result interface{}
	if err := json.Unmarshal(out.Bytes(), &result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "failed to parse python output",
			"stderr": stderr.String(),
			"stdout": out.String(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}


