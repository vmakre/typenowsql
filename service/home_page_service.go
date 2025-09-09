package service

type HomePageService interface {
	// Define methods related to home page operations if needed
}

type homePageService struct {
	// Add dependencies if needed
}

func NewHomePageService() HomePageService {
	return &homePageService{}
}
