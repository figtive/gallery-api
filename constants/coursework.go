package constants

const (
	CourseworkTypeProject = "project"
	CourseworkTypeBlog    = "blog"
)

var (
	CourseworkTableName = map[string]string{
		CourseworkTypeProject: "projects",
		CourseworkTypeBlog:    "blogs",
	}
)
