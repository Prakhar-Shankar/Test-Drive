package services

import (
    "encoding/json"
    "os"
    "backend/models"
)

func LoadProblems() ([]models.Problem, error) {
    file, err := os.ReadFile("data/problems.json")
    if err != nil {
        return nil, err
    }

    var problems []models.Problem
    err = json.Unmarshal(file, &problems)
    if err != nil {
        return nil, err
    }

    return problems, nil
}
