package main

import (
    "time"
)

type User struct {
    ID                      int         `json:"id"`
    LastName                string      `json:"last_name"`
    FirstName               string      `json:"first_name"`
    Email                   string      `json:"email"`
    Username                string      `json:"username,omitempty"`
    Locale                  string      `json:"locale,omitempty"`
    CreatedAt               time.Time   `json:"created_at"`
    UpdatedAt               time.Time   `json:"updated_at"`
    LastSignInAt            time.Time   `json:"last_sign_in_at"`
    AccountExpires          string      `json:"account_expires"`
    SignInCount             int         `json:"sign_in_count"`
    Enabled                 bool        `json:"enabled"`
    TimezoneID              string      `json:"timezone_id,omitempty"`
    UserType                string      `json:"user_type"`
    CanEnroll               bool        `json:"can_enroll,omitempty"`
    CanUnenrollUsers        bool        `json:"can_unenroll_users,omitempty"`
    CanMarkComplete         bool        `json:"can_mark_complete,omitempty"`
    CanMoveGroups           bool        `json:"can_move_groups,omitempty"`
    CanActAsInstructor      bool        `json:"can_act_as_instructor,omitempty"`
    CanManageTrainings      bool        `json:"can_manage_trainings,omitempty"`
    CanManageSessions       bool        `json:"can_manage_sessions,omitempty"`
    TutorCanEditTheirCourses bool       `json:"tutor_can_edit_their_courses,omitempty"`
    TutorCanCreateCourses   bool        `json:"tutor_can_create_courses,omitempty"`
    NumberOfEnrollments     int         `json:"number_of_enrollments"`
    NumberOfEnrollmentsAccessed int     `json:"number_of_enrollments_accessed"`
    CustomData              interface{} `json:"custom_data,omitempty"`
    MembershipType          string      `json:"membership_type,omitempty"`
    NumberOfPoints          int         `json:"number_of_points,omitempty"`
    NumberOfBadges          int         `json:"number_of_badges,omitempty"`
    SfContactID             string      `json:"sf_contact_id,omitempty"`
    SfUserID                string      `json:"sf_user_id,omitempty"`
    IsSalesforceContact     bool        `json:"is_salesforce_contact,omitempty"`
    CustomDataFieldValues   interface{} `json:"custom_data_field_values,omitempty"`
}

var users = map[int]User{
    123123: {
        ID:                      123123,
        LastName:                "Doe",
        FirstName:               "John",
        Email:                   "john.doe@example.com",
        CreatedAt:               time.Now(),
        UpdatedAt:               time.Now(),
        LastSignInAt:            time.Now(),
        AccountExpires:          "2024-04-11",
        SignInCount:             10,
        Enabled:                 true,
        UserType:                "learner",
        NumberOfEnrollments:     5,
        NumberOfEnrollmentsAccessed: 3,
    },
}
