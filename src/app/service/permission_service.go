package service

import (
	"api-gateway/app/constant"
	"api-gateway/app/domain/dao"
	"api-gateway/app/domain/dco"
	"api-gateway/app/pkg"
	"api-gateway/app/repository"
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/google/wire"
	"gorm.io/gorm"
)

type PermissionService interface {
	GetAllPermissions(c *gin.Context)
	GetPermissionById(c *gin.Context)
	AddPermission(c *gin.Context)
	UpdatePermission(c *gin.Context)
	DeletePermission(c *gin.Context)
}

type PermissionServiceImpl struct {
	PermissionRepository repository.PermissionRepository
}

func (p PermissionServiceImpl) GetAllPermissions(c *gin.Context) {
	/* GetAllPermissions is a function to get all permissions
	 * @param c is gin context
	 * @return void
	 */
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program get all permissions")

	rawData, err := p.PermissionRepository.FindAllPermissions()
	if err != nil {
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	data := mapPermissionListToPermissionResponseList(rawData)

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (p PermissionServiceImpl) GetPermissionById(c *gin.Context) {
	/* GetPermissionById is a function to get permission by id
	 * @param c is gin context
	 * @return void
	 */
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program get permission by id")

	id := c.Param("permissionID")
	permissionID, err := uuid.Parse(id)
	if err != nil {
		slog.Error("Error when parsing uuid. Error", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	rawData, err := p.PermissionRepository.FindPermissionById(permissionID)
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.DataNotFound)
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	data := mapPermissionToPermissionResponse(rawData)

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (p PermissionServiceImpl) AddPermission(c *gin.Context) {
	/* AddPermission is a function to add permission
	 * @param c is gin context
	 * @return void
	 */
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program add permission")

	var rawRequest dco.PermissionRequest
	if err := c.ShouldBindJSON(&rawRequest); err != nil {
		slog.Error("Error when binding json", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}
	request := mapPermissionRequestToPermission(rawRequest)

	// Check if permission name already exist
	_, err := p.PermissionRepository.FindPermissionByName(request.Name)
	switch err {
	case nil:
		pkg.PanicException(constant.Conflict)
	case gorm.ErrRecordNotFound:
		break
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	rawData, err := p.PermissionRepository.Save(&request)
	if err != nil {
		slog.Error("Error when saving data to database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	data := mapPermissionToPermissionResponse(rawData)

	c.JSON(http.StatusCreated, pkg.BuildResponse(constant.Success, data))
}

func (p PermissionServiceImpl) UpdatePermission(c *gin.Context) {
	/* UpdatePermission is a function to update permission by id
	 * @param c is gin context
	 * @return void
	 */
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program update permission")

	id := c.Param("permissionID")
	permissionID, err := uuid.Parse(id)
	if err != nil {
		slog.Error("Error when parsing uuid. Error", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	var rawRequest dco.PermissionRequest
	if err := c.ShouldBindJSON(&rawRequest); err != nil {
		slog.Error("Error when binding json", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	request := mapPermissionRequestToPermission(rawRequest)

	oldData, err := p.PermissionRepository.FindPermissionById(permissionID)
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		pkg.PanicException(constant.DataNotFound)
	default:
		slog.Error("Error when fetching data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	oldData.Name = request.Name
	oldData.Description = request.Description

	rawData, err := p.PermissionRepository.Save(&oldData)
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		pkg.PanicException(constant.DataNotFound)
	default:
		slog.Error("Error when saving data to database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	data := mapPermissionToPermissionResponse(rawData)

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (p PermissionServiceImpl) DeletePermission(c *gin.Context) {
	/* DeletePermission is a function to delete permission by id
	 * @param c is gin context
	 * @return void
	 */
	defer pkg.PanicHandler(c)
	slog.Info("start to execute program delete permission")

	id := c.Param("permissionID")
	permissionID, err := uuid.Parse(id)
	if err != nil {
		slog.Error("Error when parsing uuid. Error", "error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	err = p.PermissionRepository.DeletePermissionById(permissionID)
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		pkg.PanicException(constant.DataNotFound)
	default:
		slog.Error("Error when deleting data from database", "error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}

var permissionServiceSet = wire.NewSet(
	wire.Struct(new(PermissionServiceImpl), "*"),
	wire.Bind(new(PermissionService), new(*PermissionServiceImpl)),
)

func mapPermissionToPermissionResponse(permission dao.Permission) dco.PermissionResponse {
	/* mapPermissionToPermissionResponse is a function to map permission to permission response
	 * @param permission is dao.Permission
	 * @return dao.PermissionResponse
	 */

	var permissionDescription string
	if permission.Description.Valid {
		permissionDescription = permission.Description.String
	}

	return dco.PermissionResponse{
		BaseModel: dco.BaseModel{
			ID:        permission.ID,
			CreatedAt: permission.CreatedAt,
			UpdatedAt: permission.UpdatedAt,
		},
		Name:        permission.Name,
		Description: &permissionDescription,
	}
}

func mapPermissionListToPermissionResponseList(permissions []dao.Permission) []dco.PermissionResponse {
	/* mapPermissionsToPermissionResponseList is a function to map permissions to permission response list
	 * @param permissions is []dao.Permission
	 * @return []dao.PermissionResponse
	 */

	var permissionResponseList []dco.PermissionResponse
	for _, permission := range permissions {
		permissionResponseList = append(permissionResponseList, mapPermissionToPermissionResponse(permission))
	}

	return permissionResponseList
}

func mapPermissionRequestToPermission(req dco.PermissionRequest) dao.Permission {
	/* mapPermissionRequestToPermission is a function to map permission request to permission
	 * @param req is dco.PermissionRequest
	 * @return dao.Permission
	 */

	var permissionDescription sql.NullString
	if req.Description != nil {
		permissionDescription = sql.NullString{
			String: *req.Description,
			Valid:  true,
		}
	}

	return dao.Permission{
		Name:        req.Name,
		Description: permissionDescription,
	}
}
