package service

import (
	"errors"
	"github.com/aldiaprilianto/takana/dto/response"
	"github.com/aldiaprilianto/takana/entity"
	"github.com/aldiaprilianto/takana/repository"
	"github.com/gin-gonic/gin"
)

type UserService interface {
	GenerateOrganization(ctx *gin.Context) (*response.OrgHierarchyResponse, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (s *userService) GenerateOrganization(ctx *gin.Context) (*response.OrgHierarchyResponse, error) {
	orgStatus := "1" // Menentukan status organisasi yang ingin diambil

	// Mendapatkan semua organisasi dengan status yang diberikan dari repository
	allOrgs, err := s.userRepository.GetOrganizationHierarchy(orgStatus)
	if err != nil {
		return nil, err
	}

	// Membuat map untuk mengindeks organisasi berdasarkan ID
	orgMap := make(map[string]entity.Organization)
	for _, org := range allOrgs {
		orgMap[org.OrgID] = org
	}

	// Mendapatkan organisasi induk berdasarkan ID yang diberikan
	orgId := ctx.Param("org_id")

	// Membangun hirarki organisasi
	rootOrg, err := s.buildOrgHierarchy(orgMap, orgId)
	if err != nil {
		return nil, err
	}

	return rootOrg, nil
}

// Fungsi rekursif untuk membangun hirarki organisasi
func (s *userService) buildOrgHierarchy(orgMap map[string]entity.Organization, orgId string) (*response.OrgHierarchyResponse, error) {
	// Mendapatkan organisasi saat ini dari map
	currentOrg, found := orgMap[orgId]
	if !found {
		return nil, errors.New("organization not found")
	}

	// Inisialisasi response untuk organisasi saat ini
	res := &response.OrgHierarchyResponse{
		OrgID:   currentOrg.OrgID,
		OrgName: currentOrg.OrgName,
	}

	// Membuat slice untuk menyimpan child organizations
	var childOrgs []response.OrgHierarchyResponse

	// Menelusuri semua organisasi untuk mencari anak-anak dari organisasi saat ini
	for _, org := range orgMap {
		if org.OrgParentID == orgId { // Jika organisasi ini adalah anak dari organisasi saat ini
			childOrg, err := s.buildOrgHierarchy(orgMap, org.OrgID)
			if err != nil {
				return nil, err
			}
			if childOrg != nil {
				childOrgs = append(childOrgs, *childOrg)
			}
		}
	}

	// Menetapkan child organizations ke response
	res.OrgChilds = childOrgs

	return res, nil
}
