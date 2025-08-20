package data

import (
    "encoding/json"
    "os"
    "path/filepath"

    "backend/models"
)

func LoadProblems() ([]models.Problem, error) {
    // read JSON file
    data, err := os.ReadFile("data/problems.json")
    if err != nil {
        return nil, err
    }

    var problems []models.Problem
    if err := json.Unmarshal(data, &problems); err != nil {
        return nil, err
    }

    // read actual code from .go files
    for i, p := range problems {
        codeBytes, err := os.ReadFile(filepath.Clean(p.File))
        if err != nil {
            return nil, err
        }
        problems[i].Code = string(codeBytes)
    }

    return problems, nil
}
