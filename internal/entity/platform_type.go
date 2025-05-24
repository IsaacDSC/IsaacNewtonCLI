package entity

type PlatformType string

func (pt PlatformType) String() string {
	return string(pt)
}

const (
	Java       PlatformType = "Java"
	Go         PlatformType = "Go"
	Python     PlatformType = "Python"
	JavaScript PlatformType = "JavaScript"
	Typescript PlatformType = "Typescript"
	CSharp     PlatformType = "C#"
	Php        PlatformType = "Php"
	Ruby       PlatformType = "Ruby"
	Swift      PlatformType = "Swift"
	CPLusPlus  PlatformType = "C++"
)
