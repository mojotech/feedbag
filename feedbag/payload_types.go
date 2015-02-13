package feedbag

// GithubUser
type GithubUser struct {
	Id        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	Login     string `json:"login,omitempty"`
	Type      string `json:"type,omitempty"`
	AvatarUrl string `json:"avatar_url,omitempty"`
	HtmlUrl   string `json:"html_url,omitempty"`
}

type Commits []Commit
type Commit struct {
	Sha      string     `json:"sha,omitempty"`
	Message  string     `json:"message,omitempty"`
	Author   GithubUser `json:"author,omitempty"`
	Url      string     `json:"url,omitempty"`
	Distinct bool       `json:"distinct,omitempty"`
	Commiter GithubUser `json:"commiter,omitempty"`
}

type Repository struct {
	Id              int        `json:"id,omitempty"`
	Name            string     `json:"name,omitempty"`
	FullName        string     `json:"full_name,omitempty"`
	Owner           GithubUser `json:"owner,omitempty"`
	Private         bool       `json:"bool,omitempty"`
	HtmlUrl         string     `json:"html_url,omitempty"`
	Description     string     `json:"description,omitempty"`
	Fork            bool       `json:"fork,omitempty"`
	Homepage        string     `json:"homepage,omitempty"`
	StargazersCount int        `json:"stargazers_count,omitempty"`
	WatchersCount   int        `json:"watchers_count,omitempty"`
	ForksCount      int        `json:"forks_count,omitempty"`
	OpenIssues      int        `json:"open_issues,omitempty"`
}

// PullRequest
type PullRequest struct {
	Title          string     `json:"title,omitempty"`
	Body           string     `json:"body,omitempty"`
	Commits        int        `json:"commits,omitempty"`
	Number         int        `json:"number,omitempty"`
	Head           Location   `json:"head,omitempty"`
	Base           Location   `json:"base,omitempty"`
	HtmlUrl        string     `json:"html_url,omitempty"`
	State          string     `json:"state,omitempty"`
	Locked         bool       `json:"locked,omitempty"`
	User           GithubUser `json:"user,omitempty"`
	Assignee       GithubUser `json:"assignee,omitempty"`
	Merged         bool       `json:"merged,omitempty"`
	Mergable       bool       `json:"mergable,omitempty"`
	MergedBy       GithubUser `json:"merged_by,omitempty"`
	ChangedFiles   int        `json:"changed_files,omitempty"`
	ReviewComments int        `json:"review_comments,omitempty"`
	Sender         GithubUser `json:"sender,omitempty,omitempty"`
}

// Issue
type Issue struct {
	Id       int        `json:"id,omitempty"`
	HtmlUrl  string     `json:"html_url,omitempty"`
	Number   int        `json:"number,omitempty"`
	Title    string     `json:"title,omitempty"`
	Body     string     `json:"body,omitempty"`
	User     GithubUser `json:"user,omitempty"`
	Labels   Labels     `json:"labels,omitempty"`
	State    string     `json:"state,omitempty"`
	Locked   bool       `json:"locked,omitempty"`
	Assignee GithubUser `json:"assignee,omitempty"`
	Comments int        `json:"comments,omitempty"`
}

// Comment
type Comment struct {
	Id      int        `json:"id,omitempty"`
	HtmlUrl string     `json:"html_url,omitempty"`
	Body    string     `json:"body,omitempty"`
	User    GithubUser `json:"user,omitempty"`
}

type Location struct {
	Label string     `json:"label,omitempty"`
	Ref   string     `json:"ref,omitempty"`
	Sha   string     `json:"sha,omitempty"`
	User  GithubUser `json:"user,omitempty"`
	Repo  Repository `json:"repo,omitempty"`
}

type Forkee struct {
	Id         int        `json:"id,omitempty"`
	Name       string     `json:"name,omitempty"`
	Owner      GithubUser `json:"owner,omitempty"`
	Repository Repository `json:"repository,omitempty"`
	Sender     GithubUser `json:"sender,omitempty"`
}

// Labels
type Labels []Label
type Label struct {
	Url   string `json:"url,omitempty"`
	Name  string `json:"name,omitempty"`
	Color string `json:"color,omitempty"`
}
