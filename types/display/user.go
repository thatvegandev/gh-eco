package display

type User struct {
	Login           string
	Name            string
	Location        string
	Url             string
	Bio             string
	TwitterUsername string
	WebsiteUrl      string
	SelfRepo        Repo
	FollowersCount  int
	FollowingCount  int
	PinnedRepos     []Repo
	ActivityGraph   struct {
		ContributionsCount int
		Weeks              []WeeklyContribution
	}
}

type WeeklyContribution struct {
	ContributionDays []struct {
		ContributionLevel string
	}
}

type Repo struct {
	Id          string
	Name        string
	Description string
	StarsCount  int
	Url         string
	Owner       struct {
		Login string
	}
	Readme          Blob
	PrimaryLanguage struct {
		Name  string
		Color string
	}
}

type Blob struct {
	Text string
}