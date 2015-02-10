package feedbag

// GithubUser
type GithubUser struct {
	Id        int    `json:"id"`
	Login     string `json:"login"`
	Type      string `json:"type"`
	AvatarUrl string `json:"avatar_url"`
	HtmlUrl   string `json:"html_url"`
}

type Commits []Commit
type Commit struct {
	Sha      string     `json:"sha"`
	Message  string     `json:"message"`
	Auther   GithubUser `json:"auther"`
	Url      string     `json:"url"`
	Distinct bool       `json:"distinct"`
	Commiter GithubUser `json:"commiter"`
}

type Repository struct {
	Id              int        `json:"id"`
	Name            string     `json:"name"`
	FullName        string     `json:"full_name"`
	Owner           GithubUser `json:"owner"`
	Private         bool       `json:"bool"`
	HtmlUrl         string     `json:"html_url"`
	Description     string     `json:"description"`
	Fork            bool       `json:"fork"`
	Homepage        string     `json:"homepage"`
	StargazersCount int        `json:"stargazers_count"`
	WatchersCount   int        `json:"watchers_count"`
	ForksCount      int        `json:"forks_count"`
	OpenIssues      int        `json:"open_issues"`
}

// PullRequest
type PullRequest struct {
	Title          string     `json:"title"`
	Body           string     `json:"body"`
	Commits        int        `json:"commits"`
	Number         int        `json:"number"`
	HtmlUrl        string     `json:"html_url"`
	State          string     `json:"state"`
	Locked         bool       `json:"locked"`
	User           GithubUser `json:"user"`
	Assignee       GithubUser `json:"assignee"`
	Merged         bool       `json:"merged"`
	Mergable       bool       `json:"mergable"`
	MergedBy       GithubUser `json:"merged_by"`
	ChangedFiles   int        `json:"changed_files"`
	ReviewComments int        `json:"review_comments"`
	Repository     Repository `json:"repository"`
	Sender         GithubUser `json:"sender"`
}

// Issue
type Issue struct {
	Id       int        `json:"id"`
	HtmlUrl  string     `json:"html_url"`
	Number   int        `json:"number"`
	Title    string     `json:"title"`
	Body     string     `json:"body"`
	User     GithubUser `json:"user"`
	Labels   Labels     `json:"labels"`
	State    string     `json:"state"`
	Locked   bool       `json:"locked"`
	Assignee GithubUser `json:"assignee"`
	Comments int        `json:"comments"`
}

// Comment
type Comment struct {
	Id      int        `json:"id"`
	HtmlUrl string     `json:"html_url"`
	Body    string     `json:"body"`
	User    GithubUser `json:"user"`
}

type Forkee struct {
	Id         int        `json:"id"`
	Name       string     `json:"name"`
	Owner      GithubUser `json:"owner"`
	Repository Repository `json:"repository"`
	Sender     GithubUser `json:"sender"`
}

// Labels
type Labels []Label
type Label struct {
	Url   string `json:"url"`
	Name  string `json:"name"`
	Color string `json:"color"`
}
