package models

import (
	"time"
)

//
// USER
//

type User struct {
	CreatedAt time.Time `json:"joined"`
	UpdatedAt time.Time
	LastLogin time.Time `json:"last-login-at"`

	Active bool `json:"active"`

	ID int `json:"user-id" gorm:"primaryKey;unique"`

	ProfilePic   string `json:"profile-pic"`
	LogIPaddrPic string `json:"-"`

	Algorithm string `json:"algorithm"`
	PublicKey string `json:"public-key" gorm:"unique"`

	Title string `json:"title"`

	PostCount int  `json:"post-count"`
	LastPost  Post `json:"last-post" gorm:"type:jsonb"`

	PGABody   string `json:"-"`
	Signature string `json:"signature"`
}

//
// FORUM
//

type Forum struct {
	CreatedAt time.Time `json:"created-at"`
	UpdatedAt time.Time `json:"updated-at"`

	ID    int    `json:"forum-id" gorm:"primaryKey;unique"`
	Title string `json:"forum-title"`

	LastThread Thread `json:"last-thread" gorm:"type:jsonb"`
}

//
// THREAD
//

type Thread struct {
	CreatedAt time.Time `json:"created-at"`
	UpdatedAt time.Time `json:"updated-at"`

	ForumID int `json:"forum-id"`

	ID    int    `json:"thread-id" gorm:"primaryKey;unique"`
	Title string `json:"thread-title"`

	UserID   int  `json:"user-id"`
	LastPost Post `json:"last-post" gorm:"type:jsonb"`
}

//
// POST
//

type Post struct {
	CreatedAt time.Time `json:"posted-at"`
	UpdatedAt time.Time `json:"updated-at"`

	ForumID  int `json:"forum-id"`
	ThreadID int `json:"thread-id"`

	ID int `json:"post-id" gorm:"primaryKey;unique"`

	LogIPaddr string `json:"-"`
	Country   string `json:"country"`
	UserID    int    `json:"user-id"`

	ReplyTo int    `json:"reply-to"`
	Body    string `json:"body"`

	PGABody   string `json:"-"`
	Signature string `json:"signature"`
}

//
// REPORT
//

type Report struct {
	CreatedAt time.Time `json:"reported-at"`

	ID int `json:"report-id" gorm:"primaryKey;unique"`

	ReportedPost Post   `json:"post"`
	Reason       string `json:"reason"`

	ReportedBy int    `json:"reported-by"`
	PGABody    string `json:"-"`
	Signature  string `json:"-"`
}
