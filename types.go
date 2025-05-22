package main

import "time"

type GithubEventType string

var (
	AllEvent                      GithubEventType = "ALL"
	PushEvent                     GithubEventType = "PushEvent"
	CreateEvent                   GithubEventType = "CreateEvent"
	DeleteEvent                   GithubEventType = "DeleteEvent"
	ForkEvent                     GithubEventType = "ForkEvent"
	PullRequestEvent              GithubEventType = "PullRequestEvent"
	PullRequestReviewEvent        GithubEventType = "PullRequestReviewEvent"
	PullRequestReviewCommentEvent GithubEventType = "PullRequestReviewCommentEvent"
	PullRequestReviewThreadEvent  GithubEventType = "PullRequestReviewThreadEvent"
	WatchEvent                    GithubEventType = "WatchEvent"
	IssuesEvent                   GithubEventType = "IssuesEvent"
	IssueCommentEvent             GithubEventType = "IssueCommentEvent"
	MemberEvent                   GithubEventType = "MemberEvent"
	PublicEvent                   GithubEventType = "PublicEvent"
	ReleaseEvent                  GithubEventType = "ReleaseEvent"
	SponsorshipEvent              GithubEventType = "SponsorshipEvent"
	GollumEvent                   GithubEventType = "GollumEvent"
	CommitCommentEvent            GithubEventType = "CommitCommentEvent"
)

type Repo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Actor struct {
	ID         int    `json:"id"`
	Login      string `json:"login"`
	GravatarID string `json:"gravatar_id"`
	AvatarURL  string `json:"avatar_url"`
	URL        string `json:"url"`
}

type Org struct {
	ID         int    `json:"id"`
	Login      string `json:"login"`
	GravatarID string `json:"gravatar_id"`
	AvatarURL  string `json:"avatar_url"`
	URL        string `json:"url"`
}

type GithubEvent struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Public    bool      `json:"public"`
	Payload   any       `json:"payload"`
	Repo      Repo      `json:"repo"`
	Actor     Actor     `json:"actor"`
	Org       Org       `json:"org"`
	CreatedAt time.Time `json:"created_at"`
}

type AuthorPayload struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type CommitPayload struct {
	Sha      string        `json:"sha"`
	Author   AuthorPayload `json:"author"`
	Message  string        `json:"message"`
	Distinct bool          `json:"distinct"`
	URL      string        `json:"url"`
}

type PushEventPayload struct {
	PushID       int             `json:"push_id"`
	Size         int             `json:"size"`
	DistinctSize int             `json:"distinct_size"`
	Ref          string          `json:"ref"`
	Head         string          `json:"head"`
	Before       string          `json:"before"`
	Commits      []CommitPayload `json:"commits"`
}

type CreateEventPayload struct {
	Ref          string `json:"ref"`
	RefType      string `json:"ref_type"`
	MasterBranch string `json:"master_branch"`
	Description  string `json:"description"`
	PusherType   string `json:"pusher_type"`
}

type Result struct {
	Total int          `json:"total"`
	Data  []ResultData `json:"data"`
}

type ResultData struct {
	Size  int    `json:"size"`
	Type  string `json:"type"`
	Items []any  `json:"items"`
}
