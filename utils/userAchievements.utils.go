package utils

import (
    "github.com/FinanceUN/Achievements/models"
)

type UserAchievementsResponse struct {
    UserAchievement models.UserAchievement `json:"user_achievement"`
    Achievement     *models.Achievement     `json:"achievement"`
}
