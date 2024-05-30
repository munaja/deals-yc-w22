package view

type GetDetailDto struct {
	Viewer_User_Id    int `json:"-"`
	Viewer_Profile_Id int `json:"-"`
}

type GenerateListDto struct {
	Viewer_User_Id    int `json:"-"`
	Viewer_Profile_Id int `json:"-"`
}
